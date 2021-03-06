// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/model/v1beta1/extensions.proto

package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The available varieties of templates, controlling the semantics of what an adapter does with each instance.
type TemplateVariety int32

const (
	// Makes the template applicable for Mixer's check calls. Instances of such template are created during
	// check calls in Mixer and passed to the handlers based on the rule configurations.
	TEMPLATE_VARIETY_CHECK TemplateVariety = 0
	// Makes the template applicable for Mixer's report calls. Instances of such template are created during
	// report calls in Mixer and passed to the handlers based on the rule configurations.
	TEMPLATE_VARIETY_REPORT TemplateVariety = 1
	// Makes the template applicable for Mixer's quota calls. Instances of such template are created during
	// quota check calls in Mixer and passed to the handlers based on the rule configurations.
	TEMPLATE_VARIETY_QUOTA TemplateVariety = 2
	// Makes the template applicable for Mixer's attribute generation phase. Instances of such template are created during
	// pre-processing attribute generation phase and passed to the handlers based on the rule configurations.
	TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR TemplateVariety = 3
	// Makes the template applicable for Mixer's check calls. Instances of such template are created during
	// check calls in Mixer and passed to the handlers that produce values.
	TEMPLATE_VARIETY_CHECK_WITH_OUTPUT TemplateVariety = 4
)

var TemplateVariety_name = map[int32]string{
	0: "TEMPLATE_VARIETY_CHECK",
	1: "TEMPLATE_VARIETY_REPORT",
	2: "TEMPLATE_VARIETY_QUOTA",
	3: "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR",
	4: "TEMPLATE_VARIETY_CHECK_WITH_OUTPUT",
}
var TemplateVariety_value = map[string]int32{
	"TEMPLATE_VARIETY_CHECK":               0,
	"TEMPLATE_VARIETY_REPORT":              1,
	"TEMPLATE_VARIETY_QUOTA":               2,
	"TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR": 3,
	"TEMPLATE_VARIETY_CHECK_WITH_OUTPUT":   4,
}

func (TemplateVariety) EnumDescriptor() ([]byte, []int) { return fileDescriptorExtensions, []int{0} }

var E_TemplateVariety = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*TemplateVariety)(nil),
	Field:         72295727,
	Name:          "istio.mixer.adapter.model.v1beta1.template_variety",
	Tag:           "varint,72295727,opt,name=template_variety,json=templateVariety,enum=istio.mixer.adapter.model.v1beta1.TemplateVariety",
	Filename:      "mixer/adapter/model/v1beta1/extensions.proto",
}

var E_TemplateName = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         72295888,
	Name:          "istio.mixer.adapter.model.v1beta1.template_name",
	Tag:           "bytes,72295888,opt,name=template_name,json=templateName",
	Filename:      "mixer/adapter/model/v1beta1/extensions.proto",
}

