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
}
