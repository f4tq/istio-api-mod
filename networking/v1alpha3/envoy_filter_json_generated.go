// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/envoy_filter.proto

package v1alpha3 // import "github.com/f4tq/istio-api-mod/networking/v1alpha3"

import bytes "bytes"
import github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler supporting oneof fields for EnvoyFilter
func (this *EnvoyFilter) MarshalJSON() ([]byte, error) {
	str, err := EnvoyFilterMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for EnvoyFilter
func (this *EnvoyFilter) UnmarshalJSON(b []byte) error {
	return EnvoyFilterUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for EnvoyFilter_WorkloadLabelsEntry
func (this *EnvoyFilter_WorkloadLabelsEntry) MarshalJSON() ([]byte, error) {
	str, err := EnvoyFilterMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for EnvoyFilter_WorkloadLabelsEntry
func (this *EnvoyFilter_WorkloadLabelsEntry) UnmarshalJSON(b []byte) error {
	return EnvoyFilterUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for EnvoyFilter_ListenerMatch
func (this *EnvoyFilter_ListenerMatch) MarshalJSON() ([]byte, error) {
	str, err := EnvoyFilterMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for EnvoyFilter_ListenerMatch
func (this *EnvoyFilter_ListenerMatch) UnmarshalJSON(b []byte) error {
	return EnvoyFilterUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for EnvoyFilter_InsertPosition
func (this *EnvoyFilter_InsertPosition) MarshalJSON() ([]byte, error) {
	str, err := EnvoyFilterMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for EnvoyFilter_InsertPosition
func (this *EnvoyFilter_InsertPosition) UnmarshalJSON(b []byte) error {
	return EnvoyFilterUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for EnvoyFilter_Filter
func (this *EnvoyFilter_Filter) MarshalJSON() ([]byte, error) {
	str, err := EnvoyFilterMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for EnvoyFilter_Filter
func (this *EnvoyFilter_Filter) UnmarshalJSON(b []byte) error {
	return EnvoyFilterUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	EnvoyFilterMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	EnvoyFilterUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
