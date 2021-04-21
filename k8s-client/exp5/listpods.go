package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func getNamespace() string {
	if data, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace"); err == nil {
		if ns := strings.TrimSpace(string(data)); len(ns) > 0 {
			return ns
		}
	}
	//return "default"
	/*
			% kubectl get pods -n kube-system
		NAME                                     READY   STATUS    RESTARTS   AGE
		coredns-f9fd979d6-dnzc6                  1/1     Running   13         25d
		coredns-f9fd979d6-lx646                  1/1     Running   13         25d
		etcd-docker-desktop                      1/1     Running   23         24d
		kube-apiserver-docker-desktop            1/1     Running   23         24d
		kube-controller-manager-docker-desktop   1/1     Running   23         24d
		kube-proxy-d8qtd                         1/1     Running   13         24d
		kube-scheduler-docker-desktop            1/1     Running   58         24d
		storage-provisioner                      1/1     Running   1470       24d
		vpnkit-controller                        1/1     Running   13         24d
	*/
	return "kube-system"
}

type Pod struct {
	podInterface corev1.PodInterface
	app          string
}

func (p Pod) ListPods() ([]string, error) {
	ctx := context.TODO()
	pods, err := p.podInterface.List(ctx, metav1.ListOptions{
		LabelSelector: "app=" + p.app,
	})
	if err != nil {
		return nil, err
	}
	var ips []string
	for _, pod := range pods.Items {
		ips = append(ips, pod.Status.PodIP)
	}
	return ips, nil
}

func NewPod() (*Pod, error) {
	namespace := getNamespace()
	clientset := getClient()
	ctx := context.TODO()
	nodes, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	for _, n := range nodes.Items {
		fmt.Println("nodes:", n.Name)
	}
	nss, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	for _, n := range nss.Items {
		fmt.Println("ns:", n.Name)
	}
	svcs, err := clientset.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
	for _, s := range svcs.Items {
		fmt.Println("svc:", s.Name)
	}
	pi := clientset.CoreV1().Pods(namespace)
	fmt.Println(namespace)
	/*
		hostname, err := os.Hostname()
		if err != nil {
			return nil, err
		}
		pod, err := pi.Get(ctx, hostname, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		app := pod.Labels["app"]

		pods, err := pi.List(ctx, metav1.ListOptions{
			LabelSelector: "app=" + app,
		})
		if err != nil {
			return nil, err
		}
	*/
	pods, err := pi.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	//fmt.Println(pods)
	var ips []string
	for _, pod := range pods.Items {
		ips = append(ips, pod.Status.PodIP)
	}

	fmt.Println(ips)
	return &Pod{}, nil
}

func getClient() *kubernetes.Clientset {
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
	return clientset
}

func getInpodClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func main() {
	NewPod()
}
