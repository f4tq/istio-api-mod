// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/v1/attributes.proto

package v1 // import "github.com/f4tq/istio-api-mod/mixer/v1"

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

// MarshalJSON is a custom marshaler supporting oneof fields for Attributes_AttributeValue
func (this *Attributes_AttributeValue) MarshalJSON() ([]byte, error) {
	str, err := AttributesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for Attributes_AttributeValue
func (this *Attributes_AttributeValue) UnmarshalJSON(b []byte) error {
	return AttributesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	AttributesMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	AttributesUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
