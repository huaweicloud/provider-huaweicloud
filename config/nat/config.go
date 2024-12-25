package nat

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupNat = "nat"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_nat_dnat_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupNat

		r.References["nat_gateway_id"] = config.Reference{
			TerraformName: "huaweicloud_nat_gateway",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["floating_ip_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["global_eip_id"] = config.Reference{
			TerraformName: "huaweicloud_global_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"floating_ip_id", "global_eip_id", "internal_service_port", "internal_service_port_range",
				"external_service_port", "external_service_port_range", "port_id", "private_ip"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_nat_gateway", func(r *config.Resource) {
		r.ShortGroup = shortGroupNat

		r.References["vpc_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["subnet_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_nat_private_dnat_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupNat

		r.References["gateway_id"] = config.Reference{
			TerraformName: "huaweicloud_nat_private_gateway",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["transit_ip_id"] = config.Reference{
			TerraformName: "huaweicloud_nat_private_transit_ip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"backend_interface_id", "backend_private_ip"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_nat_private_gateway", func(r *config.Resource) {
		r.ShortGroup = shortGroupNat

		r.References["subnet_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_nat_private_snat_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupNat

		r.References["gateway_id"] = config.Reference{
			TerraformName: "huaweicloud_nat_private_gateway",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["subnet_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["transit_ip_id"] = config.Reference{
			TerraformName: "huaweicloud_nat_private_transit_ip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"cidr", "subnet_id"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_nat_private_transit_ip", func(r *config.Resource) {
		r.ShortGroup = shortGroupNat

		r.References["subnet_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_nat_snat_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupNat

		r.References["nat_gateway_id"] = config.Reference{
			TerraformName: "huaweicloud_nat_gateway",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["floating_ip_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["global_eip_id"] = config.Reference{
			TerraformName: "huaweicloud_global_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["subnet_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			IgnoredFields:            []string{"network_id"},
			ConditionalIgnoredFields: []string{"floating_ip_id", "global_eip_id", "subnet_id", "cidr"},
		}
	})
}
