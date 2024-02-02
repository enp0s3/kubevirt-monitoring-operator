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
	internalinterfaces "github.com/enp0s3/kubevirt-monitoring-operator/pkg/client/versioned/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// KubeVirts returns a KubeVirtInformer.
	KubeVirts() KubeVirtInformer
	// VirtualMachines returns a VirtualMachineInformer.
	VirtualMachines() VirtualMachineInformer
	// VirtualMachineInstances returns a VirtualMachineInstanceInformer.
	VirtualMachineInstances() VirtualMachineInstanceInformer
	// VirtualMachineInstanceMigrations returns a VirtualMachineInstanceMigrationInformer.
	VirtualMachineInstanceMigrations() VirtualMachineInstanceMigrationInformer
	// VirtualMachineInstancePresets returns a VirtualMachineInstancePresetInformer.
	VirtualMachineInstancePresets() VirtualMachineInstancePresetInformer
	// VirtualMachineInstanceReplicaSets returns a VirtualMachineInstanceReplicaSetInformer.
	VirtualMachineInstanceReplicaSets() VirtualMachineInstanceReplicaSetInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// KubeVirts returns a KubeVirtInformer.
func (v *version) KubeVirts() KubeVirtInformer {
	return &kubeVirtInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualMachines returns a VirtualMachineInformer.
func (v *version) VirtualMachines() VirtualMachineInformer {
	return &virtualMachineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualMachineInstances returns a VirtualMachineInstanceInformer.
func (v *version) VirtualMachineInstances() VirtualMachineInstanceInformer {
	return &virtualMachineInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualMachineInstanceMigrations returns a VirtualMachineInstanceMigrationInformer.
func (v *version) VirtualMachineInstanceMigrations() VirtualMachineInstanceMigrationInformer {
	return &virtualMachineInstanceMigrationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualMachineInstancePresets returns a VirtualMachineInstancePresetInformer.
func (v *version) VirtualMachineInstancePresets() VirtualMachineInstancePresetInformer {
	return &virtualMachineInstancePresetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualMachineInstanceReplicaSets returns a VirtualMachineInstanceReplicaSetInformer.
func (v *version) VirtualMachineInstanceReplicaSets() VirtualMachineInstanceReplicaSetInformer {
	return &virtualMachineInstanceReplicaSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
