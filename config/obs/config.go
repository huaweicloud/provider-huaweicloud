package obs

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupObs = "obs"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_obs_bucket", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
	})

	p.AddResourceConfigurator("huaweicloud_obs_bucket_acl", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs

		r.References["bucket"] = config.Reference{
			TerraformName: "huaweicloud_obs_bucket",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_obs_bucket_object", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs

		r.References["bucket"] = config.Reference{
			TerraformName: "huaweicloud_obs_bucket",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"source", "content", "website.index_document", "website.error_document", "website.routing_rules", "website.redirect_all_requests_to"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_obs_bucket_object_acl", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs

		r.References["bucket"] = config.Reference{
			TerraformName: "huaweicloud_obs_bucket",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)`,
		}
		r.References["key"] = config.Reference{
			TerraformName: "huaweicloud_obs_bucket_object",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("key",true)`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_obs_bucket_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs

		r.References["bucket"] = config.Reference{
			TerraformName: "huaweicloud_obs_bucket",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_obs_bucket_replication", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs

		r.References["bucket"] = config.Reference{
			TerraformName: "huaweicloud_obs_bucket",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)`,
		}

		r.References["destination_bucket"] = config.Reference{
			TerraformName: "huaweicloud_obs_bucket",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)`,
		}
	})

}
