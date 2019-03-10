// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rbac/v1alpha1/rbac.proto

package v1alpha1 // import "github.com/f4tq/istio-api-mod/rbac/v1alpha1"

/*
Istio RBAC (Role Based Access Control) defines ServiceRole and ServiceRoleBinding
objects.

A ServiceRole specification includes a list of rules (permissions). Each rule has
the following standard fields:

  * services: a list of services.
  * methods: HTTP methods. In the case of gRPC, this field is ignored because the value is always "POST".
  * paths: HTTP paths or gRPC methods. Note that gRPC methods should be
    presented in the form of "/packageName.serviceName/methodName" and are case sensitive.

In addition to the standard fields, operators can also use custom keys in the `constraints` field,
the supported keys are listed in the "constraints and properties" page.

Below is an example of ServiceRole object "product-viewer", which has "read" ("GET" and "HEAD")
access to "products.svc.cluster.local" service at versions "v1" and "v2". "path" is not specified,
so it applies to any path in the service.

```yaml
apiVersion: "rbac.istio.io/v1alpha1"
kind: ServiceRole
metadata:
  name: products-viewer
  namespace: default
spec:
  rules:
  - services: ["products.svc.cluster.local"]
    methods: ["GET", "HEAD"]
    constraints:
    - key: "destination.labels[version]"
      value: ["v1", "v2"]
```

A ServiceRoleBinding specification includes two parts:

 * The `roleRef` field that refers to a ServiceRole object in the same namespace.
 * A list of `subjects` that are assigned the roles.

In addition to a simple `user` field, operators can also use custom keys in the `properties` field,
the supported keys are listed in the "constraints and properties" page.

Below is an example of ServiceRoleBinding object "test-binding-products", which binds two subjects
to ServiceRole "product-viewer":

  * User "alice@yahoo.com"
  * Services in "abc" namespace.

```yaml
apiVersion: "rbac.istio.io/v1alpha1"
kind: ServiceRoleBinding
metadata:
  name: test-binding-products
  namespace: default
spec:
  subjects:
  - user: alice@yahoo.com
  - properties:
      source.namespace: "abc"
  roleRef:
    kind: ServiceRole
    name: "products-viewer"
```
*/

import bytes "bytes"
import github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler supporting oneof fields for WorkloadSelector
func (this *WorkloadSelector) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for WorkloadSelector
func (this *WorkloadSelector) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for WorkloadSelector_LabelsEntry
func (this *WorkloadSelector_LabelsEntry) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for WorkloadSelector_LabelsEntry
func (this *WorkloadSelector_LabelsEntry) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for AuthorizationPolicy
func (this *AuthorizationPolicy) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for AuthorizationPolicy
func (this *AuthorizationPolicy) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for ServiceRole
func (this *ServiceRole) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for ServiceRole
func (this *ServiceRole) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for AccessRule
func (this *AccessRule) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for AccessRule
func (this *AccessRule) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for AccessRule_Constraint
func (this *AccessRule_Constraint) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for AccessRule_Constraint
func (this *AccessRule_Constraint) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for ServiceRoleBinding
func (this *ServiceRoleBinding) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for ServiceRoleBinding
func (this *ServiceRoleBinding) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for Subject
func (this *Subject) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for Subject
func (this *Subject) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for Subject_PropertiesEntry
func (this *Subject_PropertiesEntry) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for Subject_PropertiesEntry
func (this *Subject_PropertiesEntry) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for RoleRef
func (this *RoleRef) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for RoleRef
func (this *RoleRef) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for RbacConfig
func (this *RbacConfig) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for RbacConfig
func (this *RbacConfig) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for RbacConfig_Target
func (this *RbacConfig_Target) MarshalJSON() ([]byte, error) {
	str, err := RbacMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for RbacConfig_Target
func (this *RbacConfig_Target) UnmarshalJSON(b []byte) error {
	return RbacUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	RbacMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	RbacUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
