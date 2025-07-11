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

package v1

import (
	context "context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	provisioningrequestautoscalingxk8siov1 "k8s.io/autoscaler/cluster-autoscaler/apis/provisioningrequest/autoscaling.x-k8s.io/v1"
	versioned "k8s.io/autoscaler/cluster-autoscaler/apis/provisioningrequest/client/clientset/versioned"
	internalinterfaces "k8s.io/autoscaler/cluster-autoscaler/apis/provisioningrequest/client/informers/externalversions/internalinterfaces"
	autoscalingxk8siov1 "k8s.io/autoscaler/cluster-autoscaler/apis/provisioningrequest/client/listers/autoscaling.x-k8s.io/v1"
	cache "k8s.io/client-go/tools/cache"
)

// ProvisioningRequestInformer provides access to a shared informer and lister for
// ProvisioningRequests.
type ProvisioningRequestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() autoscalingxk8siov1.ProvisioningRequestLister
}

type provisioningRequestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewProvisioningRequestInformer constructs a new informer for ProvisioningRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewProvisioningRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredProvisioningRequestInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredProvisioningRequestInformer constructs a new informer for ProvisioningRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredProvisioningRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1().ProvisioningRequests(namespace).List(context.Background(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1().ProvisioningRequests(namespace).Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1().ProvisioningRequests(namespace).List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1().ProvisioningRequests(namespace).Watch(ctx, options)
			},
		},
		&provisioningrequestautoscalingxk8siov1.ProvisioningRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *provisioningRequestInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredProvisioningRequestInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *provisioningRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&provisioningrequestautoscalingxk8siov1.ProvisioningRequest{}, f.defaultInformer)
}

func (f *provisioningRequestInformer) Lister() autoscalingxk8siov1.ProvisioningRequestLister {
	return autoscalingxk8siov1.NewProvisioningRequestLister(f.Informer().GetIndexer())
}
