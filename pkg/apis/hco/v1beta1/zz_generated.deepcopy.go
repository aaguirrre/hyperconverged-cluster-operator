// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "github.com/openshift/custom-resource-status/conditions/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertRotateConfigCA) DeepCopyInto(out *CertRotateConfigCA) {
	*out = *in
	out.Duration = in.Duration
	out.RenewBefore = in.RenewBefore
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertRotateConfigCA.
func (in *CertRotateConfigCA) DeepCopy() *CertRotateConfigCA {
	if in == nil {
		return nil
	}
	out := new(CertRotateConfigCA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertRotateConfigServer) DeepCopyInto(out *CertRotateConfigServer) {
	*out = *in
	out.Duration = in.Duration
	out.RenewBefore = in.RenewBefore
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertRotateConfigServer.
func (in *CertRotateConfigServer) DeepCopy() *CertRotateConfigServer {
	if in == nil {
		return nil
	}
	out := new(CertRotateConfigServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConverged) DeepCopyInto(out *HyperConverged) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConverged.
func (in *HyperConverged) DeepCopy() *HyperConverged {
	if in == nil {
		return nil
	}
	out := new(HyperConverged)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HyperConverged) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedCertConfig) DeepCopyInto(out *HyperConvergedCertConfig) {
	*out = *in
	out.CA = in.CA
	out.Server = in.Server
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedCertConfig.
func (in *HyperConvergedCertConfig) DeepCopy() *HyperConvergedCertConfig {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedCertConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedConfig) DeepCopyInto(out *HyperConvergedConfig) {
	*out = *in
	if in.NodePlacement != nil {
		in, out := &in.NodePlacement, &out.NodePlacement
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedConfig.
func (in *HyperConvergedConfig) DeepCopy() *HyperConvergedConfig {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedFeatureGates) DeepCopyInto(out *HyperConvergedFeatureGates) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedFeatureGates.
func (in *HyperConvergedFeatureGates) DeepCopy() *HyperConvergedFeatureGates {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedFeatureGates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedList) DeepCopyInto(out *HyperConvergedList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HyperConverged, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedList.
func (in *HyperConvergedList) DeepCopy() *HyperConvergedList {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HyperConvergedList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedObsoleteCPUs) DeepCopyInto(out *HyperConvergedObsoleteCPUs) {
	*out = *in
	if in.CPUModels != nil {
		in, out := &in.CPUModels, &out.CPUModels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedObsoleteCPUs.
func (in *HyperConvergedObsoleteCPUs) DeepCopy() *HyperConvergedObsoleteCPUs {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedObsoleteCPUs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedSpec) DeepCopyInto(out *HyperConvergedSpec) {
	*out = *in
	in.Infra.DeepCopyInto(&out.Infra)
	in.Workloads.DeepCopyInto(&out.Workloads)
	out.FeatureGates = in.FeatureGates
	in.LiveMigrationConfig.DeepCopyInto(&out.LiveMigrationConfig)
	if in.PermittedHostDevices != nil {
		in, out := &in.PermittedHostDevices, &out.PermittedHostDevices
		*out = new(PermittedHostDevices)
		(*in).DeepCopyInto(*out)
	}
	out.CertConfig = in.CertConfig
	if in.ResourceRequirements != nil {
		in, out := &in.ResourceRequirements, &out.ResourceRequirements
		*out = new(OperandResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.ScratchSpaceStorageClass != nil {
		in, out := &in.ScratchSpaceStorageClass, &out.ScratchSpaceStorageClass
		*out = new(string)
		**out = **in
	}
	if in.VddkInitImage != nil {
		in, out := &in.VddkInitImage, &out.VddkInitImage
		*out = new(string)
		**out = **in
	}
	if in.ObsoleteCPUs != nil {
		in, out := &in.ObsoleteCPUs, &out.ObsoleteCPUs
		*out = new(HyperConvergedObsoleteCPUs)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageImport != nil {
		in, out := &in.StorageImport, &out.StorageImport
		*out = new(StorageImportConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedSpec.
func (in *HyperConvergedSpec) DeepCopy() *HyperConvergedSpec {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedStatus) DeepCopyInto(out *HyperConvergedStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make(Versions, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedStatus.
func (in *HyperConvergedStatus) DeepCopy() *HyperConvergedStatus {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LiveMigrationConfigurations) DeepCopyInto(out *LiveMigrationConfigurations) {
	*out = *in
	if in.ParallelMigrationsPerCluster != nil {
		in, out := &in.ParallelMigrationsPerCluster, &out.ParallelMigrationsPerCluster
		*out = new(uint32)
		**out = **in
	}
	if in.ParallelOutboundMigrationsPerNode != nil {
		in, out := &in.ParallelOutboundMigrationsPerNode, &out.ParallelOutboundMigrationsPerNode
		*out = new(uint32)
		**out = **in
	}
	if in.BandwidthPerMigration != nil {
		in, out := &in.BandwidthPerMigration, &out.BandwidthPerMigration
		*out = new(string)
		**out = **in
	}
	if in.CompletionTimeoutPerGiB != nil {
		in, out := &in.CompletionTimeoutPerGiB, &out.CompletionTimeoutPerGiB
		*out = new(int64)
		**out = **in
	}
	if in.ProgressTimeout != nil {
		in, out := &in.ProgressTimeout, &out.ProgressTimeout
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LiveMigrationConfigurations.
func (in *LiveMigrationConfigurations) DeepCopy() *LiveMigrationConfigurations {
	if in == nil {
		return nil
	}
	out := new(LiveMigrationConfigurations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MediatedHostDevice) DeepCopyInto(out *MediatedHostDevice) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MediatedHostDevice.
func (in *MediatedHostDevice) DeepCopy() *MediatedHostDevice {
	if in == nil {
		return nil
	}
	out := new(MediatedHostDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperandResourceRequirements) DeepCopyInto(out *OperandResourceRequirements) {
	*out = *in
	if in.StorageWorkloads != nil {
		in, out := &in.StorageWorkloads, &out.StorageWorkloads
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperandResourceRequirements.
func (in *OperandResourceRequirements) DeepCopy() *OperandResourceRequirements {
	if in == nil {
		return nil
	}
	out := new(OperandResourceRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PciHostDevice) DeepCopyInto(out *PciHostDevice) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PciHostDevice.
func (in *PciHostDevice) DeepCopy() *PciHostDevice {
	if in == nil {
		return nil
	}
	out := new(PciHostDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermittedHostDevices) DeepCopyInto(out *PermittedHostDevices) {
	*out = *in
	if in.PciHostDevices != nil {
		in, out := &in.PciHostDevices, &out.PciHostDevices
		*out = make([]PciHostDevice, len(*in))
		copy(*out, *in)
	}
	if in.MediatedDevices != nil {
		in, out := &in.MediatedDevices, &out.MediatedDevices
		*out = make([]MediatedHostDevice, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermittedHostDevices.
func (in *PermittedHostDevices) DeepCopy() *PermittedHostDevices {
	if in == nil {
		return nil
	}
	out := new(PermittedHostDevices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageImportConfig) DeepCopyInto(out *StorageImportConfig) {
	*out = *in
	if in.InsecureRegistries != nil {
		in, out := &in.InsecureRegistries, &out.InsecureRegistries
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageImportConfig.
func (in *StorageImportConfig) DeepCopy() *StorageImportConfig {
	if in == nil {
		return nil
	}
	out := new(StorageImportConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Version) DeepCopyInto(out *Version) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Version.
func (in *Version) DeepCopy() *Version {
	if in == nil {
		return nil
	}
	out := new(Version)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Versions) DeepCopyInto(out *Versions) {
	{
		in := &in
		*out = make(Versions, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Versions.
func (in Versions) DeepCopy() Versions {
	if in == nil {
		return nil
	}
	out := new(Versions)
	in.DeepCopyInto(out)
	return *out
}
