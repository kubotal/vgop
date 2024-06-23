/*
Copyright © 2023 Red Hat, Inc.
Copyright 2024 Kubotal SAS.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"vgop/internal/controllers/vgmanager/dmsetup"
	"vgop/internal/controllers/vgmanager/filter"
	"vgop/internal/controllers/vgmanager/lsblk"
	"vgop/internal/controllers/vgmanager/lvm"
	"vgop/internal/controllers/vgmanager/wipefs"
	"vgop/internal/global"
	"vgop/internal/misc"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	vgopv1alpha1 "vgop/api/vgop/v1alpha1"
	"vgop/internal/controllers/vgmanager"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(vgopv1alpha1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

// NoManagedFields removes the managedFields from the object.
// This is used to reduce memory usage of the objects in the cache.
// This MUST NOT be used for SSA.
func NoManagedFields(i any) (any, error) {
	it, ok := i.(client.Object)
	if !ok {
		return nil, fmt.Errorf("unexpected object type: %T", i)
	}
	it.SetManagedFields(nil)
	return it, nil
}

func main() {
	var metricsAddr string
	var probeAddr string
	var secureMetrics bool
	var enableHTTP2 bool
	var logConfig misc.LogConfig

	flag.StringVar(&logConfig.Mode, "logMode", "json", "Log mode: 'dev' or 'json'")
	flag.StringVar(&logConfig.Level, "logLevel", "info", "Log level")
	flag.StringVar(&metricsAddr, "metrics-bind-address", "0", "The address the metric endpoint binds to. Use the port :8080. If not set, it will be 0 in order to disable the metrics server")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&secureMetrics, "metrics-secure", false, "If set the metrics endpoint is served securely")
	flag.BoolVar(&enableHTTP2, "enable-http2", false, "If set, HTTP/2 will be enabled for the metrics server")

	flag.Parse()

	logger, err := misc.HandleLog(&logConfig)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to set logging configuration: %v\n", err)
		os.Exit(2)
	}
	//opts := zap.Options{
	//	Development: true,
	//}
	//opts.BindFlags(flag.CommandLine)

	ctrl.SetLogger(logger)

	logger.Info("vgop controller start", "version", global.Version, "build", global.BuildTs, "logLevel", logConfig.Level)

	// if the enable-http2 flag is false (the default), http/2 should be disabled
	// due to its vulnerabilities. More specifically, disabling http/2 will
	// prevent from being vulnerable to the HTTP/2 Stream Cancellation and
	// Rapid Reset CVEs. For more information see:
	// - https://github.com/advisories/GHSA-qppj-fm5r-hxr3
	// - https://github.com/advisories/GHSA-4374-p667-p6c8
	disableHTTP2 := func(c *tls.Config) {
		setupLog.Info("disabling http/2")
		c.NextProtos = []string{"http/1.1"}
	}

	tlsOpts := make([]func(*tls.Config), 0)

	if !enableHTTP2 {
		tlsOpts = append(tlsOpts, disableHTTP2)
	}

	//webhookServer := webhook.NewServer(webhook.Options{
	//	TLSOpts: tlsOpts,
	//})
	operatorNamespace, err := getOperatorNamespace()
	if err != nil {
		setupLog.Error(err, "unable to get operatorNamespace")
		os.Exit(1)
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
		Metrics: metricsserver.Options{
			BindAddress:   metricsAddr,
			SecureServing: secureMetrics,
			TLSOpts:       tlsOpts,
		},
		//WebhookServer:          webhookServer,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         false, // We will be in a daemonSet
		Cache: cache.Options{
			DefaultTransform: NoManagedFields,
			ByObject: map[client.Object]cache.ByObject{
				&vgopv1alpha1.LVMVolumeGroup{}:           {Namespaces: map[string]cache.Config{operatorNamespace: {}}},
				&vgopv1alpha1.LVMVolumeGroupNodeStatus{}: {Namespaces: map[string]cache.Config{operatorNamespace: {}}},
			},
		},
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	theLvm := lvm.NewDefaultHostLVM()
	nodeName := os.Getenv("NODE_NAME")

	if err = (&vgmanager.Reconciler{
		Client:        mgr.GetClient(),
		EventRecorder: mgr.GetEventRecorderFor(vgmanager.ControllerName),
		Scheme:        mgr.GetScheme(),
		LSBLK:         lsblk.NewDefaultHostLSBLK(),
		Wipefs:        wipefs.NewDefaultHostWipefs(),
		Dmsetup:       dmsetup.NewDefaultHostDmsetup(),
		LVM:           theLvm,
		NodeName:      nodeName,
		Namespace:     operatorNamespace,
		Filters:       filter.DefaultFilters,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "LVMVolumeGroup")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

const OperatorNamespaceEnvVar = "NAMESPACE"

// GetOperatorNamespace returns the Namespace the operator should be watching for changes
func getOperatorNamespace() (string, error) {
	// The env variable NAMESPACE which specifies the Namespace the pod is running in
	// and hence will watch.

	ns, found := os.LookupEnv(OperatorNamespaceEnvVar)
	if !found {
		return "", fmt.Errorf("%s not found", OperatorNamespaceEnvVar)
	}
	return ns, nil
}
