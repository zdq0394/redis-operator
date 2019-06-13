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

package v1

import (
	core_v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisExporter) DeepCopyInto(out *RedisExporter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisExporter.
func (in *RedisExporter) DeepCopy() *RedisExporter {
	if in == nil {
		return nil
	}
	out := new(RedisExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFailover) DeepCopyInto(out *RedisFailover) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFailover.
func (in *RedisFailover) DeepCopy() *RedisFailover {
	if in == nil {
		return nil
	}
	out := new(RedisFailover)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisFailover) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFailoverList) DeepCopyInto(out *RedisFailoverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisFailover, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFailoverList.
func (in *RedisFailoverList) DeepCopy() *RedisFailoverList {
	if in == nil {
		return nil
	}
	out := new(RedisFailoverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisFailoverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFailoverSpec) DeepCopyInto(out *RedisFailoverSpec) {
	*out = *in
	in.Redis.DeepCopyInto(&out.Redis)
	in.Sentinel.DeepCopyInto(&out.Sentinel)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFailoverSpec.
func (in *RedisFailoverSpec) DeepCopy() *RedisFailoverSpec {
	if in == nil {
		return nil
	}
	out := new(RedisFailoverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFailoverStatus) DeepCopyInto(out *RedisFailoverStatus) {
	*out = *in
	if in.RedisNodes != nil {
		in, out := &in.RedisNodes, &out.RedisNodes
		*out = make([]RedisNode, len(*in))
		copy(*out, *in)
	}
	if in.SentinelNodes != nil {
		in, out := &in.SentinelNodes, &out.SentinelNodes
		*out = make([]SentinelNode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFailoverStatus.
func (in *RedisFailoverStatus) DeepCopy() *RedisFailoverStatus {
	if in == nil {
		return nil
	}
	out := new(RedisFailoverStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisNode) DeepCopyInto(out *RedisNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisNode.
func (in *RedisNode) DeepCopy() *RedisNode {
	if in == nil {
		return nil
	}
	out := new(RedisNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisPerceptron) DeepCopyInto(out *RedisPerceptron) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisPerceptron.
func (in *RedisPerceptron) DeepCopy() *RedisPerceptron {
	if in == nil {
		return nil
	}
	out := new(RedisPerceptron)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisSettings) DeepCopyInto(out *RedisSettings) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.CustomConfig != nil {
		in, out := &in.CustomConfig, &out.CustomConfig
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Storage.DeepCopyInto(&out.Storage)
	out.Exporter = in.Exporter
	out.Perceptron = in.Perceptron
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.Affinity)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PodSecurityContext)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]core_v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisSettings.
func (in *RedisSettings) DeepCopy() *RedisSettings {
	if in == nil {
		return nil
	}
	out := new(RedisSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisStorage) DeepCopyInto(out *RedisStorage) {
	*out = *in
	if in.EmptyDir != nil {
		in, out := &in.EmptyDir, &out.EmptyDir
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.EmptyDirVolumeSource)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PersistentVolumeClaim)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisStorage.
func (in *RedisStorage) DeepCopy() *RedisStorage {
	if in == nil {
		return nil
	}
	out := new(RedisStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelNode) DeepCopyInto(out *SentinelNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelNode.
func (in *SentinelNode) DeepCopy() *SentinelNode {
	if in == nil {
		return nil
	}
	out := new(SentinelNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelSettings) DeepCopyInto(out *SentinelSettings) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.CustomConfig != nil {
		in, out := &in.CustomConfig, &out.CustomConfig
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.Affinity)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PodSecurityContext)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]core_v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelSettings.
func (in *SentinelSettings) DeepCopy() *SentinelSettings {
	if in == nil {
		return nil
	}
	out := new(SentinelSettings)
	in.DeepCopyInto(out)
	return out
}
