// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: policy/v1beta1/type.proto

package v1beta1 // import "github.com/f4tq/istio-api-mod/policy/v1beta1"

/*
Describes the rules used to configure Mixer's policy and telemetry features.
*/

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

// MarshalJSON is a custom marshaler supporting oneof fields for Value
func (this *Value) MarshalJSON() ([]byte, error) {
	str, err := TypeMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for Value
func (this *Value) UnmarshalJSON(b []byte) error {
	return TypeUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	TypeMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	TypeUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)