/*
/*
@Time : 2019/12/23 3:08 下午
@Author : tianpeng.du
@File : types
@Software: GoLand
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "ip/pkg/apis/ip/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IpLister helps list Ips.
// All objects returned here must be treated as read-only.
type IpLister interface {
	// List lists all Ips in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Ip, err error)
	// Ips returns an object that can list and get Ips.
	Ips(namespace string) IpNamespaceLister
	IpListerExpansion
}

// ipLister implements the IpLister interface.
type ipLister struct {
	indexer cache.Indexer
}

// NewIpLister returns a new IpLister.
func NewIpLister(indexer cache.Indexer) IpLister {
	return &ipLister{indexer: indexer}
}

// List lists all Ips in the indexer.
func (s *ipLister) List(selector labels.Selector) (ret []*v1.Ip, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Ip))
	})
	return ret, err
}

// Ips returns an object that can list and get Ips.
func (s *ipLister) Ips(namespace string) IpNamespaceLister {
	return ipNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IpNamespaceLister helps list and get Ips.
// All objects returned here must be treated as read-only.
type IpNamespaceLister interface {
	// List lists all Ips in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Ip, err error)
	// Get retrieves the Ip from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Ip, error)
	IpNamespaceListerExpansion
}

// ipNamespaceLister implements the IpNamespaceLister
// interface.
type ipNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Ips in the indexer for a given namespace.
func (s ipNamespaceLister) List(selector labels.Selector) (ret []*v1.Ip, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Ip))
	})
	return ret, err
}

// Get retrieves the Ip from the indexer for a given namespace and name.
func (s ipNamespaceLister) Get(name string) (*v1.Ip, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("ip"), name)
	}
	return obj.(*v1.Ip), nil
}