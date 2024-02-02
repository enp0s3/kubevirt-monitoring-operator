/*
Copyright The KubeVirt Authors.

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
	"context"
	time "time"

	versioned "github.com/enp0s3/kubevirt-monitoring-operator/pkg/client/versioned"
	internalinterfaces "github.com/enp0s3/kubevirt-monitoring-operator/pkg/client/versioned/informers/externalversions/internalinterfaces"
	v1 "github.com/enp0s3/kubevirt-monitoring-operator/pkg/client/versioned/listers/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	corev1 "kubevirt.io/api/core/v1"
)

// VirtualMachineInformer provides access to a shared informer and lister for
// VirtualMachines.
type VirtualMachineInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.VirtualMachineLister
}

type virtualMachineInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVirtualMachineInformer constructs a new informer for VirtualMachine type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVirtualMachineInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVirtualMachineInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVirtualMachineInformer constructs a new informer for VirtualMachine type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVirtualMachineInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubevirtV1().VirtualMachines(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubevirtV1().VirtualMachines(namespace).Watch(context.TODO(), options)
			},
		},
		&corev1.VirtualMachine{},
		resyncPeriod,
		indexers,
	)
}

func (f *virtualMachineInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVirtualMachineInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *virtualMachineInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.VirtualMachine{}, f.defaultInformer)
}

func (f *virtualMachineInformer) Lister() v1.VirtualMachineLister {
	return v1.NewVirtualMachineLister(f.Informer().GetIndexer())
}
