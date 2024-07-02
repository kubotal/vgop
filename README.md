# VGOP: LVM Volumes Group Operator

`vgop` is an operator aimed to manage LVM volumes group in a kubernetes cluster.

Although of general usage, it was initially designed as a companion tool for 
[topolvm](https://github.com/topolvm/topolvm), to allow its fully automated deployment.

Most of the code and the logic has been borrowed from [lvm-operator](https://github.com/openshift/lvm-operator). 
In fact, this project is mostly a repackaging of a part of this wider project. 


<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Index

- [Installation](#installation)
- [Usage](#usage)
  - [VG creation at installation.](#vg-creation-at-installation)
  - [Status](#status)
- [Configuration](#configuration)
  - [nodeAffinity](#nodeaffinity)
  - [nodeSelector](#nodeselector)
  - [Storage on the control plane](#storage-on-the-control-plane)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Installation

`vgop` should be installed using helm. An OCI chart is provided, so installation (here in the `topolvm-system` namespace) 
can be performed as the following.

```
helm -n topolvm-system install vgop --create-namespace oci://quay.io/kubotal/charts/vgop --version=0.1.0-snapshot
```

## Usage

Once installed, creating LVM VolumeGroup is achieved by creating appropriate K8S resources. 

> NB: For security reasons, such resources must be created in the `vgop` deployment namespace. Otherwise, they will be ignored

For example, to create the VG 'myvg1' on a set of three workers, named 'w1', 'w2' and 'w3'

```
---
apiVersion: vgop.kubotal.io/v1alpha1
kind: LVMVolumeGroup
metadata:
  name: myvg1
  namespace: topolvm-system
spec:
  selectors:
  - deviceSelector:
      paths:
        - /dev/sdc
        - /dev/sdd
    nodeSelector:
      nodeSelectorTerms:
        - matchExpressions:
            - key: kubernetes.io/hostname
              operator: In
              values:
                - w1
                - w2
                - w3
```

In this case, the three workers are homogeneous, in terms of block devices naming and distribution.

Now, let say 'w3' has been setup differently, with only one device available. This can be handled as the following:

```
---
apiVersion: vgop.kubotal.io/v1alpha1
kind: LVMVolumeGroup
metadata:
  name: myvg1
  namespace: topolvm-system
spec:
  selectors:
  - deviceSelector:
      paths:
        - /dev/sdc
        - /dev/sdd
    nodeSelector:
      nodeSelectorTerms:
        - matchExpressions:
            - key: kubernetes.io/hostname
              operator: In
              values:
                - w1
                - w2
  - deviceSelector:
      paths:
        - /dev/sdc
    nodeSelector:
      nodeSelectorTerms:
        - matchExpressions:
            - key: kubernetes.io/hostname
              operator: In
              values:
                - w3
```

> WARNING: Currently, there is no validation webhook. This means most errors, such as using the same device several 
times, or wrong node name, or any other incoherency will not be detected upfront, and will leads some unpredictable results. 

### VG creation at installation.

It is also possible to create such resources during helm chart deployment, be adding properties in a local 
`values.yaml` file. for example:

```
volumeGroups: 
  - name: myvg1
    selectors:
      - devicePaths:
          - /dev/sdc
          - /dev/sdd
        nodes:
          - w1
          - w2
      - devicePaths:
          - /dev/sdc
        nodes:
          - w3
```

### Status

`vgop` create maintains also a set of status ressources: `LVMVolumeGroupNodeStatus`. There will be an entry per node.

You can use `jsonpath` to retrieve usable information. For example:

```
kubectl -n topolvm-system get lvmvolumegroupnodestatuses -o=jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.spec}{"\n"}{end}'
```

## Configuration

`vgop`is deployed as a `daemonset`. By default, there will be a pod on each node of the cluster, except on the control plane.

The `daemonset` configuration could be adjusted to be deployed only on the nodes hosting some storage. This could be achieved 
by setting appropriate properties in a local `values.yaml` file on the helm chart deployment

### nodeAffinity

For example, to designate the nodes by their names:  

```
nodeAffinity:
  requiredDuringSchedulingIgnoredDuringExecution:
    nodeSelectorTerms:
      - matchExpressions:
          - key: kubernetes.io/hostname
            operator: In
            values:
              - w1
              - w2
              - w1

```

### nodeSelector

Or to deploy only on nodes with a specific label:

```
nodeSelector:
  lvm.aware: "true"
```

### Storage on the control plane

Due to the 'taint/toleration' mechanism of K8S, `vgop` is not deployed on the control plane nodes. If you need to do so 
(For example for a single node test cluster), you must add some `tolerations` in a local `values.yaml` file:

```
tolerations:
  - key: node-role.kubernetes.io/control-plane
    operator: Exists
    effect: NoSchedule
  - key: node-role.kubernetes.io/master
    operator: Exists
    effect: NoSchedule
```

A usual with any helm chart, a close look inside its values.yaml file will provide more information.

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

