package main

import (
	"code-generator/generated/clientset/versioned"
	"code-generator/generated/clientset/versioned/typed/samplecontroller/v1alpha1"
	"code-generator/generated/informers/externalversions"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
	"time"
        "context"
)

func TestClient(t *testing.T) {
    config, e := clientcmd.BuildConfigFromFlags("10.30.21.238:6443", "/home/tangxu/.kube/config")
	if e != nil {
		panic(e.Error())
	}
    //注意,这里使用的是v1alpha1这个包
	client, e := v1alpha1.NewForConfig(config)
	if e != nil {
		panic(e.Error())
	}
	fooList, e := client.Foos("test").List(context.TODO(),metav1.ListOptions{})
	fmt.Println(fooList, e)

    //注意 这里的versioned包
	clientset, e := versioned.NewForConfig(config)
	factory := externalversions.NewSharedInformerFactory(clientset, 30*time.Second)
	foo, e := factory.Samplecontroller().V1alpha1().Foos().Lister().Foos("test").Get("test")
	if e != nil {
		panic(e.Error())
	}
	fmt.Println(foo, e)
}
