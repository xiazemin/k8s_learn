package main

import (
	"context"
	"flag"
	"fmt"

	appv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	var (
		k8sconfig = flag.String("k8sconfig", "./admin.conf", "kubernetes auth config") //使用kubeconfig配置文件进行集群权限认证
		config    *rest.Config
		err       error
	)
	flag.Parse()

	config, err = clientcmd.BuildConfigFromFlags("", *k8sconfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 从指定的config创建一个新的clientset
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connect kubernetes cluster success.")
	}

	ctx := context.TODO()
	// 获取pod列表 pod为名称空间级别资源需指定名称空间
	pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})

	if err != nil {
		panic(err)
	}
	// 循环打印pod的信息
	for _, pod := range pods.Items {
		fmt.Println(pod.ObjectMeta.Name, pod.Status.Phase)
	}
	//创建namespace
	nsClient := clientset.CoreV1().Namespaces()

	ns := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "testzhangsan",
		},
		Status: apiv1.NamespaceStatus{
			Phase: apiv1.NamespaceActive,
		},
	}

	opts := metav1.CreateOptions{}
	ns, err = nsClient.Create(ctx, ns, opts)

	if err != nil {
		panic(err)
	}

	fmt.Println(ns.ObjectMeta.Name, ns.Status.Phase)

	//获取指定名称空间下svc信息
	svclist, err := clientset.CoreV1().Services("kube-system").List(ctx, metav1.ListOptions{})

	for _, svc := range svclist.Items {
		fmt.Println(svc.Name, svc.Spec.ClusterIP, svc.Spec.Ports)
	}

	selecter := metav1.LabelSelector{
		MatchLabels: match,
	}
	deploy := appv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "testgolangclient",
		},

		Spec: appv1.DeploymentSpec{
			Replicas: &repl,
			Selector: &selecter,
			Template: templateSpec,
		},
	}

	podsClient, err := clientset.AppsV1().Deployments("default" /* 名称空间 */).Create(&deploy)

}
