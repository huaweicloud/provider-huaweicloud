package dcs

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupDcs = "dcs"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_dcs_account", func(r *config.Resource) {
		r.ShortGroup = shortGroupDcs

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dcs_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dcs_backup", func(r *config.Resource) {
		r.ShortGroup = shortGroupDcs

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dcs_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dcs_bigkey_analysis", func(r *config.Resource) {
		r.ShortGroup = shortGroupDcs

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dcs_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dcs_custom_template", func(r *config.Resource) {
		r.ShortGroup = shortGroupDcs
	})

	p.AddResourceConfigurator("huaweicloud_dcs_diagnosis_task", func(r *config.Resource) {
		r.ShortGroup = shortGroupDcs

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dcs_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dcs_hotkey_analysis", func(r *config.Resource) {
		r.ShortGroup = shortGroupDcs

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dcs_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dcs_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupDcs

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

	p.AddResourceConfigurator("huaweicloud_dcs_instance_restore", func(r *config.Resource) {
		r.ShortGroup = shortGroupDcs

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dcs_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["backup_id"] = config.Reference{
			TerraformName: "huaweicloud_dcs_backup",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
}
