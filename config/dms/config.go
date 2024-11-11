package dms

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupDms = "dms"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_dms_rabbitmq_exchange", func(r *config.Resource) {
		r.ShortGroup = shortGroupDms

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_rabbitmq_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_rabbitmq_exchange_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroupDms

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_rabbitmq_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_rabbitmq_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupDms

		r.References["vpc_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["network_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["security_group_id"] = config.Reference{
			TerraformName: "huaweicloud_networking_secgroup",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["public_ip_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"available_zones", "product_id"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_rabbitmq_plugin", func(r *config.Resource) {
		r.ShortGroup = shortGroupDms

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_rabbitmq_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_rabbitmq_queue", func(r *config.Resource) {
		r.ShortGroup = shortGroupDms

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_rabbitmq_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_rabbitmq_queue_message_clear", func(r *config.Resource) {
		r.ShortGroup = shortGroupDms

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_rabbitmq_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_rabbitmq_user", func(r *config.Resource) {
		r.ShortGroup = shortGroupDms

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_rabbitmq_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_rabbitmq_vhost", func(r *config.Resource) {
		r.ShortGroup = shortGroupDms

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_rabbitmq_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_rocketmq_consumer_group", func(r *config.Resource) {
		r.ShortGroup = shortGroupDms

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_rocketmq_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_rocketmq_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupDms

		r.References["vpc_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["subnet_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["security_group_id"] = config.Reference{
			TerraformName: "huaweicloud_networking_secgroup",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["publicip_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_rocketmq_migration_task", func(r *config.Resource) {
		r.ShortGroup = shortGroupDms

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_rocketmq_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_rocketmq_topic", func(r *config.Resource) {
		r.ShortGroup = shortGroupDms

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_rocketmq_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_rocketmq_user", func(r *config.Resource) {
		r.ShortGroup = shortGroupDms

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_rocketmq_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
}
