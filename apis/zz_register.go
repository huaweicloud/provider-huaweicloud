// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/huaweicloud/provider-huaweicloud/apis/cce/v1alpha1"
	v1alpha1dcs "github.com/huaweicloud/provider-huaweicloud/apis/dcs/v1alpha1"
	v1alpha1dms "github.com/huaweicloud/provider-huaweicloud/apis/dms/v1alpha1"
	v1alpha1ecs "github.com/huaweicloud/provider-huaweicloud/apis/ecs/v1alpha1"
	v1alpha1eip "github.com/huaweicloud/provider-huaweicloud/apis/eip/v1alpha1"
	v1alpha1evs "github.com/huaweicloud/provider-huaweicloud/apis/evs/v1alpha1"
	v1alpha1iam "github.com/huaweicloud/provider-huaweicloud/apis/iam/v1alpha1"
	v1alpha1modelarts "github.com/huaweicloud/provider-huaweicloud/apis/modelarts/v1alpha1"
	v1alpha1nat "github.com/huaweicloud/provider-huaweicloud/apis/nat/v1alpha1"
	v1alpha1obs "github.com/huaweicloud/provider-huaweicloud/apis/obs/v1alpha1"
	v1alpha1rds "github.com/huaweicloud/provider-huaweicloud/apis/rds/v1alpha1"
	v1alpha1swr "github.com/huaweicloud/provider-huaweicloud/apis/swr/v1alpha1"
	v1alpha1apis "github.com/huaweicloud/provider-huaweicloud/apis/v1alpha1"
	v1beta1 "github.com/huaweicloud/provider-huaweicloud/apis/v1beta1"
	v1alpha1vpc "github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1dcs.SchemeBuilder.AddToScheme,
		v1alpha1dms.SchemeBuilder.AddToScheme,
		v1alpha1ecs.SchemeBuilder.AddToScheme,
		v1alpha1eip.SchemeBuilder.AddToScheme,
		v1alpha1evs.SchemeBuilder.AddToScheme,
		v1alpha1iam.SchemeBuilder.AddToScheme,
		v1alpha1modelarts.SchemeBuilder.AddToScheme,
		v1alpha1nat.SchemeBuilder.AddToScheme,
		v1alpha1obs.SchemeBuilder.AddToScheme,
		v1alpha1rds.SchemeBuilder.AddToScheme,
		v1alpha1swr.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1vpc.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
