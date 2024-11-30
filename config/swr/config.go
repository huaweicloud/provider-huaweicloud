package swr

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupSwr = "swr"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_swr_image_auto_sync", func(r *config.Resource) {
		r.ShortGroup = shortGroupSwr

		r.References["organization"] = config.Reference{
			TerraformName: "huaweicloud_swr_organization",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)`,
		}

		r.References["repository"] = config.Reference{
			TerraformName: "huaweicloud_swr_repository",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_swr_image_permissions", func(r *config.Resource) {
		r.ShortGroup = shortGroupSwr

		r.References["users.user_id"] = config.Reference{
			TerraformName: "huaweicloud_identity_user",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_swr_image_retention_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroupSwr
	})

	p.AddResourceConfigurator("huaweicloud_swr_image_trigger", func(r *config.Resource) {
		r.ShortGroup = shortGroupSwr

		r.References["cluster_id"] = config.Reference{
			TerraformName: "huaweicloud_cce_cluster",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_swr_organization", func(r *config.Resource) {
		r.ShortGroup = shortGroupSwr
	})

	p.AddResourceConfigurator("huaweicloud_swr_organization_permissions", func(r *config.Resource) {
		r.ShortGroup = shortGroupSwr

		r.References["users.user_id"] = config.Reference{
			TerraformName: "huaweicloud_identity_user",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_swr_repository", func(r *config.Resource) {
		r.ShortGroup = shortGroupSwr
	})

	p.AddResourceConfigurator("huaweicloud_swr_repository_sharing", func(r *config.Resource) {
		r.ShortGroup = shortGroupSwr
	})
}
