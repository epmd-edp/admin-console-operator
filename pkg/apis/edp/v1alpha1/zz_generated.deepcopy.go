// +build !ignore_autogenerated

// Code generated by operator-sdk-2. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminConsole) DeepCopyInto(out *AdminConsole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminConsole.
func (in *AdminConsole) DeepCopy() *AdminConsole {
	if in == nil {
		return nil
	}
	out := new(AdminConsole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdminConsole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminConsoleDbSettings) DeepCopyInto(out *AdminConsoleDbSettings) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminConsoleDbSettings.
func (in *AdminConsoleDbSettings) DeepCopy() *AdminConsoleDbSettings {
	if in == nil {
		return nil
	}
	out := new(AdminConsoleDbSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminConsoleList) DeepCopyInto(out *AdminConsoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AdminConsole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminConsoleList.
func (in *AdminConsoleList) DeepCopy() *AdminConsoleList {
	if in == nil {
		return nil
	}
	out := new(AdminConsoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdminConsoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminConsoleSpec) DeepCopyInto(out *AdminConsoleSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	out.KeycloakSpec = in.KeycloakSpec
	out.EdpSpec = in.EdpSpec
	out.DbSpec = in.DbSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminConsoleSpec.
func (in *AdminConsoleSpec) DeepCopy() *AdminConsoleSpec {
	if in == nil {
		return nil
	}
	out := new(AdminConsoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminConsoleStatus) DeepCopyInto(out *AdminConsoleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminConsoleStatus.
func (in *AdminConsoleStatus) DeepCopy() *AdminConsoleStatus {
	if in == nil {
		return nil
	}
	out := new(AdminConsoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdpSpec) DeepCopyInto(out *EdpSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdpSpec.
func (in *EdpSpec) DeepCopy() *EdpSpec {
	if in == nil {
		return nil
	}
	out := new(EdpSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakSpec) DeepCopyInto(out *KeycloakSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakSpec.
func (in *KeycloakSpec) DeepCopy() *KeycloakSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakSpec)
	in.DeepCopyInto(out)
	return out
}
