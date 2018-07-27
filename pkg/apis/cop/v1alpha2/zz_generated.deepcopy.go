// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageExecutionCondition) DeepCopyInto(out *ImageExecutionCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageExecutionCondition.
func (in *ImageExecutionCondition) DeepCopy() *ImageExecutionCondition {
	if in == nil {
		return nil
	}
	out := new(ImageExecutionCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanningRequest) DeepCopyInto(out *ImageScanningRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanningRequest.
func (in *ImageScanningRequest) DeepCopy() *ImageScanningRequest {
	if in == nil {
		return nil
	}
	out := new(ImageScanningRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageScanningRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanningRequestList) DeepCopyInto(out *ImageScanningRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageSigningRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanningRequestList.
func (in *ImageScanningRequestList) DeepCopy() *ImageScanningRequestList {
	if in == nil {
		return nil
	}
	out := new(ImageScanningRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageScanningRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanningRequestSpec) DeepCopyInto(out *ImageScanningRequestSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanningRequestSpec.
func (in *ImageScanningRequestSpec) DeepCopy() *ImageScanningRequestSpec {
	if in == nil {
		return nil
	}
	out := new(ImageScanningRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanningRequestStatus) DeepCopyInto(out *ImageScanningRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ImageExecutionCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ScanResult = in.ScanResult
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanningRequestStatus.
func (in *ImageScanningRequestStatus) DeepCopy() *ImageScanningRequestStatus {
	if in == nil {
		return nil
	}
	out := new(ImageScanningRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSigningRequest) DeepCopyInto(out *ImageSigningRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSigningRequest.
func (in *ImageSigningRequest) DeepCopy() *ImageSigningRequest {
	if in == nil {
		return nil
	}
	out := new(ImageSigningRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageSigningRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSigningRequestList) DeepCopyInto(out *ImageSigningRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageSigningRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSigningRequestList.
func (in *ImageSigningRequestList) DeepCopy() *ImageSigningRequestList {
	if in == nil {
		return nil
	}
	out := new(ImageSigningRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageSigningRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSigningRequestSpec) DeepCopyInto(out *ImageSigningRequestSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSigningRequestSpec.
func (in *ImageSigningRequestSpec) DeepCopy() *ImageSigningRequestSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSigningRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSigningRequestStatus) DeepCopyInto(out *ImageSigningRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ImageExecutionCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSigningRequestStatus.
func (in *ImageSigningRequestStatus) DeepCopy() *ImageSigningRequestStatus {
	if in == nil {
		return nil
	}
	out := new(ImageSigningRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageVerificationRequest) DeepCopyInto(out *ImageVerificationRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageVerificationRequest.
func (in *ImageVerificationRequest) DeepCopy() *ImageVerificationRequest {
	if in == nil {
		return nil
	}
	out := new(ImageVerificationRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageVerificationRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageVerificationRequestList) DeepCopyInto(out *ImageVerificationRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageVerificationRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageVerificationRequestList.
func (in *ImageVerificationRequestList) DeepCopy() *ImageVerificationRequestList {
	if in == nil {
		return nil
	}
	out := new(ImageVerificationRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageVerificationRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageVerificationRequestSpec) DeepCopyInto(out *ImageVerificationRequestSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageVerificationRequestSpec.
func (in *ImageVerificationRequestSpec) DeepCopy() *ImageVerificationRequestSpec {
	if in == nil {
		return nil
	}
	out := new(ImageVerificationRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageVerificationRequestStatus) DeepCopyInto(out *ImageVerificationRequestStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageVerificationRequestStatus.
func (in *ImageVerificationRequestStatus) DeepCopy() *ImageVerificationRequestStatus {
	if in == nil {
		return nil
	}
	out := new(ImageVerificationRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanResult) DeepCopyInto(out *ScanResult) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanResult.
func (in *ScanResult) DeepCopy() *ScanResult {
	if in == nil {
		return nil
	}
	out := new(ScanResult)
	in.DeepCopyInto(out)
	return out
}
