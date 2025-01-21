package cc

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupCc = "cc"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_cc_authorization", func(r *config.Resource) {
		r.ShortGroup = shortGroupCc

		r.References["connection_id"] = config.Reference{
			TerraformName: "huaweicloud_cc_connection",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_cc_bandwidth_package", func(r *config.Resource) {
		r.ShortGroup = shortGroupCc
	})

	p.AddResourceConfigurator("huaweicloud_cc_central_network", func(r *config.Resource) {
		r.ShortGroup = shortGroupCc
	})

	p.AddResourceConfigurator("huaweicloud_cc_central_network_attachment", func(r *config.Resource) {
		r.ShortGroup = shortGroupCc

		r.References["central_network_id"] = config.Reference{
			TerraformName: "huaweicloud_cc_central_network",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_cc_central_network_connection_bandwidth_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroupCc

		r.References["central_network_id"] = config.Reference{
			TerraformName: "huaweicloud_cc_central_network",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["connection_id"] = config.Reference{
			TerraformName: "huaweicloud_cc_connection",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["global_connection_bandwidth_id"] = config.Reference{
			TerraformName: "huaweicloud_cc_global_connection_bandwidth",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_cc_central_network_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroupCc

		r.References["central_network_id"] = config.Reference{
			TerraformName: "huaweicloud_cc_central_network",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_cc_central_network_policy_apply", func(r *config.Resource) {
		r.ShortGroup = shortGroupCc

		r.References["central_network_id"] = config.Reference{
			TerraformName: "huaweicloud_cc_central_network",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["policy_id"] = config.Reference{
			TerraformName: "huaweicloud_cc_central_network_policy",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_cc_connection", func(r *config.Resource) {
		r.ShortGroup = shortGroupCc
	})

	p.AddResourceConfigurator("huaweicloud_cc_global_connection_bandwidth", func(r *config.Resource) {
		r.ShortGroup = shortGroupCc
	})

	p.AddResourceConfigurator("huaweicloud_cc_global_connection_bandwidth_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroupCc

		r.References["gcb_id"] = config.Reference{
			TerraformName: "huaweicloud_cc_global_connection_bandwidth",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_cc_inter_region_bandwidth", func(r *config.Resource) {
		r.ShortGroup = shortGroupCc

		r.References["connection_id"] = config.Reference{
			TerraformName: "huaweicloud_cc_connection",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["bandwidth_package_id"] = config.Reference{
			TerraformName: "huaweicloud_cc_bandwidth_package",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_cc_network_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupCc

		r.References["connection_id"] = config.Reference{
			TerraformName: "huaweicloud_cc_connection",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
}
