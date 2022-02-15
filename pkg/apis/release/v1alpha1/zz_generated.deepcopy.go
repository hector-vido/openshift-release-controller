//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobRunCoordinates) DeepCopyInto(out *JobRunCoordinates) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobRunCoordinates.
func (in *JobRunCoordinates) DeepCopy() *JobRunCoordinates {
	if in == nil {
		return nil
	}
	out := new(JobRunCoordinates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobRunResult) DeepCopyInto(out *JobRunResult) {
	*out = *in
	out.Coordinates = in.Coordinates
	in.StartTime.DeepCopyInto(&out.StartTime)
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobRunResult.
func (in *JobRunResult) DeepCopy() *JobRunResult {
	if in == nil {
		return nil
	}
	out := new(JobRunResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobStatus) DeepCopyInto(out *JobStatus) {
	*out = *in
	if in.JobRunResults != nil {
		in, out := &in.JobRunResults, &out.JobRunResults
		*out = make([]JobRunResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobStatus.
func (in *JobStatus) DeepCopy() *JobStatus {
	if in == nil {
		return nil
	}
	out := new(JobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PayloadCoordinates) DeepCopyInto(out *PayloadCoordinates) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PayloadCoordinates.
func (in *PayloadCoordinates) DeepCopy() *PayloadCoordinates {
	if in == nil {
		return nil
	}
	out := new(PayloadCoordinates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePayload) DeepCopyInto(out *ReleasePayload) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePayload.
func (in *ReleasePayload) DeepCopy() *ReleasePayload {
	if in == nil {
		return nil
	}
	out := new(ReleasePayload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleasePayload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePayloadCondition) DeepCopyInto(out *ReleasePayloadCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePayloadCondition.
func (in *ReleasePayloadCondition) DeepCopy() *ReleasePayloadCondition {
	if in == nil {
		return nil
	}
	out := new(ReleasePayloadCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePayloadList) DeepCopyInto(out *ReleasePayloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReleasePayload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePayloadList.
func (in *ReleasePayloadList) DeepCopy() *ReleasePayloadList {
	if in == nil {
		return nil
	}
	out := new(ReleasePayloadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleasePayloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePayloadOverride) DeepCopyInto(out *ReleasePayloadOverride) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePayloadOverride.
func (in *ReleasePayloadOverride) DeepCopy() *ReleasePayloadOverride {
	if in == nil {
		return nil
	}
	out := new(ReleasePayloadOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePayloadSpec) DeepCopyInto(out *ReleasePayloadSpec) {
	*out = *in
	out.PayloadCoordinates = in.PayloadCoordinates
	out.PayloadOverride = in.PayloadOverride
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePayloadSpec.
func (in *ReleasePayloadSpec) DeepCopy() *ReleasePayloadSpec {
	if in == nil {
		return nil
	}
	out := new(ReleasePayloadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePayloadStatus) DeepCopyInto(out *ReleasePayloadStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ReleasePayloadCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BlockingJobResults != nil {
		in, out := &in.BlockingJobResults, &out.BlockingJobResults
		*out = make([]JobStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InformingJobResults != nil {
		in, out := &in.InformingJobResults, &out.InformingJobResults
		*out = make([]JobStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AnalysisJobResults != nil {
		in, out := &in.AnalysisJobResults, &out.AnalysisJobResults
		*out = make([]JobStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePayloadStatus.
func (in *ReleasePayloadStatus) DeepCopy() *ReleasePayloadStatus {
	if in == nil {
		return nil
	}
	out := new(ReleasePayloadStatus)
	in.DeepCopyInto(out)
	return out
}
