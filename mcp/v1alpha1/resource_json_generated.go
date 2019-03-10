// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mcp/v1alpha1/resource.proto

package v1alpha1 // import "github.com/f4tq/istio-api-mod/mcp/v1alpha1"

/*
This package defines the common, core types used by the Mesh Configuration Protocol.
*/

import bytes "bytes"
import github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler supporting oneof fields for Resource
func (this *Resource) MarshalJSON() ([]byte, error) {
	str, err := ResourceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for Resource
func (this *Resource) UnmarshalJSON(b []byte) error {
	return ResourceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ResourceMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	ResourceUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
