// Copyright 2019 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	internalinterfaces "github.com/f4tq/istio-api-mod/pkg/kube/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AttributeManifests returns a AttributeManifestInformer.
	AttributeManifests() AttributeManifestInformer
	// HTTPAPISpecs returns a HTTPAPISpecInformer.
	HTTPAPISpecs() HTTPAPISpecInformer
	// HTTPAPISpecBindings returns a HTTPAPISpecBindingInformer.
	HTTPAPISpecBindings() HTTPAPISpecBindingInformer
	// Handlers returns a HandlerInformer.
	Handlers() HandlerInformer
	// Instances returns a InstanceInformer.
	Instances() InstanceInformer
	// QuotaSpecs returns a QuotaSpecInformer.
	QuotaSpecs() QuotaSpecInformer
	// QuotaSpecBindings returns a QuotaSpecBindingInformer.
	QuotaSpecBindings() QuotaSpecBindingInformer
	// Rules returns a RuleInformer.
	Rules() RuleInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AttributeManifests returns a AttributeManifestInformer.
func (v *version) AttributeManifests() AttributeManifestInformer {
	return &attributeManifestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HTTPAPISpecs returns a HTTPAPISpecInformer.
func (v *version) HTTPAPISpecs() HTTPAPISpecInformer {
	return &hTTPAPISpecInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HTTPAPISpecBindings returns a HTTPAPISpecBindingInformer.
func (v *version) HTTPAPISpecBindings() HTTPAPISpecBindingInformer {
	return &hTTPAPISpecBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Handlers returns a HandlerInformer.
func (v *version) Handlers() HandlerInformer {
	return &handlerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Instances returns a InstanceInformer.
func (v *version) Instances() InstanceInformer {
	return &instanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// QuotaSpecs returns a QuotaSpecInformer.
func (v *version) QuotaSpecs() QuotaSpecInformer {
	return &quotaSpecInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// QuotaSpecBindings returns a QuotaSpecBindingInformer.
func (v *version) QuotaSpecBindings() QuotaSpecBindingInformer {
	return &quotaSpecBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Rules returns a RuleInformer.
func (v *version) Rules() RuleInformer {
	return &ruleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
