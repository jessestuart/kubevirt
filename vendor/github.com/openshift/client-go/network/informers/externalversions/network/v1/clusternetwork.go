// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	network_v1 "github.com/openshift/api/network/v1"
	versioned "github.com/openshift/client-go/network/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/network/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/network/listers/network/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterNetworkInformer provides access to a shared informer and lister for
// ClusterNetworks.
type ClusterNetworkInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterNetworkLister
}

type clusterNetworkInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterNetworkInformer constructs a new informer for ClusterNetwork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterNetworkInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterNetworkInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterNetworkInformer constructs a new informer for ClusterNetwork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterNetworkInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1().ClusterNetworks().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1().ClusterNetworks().Watch(options)
			},
		},
		&network_v1.ClusterNetwork{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterNetworkInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterNetworkInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterNetworkInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&network_v1.ClusterNetwork{}, f.defaultInformer)
}

func (f *clusterNetworkInformer) Lister() v1.ClusterNetworkLister {
	return v1.NewClusterNetworkLister(f.Informer().GetIndexer())
}
