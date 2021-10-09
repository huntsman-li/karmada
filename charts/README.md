# Karmada

Karmada (Kubernetes Armada) is a Kubernetes management system that enables you to run your cloud-native applications across multiple Kubernetes clusters and clouds, with no changes to your applications. By speaking Kubernetes-native APIs and providing advanced scheduling capabilities, Karmada enables truly open, multi-cloud Kubernetes.

Karmada aims to provide turnkey automation for multi-cluster application management in multi-cloud and hybrid cloud scenarios, with key features such as centralized multi-cloud management, high availability, failure recovery, and traffic scheduling.

## TL;DR

Switch to the `root` directory of the repo.
```console
$ helm install karmada -n karmada-system --create-namespace ./charts
```

## Prerequisites

- Kubernetes 1.16+
- helm v3+

## Installing the Chart

To install the chart with the release name `karmada` in namespace `karmada-system`:

Switch to the `root` directory of the repo.
```console
$ helm install karmada -n karmada-system --create-namespace ./charts
```

Get kubeconfig from the cluster:

```console
$ kubectl get secret -n karmada-system karmada-kubeconfig -o jsonpath={.data.kubeconfig} | base64 -d
```

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart
To uninstall/delete the `karmada` helm release in namespace `karmada-system`:

```console
$ helm uninstall karmada -n karmada-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Example
### 1. Install agent
Edited values.yaml
```YAML
installMode: "agent"
agent:
  clusterName: "member"
  ## kubeconfig of the karmada
  kubeconfig:
    caCrt: |
      -----BEGIN CERTIFICATE-----
      XXXXXXXXXXXXXXXXXXXXXXXXXXX
      -----END CERTIFICATE-----
    crt: |
      -----BEGIN CERTIFICATE-----
      XXXXXXXXXXXXXXXXXXXXXXXXXXX
      -----END CERTIFICATE-----
    key: |
      -----BEGIN RSA PRIVATE KEY-----
      XXXXXXXXXXXXXXXXXXXXXXXXXXX
      -----END RSA PRIVATE KEY-----
    server: "https://apiserver.karmada"
```
Execute command (switch to the `root` directory of the repo, and sets the `current-context` in a kubeconfig file)
```console
$ kubectl config use-context member
$ helm install karmada-agent -n karmada-system --create-namespace ./charts
```
### 2. Install component
Edited values.yaml
```YAML
installMode: "component"
components: [
  "schedulerEstimator"
]
schedulerEstimator:
  clusterName: "member"
  ## kubeconfig of the member cluster
  kubeconfig:
    caCrt: |
      -----BEGIN CERTIFICATE-----
      XXXXXXXXXXXXXXXXXXXXXXXXXXX
      -----END CERTIFICATE-----
    crt: |
      -----BEGIN CERTIFICATE-----
      XXXXXXXXXXXXXXXXXXXXXXXXXXX
      -----END CERTIFICATE-----
    key: |
      -----BEGIN RSA PRIVATE KEY-----
      XXXXXXXXXXXXXXXXXXXXXXXXXXX
      -----END RSA PRIVATE KEY-----
    server: "https://apiserver.member"
