// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/karmada-io/karmada/pkg/apis/query/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterCacheLister helps list ClusterCaches.
// All objects returned here must be treated as read-only.
type ClusterCacheLister interface {
	// List lists all ClusterCaches in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterCache, err error)
	// ClusterCaches returns an object that can list and get ClusterCaches.
	ClusterCaches(namespace string) ClusterCacheNamespaceLister
	ClusterCacheListerExpansion
}

// clusterCacheLister implements the ClusterCacheLister interface.
type clusterCacheLister struct {
	indexer cache.Indexer
}

// NewClusterCacheLister returns a new ClusterCacheLister.
func NewClusterCacheLister(indexer cache.Indexer) ClusterCacheLister {
	return &clusterCacheLister{indexer: indexer}
}

// List lists all ClusterCaches in the indexer.
func (s *clusterCacheLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterCache, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterCache))
	})
	return ret, err
}

// ClusterCaches returns an object that can list and get ClusterCaches.
func (s *clusterCacheLister) ClusterCaches(namespace string) ClusterCacheNamespaceLister {
	return clusterCacheNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterCacheNamespaceLister helps list and get ClusterCaches.
// All objects returned here must be treated as read-only.
type ClusterCacheNamespaceLister interface {
	// List lists all ClusterCaches in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterCache, err error)
	// Get retrieves the ClusterCache from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterCache, error)
	ClusterCacheNamespaceListerExpansion
}

// clusterCacheNamespaceLister implements the ClusterCacheNamespaceLister
// interface.
type clusterCacheNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterCaches in the indexer for a given namespace.
func (s clusterCacheNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterCache, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterCache))
	})
	return ret, err
}

// Get retrieves the ClusterCache from the indexer for a given namespace and name.
func (s clusterCacheNamespaceLister) Get(name string) (*v1alpha1.ClusterCache, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clustercache"), name)
	}
	return obj.(*v1alpha1.ClusterCache), nil
}
