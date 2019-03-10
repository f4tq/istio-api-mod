// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/service_entry.proto

package v1alpha3 // import "github.com/f4tq/istio-api-mod/networking/v1alpha3"

import bytes "bytes"
import github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler supporting oneof fields for ServiceEntry
func (this *ServiceEntry) MarshalJSON() ([]byte, error) {
	str, err := ServiceEntryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for ServiceEntry
func (this *ServiceEntry) UnmarshalJSON(b []byte) error {
	return ServiceEntryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for ServiceEntry_Endpoint
func (this *ServiceEntry_Endpoint) MarshalJSON() ([]byte, error) {
	str, err := ServiceEntryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for ServiceEntry_Endpoint
func (this *ServiceEntry_Endpoint) UnmarshalJSON(b []byte) error {
	return ServiceEntryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for ServiceEntry_Endpoint_PortsEntry
func (this *ServiceEntry_Endpoint_PortsEntry) MarshalJSON() ([]byte, error) {
	str, err := ServiceEntryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for ServiceEntry_Endpoint_PortsEntry
func (this *ServiceEntry_Endpoint_PortsEntry) UnmarshalJSON(b []byte) error {
	return ServiceEntryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for ServiceEntry_Endpoint_LabelsEntry
func (this *ServiceEntry_Endpoint_LabelsEntry) MarshalJSON() ([]byte, error) {
	str, err := ServiceEntryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for ServiceEntry_Endpoint_LabelsEntry
func (this *ServiceEntry_Endpoint_LabelsEntry) UnmarshalJSON(b []byte) error {
	return ServiceEntryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ServiceEntryMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	ServiceEntryUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
