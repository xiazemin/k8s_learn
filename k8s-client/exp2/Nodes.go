package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// kubernetesの設定ファイルのパスを組み立てる
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	// BuildConfigFromFlags is a helper function that builds configs from a master url or
	// a kubeconfig filepath.
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// NewForConfig creates a new Clientset for the given config.
	// https://godoc.org/k8s.io/client-go/kubernetes#NewForConfig
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.TODO()
	// https://godoc.org/k8s.io/client-go/kubernetes/typed/core/v1
	nodes, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get nodes:", err)
	}

	for i, node := range nodes.Items {
		fmt.Printf("[%d] %s\n", i, node.GetName())
	}
}
