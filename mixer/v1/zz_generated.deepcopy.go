// +build !ignore_autogenerated

// Copyright 2019 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	time "time"

	types "github.com/gogo/protobuf/types"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attributes) DeepCopyInto(out *Attributes) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]*Attributes_AttributeValue, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(Attributes_AttributeValue)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attributes.
func (in *Attributes) DeepCopy() *Attributes {
	if in == nil {
		return nil
	}
	out := new(Attributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attributes_AttributeValue) DeepCopyInto(out *Attributes_AttributeValue) {
	*out = *in
	if in.Value == nil {
		out.Value = nil
	} else {
		out.Value = in.Value.DeepCopyisAttributes_AttributeValue_Value()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attributes_AttributeValue.
func (in *Attributes_AttributeValue) DeepCopy() *Attributes_AttributeValue {
	if in == nil {
		return nil
	}
	out := new(Attributes_AttributeValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attributes_AttributeValue_BoolValue) DeepCopyInto(out *Attributes_AttributeValue_BoolValue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attributes_AttributeValue_BoolValue.
func (in *Attributes_AttributeValue_BoolValue) DeepCopy() *Attributes_AttributeValue_BoolValue {
	if in == nil {
		return nil
	}
	out := new(Attributes_AttributeValue_BoolValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attributes_AttributeValue_BytesValue) DeepCopyInto(out *Attributes_AttributeValue_BytesValue) {
	*out = *in
	if in.BytesValue != nil {
		in, out := &in.BytesValue, &out.BytesValue
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attributes_AttributeValue_BytesValue.
func (in *Attributes_AttributeValue_BytesValue) DeepCopy() *Attributes_AttributeValue_BytesValue {
	if in == nil {
		return nil
	}
	out := new(Attributes_AttributeValue_BytesValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attributes_AttributeValue_DoubleValue) DeepCopyInto(out *Attributes_AttributeValue_DoubleValue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attributes_AttributeValue_DoubleValue.
func (in *Attributes_AttributeValue_DoubleValue) DeepCopy() *Attributes_AttributeValue_DoubleValue {
	if in == nil {
		return nil
	}
	out := new(Attributes_AttributeValue_DoubleValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attributes_AttributeValue_DurationValue) DeepCopyInto(out *Attributes_AttributeValue_DurationValue) {
	*out = *in
	if in.DurationValue != nil {
		in, out := &in.DurationValue, &out.DurationValue
		if *in == nil {
			*out = nil
		} else {
			*out = new(types.Duration)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attributes_AttributeValue_DurationValue.
func (in *Attributes_AttributeValue_DurationValue) DeepCopy() *Attributes_AttributeValue_DurationValue {
	if in == nil {
		return nil
	}
	out := new(Attributes_AttributeValue_DurationValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attributes_AttributeValue_Int64Value) DeepCopyInto(out *Attributes_AttributeValue_Int64Value) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attributes_AttributeValue_Int64Value.
func (in *Attributes_AttributeValue_Int64Value) DeepCopy() *Attributes_AttributeValue_Int64Value {
	if in == nil {
		return nil
	}
	out := new(Attributes_AttributeValue_Int64Value)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attributes_AttributeValue_StringMapValue) DeepCopyInto(out *Attributes_AttributeValue_StringMapValue) {
	*out = *in
	if in.StringMapValue != nil {
		in, out := &in.StringMapValue, &out.StringMapValue
		if *in == nil {
			*out = nil
		} else {
			*out = new(Attributes_StringMap)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attributes_AttributeValue_StringMapValue.
func (in *Attributes_AttributeValue_StringMapValue) DeepCopy() *Attributes_AttributeValue_StringMapValue {
	if in == nil {
		return nil
	}
	out := new(Attributes_AttributeValue_StringMapValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attributes_AttributeValue_StringValue) DeepCopyInto(out *Attributes_AttributeValue_StringValue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attributes_AttributeValue_StringValue.
func (in *Attributes_AttributeValue_StringValue) DeepCopy() *Attributes_AttributeValue_StringValue {
	if in == nil {
		return nil
	}
	out := new(Attributes_AttributeValue_StringValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attributes_AttributeValue_TimestampValue) DeepCopyInto(out *Attributes_AttributeValue_TimestampValue) {
	*out = *in
	if in.TimestampValue != nil {
		in, out := &in.TimestampValue, &out.TimestampValue
		if *in == nil {
			*out = nil
		} else {
			*out = new(types.Timestamp)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attributes_AttributeValue_TimestampValue.
func (in *Attributes_AttributeValue_TimestampValue) DeepCopy() *Attributes_AttributeValue_TimestampValue {
	if in == nil {
		return nil
	}
	out := new(Attributes_AttributeValue_TimestampValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attributes_StringMap) DeepCopyInto(out *Attributes_StringMap) {
	*out = *in
	if in.Entries != nil {
		in, out := &in.Entries, &out.Entries
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attributes_StringMap.
func (in *Attributes_StringMap) DeepCopy() *Attributes_StringMap {
	if in == nil {
		return nil
	}
	out := new(Attributes_StringMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheckRequest) DeepCopyInto(out *CheckRequest) {
	*out = *in
	in.Attributes.DeepCopyInto(&out.Attributes)
	if in.Quotas != nil {
		in, out := &in.Quotas, &out.Quotas
		*out = make(map[string]CheckRequest_QuotaParams, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheckRequest.
func (in *CheckRequest) DeepCopy() *CheckRequest {
	if in == nil {
		return nil
	}
	out := new(CheckRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheckRequest_QuotaParams) DeepCopyInto(out *CheckRequest_QuotaParams) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheckRequest_QuotaParams.
func (in *CheckRequest_QuotaParams) DeepCopy() *CheckRequest_QuotaParams {
	if in == nil {
		return nil
	}
	out := new(CheckRequest_QuotaParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheckResponse) DeepCopyInto(out *CheckResponse) {
	*out = *in
	in.Precondition.DeepCopyInto(&out.Precondition)
	if in.Quotas != nil {
		in, out := &in.Quotas, &out.Quotas
		*out = make(map[string]CheckResponse_QuotaResult, len(*in))
		for key, val := range *in {
			newVal := new(CheckResponse_QuotaResult)
			val.DeepCopyInto(newVal)
			(*out)[key] = *newVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheckResponse.
func (in *CheckResponse) DeepCopy() *CheckResponse {
	if in == nil {
		return nil
	}
	out := new(CheckResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheckResponse_PreconditionResult) DeepCopyInto(out *CheckResponse_PreconditionResult) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.ReferencedAttributes != nil {
		in, out := &in.ReferencedAttributes, &out.ReferencedAttributes
		if *in == nil {
			*out = nil
		} else {
			*out = new(ReferencedAttributes)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.RouteDirective != nil {
		in, out := &in.RouteDirective, &out.RouteDirective
		if *in == nil {
			*out = nil
		} else {
			*out = new(RouteDirective)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheckResponse_PreconditionResult.
func (in *CheckResponse_PreconditionResult) DeepCopy() *CheckResponse_PreconditionResult {
	if in == nil {
		return nil
	}
	out := new(CheckResponse_PreconditionResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheckResponse_QuotaResult) DeepCopyInto(out *CheckResponse_QuotaResult) {
	*out = *in
	in.ReferencedAttributes.DeepCopyInto(&out.ReferencedAttributes)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheckResponse_QuotaResult.
func (in *CheckResponse_QuotaResult) DeepCopy() *CheckResponse_QuotaResult {
	if in == nil {
		return nil
	}
	out := new(CheckResponse_QuotaResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompressedAttributes) DeepCopyInto(out *CompressedAttributes) {
	*out = *in
	if in.Words != nil {
		in, out := &in.Words, &out.Words
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Strings != nil {
		in, out := &in.Strings, &out.Strings
		*out = make(map[int32]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Int64S != nil {
		in, out := &in.Int64S, &out.Int64S
		*out = make(map[int32]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Doubles != nil {
		in, out := &in.Doubles, &out.Doubles
		*out = make(map[int32]float64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Bools != nil {
		in, out := &in.Bools, &out.Bools
		*out = make(map[int32]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Timestamps != nil {
		in, out := &in.Timestamps, &out.Timestamps
		*out = make(map[int32]time.Time, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Durations != nil {
		in, out := &in.Durations, &out.Durations
		*out = make(map[int32]time.Duration, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Bytes != nil {
		in, out := &in.Bytes, &out.Bytes
		*out = make(map[int32][]byte, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = make([]byte, len(val))
				copy((*out)[key], val)
			}
		}
	}
	if in.StringMaps != nil {
		in, out := &in.StringMaps, &out.StringMaps
		*out = make(map[int32]StringMap, len(*in))
		for key, val := range *in {
			newVal := new(StringMap)
			val.DeepCopyInto(newVal)
			(*out)[key] = *newVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompressedAttributes.
func (in *CompressedAttributes) DeepCopy() *CompressedAttributes {
	if in == nil {
		return nil
	}
	out := new(CompressedAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderOperation) DeepCopyInto(out *HeaderOperation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderOperation.
func (in *HeaderOperation) DeepCopy() *HeaderOperation {
	if in == nil {
		return nil
	}
	out := new(HeaderOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReferencedAttributes) DeepCopyInto(out *ReferencedAttributes) {
	*out = *in
	if in.Words != nil {
		in, out := &in.Words, &out.Words
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AttributeMatches != nil {
		in, out := &in.AttributeMatches, &out.AttributeMatches
		*out = make([]ReferencedAttributes_AttributeMatch, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReferencedAttributes.
func (in *ReferencedAttributes) DeepCopy() *ReferencedAttributes {
	if in == nil {
		return nil
	}
	out := new(ReferencedAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReferencedAttributes_AttributeMatch) DeepCopyInto(out *ReferencedAttributes_AttributeMatch) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReferencedAttributes_AttributeMatch.
func (in *ReferencedAttributes_AttributeMatch) DeepCopy() *ReferencedAttributes_AttributeMatch {
	if in == nil {
		return nil
	}
	out := new(ReferencedAttributes_AttributeMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportRequest) DeepCopyInto(out *ReportRequest) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make([]CompressedAttributes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultWords != nil {
		in, out := &in.DefaultWords, &out.DefaultWords
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportRequest.
func (in *ReportRequest) DeepCopy() *ReportRequest {
	if in == nil {
		return nil
	}
	out := new(ReportRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportResponse) DeepCopyInto(out *ReportResponse) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportResponse.
func (in *ReportResponse) DeepCopy() *ReportResponse {
	if in == nil {
		return nil
	}
	out := new(ReportResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteDirective) DeepCopyInto(out *RouteDirective) {
	*out = *in
	if in.RequestHeaderOperations != nil {
		in, out := &in.RequestHeaderOperations, &out.RequestHeaderOperations
		*out = make([]HeaderOperation, len(*in))
		copy(*out, *in)
	}
	if in.ResponseHeaderOperations != nil {
		in, out := &in.ResponseHeaderOperations, &out.ResponseHeaderOperations
		*out = make([]HeaderOperation, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteDirective.
func (in *RouteDirective) DeepCopy() *RouteDirective {
	if in == nil {
		return nil
	}
	out := new(RouteDirective)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringMap) DeepCopyInto(out *StringMap) {
	*out = *in
	if in.Entries != nil {
		in, out := &in.Entries, &out.Entries
		*out = make(map[int32]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringMap.
func (in *StringMap) DeepCopy() *StringMap {
	if in == nil {
		return nil
	}
	out := new(StringMap)
	in.DeepCopyInto(out)
	return out
}