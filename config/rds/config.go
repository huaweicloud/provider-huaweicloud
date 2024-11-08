package rds

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupRds = "rds"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_rds_backup", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_cross_region_backup_strategy", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_database_logs_shrinking", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_extend_log_link", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

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

		r.References["param_group_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_parametergroup",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"restore", "period"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_instance_eip_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["public_ip"] = config.Reference{
			TerraformName: "huaweicloud_vpc_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("address",true)`,
		}

		r.References["public_ip_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_eip",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_lts_log", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_mysql_account", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_mysql_binlog", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_mysql_database", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_mysql_database_privilege", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_mysql_database_table_restore", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"restore_tables", "databases"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_parametergroup", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds
	})

	p.AddResourceConfigurator("huaweicloud_rds_pg_account", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_pg_account_privileges", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_pg_account_roles", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_pg_database", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_pg_database_privilege", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_pg_hba", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_pg_plugin", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_pg_plugin_parameter", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_pg_plugin_update", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_pg_sql_limit", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"query_id", "query_string"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_primary_standby_switch", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_read_replica_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["primary_instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["security_group_id"] = config.Reference{
			TerraformName: "huaweicloud_networking_secgroup",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_restore", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["target_instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["source_instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"restore_time", "backup_id"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_sql_audit", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_sqlserver_account", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_sqlserver_database", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_rds_sqlserver_database_privilege", func(r *config.Resource) {
		r.ShortGroup = shortGroupRds

		r.References["instance_id"] = config.Reference{
			TerraformName: "huaweicloud_rds_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
}
