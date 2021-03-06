package v1alpha1

import (
	"github.com/f4tq/istio-api-mod/util"
)

// DeepCopyInto for RbacConfig type.
func (in *RbacConfig) DeepCopyInto(out *RbacConfig) {
	util.DeepCopyInto(in, out)
}
// DeepCopyInto for ServiceRole type.
func (in *ServiceRole) DeepCopyInto(out *ServiceRole) {
	util.DeepCopyInto(in, out)
}
// DeepCopyInto for ServiceRoleBinding type.
func (in *ServiceRoleBinding) DeepCopyInto(out *ServiceRoleBinding) {
	util.DeepCopyInto(in, out)
}
