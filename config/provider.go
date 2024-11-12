/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/huaweicloud/provider-huaweicloud/config/cce"
	"github.com/huaweicloud/provider-huaweicloud/config/dcs"
	"github.com/huaweicloud/provider-huaweicloud/config/dms"
	"github.com/huaweicloud/provider-huaweicloud/config/ecs"
	"github.com/huaweicloud/provider-huaweicloud/config/eip"
	"github.com/huaweicloud/provider-huaweicloud/config/iam"
	"github.com/huaweicloud/provider-huaweicloud/config/obs"
	"github.com/huaweicloud/provider-huaweicloud/config/rds"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/huaweicloud/provider-huaweicloud/config/vpc"
)

const (
	resourcePrefix = "huaweicloud"
	modulePath     = "github.com/huaweicloud/provider-huaweicloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("huaweicloud.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		cce.Configure,
		dcs.Configure,
		dms.Configure,
		ecs.Configure,
		iam.Configure,
		obs.Configure,
		vpc.Configure,
		eip.Configure,
		rds.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
