//go:build !ignore_autogenerated

/*
Copyright 2024.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CategrafMonitor) DeepCopyInto(out *CategrafMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CategrafMonitor.
func (in *CategrafMonitor) DeepCopy() *CategrafMonitor {
	if in == nil {
		return nil
	}
	out := new(CategrafMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CategrafMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CategrafMonitorList) DeepCopyInto(out *CategrafMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CategrafMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CategrafMonitorList.
func (in *CategrafMonitorList) DeepCopy() *CategrafMonitorList {
	if in == nil {
		return nil
	}
	out := new(CategrafMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CategrafMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CategrafMonitorSpec) DeepCopyInto(out *CategrafMonitorSpec) {
	*out = *in
	if in.OtherFile != nil {
		in, out := &in.OtherFile, &out.OtherFile
		*out = make([]OtherFile, len(*in))
		copy(*out, *in)
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = make([]Certificate, len(*in))
		copy(*out, *in)
	}
	if in.MonitorLite != nil {
		in, out := &in.MonitorLite, &out.MonitorLite
		*out = make([]MonitorLite, len(*in))
		copy(*out, *in)
	}
	if in.MonitorSuper != nil {
		in, out := &in.MonitorSuper, &out.MonitorSuper
		*out = make([]MonitorSuper, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CategrafMonitorSpec.
func (in *CategrafMonitorSpec) DeepCopy() *CategrafMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(CategrafMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CategrafMonitorStatus) DeepCopyInto(out *CategrafMonitorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CategrafMonitorStatus.
func (in *CategrafMonitorStatus) DeepCopy() *CategrafMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(CategrafMonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorLite) DeepCopyInto(out *MonitorLite) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorLite.
func (in *MonitorLite) DeepCopy() *MonitorLite {
	if in == nil {
		return nil
	}
	out := new(MonitorLite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorSuper) DeepCopyInto(out *MonitorSuper) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]Instance, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorSuper.
func (in *MonitorSuper) DeepCopy() *MonitorSuper {
	if in == nil {
		return nil
	}
	out := new(MonitorSuper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OtherFile) DeepCopyInto(out *OtherFile) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OtherFile.
func (in *OtherFile) DeepCopy() *OtherFile {
	if in == nil {
		return nil
	}
	out := new(OtherFile)
	in.DeepCopyInto(out)
	return out
}
