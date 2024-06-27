# VGOP: LVM Volumes Group Operator

`vgop` is an operator aimed to manage LVM volumes group in a kubernetes cluster.

Although of general usage, it was initially designed as a companion tool for [topolvm](https://github.com/topolvm/topolvm), to allow its fully automated deployment.

Most of the code and the logic has been borrowed from [lvm-operator](https://github.com/openshift/lvm-operator). 
In fact, this project is mostly a repackaging of a part of this wider project.  

## Installation

`vgop` should be installed using helm. An OCI chart is provided, so installation (here in the `topolvm-system` namespace) 
can be performed as the following.

```
helm -n topolvm-system install vgop --create-namespace oci://quay.io/kubotal/charts/vgop --version=0.1.0-snapshot
```

## Usage

Once installed, creating LVM VolumeGroup is achieved by creating appropriate K8S resources. For example:

```
---
apiVersion: vgop.kubotal.io/v1alpha1
kind: LVMVolumeGroup
metadata:
  name: myvg1
  namespace: topolvm-system
spec:
  deviceSelector:
    paths:
      - /dev/sdc
      - /dev/sdd
  nodeSelector:
    nodeSelectorTerms:
      - matchExpressions:
          - key: kubernetes.io/hostname
            operator: In
            values:
              - n0
```

or

```
---
apiVersion: vgop.kubotal.io/v1alpha1
kind: LVMVolumeGroup
metadata:
  name: myvg2
  namespace: topolvm-system
spec:
  deviceSelector:
    paths:
      - /dev/sde
  nodeSelector:
    nodeSelectorTerms:
      - matchFields:
          - key: metadata.name
            operator: In
            values:
              - n0
```

## Configuration

`vgop`is deployed as a daemonset. By default, there will be a pod on each node of the cluster, except on the control plane.

daemonset configuration should be adjusted to deploy a pod only on the node hosting some storage.

### Storage on the control plane

Due to the 'taint/toleration' mechanism of K8S, `vgop` is not deployed on the control plane nodes. If you need to do so 
(For example for a single node test cluster), you must add some 'tolerations' in a local value file:

```
tolerations:
  - key: node-role.kubernetes.io/control-plane
    operator: Exists
    effect: NoSchedule
  - key: node-role.kubernetes.io/master
    operator: Exists
    effect: NoSchedule
```

### Deploy only on specific nodes

TODO

## License

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

