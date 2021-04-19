package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	appsv1beta1 "k8s.io/api/apps/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//kubelet.kubeconfig  是文件对应地址
	kubeconfig := flag.String("kubeconfig", "kubelet.kubeconfig", "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	// 解析到config
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// 创建连接
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	deploymentsClient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)

	//创建deployment
	go createDeployment(deploymentsClient)

	//监听deployment
	startWatchDeployment(deploymentsClient)
}

//监听Deployment变化
func startWatchDeployment(deploymentsClient v1beta1.DeploymentInterface) {
	w, _ := deploymentsClient.Watch(context.TODO(), metav1.ListOptions{})
	for {
		select {
		case e, _ := <-w.ResultChan():
			fmt.Println(e.Type, e.Object)
		}
	}
}

//创建deployemnt，需要谨慎按照部署的k8s版本来使用api接口
func createDeployment(deploymentsClient v1beta1.DeploymentInterface) {
	var r apiv1.ResourceRequirements
	//资源分配会遇到无法设置值的问题，故采用json反解析
	j := `{"limits": {"cpu":"2000m", "memory": "1Gi"}, "requests": {"cpu":"2000m", "memory": "1Gi"}}`
	json.Unmarshal([]byte(j), &r)
	deployment := &appsv1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "engine",
			Labels: map[string]string{
				"app": "engine",
			},
		},
		Spec: appsv1beta1.DeploymentSpec{
			Replicas: int32Ptr2(1),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "engine",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{Name: "engine",
							Image:     "my.com/engine:v2",
							Resources: r,
						},
					},
				},
			},
		},
	}

	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, v1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

func int32Ptr2(i int32) *int32 { return &i }
