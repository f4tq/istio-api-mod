package v1alpha1

import (
	"github.com/f4tq/istio-api-mod/util"
)

// DeepCopyInto for Policy type.
func (in *Policy) DeepCopyInto(out *Policy) {
	util.DeepCopyInto(in, out)
}
