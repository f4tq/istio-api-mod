// Copyright 2019 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha3

import (
	time "time"

	networking_v1alpha3 "github.com/f4tq/istio-api-mod/pkg/kube/apis/networking/v1alpha3"
	versioned "github.com/f4tq/istio-api-mod/pkg/kube/clientset/versioned"
	internalinterfaces "github.com/f4tq/istio-api-mod/pkg/kube/informers/externalversions/internalinterfaces"
	v1alpha3 "github.com/f4tq/istio-api-mod/pkg/kube/listers/networking/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SidecarInformer provides access to a shared informer and lister for
// Sidecars.
type SidecarInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha3.SidecarLister
}

type sidecarInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSidecarInformer constructs a new informer for Sidecar type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSidecarInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSidecarInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSidecarInformer constructs a new informer for Sidecar type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSidecarInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha3().Sidecars(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha3().Sidecars(namespace).Watch(options)
			},
		},
		&networking_v1alpha3.Sidecar{},
		resyncPeriod,
		indexers,
	)
}

func (f *sidecarInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSidecarInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sidecarInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networking_v1alpha3.Sidecar{}, f.defaultInformer)
}

func (f *sidecarInformer) Lister() v1alpha3.SidecarLister {
	return v1alpha3.NewSidecarLister(f.Informer().GetIndexer())
}
