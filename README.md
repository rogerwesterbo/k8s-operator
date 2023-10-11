# k8s-operator

Showing how to create kubernetes operators in Go

# Prerequisites:

- A kubernetes cluster (locally or from a cloud service)
  - Kind (https://kind.sigs.k8s.io/)
  - K3d ()https://k3d.io/
  - Microk8s (https://microk8s.io/)
  - Minikube (https://minikube.sigs.k8s.io/)
- Golang (https://go.dev)
  - Other sdks: https://kubernetes.io/docs/reference/using-api/client-libraries/

## Create a local cluster with kind:

`kind create cluster --name testingk8soperator --config ./kind/kind-config.yaml`

### Testing getting nodes:

`kubectl get nodes`

Should show somethinkg like this:

```
NAME                           STATUS   ROLES           AGE   VERSION
k8soperatork8s-control-plane   Ready    control-plane   43s   v1.27.3
k8soperatork8s-worker          Ready    <none>          22s   v1.27.3
k8soperatork8s-worker2         Ready    <none>          22s   v1.27.3
```

## Creating a golang project

`go mod init a-cool-domain.io/k8s`

create a main.go file

```
package main

func main() {
	println("Hello, World!")
}
```
