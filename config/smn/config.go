package smn

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupSmn = "smn"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_smn_logtank", func(r *config.Resource) {
		r.ShortGroup = shortGroupSmn

		r.References["topic_urn"] = config.Reference{
			TerraformName: "huaweicloud_smn_topic",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_smn_message_detection", func(r *config.Resource) {
		r.ShortGroup = shortGroupSmn

		r.References["topic_urn"] = config.Reference{
			TerraformName: "huaweicloud_smn_topic",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_smn_message_publish", func(r *config.Resource) {
		r.ShortGroup = shortGroupSmn

		r.References["topic_urn"] = config.Reference{
			TerraformName: "huaweicloud_smn_topic",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"message", "message_structure", "message_template_name"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_smn_message_template", func(r *config.Resource) {
		r.ShortGroup = shortGroupSmn
	})

	p.AddResourceConfigurator("huaweicloud_smn_subscription", func(r *config.Resource) {
		r.ShortGroup = shortGroupSmn

		r.References["topic_urn"] = config.Reference{
			TerraformName: "huaweicloud_smn_topic",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_smn_subscription_filter_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroupSmn

		r.References["subscription_urn"] = config.Reference{
			TerraformName: "huaweicloud_smn_subscription",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_smn_topic", func(r *config.Resource) {
		r.ShortGroup = shortGroupSmn

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"users_publish_allowed", "services_publish_allowed", "access_policy"},
		}
	})
}
