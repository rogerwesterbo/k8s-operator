package main

import (
	"os"
	"os/signal"
	"syscall"

	"a-cool-domain.io/k8s/dynamicwatcher"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	println("Hello, kubernetes operator!")

	sigs := make(chan os.Signal, 1)                                    // Create channel to receive os signals
	stop := make(chan struct{})                                        // Create channel to receive stop signal
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT) // Register the sigs channel to receieve SIGTERM

	println("Getting kubernetes config")
	config := config.GetConfigOrDie()

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	resources := []schema.GroupVersionResource{
		{Group: "", Version: "v1", Resource: "namespaces"},
		{Group: "", Version: "v1", Resource: "pods"},
	}
	go func() {
		println("Starting dynamic watcher")
		dynamicwatcher.Listen(dynamicClient, resources, stop)
		sig := <-sigs
		println("Received signal: ", sig)
		stop <- struct{}{}
	}()

	<-stop
	println("Stopping operator")
}
