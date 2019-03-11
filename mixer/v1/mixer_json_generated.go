// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/v1/mixer.proto

package v1 // import "github.com/f4tq/istio-api-mod/mixer/v1"

/*
This package defines the Mixer API that the sidecar proxy uses to perform
precondition checks, manage quotas, and report telemetry.
*/

import bytes "bytes"
import github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/googleapis/google/rpc"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler supporting oneof fields for ReferencedAttributes
func (this *ReferencedAttributes) MarshalJSON() ([]byte, error) {
	str, err := MixerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for ReferencedAttributes
func (this *ReferencedAttributes) UnmarshalJSON(b []byte) error {
	return MixerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for HeaderOperation
func (this *HeaderOperation) MarshalJSON() ([]byte, error) {
	str, err := MixerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for HeaderOperation
func (this *HeaderOperation) UnmarshalJSON(b []byte) error {
	return MixerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	MixerMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	MixerUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
