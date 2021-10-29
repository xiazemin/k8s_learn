package main

import (
	"flag"
	"time"

	clientset "ip/pkg/client/clientset/versioned"
	"ip/pkg/client/informers/externalversions"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"k8s.io/sample-controller/pkg/signals"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	// kubeClient, err := kubernetes.NewForConfig(cfg)
	// if err != nil {
	// 	klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	// }

	exampleClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %s", err.Error())
	}

	informer := externalversions.NewSharedInformerFactoryWithOptions(exampleClient, 10*time.Second, externalversions.WithNamespace("default"))
	go informer.Start(stopCh)

	IpCrdInformer := informer.Rocdu().V1().Ips()
	cache.WaitForCacheSync(nil, IpCrdInformer.Informer().HasSynced)

	// controller := NewController(kubeClient, exampleClient,
	// 	kubeInformerFactory.Apps().V1().Deployments(),
	// 	exampleInformerFactory.Samplecontroller().V1alpha1().Foos())

	// // notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(stopCh)
	// // Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	// kubeInformerFactory.Start(stopCh)
	// exampleInformerFactory.Start(stopCh)

	// if err = controller.Run(2, stopCh); err != nil {
	// 	klog.Fatalf("Error running controller: %s", err.Error())
	// }
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
