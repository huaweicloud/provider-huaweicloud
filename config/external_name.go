/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	// cc
	"huaweicloud_cc_authorization":                                  config.IdentifierFromProvider,
	"huaweicloud_cc_bandwidth_package":                              config.IdentifierFromProvider,
	"huaweicloud_cc_central_network":                                config.IdentifierFromProvider,
	"huaweicloud_cc_central_network_attachment":                     config.IdentifierFromProvider,
	"huaweicloud_cc_central_network_connection_bandwidth_associate": config.IdentifierFromProvider,
	"huaweicloud_cc_central_network_policy":                         config.IdentifierFromProvider,
	"huaweicloud_cc_central_network_policy_apply":                   config.IdentifierFromProvider,
	"huaweicloud_cc_connection":                                     config.IdentifierFromProvider,
	"huaweicloud_cc_global_connection_bandwidth":                    config.IdentifierFromProvider,
	"huaweicloud_cc_global_connection_bandwidth_associate":          config.IdentifierFromProvider,
	"huaweicloud_cc_inter_region_bandwidth":                         config.IdentifierFromProvider,
	"huaweicloud_cc_network_instance":                               config.IdentifierFromProvider,

	// cce
	"huaweicloud_cce_addon":               config.IdentifierFromProvider,
	"huaweicloud_cce_chart":               config.IdentifierFromProvider,
	"huaweicloud_cce_cluster":             config.IdentifierFromProvider,
	"huaweicloud_cce_cluster_log_config":  config.IdentifierFromProvider,
	"huaweicloud_cce_cluster_upgrade":     config.IdentifierFromProvider,
	"huaweicloud_cce_namespace":           config.IdentifierFromProvider,
	"huaweicloud_cce_node":                config.IdentifierFromProvider,
	"huaweicloud_cce_node_attach":         config.IdentifierFromProvider,
	"huaweicloud_cce_node_pool":           config.IdentifierFromProvider,
	"huaweicloud_cce_node_pool_nodes_add": config.IdentifierFromProvider,
	"huaweicloud_cce_pvc":                 config.IdentifierFromProvider,

	// dcs
	"huaweicloud_dcs_account":          config.IdentifierFromProvider,
	"huaweicloud_dcs_backup":           config.IdentifierFromProvider,
	"huaweicloud_dcs_bigkey_analysis":  config.IdentifierFromProvider,
	"huaweicloud_dcs_custom_template":  config.IdentifierFromProvider,
	"huaweicloud_dcs_diagnosis_task":   config.IdentifierFromProvider,
	"huaweicloud_dcs_hotkey_analysis":  config.IdentifierFromProvider,
	"huaweicloud_dcs_instance":         config.IdentifierFromProvider,
	"huaweicloud_dcs_instance_restore": config.IdentifierFromProvider,

	// dms
	"huaweicloud_dms_rabbitmq_exchange":            config.IdentifierFromProvider,
	"huaweicloud_dms_rabbitmq_exchange_associate":  config.IdentifierFromProvider,
	"huaweicloud_dms_rabbitmq_instance":            config.IdentifierFromProvider,
	"huaweicloud_dms_rabbitmq_plugin":              config.IdentifierFromProvider,
	"huaweicloud_dms_rabbitmq_queue":               config.IdentifierFromProvider,
	"huaweicloud_dms_rabbitmq_queue_message_clear": config.IdentifierFromProvider,
	"huaweicloud_dms_rabbitmq_user":                config.IdentifierFromProvider,
	"huaweicloud_dms_rabbitmq_vhost":               config.IdentifierFromProvider,
	"huaweicloud_dms_rocketmq_consumer_group":      config.IdentifierFromProvider,
	"huaweicloud_dms_rocketmq_instance":            config.IdentifierFromProvider,
	"huaweicloud_dms_rocketmq_migration_task":      config.IdentifierFromProvider,
	"huaweicloud_dms_rocketmq_topic":               config.IdentifierFromProvider,
	"huaweicloud_dms_rocketmq_user":                config.IdentifierFromProvider,

	// ecs
	"huaweicloud_compute_auto_launch_group": config.IdentifierFromProvider,
	"huaweicloud_compute_eip_associate":     config.IdentifierFromProvider,
	"huaweicloud_compute_instance":          config.IdentifierFromProvider,
	"huaweicloud_compute_interface_attach":  config.IdentifierFromProvider,
	"huaweicloud_compute_servergroup":       config.IdentifierFromProvider,
	"huaweicloud_compute_volume_attach":     config.IdentifierFromProvider,

	// eip
	"huaweicloud_global_eip":                config.IdentifierFromProvider,
	"huaweicloud_global_eip_associate":      config.IdentifierFromProvider,
	"huaweicloud_global_internet_bandwidth": config.IdentifierFromProvider,
	"huaweicloud_vpc_bandwidth":             config.IdentifierFromProvider,
	"huaweicloud_vpc_bandwidth_associate":   config.IdentifierFromProvider,
	"huaweicloud_vpc_eip":                   config.IdentifierFromProvider,
	"huaweicloud_vpc_eip_associate":         config.IdentifierFromProvider,
	"huaweicloud_vpc_eipv3_associate":       config.IdentifierFromProvider,
	"huaweicloud_vpc_internet_gateway":      config.IdentifierFromProvider,

	// evs
	"huaweicloud_evs_snapshot":                 config.IdentifierFromProvider,
	"huaweicloud_evs_snapshot_rollback":        config.IdentifierFromProvider,
	"huaweicloud_evs_volume":                   config.IdentifierFromProvider,
	"huaweicloud_evs_volume_transfer":          config.IdentifierFromProvider,
	"huaweicloud_evs_volume_transfer_accepter": config.IdentifierFromProvider,

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

	// modelarts
	"huaweicloud_modelarts_authorization":          config.IdentifierFromProvider,
	"huaweicloud_modelarts_dataset":                config.IdentifierFromProvider,
	"huaweicloud_modelarts_dataset_version":        config.IdentifierFromProvider,
	"huaweicloud_modelarts_model":                  config.IdentifierFromProvider,
	"huaweicloud_modelarts_network":                config.IdentifierFromProvider,
	"huaweicloud_modelarts_notebook":               config.IdentifierFromProvider,
	"huaweicloud_modelarts_notebook_mount_storage": config.IdentifierFromProvider,
	"huaweicloud_modelarts_resource_pool":          config.IdentifierFromProvider,
	"huaweicloud_modelarts_service":                config.IdentifierFromProvider,
	"huaweicloud_modelarts_workspace":              config.IdentifierFromProvider,

	// nat
	"huaweicloud_nat_dnat_rule":          config.IdentifierFromProvider,
	"huaweicloud_nat_gateway":            config.IdentifierFromProvider,
	"huaweicloud_nat_private_dnat_rule":  config.IdentifierFromProvider,
	"huaweicloud_nat_private_gateway":    config.IdentifierFromProvider,
	"huaweicloud_nat_private_snat_rule":  config.IdentifierFromProvider,
	"huaweicloud_nat_private_transit_ip": config.IdentifierFromProvider,
	"huaweicloud_nat_snat_rule":          config.IdentifierFromProvider,

	// obs
	"huaweicloud_obs_bucket":             config.IdentifierFromProvider,
	"huaweicloud_obs_bucket_acl":         config.IdentifierFromProvider,
	"huaweicloud_obs_bucket_object":      config.IdentifierFromProvider,
	"huaweicloud_obs_bucket_object_acl":  config.IdentifierFromProvider,
	"huaweicloud_obs_bucket_policy":      config.IdentifierFromProvider,
	"huaweicloud_obs_bucket_replication": config.IdentifierFromProvider,

	// rds
	"huaweicloud_rds_backup":                       config.IdentifierFromProvider,
	"huaweicloud_rds_cross_region_backup_strategy": config.IdentifierFromProvider,
	"huaweicloud_rds_database_logs_shrinking":      config.IdentifierFromProvider,
	"huaweicloud_rds_extend_log_link":              config.IdentifierFromProvider,
	"huaweicloud_rds_instance":                     config.IdentifierFromProvider,
	"huaweicloud_rds_instance_eip_associate":       config.IdentifierFromProvider,
	"huaweicloud_rds_lts_log":                      config.IdentifierFromProvider,
	"huaweicloud_rds_mysql_account":                config.IdentifierFromProvider,
	"huaweicloud_rds_mysql_binlog":                 config.IdentifierFromProvider,
	"huaweicloud_rds_mysql_database":               config.IdentifierFromProvider,
	"huaweicloud_rds_mysql_database_privilege":     config.IdentifierFromProvider,
	"huaweicloud_rds_mysql_database_table_restore": config.IdentifierFromProvider,
	"huaweicloud_rds_parametergroup":               config.IdentifierFromProvider,
	"huaweicloud_rds_pg_account":                   config.IdentifierFromProvider,
	"huaweicloud_rds_pg_account_privileges":        config.IdentifierFromProvider,
	"huaweicloud_rds_pg_account_roles":             config.IdentifierFromProvider,
	"huaweicloud_rds_pg_database":                  config.IdentifierFromProvider,
	"huaweicloud_rds_pg_database_privilege":        config.IdentifierFromProvider,
	"huaweicloud_rds_pg_hba":                       config.IdentifierFromProvider,
	"huaweicloud_rds_pg_plugin":                    config.IdentifierFromProvider,
	"huaweicloud_rds_pg_plugin_parameter":          config.IdentifierFromProvider,
	"huaweicloud_rds_pg_plugin_update":             config.IdentifierFromProvider,
	"huaweicloud_rds_pg_sql_limit":                 config.IdentifierFromProvider,
	"huaweicloud_rds_primary_standby_switch":       config.IdentifierFromProvider,
	"huaweicloud_rds_read_replica_instance":        config.IdentifierFromProvider,
	"huaweicloud_rds_restore":                      config.IdentifierFromProvider,
	"huaweicloud_rds_sql_audit":                    config.IdentifierFromProvider,
	"huaweicloud_rds_sqlserver_account":            config.IdentifierFromProvider,
	"huaweicloud_rds_sqlserver_database":           config.IdentifierFromProvider,
	"huaweicloud_rds_sqlserver_database_privilege": config.IdentifierFromProvider,

	// smn
	"huaweicloud_smn_logtank":                    config.IdentifierFromProvider,
	"huaweicloud_smn_message_detection":          config.IdentifierFromProvider,
	"huaweicloud_smn_message_publish":            config.IdentifierFromProvider,
	"huaweicloud_smn_message_template":           config.IdentifierFromProvider,
	"huaweicloud_smn_subscription":               config.IdentifierFromProvider,
	"huaweicloud_smn_subscription_filter_policy": config.IdentifierFromProvider,
	"huaweicloud_smn_topic":                      config.IdentifierFromProvider,

	// swr
	"huaweicloud_swr_image_auto_sync":          config.IdentifierFromProvider,
	"huaweicloud_swr_image_permissions":        config.IdentifierFromProvider,
	"huaweicloud_swr_image_retention_policy":   config.IdentifierFromProvider,
	"huaweicloud_swr_image_trigger":            config.IdentifierFromProvider,
	"huaweicloud_swr_organization":             config.IdentifierFromProvider,
	"huaweicloud_swr_organization_permissions": config.IdentifierFromProvider,
	"huaweicloud_swr_repository":               config.IdentifierFromProvider,
	"huaweicloud_swr_repository_sharing":       config.IdentifierFromProvider,

	// vpc
	"huaweicloud_vpc":                             config.IdentifierFromProvider,
	"huaweicloud_vpc_subnet":                      config.IdentifierFromProvider,
	"huaweicloud_networking_secgroup":             config.IdentifierFromProvider,
	"huaweicloud_networking_secgroup_rule":        config.IdentifierFromProvider,
	"huaweicloud_networking_vip":                  config.IdentifierFromProvider,
	"huaweicloud_networking_vip_associate":        config.IdentifierFromProvider,
	"huaweicloud_vpc_address_group":               config.IdentifierFromProvider,
	"huaweicloud_vpc_flow_log":                    config.IdentifierFromProvider,
	"huaweicloud_vpc_network_acl":                 config.IdentifierFromProvider,
	"huaweicloud_vpc_network_interface":           config.IdentifierFromProvider,
	"huaweicloud_vpc_peering_connection":          config.IdentifierFromProvider,
	"huaweicloud_vpc_peering_connection_accepter": config.IdentifierFromProvider,
	"huaweicloud_vpc_route":                       config.IdentifierFromProvider,
	"huaweicloud_vpc_route_table":                 config.IdentifierFromProvider,
	"huaweicloud_vpc_sub_network_interface":       config.IdentifierFromProvider,
	"huaweicloud_vpc_traffic_mirror_filter":       config.IdentifierFromProvider,
	"huaweicloud_vpc_traffic_mirror_filter_rule":  config.IdentifierFromProvider,
	"huaweicloud_vpc_traffic_mirror_session":      config.IdentifierFromProvider,
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