```
Execute command (switch to the `root` directory of the repo, and sets the `current-context` in a kubeconfig file)
```console
$ kubectl config use-context host
$ helm install karmada-scheduler-estimator -n karmada-system ./charts
```
## Configuration
| Name                      | Description                                     | Value |
| ------------------------- | ----------------------------------------------- | ----- |
| `installMode` | InstallMode "host", "agent" and "component" are provided, "host" means install karmada in the control-cluster, "agent" means install agent client in the member cluster, "component" means install selected components in the control-cluster   | `"host"`|
| `clusterDomain` | Default cluster domain for karmada | `"cluster.local"`  |
| `components` | Selected components list, selectable values: "schedulerEstimator"  | `[]`  |
|`certs.mode`| Mode "auto" and "custom" are provided, "auto" means auto generate certificate, "custom" means use user certificate |`"auto"`|
|`certs.auto.expiry`| Expiry of the certificate |`"43800h"`|
|`certs.auto.hosts`| Hosts of the certificate |`["kubernetes.default.svc","*.etcd.karmada-system.svc.cluster.local","*.karmada-system.svc.cluster.local","*.karmada-system.svc","localhost","127.0.0.1"]`|
|`certs.custom.caCrt`|CA CRT of the certificate|`""`|
|`certs.custom.crt`|CRT of the certificate|`""`|
|`certs.custom.key`|KEY of the certificate|`""`|
|`etcd.mode`| Mode "external" and "internal" are provided, "external" means use external ectd, "internal" means install a etcd in the cluster |`"internal"`|
|`etcd.external.servers`| Servers of etcd |`""`|
|`etcd.external.registryPrefix`| Use to registry prefix of etcd |`"/registry/karmada"`|
|`etcd.external.certs.caCrt`| CA CRT of the etcd certificate |`""`|
|`etcd.external.certs.crt`| CRT of the etcd certificate |`""`|
|`etcd.external.certs.key`| KEY of the etcd certificate |`""`|
|`etcd.internal.replicaCount`| Target replicas of the etcd |`1`|
|`etcd.internal.image.repository`| Image of the etcd |`"k8s.gcr.io/etcd"`|
|`etcd.internal.image.pullPolicy`| Image pull policy of the etcd |`"IfNotPresent"`|
|`etcd.internal.image.tag`| Image tag of the etcd |`"3.4.13-0"`|
|`agent.clusterName`| Name of the member cluster |`""`|
|`agent.kubeconfig.caCrt`| CA CRT of the karmada certificate |`""`|
|`agent.kubeconfig.crt`| CRT of the karmada certificate |`""`|
|`agent.kubeconfig.key`| KEY of the karmada certificate |`""`|
|`agent.kubeconfig.server`| API-server of the karmada |`""`|
|`agent.labels`| Labels of the agent deployment |`{"app": "karmada-agent"}`|
|`agent.replicaCount`| Target replicas of the agent |`1`|
|`agent.podLabels`| Labels of the agent pods |`{}`|
|`agent.podAnnotations`| Annotaions of the agent pods |`{}`|
|`agent.imagePullSecrets`| Image pull secret of the agent |`[]`|
|`agent.image.repository`| Image of the agent |`"swr.ap-southeast-1.myhuaweicloud.com/karmada/karmada-agent"`|
|`agent.image.tag`| Image tag of the agent |`"latest"`|
|`agent.image.pullPolicy`| Image pull policy of the agent |`"IfNotPresent"`|
|`agent.resources`| Resource quota of the agent |`{}`|
|`agent.nodeSelector`| Node selector of the agent |`{}`|
|`agent.affinity`| Affinity of the agent |`{}`|
|`agent.tolerations`| Tolerations of the agent |`[]`|
|`scheduler.labels`| Labels of the schedeler deployment |`{"app": "karmada-scheduler"}`|
|`scheduler.replicaCount`| Target replicas of the scheduler |`1`|
|`scheduler.podLabels`| Labels of the scheduler pods |`{}`|
|`scheduler.podAnnotations`| Annotaions of the scheduler pods |`{}`|
|`scheduler.imagePullSecrets`| Image pull secret of the scheduler |`[]`|
|`scheduler.image.repository`| Image of the scheduler |`"swr.ap-southeast-1.myhuaweicloud.com/karmada/karmada-scheduler"`|
|`scheduler.image.tag`| Image tag of the scheduler |`"latest"`|
|`scheduler.image.pullPolicy`| Image pull policy of the scheduler |`"IfNotPresent"`|
|`scheduler.resources`| Resource quota of the scheduler |`{}`|
|`scheduler.nodeSelector`| Node selector of the scheduler |`{}`|
|`scheduler.affinity`| Affinity of the scheduler |`{}`|
|`scheduler.tolerations`| Tolerations of the scheduler |`[]`|
|`webhook.labels`| Labels of the webhook deployment |`{"app": "karmada-webhook"}`|
|`webhook.replicaCount`| Target replicas of the webhook |`1`|
|`webhook.podLabels`| Labels of the webhook pods |`{}`|
|`webhook.podAnnotations`| Annotaions of the webhook pods |`{}`|
|`webhook.imagePullSecrets`| Image pull secret of the webhook |`[]`|
|`webhook.image.repository`| Image of the webhook |`"swr.ap-southeast-1.myhuaweicloud.com/karmada/karmada-webhook"`|
|`webhook.image.tag`| Image tag of the webhook |`"latest"`|
|`webhook.image.pullPolicy`| Image pull policy of the webhook |`"IfNotPresent"`|
|`webhook.resources`| Resource quota of the webhook |`{}`|
|`webhook.nodeSelector`| Node selector of the webhook |`{}`|
|`webhook.affinity`| Affinity of the webhook |`{}`|
|`webhook.tolerations`| Tolerations of the webhook |`[]`|
|`controllerManager.labels`| Labels of the karmada-controller-manager deployment |`{"app": "karmada-controller-manager"}`|
|`controllerManager.replicaCount`| Target replicas of the karmada-controller-manager |`1`|
|`controllerManager.podLabels`| Labels of the karmada-controller-manager pods |`{}`|
|`controllerManager.podAnnotations`| Annotaions of the karmada-controller-manager pods |`{}`|
|`controllerManager.imagePullSecrets`| Image pull secret of the karmada-controller-manager |`[]`|
|`controllerManager.image.repository`| Image of the karmada-controller-manager |`"swr.ap-southeast-1.myhuaweicloud.com/karmada/karmada-controller-manager"`|
|`controllerManager.image.tag`| Image tag of the karmada-controller-manager |`"latest"`|
|`controllerManager.image.pullPolicy`| Image pull policy of the karmada-controller-manager |`"IfNotPresent"`|
|`controllerManager.resources`| Resource quota of the karmada-controller-manager |`{}`|
|`controllerManager.nodeSelector`| Node selector of the karmada-controller-manager |`{}`|
|`controllerManager.affinity`| Affinity of the karmada-controller-manager |`{}`|
|`controllerManager.tolerations`| Tolerations of the karmada-controller-manager |`[]`|
|`apiServer.labels`| Labels of the karmada-apiserver deployment |`{"app": "karmada-apiserver"}`|
|`apiServer.replicaCount`| Target replicas of the karmada-apiserver |`1`|
|`apiServer.podLabels`| Labels of the karmada-apiserver pods |`{}`|
|`apiServer.podAnnotations`| Annotaions of the karmada-apiserver pods |`{}`|
|`apiServer.imagePullSecrets`| Image pull secret of the karmada-apiserver |`[]`|
|`apiServer.image.repository`| Image of the karmada-apiserver |`"k8s.gcr.io/kube-apiserver"`|
|`apiServer.image.tag`| Image tag of the karmada-apiserver |`"v1.19.1"`|
|`apiServer.image.pullPolicy`| Image pull policy of the karmada-apiserver |`"IfNotPresent"`|
|`apiServer.resources`| Resource quota of the karmada-apiserver |`{}`|
|`apiServer.nodeSelector`| Node selector of the karmada-apiserver |`{}`|
|`apiServer.affinity`| Affinity of the karmada-apiserver |`{}`|
|`apiServer.tolerations`| Tolerations of the karmada-apiserver |`[]`|
|`kubeControllerManager.labels`| Labels of the kube-controller-manager deployment |`{"app": "kube-controller-manager"}`|
|`kubeControllerManager.replicaCount`| Target replicas of the kube-controller-manager |`1`|
|`kubeControllerManager.podLabels`| Labels of the kube-controller-manager pods |`{}`|
|`kubeControllerManager.podAnnotations`| Annotaions of the kube-controller-manager pods |`{}`|
|`kubeControllerManager.imagePullSecrets`| Image pull secret of the kube-controller-manager |`[]`|
|`kubeControllerManager.image.repository`| Image of the kube-controller-manager |`"k8s.gcr.io/kube-controller-manager"`|
|`kubeControllerManager.image.tag`| Image tag of the kube-controller-manager |`"v1.19.1"`|
|`kubeControllerManager.image.pullPolicy`| Image pull policy of the kube-controller-manager |`"IfNotPresent"`|
|`kubeControllerManager.resources`| Resource quota of the kube-controller-manager |`{}`|
|`kubeControllerManager.nodeSelector`| Node selector of the kube-controller-manager |`{}`|
|`kubeControllerManager.affinity`| Affinity of the kube-controller-manager |`{}`|
|`kubeControllerManager.tolerations`| Tolerations of the kube-controller-manager |`[]`|
|`schedulerEstimator.clusterName`| Name of the member cluster |`""`|
|`schedulerEstimator.kubeconfig.caCrt`| CA CRT of the certificate |`""`|
|`schedulerEstimator.kubeconfig.crt`| CRT of the certificate |`""`|
|`schedulerEstimator.kubeconfig.key`| KEY of the certificate |`""`|
|`schedulerEstimator.kubeconfig.server`| API-server of the member cluster |`""`|
|`schedulerEstimator.labels`| Labels of the scheduler-estimator deployment |`{}`|
|`schedulerEstimator.replicaCount`| Target replicas of the scheduler-estimator |`1`|
|`schedulerEstimator.podLabels`| Labels of the scheduler-estimator pods |`{}`|
|`schedulerEstimator.podAnnotations`| Annotaions of the scheduler-estimator pods |`{}`|
|`schedulerEstimator.imagePullSecrets`| Image pull secret of the scheduler-estimator |`[]`|
|`schedulerEstimator.image.repository`| Image of the scheduler-estimator |`"swr.ap-southeast-1.myhuaweicloud.com/karmada/karmada-scheduler-estimator"`|
|`schedulerEstimator.image.tag`| Image tag of the scheduler-estimator |`"latest"`|
|`schedulerEstimator.image.pullPolicy`| Image pull policy of the scheduler-estimator |`"IfNotPresent"`|
|`schedulerEstimator.resources`| Resource quota of the scheduler-estimator |`{}`|
|`schedulerEstimator.nodeSelector`| Node selector of the scheduler-estimator |`{}`|
|`schedulerEstimator.affinity`| Affinity of the scheduler-estimator |`{}`|
|`schedulerEstimator.tolerations`| Tolerations of the scheduler-estimator |`[]`|