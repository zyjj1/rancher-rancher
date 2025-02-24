/*
Copyright 2025 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1

import (
	v1 "github.com/rancher/rancher/pkg/apis/ext.cattle.io/v1"
	"github.com/rancher/wrangler/v3/pkg/generic"
)

// TokenPrincipalController interface for managing TokenPrincipal resources.
type TokenPrincipalController interface {
	generic.ControllerInterface[*v1.TokenPrincipal, *v1.TokenPrincipalList]
}

// TokenPrincipalClient interface for managing TokenPrincipal resources in Kubernetes.
type TokenPrincipalClient interface {
	generic.ClientInterface[*v1.TokenPrincipal, *v1.TokenPrincipalList]
}

// TokenPrincipalCache interface for retrieving TokenPrincipal resources in memory.
type TokenPrincipalCache interface {
	generic.CacheInterface[*v1.TokenPrincipal]
}
