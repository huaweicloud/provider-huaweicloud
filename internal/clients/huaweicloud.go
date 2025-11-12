/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/huaweicloud/provider-huaweicloud/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal huaweicloud credentials as JSON"

	// provider config
	keyRegion               = "region"
	keyAccessKey            = "access_key"
	keySecretKey            = "secret_key"
	keySecurityToken        = "security_token"
	keyDomainID             = "domain_id"
	keyDomainName           = "domain_name"
	keyUserName             = "user_name"
	keyUserID               = "user_id"
	keyPassword             = "password"
	keyAssumeRole           = "assume_role"
	keyAssumeRoleWithOIDC   = "assume_role_with_oidc"
	keyProjectID            = "project_id"
	keyProjectName          = "project_name"
	keyTenantID             = "tenant_id"
	keyTenantName           = "tenant_name"
	keyToken                = "token"
	keyInsecure             = "insecure"
	keyCACertFile           = "cacert_file"
	keyCert                 = "cert"
	keyKey                  = "key"
	keyAgencyName           = "agency_name"
	keyAgencyDomainName     = "agency_domain_name"
	keyDelegatedProject     = "delegated_project"
	keyAuthURL              = "auth_url"
	keyCloud                = "cloud"
	keyEndpoints            = "endpoints"
	keyRegional             = "regional"
	keySharedConfigFile     = "shared_config_file"
	keyProfile              = "profile"
	keyEnterpriseProjectID  = "enterprise_project_id"
	keyMaxRetries           = "max_retries"
	keyEnableForceNew       = "enable_force_new"
	keySigningAlgorithm     = "signing_algorithm"
	keySkipCheckWebsiteType = "skip_check_website_type"
	keySkipCheckUpgrade     = "skip_check_upgrade"
	keyDefaultTags          = "default_tags"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]interface{}{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]interface{}{}
		if v, ok := creds[keyRegion]; ok {
			ps.Configuration[keyRegion] = v
		}
		if v, ok := creds[keyAccessKey]; ok {
			ps.Configuration[keyAccessKey] = v
		}
		if v, ok := creds[keySecretKey]; ok {
			ps.Configuration[keySecretKey] = v
		}
		if v, ok := creds[keySecurityToken]; ok {
			ps.Configuration[keySecurityToken] = v
		}
		if v, ok := creds[keyDomainID]; ok {
			ps.Configuration[keyDomainID] = v
		}
		if v, ok := creds[keyDomainName]; ok {
			ps.Configuration[keyDomainName] = v
		}
		if v, ok := creds[keyUserName]; ok {
			ps.Configuration[keyUserName] = v
		}
		if v, ok := creds[keyUserID]; ok {
			ps.Configuration[keyUserID] = v
		}
		if v, ok := creds[keyPassword]; ok {
			ps.Configuration[keyPassword] = v
		}
		if v, ok := creds[keyAssumeRole]; ok {
			ps.Configuration[keyAssumeRole] = v
		}
		if v, ok := creds[keyAssumeRoleWithOIDC]; ok {
			ps.Configuration[keyAssumeRoleWithOIDC] = v
		}
		if v, ok := creds[keyProjectID]; ok {
			ps.Configuration[keyProjectID] = v
		}
		if v, ok := creds[keyProjectName]; ok {
			ps.Configuration[keyProjectName] = v
		}
		if v, ok := creds[keyTenantID]; ok {
			ps.Configuration[keyTenantID] = v
		}
		if v, ok := creds[keyTenantName]; ok {
			ps.Configuration[keyTenantName] = v
		}
		if v, ok := creds[keyToken]; ok {
			ps.Configuration[keyToken] = v
		}
		if v, ok := creds[keyInsecure]; ok {
			ps.Configuration[keyInsecure] = v
		}
		if v, ok := creds[keyCACertFile]; ok {
			ps.Configuration[keyCACertFile] = v
		}
		if v, ok := creds[keyCert]; ok {
			ps.Configuration[keyCert] = v
		}
		if v, ok := creds[keyKey]; ok {
			ps.Configuration[keyKey] = v
		}
		if v, ok := creds[keyAgencyName]; ok {
			ps.Configuration[keyAgencyName] = v
		}
		if v, ok := creds[keyAgencyDomainName]; ok {
			ps.Configuration[keyAgencyDomainName] = v
		}
		if v, ok := creds[keyDelegatedProject]; ok {
			ps.Configuration[keyDelegatedProject] = v
		}
		if v, ok := creds[keyAuthURL]; ok {
			ps.Configuration[keyAuthURL] = v
		}
		if v, ok := creds[keyCloud]; ok {
			ps.Configuration[keyCloud] = v
		}
		if v, ok := creds[keyEndpoints]; ok {
			ps.Configuration[keyEndpoints] = v
		}
		if v, ok := creds[keyRegional]; ok {
			ps.Configuration[keyRegional] = v
		}
		if v, ok := creds[keySharedConfigFile]; ok {
			ps.Configuration[keySharedConfigFile] = v
		}
		if v, ok := creds[keyProfile]; ok {
			ps.Configuration[keyProfile] = v
		}
		if v, ok := creds[keyEnterpriseProjectID]; ok {
			ps.Configuration[keyEnterpriseProjectID] = v
		}
		if v, ok := creds[keyMaxRetries]; ok {
			ps.Configuration[keyMaxRetries] = v
		}
		if v, ok := creds[keyEnableForceNew]; ok {
			ps.Configuration[keyEnableForceNew] = v
		}
		if v, ok := creds[keySigningAlgorithm]; ok {
			ps.Configuration[keySigningAlgorithm] = v
		}
		if v, ok := creds[keySkipCheckWebsiteType]; ok {
			ps.Configuration[keySkipCheckWebsiteType] = v
		}
		if v, ok := creds[keySkipCheckUpgrade]; ok {
			ps.Configuration[keySkipCheckUpgrade] = v
		}
		if v, ok := creds[keyDefaultTags]; ok {
			ps.Configuration[keyDefaultTags] = v
		}
		return ps, nil
	}
}
