//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v3alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Access) DeepCopyInto(out *Access) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Access.
func (in *Access) DeepCopy() *Access {
	if in == nil {
		return nil
	}
	out := new(Access)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	if in.RepositoryContexts != nil {
		in, out := &in.RepositoryContexts, &out.RepositoryContexts
		*out = make([]OCIRepositoryContext, len(*in))
		copy(*out, *in)
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]SourceDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ComponentReferences != nil {
		in, out := &in.ComponentReferences, &out.ComponentReferences
		*out = make([]ComponentReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Label, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDescriptor) DeepCopyInto(out *ComponentDescriptor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDescriptor.
func (in *ComponentDescriptor) DeepCopy() *ComponentDescriptor {
	if in == nil {
		return nil
	}
	out := new(ComponentDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentDescriptor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDescriptorList) DeepCopyInto(out *ComponentDescriptorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComponentDescriptor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDescriptorList.
func (in *ComponentDescriptorList) DeepCopy() *ComponentDescriptorList {
	if in == nil {
		return nil
	}
	out := new(ComponentDescriptorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentDescriptorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDescriptorSpec) DeepCopyInto(out *ComponentDescriptorSpec) {
	*out = *in
	in.Meta.DeepCopyInto(&out.Meta)
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = make([]Label, len(*in))
		copy(*out, *in)
	}
	in.IdentityAttribute.DeepCopyInto(&out.IdentityAttribute)
	in.Component.DeepCopyInto(&out.Component)
	out.RepositoryContext = in.RepositoryContext
	out.Access = in.Access
	out.DigestSpec = in.DigestSpec
	out.SignatureSpec = in.SignatureSpec
	out.Signature = in.Signature
	in.Source.DeepCopyInto(&out.Source)
	in.ComponentReference.DeepCopyInto(&out.ComponentReference)
	in.Resource.DeepCopyInto(&out.Resource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDescriptorSpec.
func (in *ComponentDescriptorSpec) DeepCopy() *ComponentDescriptorSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentDescriptorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDescriptorStatus) DeepCopyInto(out *ComponentDescriptorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDescriptorStatus.
func (in *ComponentDescriptorStatus) DeepCopy() *ComponentDescriptorStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentDescriptorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentReference) DeepCopyInto(out *ComponentReference) {
	*out = *in
	in.ExtraIdentity.DeepCopyInto(&out.ExtraIdentity)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Label, len(*in))
		copy(*out, *in)
	}
	out.Digest = in.Digest
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentReference.
func (in *ComponentReference) DeepCopy() *ComponentReference {
	if in == nil {
		return nil
	}
	out := new(ComponentReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DigestSpec) DeepCopyInto(out *DigestSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DigestSpec.
func (in *DigestSpec) DeepCopy() *DigestSpec {
	if in == nil {
		return nil
	}
	out := new(DigestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericAccess) DeepCopyInto(out *GenericAccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericAccess.
func (in *GenericAccess) DeepCopy() *GenericAccess {
	if in == nil {
		return nil
	}
	out := new(GenericAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericResource) DeepCopyInto(out *GenericResource) {
	*out = *in
	out.Access = in.Access
	out.Digest = in.Digest
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Label, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericResource.
func (in *GenericResource) DeepCopy() *GenericResource {
	if in == nil {
		return nil
	}
	out := new(GenericResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubAccess) DeepCopyInto(out *GithubAccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubAccess.
func (in *GithubAccess) DeepCopy() *GithubAccess {
	if in == nil {
		return nil
	}
	out := new(GithubAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPAccess) DeepCopyInto(out *HTTPAccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPAccess.
func (in *HTTPAccess) DeepCopy() *HTTPAccess {
	if in == nil {
		return nil
	}
	out := new(HTTPAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityAttribute) DeepCopyInto(out *IdentityAttribute) {
	*out = *in
	if in.IdentityAttributeKeys != nil {
		in, out := &in.IdentityAttributeKeys, &out.IdentityAttributeKeys
		*out = make([]IdentityAttributeKey, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityAttribute.
func (in *IdentityAttribute) DeepCopy() *IdentityAttribute {
	if in == nil {
		return nil
	}
	out := new(IdentityAttribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Label) DeepCopyInto(out *Label) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Label.
func (in *Label) DeepCopy() *Label {
	if in == nil {
		return nil
	}
	out := new(Label)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalFilesystemBlobAccess) DeepCopyInto(out *LocalFilesystemBlobAccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalFilesystemBlobAccess.
func (in *LocalFilesystemBlobAccess) DeepCopy() *LocalFilesystemBlobAccess {
	if in == nil {
		return nil
	}
	out := new(LocalFilesystemBlobAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalOciBlobAccess) DeepCopyInto(out *LocalOciBlobAccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalOciBlobAccess.
func (in *LocalOciBlobAccess) DeepCopy() *LocalOciBlobAccess {
	if in == nil {
		return nil
	}
	out := new(LocalOciBlobAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Meta) DeepCopyInto(out *Meta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Label, len(*in))
		copy(*out, *in)
	}
	in.Provider.DeepCopyInto(&out.Provider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Meta.
func (in *Meta) DeepCopy() *Meta {
	if in == nil {
		return nil
	}
	out := new(Meta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoneAccess) DeepCopyInto(out *NoneAccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoneAccess.
func (in *NoneAccess) DeepCopy() *NoneAccess {
	if in == nil {
		return nil
	}
	out := new(NoneAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIBlobAccess) DeepCopyInto(out *OCIBlobAccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIBlobAccess.
func (in *OCIBlobAccess) DeepCopy() *OCIBlobAccess {
	if in == nil {
		return nil
	}
	out := new(OCIBlobAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIImageAccess) DeepCopyInto(out *OCIImageAccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIImageAccess.
func (in *OCIImageAccess) DeepCopy() *OCIImageAccess {
	if in == nil {
		return nil
	}
	out := new(OCIImageAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIImageResource) DeepCopyInto(out *OCIImageResource) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Label, len(*in))
		copy(*out, *in)
	}
	out.Access = in.Access
	in.ExtraIdentity.DeepCopyInto(&out.ExtraIdentity)
	out.Digest = in.Digest
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIImageResource.
func (in *OCIImageResource) DeepCopy() *OCIImageResource {
	if in == nil {
		return nil
	}
	out := new(OCIImageResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIRepositoryContext) DeepCopyInto(out *OCIRepositoryContext) {
	*out = *in
	out.RepositoryContext = in.RepositoryContext
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIRepositoryContext.
func (in *OCIRepositoryContext) DeepCopy() *OCIRepositoryContext {
	if in == nil {
		return nil
	}
	out := new(OCIRepositoryContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Object) DeepCopyInto(out *Object) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Object.
func (in *Object) DeepCopy() *Object {
	if in == nil {
		return nil
	}
	out := new(Object)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Provider) DeepCopyInto(out *Provider) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Label, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Provider.
func (in *Provider) DeepCopy() *Provider {
	if in == nil {
		return nil
	}
	out := new(Provider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryContext) DeepCopyInto(out *RepositoryContext) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryContext.
func (in *RepositoryContext) DeepCopy() *RepositoryContext {
	if in == nil {
		return nil
	}
	out := new(RepositoryContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceType) DeepCopyInto(out *ResourceType) {
	*out = *in
	out.Access = in.Access
	if in.SrcRefs != nil {
		in, out := &in.SrcRefs, &out.SrcRefs
		*out = make([]SourceReferences, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ExtraIdentity.DeepCopyInto(&out.ExtraIdentity)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Label, len(*in))
		copy(*out, *in)
	}
	out.Digest = in.Digest
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceType.
func (in *ResourceType) DeepCopy() *ResourceType {
	if in == nil {
		return nil
	}
	out := new(ResourceType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Signature) DeepCopyInto(out *Signature) {
	*out = *in
	out.Digest = in.Digest
	out.Signature = in.Signature
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Signature.
func (in *Signature) DeepCopy() *Signature {
	if in == nil {
		return nil
	}
	out := new(Signature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureSpec) DeepCopyInto(out *SignatureSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureSpec.
func (in *SignatureSpec) DeepCopy() *SignatureSpec {
	if in == nil {
		return nil
	}
	out := new(SignatureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceDefinition) DeepCopyInto(out *SourceDefinition) {
	*out = *in
	out.Access = in.Access
	in.ExtraIdentity.DeepCopyInto(&out.ExtraIdentity)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Label, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceDefinition.
func (in *SourceDefinition) DeepCopy() *SourceDefinition {
	if in == nil {
		return nil
	}
	out := new(SourceDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceReferences) DeepCopyInto(out *SourceReferences) {
	*out = *in
	in.IdentitySelector.DeepCopyInto(&out.IdentitySelector)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Label, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceReferences.
func (in *SourceReferences) DeepCopy() *SourceReferences {
	if in == nil {
		return nil
	}
	out := new(SourceReferences)
	in.DeepCopyInto(out)
	return out
}
