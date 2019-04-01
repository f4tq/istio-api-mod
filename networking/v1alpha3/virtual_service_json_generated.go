// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/virtual_service.proto

package v1alpha3 // import "github.com/f4tq/istio-api-mod/networking/v1alpha3"

/*
Configuration affecting traffic routing. Here are a few terms useful to define
in the context of traffic routing.

`Service` a unit of application behavior bound to a unique name in a
service registry. Services consist of multiple network *endpoints*
implemented by workload instances running on pods, containers, VMs etc.

`Service versions (a.k.a. subsets)` - In a continuous deployment
scenario, for a given service, there can be distinct subsets of
instances running different variants of the application binary. These
variants are not necessarily different API versions. They could be
iterative changes to the same service, deployed in different
environments (prod, staging, dev, etc.). Common scenarios where this
occurs include A/B testing, canary rollouts, etc. The choice of a
particular version can be decided based on various criterion (headers,
url, etc.) and/or by weights assigned to each version. Each service has
a default version consisting of all its instances.

`Source` - A downstream client calling a service.

`Host` - The address used by a client when attempting to connect to a
service.

`Access model` - Applications address only the destination service
(Host) without knowledge of individual service versions (subsets). The
actual choice of the version is determined by the proxy/sidecar, enabling the
application code to decouple itself from the evolution of dependent
services.
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

// message: VirtualService
// skipping message: VirtualService
// message: Destination
// skipping message: Destination
// message: HTTPRoute
// Generating Marshal for message: HTTPRoute
// MarshalJSON is a custom marshaler supporting oneof fields for HTTPRoute
func (this *HTTPRoute) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for HTTPRoute
func (this *HTTPRoute) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: AppendHeadersEntry
// skipping message: AppendHeadersEntry
// message: AppendResponseHeadersEntry
// skipping message: AppendResponseHeadersEntry
// message: AppendRequestHeadersEntry
// skipping message: AppendRequestHeadersEntry
// message: Headers
// Generating Marshal for message: Headers
// MarshalJSON is a custom marshaler supporting oneof fields for Headers
func (this *Headers) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for Headers
func (this *Headers) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: HeaderOperations
// Generating Marshal for message: HeaderOperations
// MarshalJSON is a custom marshaler supporting oneof fields for Headers_HeaderOperations
func (this *Headers_HeaderOperations) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for Headers_HeaderOperations
func (this *Headers_HeaderOperations) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: SetEntry
// skipping message: SetEntry
// message: AddEntry
// skipping message: AddEntry
// message: TLSRoute
// skipping message: TLSRoute
// message: TCPRoute
// skipping message: TCPRoute
// message: HTTPMatchRequest
// Generating Marshal for message: HTTPMatchRequest
// MarshalJSON is a custom marshaler supporting oneof fields for HTTPMatchRequest
func (this *HTTPMatchRequest) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for HTTPMatchRequest
func (this *HTTPMatchRequest) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: HeadersEntry
// skipping message: HeadersEntry
// message: SourceLabelsEntry
// skipping message: SourceLabelsEntry
// message: HTTPRouteDestination
// Generating Marshal for message: HTTPRouteDestination
// MarshalJSON is a custom marshaler supporting oneof fields for HTTPRouteDestination
func (this *HTTPRouteDestination) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for HTTPRouteDestination
func (this *HTTPRouteDestination) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: AppendResponseHeadersEntry
// skipping message: AppendResponseHeadersEntry
// message: AppendRequestHeadersEntry
// skipping message: AppendRequestHeadersEntry
// message: RouteDestination
// skipping message: RouteDestination
// message: L4MatchAttributes
// Generating Marshal for message: L4MatchAttributes
// MarshalJSON is a custom marshaler supporting oneof fields for L4MatchAttributes
func (this *L4MatchAttributes) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for L4MatchAttributes
func (this *L4MatchAttributes) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: SourceLabelsEntry
// skipping message: SourceLabelsEntry
// message: TLSMatchAttributes
// Generating Marshal for message: TLSMatchAttributes
// MarshalJSON is a custom marshaler supporting oneof fields for TLSMatchAttributes
func (this *TLSMatchAttributes) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for TLSMatchAttributes
func (this *TLSMatchAttributes) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: SourceLabelsEntry
// skipping message: SourceLabelsEntry
// message: HTTPRedirect
// skipping message: HTTPRedirect
// message: HTTPRewrite
// skipping message: HTTPRewrite
// message: StringMatch
// Generating Marshal for message: StringMatch
// MarshalJSON is a custom marshaler supporting oneof fields for StringMatch
func (this *StringMatch) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for StringMatch
func (this *StringMatch) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: HTTPRetry
// skipping message: HTTPRetry
// message: CorsPolicy
// skipping message: CorsPolicy
// message: HTTPFaultInjection
// Generating Marshal for message: HTTPFaultInjection
// MarshalJSON is a custom marshaler supporting oneof fields for HTTPFaultInjection
func (this *HTTPFaultInjection) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for HTTPFaultInjection
func (this *HTTPFaultInjection) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: Delay
// Generating Marshal for message: Delay
// MarshalJSON is a custom marshaler supporting oneof fields for HTTPFaultInjection_Delay
func (this *HTTPFaultInjection_Delay) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for HTTPFaultInjection_Delay
func (this *HTTPFaultInjection_Delay) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: Abort
// Generating Marshal for message: Abort
// MarshalJSON is a custom marshaler supporting oneof fields for HTTPFaultInjection_Abort
func (this *HTTPFaultInjection_Abort) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for HTTPFaultInjection_Abort
func (this *HTTPFaultInjection_Abort) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: PortSelector
// Generating Marshal for message: PortSelector
// MarshalJSON is a custom marshaler supporting oneof fields for PortSelector
func (this *PortSelector) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for PortSelector
func (this *PortSelector) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// message: Percent
// skipping message: Percent
var (
	VirtualServiceMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	VirtualServiceUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
