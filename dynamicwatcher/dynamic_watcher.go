package dynamicwatcher

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/tools/cache"
)

func Listen(dynamicClient dynamic.Interface, resources []schema.GroupVersionResource, stop <-chan struct{}) {
	dynInformer := dynamicinformer.NewDynamicSharedInformerFactory(dynamicClient, 0)

	for _, resource := range resources {
		informer := dynInformer.ForResource(resource).Informer()

		_, err := informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
			AddFunc:    addResource,
			UpdateFunc: updateResource,
			DeleteFunc: deleteResource,
		})
		if err != nil {
			panic(err)
		}
	}
	dynInformer.Start(stop)
}

func addResource(obj interface{}) {
	rawData := obj.(*unstructured.Unstructured)
	_, _ = fmt.Printf("add resource\n Name: %s\n Namespace: %s\n Kind: %s\n ApiVersion: %s\n\n", rawData.GetName(), rawData.GetNamespace(), rawData.GetKind(), rawData.GetAPIVersion())
}

func updateResource(oldObj, newObj interface{}) {
	rawData := newObj.(*unstructured.Unstructured)

	_, _ = fmt.Printf("update resource\n Name: %s\n Namespace: %s\n Kind: %s\n ApiVersion: %s\n\n", rawData.GetName(), rawData.GetNamespace(), rawData.GetKind(), rawData.GetAPIVersion())
}

func deleteResource(obj interface{}) {
	rawData := obj.(*unstructured.Unstructured)

	_, _ = fmt.Printf("delete resource\n Name: %s\n Namespace: %s\n Kind: %s\n ApiVersion: %s\n\n", rawData.GetName(), rawData.GetNamespace(), rawData.GetKind(), rawData.GetAPIVersion())
}
