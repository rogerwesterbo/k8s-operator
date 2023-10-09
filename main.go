package main

import (
	"context"
	"fmt"
	"os"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	println("Hello, kubernetes operator!")

	println("Getting kubernetes config, or die ...")
	config, err := config.GetConfig()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error getting kubernetes config (is container running in a cluster?!)")
		panic(err)
	}

	println("Puh, did not die. Kubernetes config apprihended.")

	k8sclient, err := kubernetes.NewForConfig(config)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error getting kubernetes client")
		panic(err)
	}

	namespaces, err := k8sclient.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error getting namespaces")
		panic(err)
	}

	println("namespaces")
	for _, ns := range namespaces.Items {
		println(ns.Name)
	}

	println("Stopping operator")
}
