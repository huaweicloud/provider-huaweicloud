package kafka

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupKafka = "kafka"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_dms_kafka_background_task_delete", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_consumer_group_topic_batch_delete", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_consumer_group", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_instance_batch_action", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_instance_rebalance_log", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_instance_restart", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

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
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_message_diagnosis_task", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_message_offset_reset", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_message_produce", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_partition_reassign", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_permissions", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_smart_connect_task_action", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["task_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_smart_connect_task",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_smart_connect_task", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["connector_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("connector_id",true)`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_smart_connect", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_smart_connector_validate", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_topic_quota", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_topic", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_user_client_quota", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_user_password_reset", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafka_user", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_dms_kafkav2_smart_connect_task", func(r *config.Resource) {
		r.ShortGroup = shortGroupKafka

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_dms_kafka_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
}
