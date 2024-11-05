package vpc

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupVpc = "vpc"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_vpc", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.TerraformResource.Schema["name"].Sensitive = true

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return map[string][]byte{
				xpv1.ResourceCredentialsSecretPasswordKey: []byte(attr["cidr"].(string)),
			}, nil
		}
	})

	p.AddResourceConfigurator("huaweicloud_vpc_subnet", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc

		r.References["vpc_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_networking_secgroup", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
	})

	p.AddResourceConfigurator("huaweicloud_networking_secgroup_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc

		r.References["security_group_id "] = config.Reference{
			TerraformName: "huaweicloud_networking_secgroup",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_networking_vip", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc

		r.References["network_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_networking_vip_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc

		r.References["vip_id"] = config.Reference{
			TerraformName: "huaweicloud_networking_vip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_vpc_address_group", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
	})

	p.AddResourceConfigurator("huaweicloud_vpc_flow_log", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
	})

	p.AddResourceConfigurator("huaweicloud_vpc_network_acl", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
	})

	p.AddResourceConfigurator("huaweicloud_vpc_network_interface", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc

		r.References["subnet_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_vpc_peering_connection", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc

		r.References["vpc_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["peer_vpc_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_vpc_peering_connection_accepter", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc

		r.References["vpc_peering_connection_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_peering_connection",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_vpc_route", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc

		r.References["vpc_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["route_table_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_route_table",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_vpc_route_table", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc

		r.References["vpc_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_vpc_sub_network_interface", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc

		r.References["subnet_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_vpc_traffic_mirror_filter", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
	})

	p.AddResourceConfigurator("huaweicloud_vpc_traffic_mirror_filter_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc

		r.References["traffic_mirror_filter_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_traffic_mirror_filter",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_vpc_traffic_mirror_session", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc

		r.References["traffic_mirror_filter_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_traffic_mirror_filter",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
}
