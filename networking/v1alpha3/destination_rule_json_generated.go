// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/destination_rule.proto

package v1alpha3 // import "github.com/f4tq/istio-api-mod/networking/v1alpha3"

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

// message: DestinationRule
// skipping message: DestinationRule
// message: TrafficPolicy
// Generating Marshal for message: TrafficPolicy
// MarshalJSON is a custom marshaler supporting oneof fields for TrafficPolicy
func (this *TrafficPolicy) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for TrafficPolicy
func (this *TrafficPolicy) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: PortTrafficPolicy
// skipping message: PortTrafficPolicy
// message: Subset
// Generating Marshal for message: Subset
// MarshalJSON is a custom marshaler supporting oneof fields for Subset
func (this *Subset) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for Subset
func (this *Subset) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: LabelsEntry
// skipping message: LabelsEntry
// message: LoadBalancerSettings
// Generating Marshal for message: LoadBalancerSettings
// MarshalJSON is a custom marshaler supporting oneof fields for LoadBalancerSettings
func (this *LoadBalancerSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for LoadBalancerSettings
func (this *LoadBalancerSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: ConsistentHashLB
// Generating Marshal for message: ConsistentHashLB
// MarshalJSON is a custom marshaler supporting oneof fields for LoadBalancerSettings_ConsistentHashLB
func (this *LoadBalancerSettings_ConsistentHashLB) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for LoadBalancerSettings_ConsistentHashLB
func (this *LoadBalancerSettings_ConsistentHashLB) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: HTTPCookie
// skipping message: HTTPCookie
// message: ConnectionPoolSettings
// Generating Marshal for message: ConnectionPoolSettings
// MarshalJSON is a custom marshaler supporting oneof fields for ConnectionPoolSettings
func (this *ConnectionPoolSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for ConnectionPoolSettings
func (this *ConnectionPoolSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: TCPSettings
// Generating Marshal for message: TCPSettings
// MarshalJSON is a custom marshaler supporting oneof fields for ConnectionPoolSettings_TCPSettings
func (this *ConnectionPoolSettings_TCPSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for ConnectionPoolSettings_TCPSettings
func (this *ConnectionPoolSettings_TCPSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: TcpKeepalive
// skipping message: TcpKeepalive
// message: HTTPSettings
// skipping message: HTTPSettings
// message: OutlierDetection
// skipping message: OutlierDetection
// message: TLSSettings
// Generating Marshal for message: TLSSettings
// MarshalJSON is a custom marshaler supporting oneof fields for TLSSettings
func (this *TLSSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for TLSSettings
func (this *TLSSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	DestinationRuleMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	DestinationRuleUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
