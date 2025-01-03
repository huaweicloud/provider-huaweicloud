package ecs

import "github.com/crossplane/upjet/pkg/config"

const shortGroupEcs = "ecs"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_compute_auto_launch_group", func(r *config.Resource) {
		r.ShortGroup = shortGroupEcs
	})

	p.AddResourceConfigurator("huaweicloud_compute_eip_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroupEcs
		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_compute_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"bandwidth_id", "public_ip"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_compute_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupEcs
		r.LateInitializer = config.LateInitializer{
			IgnoredFields:            []string{"security_groups"},
			ConditionalIgnoredFields: []string{"eip_id", "eip_type", "bandwidth", "bandwidth.size", "bandwidth.charge_mode", "bandwidth.id"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_compute_interface_attach", func(r *config.Resource) {
		r.ShortGroup = shortGroupEcs
		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_compute_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"port_id", "network_id"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_compute_servergroup", func(r *config.Resource) {
		r.ShortGroup = shortGroupEcs
	})

	p.AddResourceConfigurator("huaweicloud_compute_volume_attach", func(r *config.Resource) {
		r.ShortGroup = shortGroupEcs
		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_compute_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
}
