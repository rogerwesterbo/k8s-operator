package main

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	println("Hello, kubernetes operator!")

	println("Getting kubernetes config")
	config := config.GetConfigOrDie()
	k8sclient := kubernetes.NewForConfigOrDie(config)

	namespaces, err := k8sclient.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}

	println("namespaces")
	for _, ns := range namespaces.Items {
		println(ns.Name)
	}

	println("Stopping operator")
}
