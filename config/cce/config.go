package cce

import "github.com/crossplane/upjet/pkg/config"

const shortGroupCce = "cce"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_cce_addon", func(r *config.Resource) {
		r.ShortGroup = shortGroupCce

		r.References["cluster_id"] = config.Reference{
			TerraformName: "huaweicloud_cce_cluster",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"values.basic", "values.basic_json", "values.custom", "values.custom_json", "values.flavor", "values.flavor_json"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_cce_chart", func(r *config.Resource) {
		r.ShortGroup = shortGroupCce
	})

	p.AddResourceConfigurator("huaweicloud_cce_cluster", func(r *config.Resource) {
		r.ShortGroup = shortGroupCce

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

		r.References["eip"] = config.Reference{
			TerraformName: "huaweicloud_vpc_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("address",true)`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"delete_efs", "delete_eni", "delete_evs", "delete_net", "delete_obs",
				"delete_sfs", "delete_all", "multi_az", "masters", "extend_params"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_cce_cluster_log_config", func(r *config.Resource) {
		r.ShortGroup = shortGroupCce

		r.References["cluster_id"] = config.Reference{
			TerraformName: "huaweicloud_cce_cluster",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_cce_cluster_upgrade", func(r *config.Resource) {
		r.ShortGroup = shortGroupCce

		r.References["cluster_id"] = config.Reference{
			TerraformName: "huaweicloud_cce_cluster",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_cce_namespace", func(r *config.Resource) {
		r.ShortGroup = shortGroupCce

		r.References["cluster_id"] = config.Reference{
			TerraformName: "huaweicloud_cce_cluster",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"name", "prefix"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_cce_node", func(r *config.Resource) {
		r.ShortGroup = shortGroupCce

		r.References["cluster_id"] = config.Reference{
			TerraformName: "huaweicloud_cce_cluster",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["subnet_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["eip_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["ecs_group_id"] = config.Reference{
			TerraformName: "huaweicloud_compute_servergroup",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"eip_id", "iptype", "bandwidth_charge_mode", "bandwidth_size",
				"sharetype", "extend_params", "max_pods", "public_key", "preinstall", "postinstall", "extend_param",
				"billing_mode", "order_id", "product_id", "ecs_performance_type", "password", "key_pair"},
			IgnoredFields: []string{"eip_ids"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_cce_node_attach", func(r *config.Resource) {
		r.ShortGroup = shortGroupCce

		r.References["cluster_id"] = config.Reference{
			TerraformName: "huaweicloud_cce_cluster",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["server_id"] = config.Reference{
			TerraformName: "huaweicloud_compute_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"lvm_config", "storage", "password", "key_pair"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_cce_node_pool", func(r *config.Resource) {
		r.ShortGroup = shortGroupCce

		r.References["cluster_id"] = config.Reference{
			TerraformName: "huaweicloud_cce_cluster",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["subnet_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["ecs_group_id"] = config.Reference{
			TerraformName: "huaweicloud_compute_servergroup",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"extend_params", "max_pods", "preinstall", "postinstall",
				"extend_param", "password", "key_pair"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_cce_node_pool_nodes_add", func(r *config.Resource) {
		r.ShortGroup = shortGroupCce

		r.References["cluster_id"] = config.Reference{
			TerraformName: "huaweicloud_cce_cluster",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["nodepool_id"] = config.Reference{
			TerraformName: "huaweicloud_cce_node_pool",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_cce_pvc", func(r *config.Resource) {
		r.ShortGroup = shortGroupCce

		r.References["cluster_id"] = config.Reference{
			TerraformName: "huaweicloud_cce_cluster",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
}
