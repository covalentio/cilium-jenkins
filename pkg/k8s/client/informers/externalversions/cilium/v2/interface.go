// Copyright 2017 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was automatically generated by informer-gen

package v2

import (
	internalinterfaces "github.com/cilium/cilium/pkg/k8s/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CiliumNetworkPolicies returns a CiliumNetworkPolicyInformer.
	CiliumNetworkPolicies() CiliumNetworkPolicyInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// CiliumNetworkPolicies returns a CiliumNetworkPolicyInformer.
func (v *version) CiliumNetworkPolicies() CiliumNetworkPolicyInformer {
	return &ciliumNetworkPolicyInformer{factory: v.SharedInformerFactory}
}
