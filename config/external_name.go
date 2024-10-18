/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	// cce
	"huaweicloud_cce_cluster": config.IdentifierFromProvider,
	"huaweicloud_cce_node":    config.IdentifierFromProvider,

	// iam
	"huaweicloud_identity_access_key":            config.IdentifierFromProvider,
	"huaweicloud_identity_acl":                   config.IdentifierFromProvider,
	"huaweicloud_identity_agency":                config.IdentifierFromProvider,
	"huaweicloud_identity_group":                 config.IdentifierFromProvider,
	"huaweicloud_identity_group_membership":      config.IdentifierFromProvider,
	"huaweicloud_identity_group_role_assignment": config.IdentifierFromProvider,
	"huaweicloud_identity_login_policy":          config.IdentifierFromProvider,
	"huaweicloud_identity_password_policy":       config.IdentifierFromProvider,
	"huaweicloud_identity_project":               config.IdentifierFromProvider,
	"huaweicloud_identity_protection_policy":     config.IdentifierFromProvider,
	"huaweicloud_identity_provider":              config.IdentifierFromProvider,
	"huaweicloud_identity_provider_conversion":   config.IdentifierFromProvider,
	"huaweicloud_identity_role":                  config.IdentifierFromProvider,
	"huaweicloud_identity_role_assignment":       config.IdentifierFromProvider,
	"huaweicloud_identity_user":                  config.IdentifierFromProvider,
	"huaweicloud_identity_user_role_assignment":  config.IdentifierFromProvider,
	"huaweicloud_identity_user_token":            config.IdentifierFromProvider,
	"huaweicloud_identity_virtual_mfa_device":    config.IdentifierFromProvider,

	// obs
	"huaweicloud_obs_bucket":            config.IdentifierFromProvider,
	"huaweicloud_obs_bucket_acl":        config.IdentifierFromProvider,
	"huaweicloud_obs_bucket_object":     config.IdentifierFromProvider,
	"huaweicloud_obs_bucket_object_acl": config.IdentifierFromProvider,

	// vpc
	"huaweicloud_vpc":                      config.IdentifierFromProvider,
	"huaweicloud_vpc_subnet":               config.IdentifierFromProvider,
	"huaweicloud_networking_secgroup":      config.IdentifierFromProvider,
	"huaweicloud_networking_secgroup_rule": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
