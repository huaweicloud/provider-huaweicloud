package eip

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupEip = "eip"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_global_eip", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
		r.Kind = "GlobalEip"
	})
	p.AddResourceConfigurator("huaweicloud_global_eip_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
		r.Kind = "GlobalEipAssociate"
		r.References["global_eip_id"] = config.Reference{
			TerraformName: "huaweicloud_global_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
	p.AddResourceConfigurator("huaweicloud_global_internet_bandwidth", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
		r.Kind = "GlobalInternetBandwidth"
	})
	p.AddResourceConfigurator("huaweicloud_vpc_bandwidth", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
		r.Kind = "VpcBandwidth"
	})
	p.AddResourceConfigurator("huaweicloud_vpc_bandwidth_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
		r.Kind = "VpcBandwidthAssociate"
		r.References["bandwidth_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_bandwidth",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
		r.References["eip_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
	p.AddResourceConfigurator("huaweicloud_vpc_eip", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
		r.Kind = "VpcEip"
		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"publicip.ip_address", "period_unit", "period", "auto_renew", "auto_pay",
				"bandwidth.name", "bandwidth.id"},
		}
	})
	p.AddResourceConfigurator("huaweicloud_vpc_eip_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
		r.Kind = "VpcEipAssociate"
		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"fixed_ip", "network_id", "port_id"},
		}
	})
	p.AddResourceConfigurator("huaweicloud_vpc_eipv3_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
		r.Kind = "VpcEipV3Associate"
		r.References["publicip_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
	p.AddResourceConfigurator("huaweicloud_vpc_internet_gateway", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
		r.Kind = "VpcInternetGateway"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
}