func init() {
	proto.RegisterEnum("istio.mixer.adapter.model.v1beta1.TemplateVariety", TemplateVariety_name, TemplateVariety_value)
	proto.RegisterExtension(E_TemplateVariety)
	proto.RegisterExtension(E_TemplateName)
}
func (x TemplateVariety) String() string {
	s, ok := TemplateVariety_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() {
	proto.RegisterFile("mixer/adapter/model/v1beta1/extensions.proto", fileDescriptorExtensions)
}

var fileDescriptorExtensions = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x8a, 0xd3, 0x40,
	0x1c, 0xc7, 0x33, 0x2a, 0x82, 0x83, 0xba, 0x21, 0x07, 0x95, 0x55, 0x86, 0x75, 0x11, 0x59, 0xc4,
	0x9d, 0x61, 0xab, 0x17, 0x7b, 0x4b, 0xeb, 0x68, 0x8b, 0xda, 0xd4, 0x61, 0x52, 0xd1, 0x4b, 0x98,
	0x34, 0xd3, 0x38, 0x90, 0x74, 0x62, 0x32, 0x2d, 0xf5, 0x22, 0x3e, 0x82, 0x8f, 0xe1, 0xcd, 0xab,
	0x07, 0x1f, 0x40, 0xf4, 0xd2, 0xa3, 0x47, 0x1b, 0x2f, 0x1e, 0xfb, 0x08, 0x62, 0x9a, 0x0a, 0xda,
	0x62, 0xaf, 0xf3, 0xe5, 0x33, 0xdf, 0x3f, 0xfc, 0xe0, 0xed, 0x54, 0xcd, 0x64, 0x4e, 0x44, 0x24,
	0x32, 0x23, 0x73, 0x92, 0xea, 0x48, 0x26, 0x64, 0x7a, 0x12, 0x4a, 0x23, 0x4e, 0x88, 0x9c, 0x19,
	0x39, 0x2e, 0x94, 0x1e, 0x17, 0x38, 0xcb, 0xb5, 0xd1, 0xce, 0x75, 0x55, 0x18, 0xa5, 0x71, 0xc5,
	0xe0, 0x9a, 0xc1, 0x15, 0x83, 0x6b, 0x66, 0xff, 0x20, 0xd6, 0x3a, 0x4e, 0x24, 0xa9, 0x80, 0x70,
	0x32, 0x22, 0x91, 0x2c, 0x86, 0xb9, 0xca, 0x8c, 0xce, 0x57, 0x9f, 0xdc, 0xfa, 0x08, 0xe0, 0x1e,
	0x97, 0x69, 0x96, 0x08, 0x23, 0x07, 0x22, 0x57, 0xd2, 0xbc, 0x76, 0xf6, 0xe1, 0x25, 0x4e, 0x9f,
	0xf4, 0x1f, 0xbb, 0x9c, 0x06, 0x03, 0x97, 0x75, 0x29, 0x7f, 0x1e, 0xb4, 0x3b, 0xb4, 0xfd, 0xc8,
	0xb6, 0x9c, 0xab, 0xf0, 0xf2, 0x86, 0xc6, 0x68, 0xdf, 0x63, 0xdc, 0x06, 0x5b, 0xc1, 0xa7, 0xbe,
	0xc7, 0x5d, 0xfb, 0x94, 0x73, 0x04, 0x6f, 0x6c, 0x68, 0x2e, 0xe7, 0xac, 0xdb, 0xf2, 0x39, 0x0d,
	0x1e, 0xd2, 0x1e, 0x65, 0x2e, 0xf7, 0x98, 0x7d, 0xda, 0xb9, 0x09, 0x0f, 0xb7, 0xdb, 0x07, 0xcf,
	0xba, 0xbc, 0x13, 0x78, 0x3e, 0xef, 0xfb, 0xdc, 0x3e, 0xd3, 0x7c, 0x03, 0x6d, 0x53, 0x27, 0x0f,
	0xa6, 0x75, 0xf4, 0x6b, 0x78, 0xd5, 0x18, 0xaf, 0x1b, 0xe3, 0x07, 0x2a, 0x91, 0x5e, 0x66, 0x7e,
	0xef, 0x76, 0xe5, 0xc3, 0x97, 0x4f, 0x87, 0x07, 0xe0, 0xe8, 0x62, 0xa3, 0x81, 0x77, 0x6e, 0x87,
	0xff, 0x59, 0x85, 0xed, 0x99, 0xbf, 0x1f, 0x9a, 0xf7, 0xe1, 0x85, 0x3f, 0xfe, 0x63, 0x91, 0xca,
	0x1d, 0xe6, 0xf3, 0xaf, 0x95, 0xf9, 0x39, 0x76, 0x7e, 0x4d, 0xf5, 0x44, 0x2a, 0x5b, 0xf1, 0x7c,
	0x81, 0xac, 0x6f, 0x0b, 0x64, 0x2d, 0x17, 0x08, 0xbc, 0x2d, 0x11, 0x78, 0x5f, 0x22, 0xf0, 0xb9,
	0x44, 0x60, 0x5e, 0x22, 0xf0, 0xbd, 0x44, 0xe0, 0x67, 0x89, 0xac, 0x65, 0x89, 0xc0, 0xbb, 0x1f,
	0xc8, 0x7a, 0x71, 0x2f, 0x56, 0xe6, 0xe5, 0x24, 0xc4, 0x43, 0x9d, 0x92, 0xd1, 0x5d, 0xf3, 0x8a,
	0x54, 0x3d, 0x8e, 0x45, 0xa6, 0x8e, 0x53, 0x1d, 0x91, 0xff, 0xdc, 0x4f, 0x78, 0xb6, 0x4a, 0x75,
	0xe7, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0xd0, 0x4e, 0x07, 0x65, 0x02, 0x00, 0x00,
}
