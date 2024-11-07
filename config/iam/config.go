package iam

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupObs = "iam"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_identity_access_key", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
	})
	p.AddResourceConfigurator("huaweicloud_identity_acl", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
	})
	p.AddResourceConfigurator("huaweicloud_identity_agency", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"delegated_domain_name", "delegated_service_name"},
		}
	})
	p.AddResourceConfigurator("huaweicloud_identity_group", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
	})
	p.AddResourceConfigurator("huaweicloud_identity_group_membership", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
	})
	p.AddResourceConfigurator("huaweicloud_identity_group_role_assignment", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
		r.References["group_id"] = config.Reference{
			TerraformName: "huaweicloud_identity_group",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
		r.References["role_id"] = config.Reference{
			TerraformName: "huaweicloud_identity_role",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"domain_id", "project_id", "enterprise_project_id"},
		}
	})
	p.AddResourceConfigurator("huaweicloud_identity_login_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
	})
	p.AddResourceConfigurator("huaweicloud_identity_password_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
	})
	p.AddResourceConfigurator("huaweicloud_identity_project", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
	})
	p.AddResourceConfigurator("huaweicloud_identity_protection_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
	})
	p.AddResourceConfigurator("huaweicloud_identity_provider", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
	})
	p.AddResourceConfigurator("huaweicloud_identity_provider_conversion", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
		r.References["provider_id "] = config.Reference{
			TerraformName: "huaweicloud_identity_provider",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
	p.AddResourceConfigurator("huaweicloud_identity_role", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
	})
	p.AddResourceConfigurator("huaweicloud_identity_role_assignment", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
		r.References["group_id"] = config.Reference{
			TerraformName: "huaweicloud_identity_group",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
		r.References["role_id"] = config.Reference{
			TerraformName: "huaweicloud_identity_role",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
	p.AddResourceConfigurator("huaweicloud_identity_user", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
	})
	p.AddResourceConfigurator("huaweicloud_identity_user_role_assignment", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
		r.References["user_id"] = config.Reference{
			TerraformName: "huaweicloud_identity_user",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
		r.References["role_id"] = config.Reference{
			TerraformName: "huaweicloud_identity_role",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
	p.AddResourceConfigurator("huaweicloud_identity_user_token", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
	})
	p.AddResourceConfigurator("huaweicloud_identity_virtual_mfa_device", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
		r.References["user_id"] = config.Reference{
			TerraformName: "huaweicloud_identity_user",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
}
