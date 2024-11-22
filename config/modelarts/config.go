package modelarts

import "github.com/crossplane/upjet/pkg/config"

const shortGroupModelArts = "modelarts"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_modelarts_authorization", func(r *config.Resource) {
		r.ShortGroup = shortGroupModelArts

		r.References["user_id"] = config.Reference{
			TerraformName: "huaweicloud_identity_user",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_modelarts_dataset", func(r *config.Resource) {
		r.ShortGroup = shortGroupModelArts
	})

	p.AddResourceConfigurator("huaweicloud_modelarts_dataset_version", func(r *config.Resource) {
		r.ShortGroup = shortGroupModelArts

		r.References["dataset_id"] = config.Reference{
			TerraformName: "huaweicloud_modelarts_dataset",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_modelarts_model", func(r *config.Resource) {
		r.ShortGroup = shortGroupModelArts
	})

	p.AddResourceConfigurator("huaweicloud_modelarts_network", func(r *config.Resource) {
		r.ShortGroup = shortGroupModelArts

		r.References["workspace_id"] = config.Reference{
			TerraformName: "huaweicloud_modelarts_workspace",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_modelarts_notebook", func(r *config.Resource) {
		r.ShortGroup = shortGroupModelArts

		r.References["workspace_id"] = config.Reference{
			TerraformName: "huaweicloud_modelarts_workspace",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["pool_id"] = config.Reference{
			TerraformName: "huaweicloud_modelarts_resource_pool",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_modelarts_notebook_mount_storage", func(r *config.Resource) {
		r.ShortGroup = shortGroupModelArts

		r.References["notebook_id"] = config.Reference{
			TerraformName: "huaweicloud_modelarts_notebook",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_modelarts_resource_pool", func(r *config.Resource) {
		r.ShortGroup = shortGroupModelArts

		r.References["workspace_id"] = config.Reference{
			TerraformName: "huaweicloud_modelarts_workspace",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["network_id"] = config.Reference{
			TerraformName: "huaweicloud_modelarts_network",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["vpc_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["subnet_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"network_id", "vpc_id"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_modelarts_service", func(r *config.Resource) {
		r.ShortGroup = shortGroupModelArts

		r.References["workspace_id"] = config.Reference{
			TerraformName: "huaweicloud_modelarts_workspace",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["vpc_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["subnet_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["security_group_id"] = config.Reference{
			TerraformName: "huaweicloud_networking_secgroup",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_modelarts_workspace", func(r *config.Resource) {
		r.ShortGroup = shortGroupModelArts
	})
}
