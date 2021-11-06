/*
/*
@Time : 2019/12/23 3:08 下午
@Author : tianpeng.du
@File : types
@Software: GoLand
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	ipv1 "ip/pkg/apis/ip/v1"
	versioned "ip/pkg/client/clientset/versioned"
	internalinterfaces "ip/pkg/client/informers/externalversions/internalinterfaces"
	v1 "ip/pkg/client/listers/ip/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// IpInformer provides access to a shared informer and lister for
// Ips.
type IpInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.IpLister
}

type ipInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewIpInformer constructs a new informer for Ip type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIpInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIpInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredIpInformer constructs a new informer for Ip type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIpInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RocduV1().Ips(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RocduV1().Ips(namespace).Watch(context.TODO(), options)
			},
		},
		&ipv1.Ip{},
		resyncPeriod,
		indexers,
	)
}

func (f *ipInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIpInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ipInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ipv1.Ip{}, f.defaultInformer)
}

func (f *ipInformer) Lister() v1.IpLister {
	return v1.NewIpLister(f.Informer().GetIndexer())
}