---
title: "CNI plugin configuration"
linkTitle: "CNI"
weight: 20
description: >
 EKS Anywhere cluster yaml cni plugin specification reference
---

### Specifying CNI Plugin in EKS Anywhere cluster spec

EKS Anywhere currently supports two CNI plugins: Cilium and Kindnet. Only one of them can be selected
for a cluster, and the plugin cannot be changed once the cluster is created.
Up until the 0.7.x releases, the plugin had to be specified using the `cni` field on cluster spec.
Starting with release 0.8, the plugin should be specified using the new `cniConfig` field as follows:

- For selecting Cilium as the CNI plugin:
    ```yaml
    apiVersion: anywhere.eks.amazonaws.com/v1alpha1
    kind: Cluster
    metadata:
      name: my-cluster-name
    spec:
      clusterNetwork:
        pods:
          cidrBlocks:
          - 192.168.0.0/16
        services:
          cidrBlocks:
          - 10.96.0.0/12
        cniConfig:
          cilium: {}
    ```
    EKS Anywhere selects this as the default plugin when generating a cluster config.

- Or for selecting Kindnetd as the CNI plugin:
    ```yaml
    apiVersion: anywhere.eks.amazonaws.com/v1alpha1
    kind: Cluster
    metadata:
      name: my-cluster-name
    spec:
      clusterNetwork:
        pods:
          cidrBlocks:
          - 192.168.0.0/16
        services:
          cidrBlocks:
          - 10.96.0.0/12
        cniConfig:
          kindnetd: {}
    ```

> NOTE: EKS Anywhere allows specifying only 1 plugin for a cluster and does not allow switching the plugins
after the cluster is created.

### Configuration options for Cilium plugin

Cilium accepts policy enforcement modes from the users to determine the allowed traffic between pods.
The allowed values for this mode are: `default`, `always` and `never`.
Please refer the official [Cilium documentation](https://docs.cilium.io/en/stable/policy/intro/) for more details on how each mode affects
the communication within the cluster and choose a mode accordingly.
You can choose to not set this field so that cilium will be launched with the `default` mode.
Starting release 0.8, Cilium's policy enforcement mode can be set through the cluster spec
as follows:

```yaml
apiVersion: anywhere.eks.amazonaws.com/v1alpha1
kind: Cluster
metadata:
  name: my-cluster-name
spec:
  clusterNetwork:
    pods:
      cidrBlocks:
      - 192.168.0.0/16
    services:
      cidrBlocks:
      - 10.96.0.0/12
    cniConfig:
      cilium: 
        policyEnforcementMode: "always"
```

Please note that if the `always` mode is selected, all communication between pods is blocked unless
NetworkPolicy objects allowing communication are created.
In order to ensure that the cluster gets created successfully, EKS Anywhere will create the required
NetworkPolicy objects for all its core components. But it is up to the user to create the NetworkPolicy
objects needed for the user workloads once the cluster is created.

#### Network policies created by EKS Anywhere for "always" mode

As mentioned above, if Cilium is configured with `policyEnforcementMode` set to `always`,
EKS Anywhere creates NetworkPolicy objects to enable communication between
its core components. These policies are created based on the type of cluster as follows:

1. For self-managed/management cluster, EKS Anywhere will create NetworkPolicy resources in the following namespaces allowing all ingress/egress traffic by default:
    - kube-system
    - eksa-system
    - All core Cluster API namespaces:
        + capi-system
        + capi-kubeadm-bootstrap-system
        + capi-kubeadm-control-plane-system
        + etcdadm-bootstrap-provider-system
        + etcdadm-controller-system
        + cert-manager
    - Infrastruture provider's namespace (for instance, capd-system OR capv-system)
    - If Gitops is enabled, then the gitops namespace (flux-system by default)
    
    This is the NetworkPolicy that will be created in these namespaces for the self-managed cluster:
    ```yaml
    apiVersion: networking.k8s.io/v1
    kind: NetworkPolicy
    metadata:
      name: allow-all-ingress-egress
      namespace: test
    spec:
      podSelector: {}
      ingress:
      - {}
      egress:
      - {}
      policyTypes:
      - Ingress
      - Egress
    ```

2. For a workload cluster managed by another EKS Anywhere cluster, EKS Anywhere will create NetworkPolicy resource only in the following namespace by default:
    - kube-system
    
    For the workload clusters using Kubernetes version 1.21 and higher, the ingress/egress of pods in the kube-system namespace will be limited
    to other pods only in the kube-system namespace by using the following NetworkPolicy:
    
    ```yaml
    apiVersion: networking.k8s.io/v1
    kind: NetworkPolicy
    metadata:
      name: allow-all-ingress-egress
      namespace: test
    spec:
      podSelector: {}
      ingress:
      - from:
        - namespaceSelector:
            matchLabels:
              kubernetes.io/metadata.name: kube-system
      egress:
      - to:
        - namespaceSelector:
            matchLabels:
              kubernetes.io/metadata.name: kube-system
      policyTypes:
      - Ingress
      - Egress
    ```
    
    For workload clusters using Kubernetes version 1.20, the NetworkPolicy in kube-system will
    allow ingress/egress from all pods. This is because Kubernetes versions prior to 1.21 do not
    set the default labels on the namespaces so EKS Anywhere cannot use a namespace selector.
    This NetworkPolicy will ensure that the cluster gets created successfully. Later the cluster admin can edit/replace it if required.

#### Switching the Cilium policy enforcement mode

The policy enforcement mode for Cilium can be changed as a part of cluster upgrade
through the cli upgrade command.
1. Switching to `always` mode: When switching from `default`/`never` to `always` mode,
EKS Anywhere will create the required NetworkPolicy objects for its core components (listed above).
   This will ensure that the cluster gets upgraded successfully. But it is up to the user to create
   the NetworkPolicy objects required for the user workloads.
   
2. Switching from `always` mode: When switching from `always` to `default` mode, EKS Anywhere
will not delete any of the existing NetworkPolicy objects, including the ones required
   for EKS Anywhere components (listed above). The user must delete NetworkPolicy objects as needed.