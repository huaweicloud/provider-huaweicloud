package eip

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupEip = "eip"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_global_eip", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
	})
	p.AddResourceConfigurator("huaweicloud_global_eip_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
		r.References["global_eip_id"] = config.Reference{
			TerraformName: "huaweicloud_global_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
	p.AddResourceConfigurator("huaweicloud_global_internet_bandwidth", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
	})
	p.AddResourceConfigurator("huaweicloud_vpc_bandwidth", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
	})
	p.AddResourceConfigurator("huaweicloud_vpc_bandwidth_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
		r.ShortGroup = shortGroupEip
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
	})
	p.AddResourceConfigurator("huaweicloud_vpc_eip_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
	})
	p.AddResourceConfigurator("huaweicloud_vpc_eipv3_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
		r.References["publicip_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
	p.AddResourceConfigurator("huaweicloud_vpc_internet_gateway", func(r *config.Resource) {
		r.ShortGroup = shortGroupEip
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
