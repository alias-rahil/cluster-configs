/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	clusterconfigsiov1alpha1 "github.com/alias-rahil/cluster-configs/pkg/apis/clusterconfigs.io/v1alpha1"
	versioned "github.com/alias-rahil/cluster-configs/pkg/client/clientset/versioned"
	internalinterfaces "github.com/alias-rahil/cluster-configs/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/alias-rahil/cluster-configs/pkg/client/listers/clusterconfigs.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterSecretInformer provides access to a shared informer and lister for
// ClusterSecrets.
type ClusterSecretInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterSecretLister
}

type clusterSecretInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterSecretInformer constructs a new informer for ClusterSecret type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterSecretInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterSecretInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterSecretInformer constructs a new informer for ClusterSecret type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterSecretInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterconfigsV1alpha1().ClusterSecrets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterconfigsV1alpha1().ClusterSecrets(namespace).Watch(context.TODO(), options)
			},
		},
		&clusterconfigsiov1alpha1.ClusterSecret{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterSecretInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterSecretInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterSecretInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&clusterconfigsiov1alpha1.ClusterSecret{}, f.defaultInformer)
}

func (f *clusterSecretInformer) Lister() v1alpha1.ClusterSecretLister {
	return v1alpha1.NewClusterSecretLister(f.Informer().GetIndexer())
}
