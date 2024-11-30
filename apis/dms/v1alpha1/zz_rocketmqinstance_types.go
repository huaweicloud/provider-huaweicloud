// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConfigsInitParameters struct {

	// Specifies the config name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the config value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConfigsObservation struct {

	// Specifies the config name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the config value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConfigsParameters struct {

	// Specifies the config name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the config value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type CrossVPCAccessesInitParameters struct {
}

type CrossVPCAccessesObservation struct {

	// The advertised IP Address or domain name.
	AdvertisedIP *string `json:"advertisedIp,omitempty" tf:"advertised_ip,omitempty"`

	LisenterIP *string `json:"lisenterIp,omitempty" tf:"lisenter_ip,omitempty"`

	// The listener IP address.
	ListenerIP *string `json:"listenerIp,omitempty" tf:"listener_ip,omitempty"`

	// The port number.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The port ID associated with the address.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`
}

type CrossVPCAccessesParameters struct {
}

type RocketmqInstanceInitParameters struct {

	// Specifies whether auto renew is enabled. Valid values are "true" and "false".
	AutoRenew *string `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// Specifies the list of availability zone names, where
	// instance brokers reside and which has available resources.
	// Specifies the list of availability zone names
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Specifies the broker numbers. It's required when instance architecture is
	// cluster. Defaults to 1 when instance architecture is single node.
	// Specifies the broker numbers.
	BrokerNum *float64 `json:"brokerNum,omitempty" tf:"broker_num,omitempty"`

	// Specifies the charging mode of the instance. Valid values are prePaid
	// and postPaid, defaults to postPaid. Changing this creates a new resource.
	ChargingMode *string `json:"chargingMode,omitempty" tf:"charging_mode,omitempty"`

	// Specifies the instance configs.
	// The configs structure is documented below.
	// Specifies the instance configs.
	Configs []ConfigsInitParameters `json:"configs,omitempty" tf:"configs,omitempty"`

	// Specifies the description of the DMS RocketMQ instance.
	// The description can contain a maximum of 1,024 characters.
	// Specifies the description of the DMS RocketMQ instance.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether access control is enabled.
	// Specifies whether access control is enabled.
	EnableACL *bool `json:"enableAcl,omitempty" tf:"enable_acl,omitempty"`

	// Specifies whether to enable public access. By default, public access is disabled.
	// Specifies whether to enable public access.
	EnablePublicip *bool `json:"enablePublicip,omitempty" tf:"enable_publicip,omitempty"`

	// Specifies the version of the RocketMQ engine.
	// Valid values are 4.8.0 and 5.x.
	// Changing this parameter will create a new resource.
	// Specifies the version of the RocketMQ engine.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Specifies the enterprise project id of the instance.
	// Specifies the enterprise project id of the instance.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies the flavor ID.
	// Specifies a product ID
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// Specifies whether to support IPv6. Defaults to false.
	// Changing this parameter will create a new resource.
	// Specifies whether to support IPv6
	IPv6Enable *bool `json:"ipv6Enable,omitempty" tf:"ipv6_enable,omitempty"`

	// Specifies the name of the DMS RocketMQ instance.
	// An instance name starts with a letter, consists of 4 to 64 characters, and can contain only letters,
	// digits, underscores (_), and hyphens (-).
	// Specifies the name of the DMS RocketMQ instance
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the charging period of the instance. If period_unit is set to month
	// , the value ranges from 1 to 9. If period_unit is set to year, the value ranges from 1 to 3. This parameter is
	// mandatory if charging_mode is set to prePaid. Changing this creates a new resource.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the charging period unit of the instance.
	// Valid values are month and year. This parameter is mandatory if charging_mode is set to prePaid.
	// Changing this creates a new resource.
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// Specifies the ID of the EIP bound to the instance. Use commas (,) to separate
	// multiple EIP IDs. It is mandatory if enable_publicip is true and should be empty when enable_publicip is false.
	// Specifies the ID of the EIP bound to the instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/eip/v1alpha1.VpcEip
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	PublicipID *string `json:"publicipId,omitempty" tf:"publicip_id,omitempty"`

	// Reference to a VpcEip in eip to populate publicipId.
	// +kubebuilder:validation:Optional
	PublicipIDRef *v1.Reference `json:"publicipIdRef,omitempty" tf:"-"`

	// Selector for a VpcEip in eip to populate publicipId.
	// +kubebuilder:validation:Optional
	PublicipIDSelector *v1.Selector `json:"publicipIdSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies whether access control is enabled.
	RetentionPolicy *bool `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// Specifies whether the RocketMQ SASL_SSL is enabled. Defaults to false.
	// Changing this parameter will create a new resource.
	// Specifies whether the RocketMQ SASL_SSL is enabled.
	SSLEnable *bool `json:"sslEnable,omitempty" tf:"ssl_enable,omitempty"`

	// Specifies the ID of a security group.
	// Specifies the ID of a security group
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.Secgroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a Secgroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a Secgroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Specifies the message storage capacity, Unit: GB.
	// When engine_version is 4.8.0, value ranges from 300 to 30,000.
	// When engine_version is 5.x, value ranges from 200 to 60,000.
	// Specifies the message storage capacity, Unit: GB.
	StorageSpace *float64 `json:"storageSpace,omitempty" tf:"storage_space,omitempty"`

	// Specifies the storage I/O specification.
	// The options are as follows:
	// Specifies the storage I/O specification
	StorageSpecCode *string `json:"storageSpecCode,omitempty" tf:"storage_spec_code,omitempty"`

	// Specifies the ID of a subnet.
	// Changing this parameter will create a new resource.
	// Specifies the ID of a subnet
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Specifies the key/value pairs to associate with the instance.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ID of a VPC.
	// Changing this parameter will create a new resource.
	// Specifies the ID of a VPC
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type RocketmqInstanceObservation struct {

	// Specifies whether auto renew is enabled. Valid values are "true" and "false".
	AutoRenew *string `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// Specifies the list of availability zone names, where
	// instance brokers reside and which has available resources.
	// Specifies the list of availability zone names
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Indicates the service data address.
	// Indicates the service data address.
	BrokerAddress *string `json:"brokerAddress,omitempty" tf:"broker_address,omitempty"`

	// Specifies the broker numbers. It's required when instance architecture is
	// cluster. Defaults to 1 when instance architecture is single node.
	// Specifies the broker numbers.
	BrokerNum *float64 `json:"brokerNum,omitempty" tf:"broker_num,omitempty"`

	// Specifies the charging mode of the instance. Valid values are prePaid
	// and postPaid, defaults to postPaid. Changing this creates a new resource.
	ChargingMode *string `json:"chargingMode,omitempty" tf:"charging_mode,omitempty"`

	// Specifies the instance configs.
	// The configs structure is documented below.
	// Specifies the instance configs.
	Configs []ConfigsObservation `json:"configs,omitempty" tf:"configs,omitempty"`

	// Indicates the Access information of cross-VPC. The structure is documented below.
	CrossVPCAccesses []CrossVPCAccessesObservation `json:"crossVpcAccesses,omitempty" tf:"cross_vpc_accesses,omitempty"`

	// Specifies the description of the DMS RocketMQ instance.
	// The description can contain a maximum of 1,024 characters.
	// Specifies the description of the DMS RocketMQ instance.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether access control is enabled.
	// Specifies whether access control is enabled.
	EnableACL *bool `json:"enableAcl,omitempty" tf:"enable_acl,omitempty"`

	// Specifies whether to enable public access. By default, public access is disabled.
	// Specifies whether to enable public access.
	EnablePublicip *bool `json:"enablePublicip,omitempty" tf:"enable_publicip,omitempty"`

	// Specifies the version of the RocketMQ engine.
	// Valid values are 4.8.0 and 5.x.
	// Changing this parameter will create a new resource.
	// Specifies the version of the RocketMQ engine.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Specifies the enterprise project id of the instance.
	// Specifies the enterprise project id of the instance.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies the flavor ID.
	// Specifies a product ID
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// Specifies a resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies whether to support IPv6. Defaults to false.
	// Changing this parameter will create a new resource.
	// Specifies whether to support IPv6
	IPv6Enable *bool `json:"ipv6Enable,omitempty" tf:"ipv6_enable,omitempty"`

	// Indicates the time at which the maintenance window starts. The format is HH:mm:ss.
	// Indicates the time at which the maintenance window starts. The format is HH:mm:ss.
	MaintainBegin *string `json:"maintainBegin,omitempty" tf:"maintain_begin,omitempty"`

	// Indicates the time at which the maintenance window ends. The format is HH:mm:ss.
	// Indicates the time at which the maintenance window ends. The format is HH:mm:ss.
	MaintainEnd *string `json:"maintainEnd,omitempty" tf:"maintain_end,omitempty"`

	// Specifies the name of the DMS RocketMQ instance.
	// An instance name starts with a letter, consists of 4 to 64 characters, and can contain only letters,
	// digits, underscores (_), and hyphens (-).
	// Specifies the name of the DMS RocketMQ instance
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the metadata address.
	// Indicates the metadata address.
	NamesrvAddress *string `json:"namesrvAddress,omitempty" tf:"namesrv_address,omitempty"`

	// Indicates whether billing based on new specifications is enabled.
	// Indicates whether billing based on new specifications is enabled.
	NewSpecBillingEnable *bool `json:"newSpecBillingEnable,omitempty" tf:"new_spec_billing_enable,omitempty"`

	// Indicates the node quantity.
	// Indicates the node quantity.
	NodeNum *float64 `json:"nodeNum,omitempty" tf:"node_num,omitempty"`

	// Specifies the charging period of the instance. If period_unit is set to month
	// , the value ranges from 1 to 9. If period_unit is set to year, the value ranges from 1 to 3. This parameter is
	// mandatory if charging_mode is set to prePaid. Changing this creates a new resource.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the charging period unit of the instance.
	// Valid values are month and year. This parameter is mandatory if charging_mode is set to prePaid.
	// Changing this creates a new resource.
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// Indicates the public network service data address.
	// Indicates the public network service data address.
	PublicBrokerAddress *string `json:"publicBrokerAddress,omitempty" tf:"public_broker_address,omitempty"`

	// Indicates the public network metadata address.
	// Indicates the public network metadata address.
	PublicNamesrvAddress *string `json:"publicNamesrvAddress,omitempty" tf:"public_namesrv_address,omitempty"`

	// Indicates the public IP address.
	// Indicates the public IP address.
	PublicipAddress *string `json:"publicipAddress,omitempty" tf:"publicip_address,omitempty"`

	// Specifies the ID of the EIP bound to the instance. Use commas (,) to separate
	// multiple EIP IDs. It is mandatory if enable_publicip is true and should be empty when enable_publicip is false.
	// Specifies the ID of the EIP bound to the instance.
	PublicipID *string `json:"publicipId,omitempty" tf:"publicip_id,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates the resource specifications.
	// Indicates the resource specifications.
	ResourceSpecCode *string `json:"resourceSpecCode,omitempty" tf:"resource_spec_code,omitempty"`

	// Specifies whether access control is enabled.
	RetentionPolicy *bool `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// Specifies whether the RocketMQ SASL_SSL is enabled. Defaults to false.
	// Changing this parameter will create a new resource.
	// Specifies whether the RocketMQ SASL_SSL is enabled.
	SSLEnable *bool `json:"sslEnable,omitempty" tf:"ssl_enable,omitempty"`

	// Specifies the ID of a security group.
	// Specifies the ID of a security group
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Indicates the instance specification. For a cluster DMS RocketMQ instance, VM specifications
	// and the number of nodes are returned.
	// Indicates the instance specification. For a cluster DMS RocketMQ instance, VM specifications
	// and the number of nodes are returned.
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`

	// Indicates the status of the DMS RocketMQ instance.
	// Indicates the status of the DMS RocketMQ instance.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the message storage capacity, Unit: GB.
	// When engine_version is 4.8.0, value ranges from 300 to 30,000.
	// When engine_version is 5.x, value ranges from 200 to 60,000.
	// Specifies the message storage capacity, Unit: GB.
	StorageSpace *float64 `json:"storageSpace,omitempty" tf:"storage_space,omitempty"`

	// Specifies the storage I/O specification.
	// The options are as follows:
	// Specifies the storage I/O specification
	StorageSpecCode *string `json:"storageSpecCode,omitempty" tf:"storage_spec_code,omitempty"`

	// Specifies the ID of a subnet.
	// Changing this parameter will create a new resource.
	// Specifies the ID of a subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Specifies the key/value pairs to associate with the instance.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Indicates the DMS RocketMQ instance type. Value: cluster.
	// Indicates the DMS RocketMQ instance type. Value: cluster.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Indicates the used message storage space. Unit: GB.
	// Indicates the used message storage space. Unit: GB.
	UsedStorageSpace *float64 `json:"usedStorageSpace,omitempty" tf:"used_storage_space,omitempty"`

	// Specifies the ID of a VPC.
	// Changing this parameter will create a new resource.
	// Specifies the ID of a VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type RocketmqInstanceParameters struct {

	// Specifies whether auto renew is enabled. Valid values are "true" and "false".
	// +kubebuilder:validation:Optional
	AutoRenew *string `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// Specifies the list of availability zone names, where
	// instance brokers reside and which has available resources.
	// Specifies the list of availability zone names
	// +kubebuilder:validation:Optional
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Specifies the broker numbers. It's required when instance architecture is
	// cluster. Defaults to 1 when instance architecture is single node.
	// Specifies the broker numbers.
	// +kubebuilder:validation:Optional
	BrokerNum *float64 `json:"brokerNum,omitempty" tf:"broker_num,omitempty"`

	// Specifies the charging mode of the instance. Valid values are prePaid
	// and postPaid, defaults to postPaid. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	ChargingMode *string `json:"chargingMode,omitempty" tf:"charging_mode,omitempty"`

	// Specifies the instance configs.
	// The configs structure is documented below.
	// Specifies the instance configs.
	// +kubebuilder:validation:Optional
	Configs []ConfigsParameters `json:"configs,omitempty" tf:"configs,omitempty"`

	// Specifies the description of the DMS RocketMQ instance.
	// The description can contain a maximum of 1,024 characters.
	// Specifies the description of the DMS RocketMQ instance.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether access control is enabled.
	// Specifies whether access control is enabled.
	// +kubebuilder:validation:Optional
	EnableACL *bool `json:"enableAcl,omitempty" tf:"enable_acl,omitempty"`

	// Specifies whether to enable public access. By default, public access is disabled.
	// Specifies whether to enable public access.
	// +kubebuilder:validation:Optional
	EnablePublicip *bool `json:"enablePublicip,omitempty" tf:"enable_publicip,omitempty"`

	// Specifies the version of the RocketMQ engine.
	// Valid values are 4.8.0 and 5.x.
	// Changing this parameter will create a new resource.
	// Specifies the version of the RocketMQ engine.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Specifies the enterprise project id of the instance.
	// Specifies the enterprise project id of the instance.
	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies the flavor ID.
	// Specifies a product ID
	// +kubebuilder:validation:Optional
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// Specifies whether to support IPv6. Defaults to false.
	// Changing this parameter will create a new resource.
	// Specifies whether to support IPv6
	// +kubebuilder:validation:Optional
	IPv6Enable *bool `json:"ipv6Enable,omitempty" tf:"ipv6_enable,omitempty"`

	// Specifies the name of the DMS RocketMQ instance.
	// An instance name starts with a letter, consists of 4 to 64 characters, and can contain only letters,
	// digits, underscores (_), and hyphens (-).
	// Specifies the name of the DMS RocketMQ instance
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the charging period of the instance. If period_unit is set to month
	// , the value ranges from 1 to 9. If period_unit is set to year, the value ranges from 1 to 3. This parameter is
	// mandatory if charging_mode is set to prePaid. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the charging period unit of the instance.
	// Valid values are month and year. This parameter is mandatory if charging_mode is set to prePaid.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// Specifies the ID of the EIP bound to the instance. Use commas (,) to separate
	// multiple EIP IDs. It is mandatory if enable_publicip is true and should be empty when enable_publicip is false.
	// Specifies the ID of the EIP bound to the instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/eip/v1alpha1.VpcEip
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PublicipID *string `json:"publicipId,omitempty" tf:"publicip_id,omitempty"`

	// Reference to a VpcEip in eip to populate publicipId.
	// +kubebuilder:validation:Optional
	PublicipIDRef *v1.Reference `json:"publicipIdRef,omitempty" tf:"-"`

	// Selector for a VpcEip in eip to populate publicipId.
	// +kubebuilder:validation:Optional
	PublicipIDSelector *v1.Selector `json:"publicipIdSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies whether access control is enabled.
	// +kubebuilder:validation:Optional
	RetentionPolicy *bool `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// Specifies whether the RocketMQ SASL_SSL is enabled. Defaults to false.
	// Changing this parameter will create a new resource.
	// Specifies whether the RocketMQ SASL_SSL is enabled.
	// +kubebuilder:validation:Optional
	SSLEnable *bool `json:"sslEnable,omitempty" tf:"ssl_enable,omitempty"`

	// Specifies the ID of a security group.
	// Specifies the ID of a security group
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.Secgroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a Secgroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a Secgroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Specifies the message storage capacity, Unit: GB.
	// When engine_version is 4.8.0, value ranges from 300 to 30,000.
	// When engine_version is 5.x, value ranges from 200 to 60,000.
	// Specifies the message storage capacity, Unit: GB.
	// +kubebuilder:validation:Optional
	StorageSpace *float64 `json:"storageSpace,omitempty" tf:"storage_space,omitempty"`

	// Specifies the storage I/O specification.
	// The options are as follows:
	// Specifies the storage I/O specification
	// +kubebuilder:validation:Optional
	StorageSpecCode *string `json:"storageSpecCode,omitempty" tf:"storage_spec_code,omitempty"`

	// Specifies the ID of a subnet.
	// Changing this parameter will create a new resource.
	// Specifies the ID of a subnet
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Specifies the key/value pairs to associate with the instance.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ID of a VPC.
	// Changing this parameter will create a new resource.
	// Specifies the ID of a VPC
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// RocketmqInstanceSpec defines the desired state of RocketmqInstance
type RocketmqInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RocketmqInstanceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RocketmqInstanceInitParameters `json:"initProvider,omitempty"`
}

// RocketmqInstanceStatus defines the observed state of RocketmqInstance.
type RocketmqInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RocketmqInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RocketmqInstance is the Schema for the RocketmqInstances API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type RocketmqInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZones) || (has(self.initProvider) && has(self.initProvider.availabilityZones))",message="spec.forProvider.availabilityZones is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engineVersion) || (has(self.initProvider) && has(self.initProvider.engineVersion))",message="spec.forProvider.engineVersion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.flavorId) || (has(self.initProvider) && has(self.initProvider.flavorId))",message="spec.forProvider.flavorId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageSpace) || (has(self.initProvider) && has(self.initProvider.storageSpace))",message="spec.forProvider.storageSpace is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageSpecCode) || (has(self.initProvider) && has(self.initProvider.storageSpecCode))",message="spec.forProvider.storageSpecCode is a required parameter"
	Spec   RocketmqInstanceSpec   `json:"spec"`
	Status RocketmqInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RocketmqInstanceList contains a list of RocketmqInstances
type RocketmqInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RocketmqInstance `json:"items"`
}

// Repository type metadata.
var (
	RocketmqInstance_Kind             = "RocketmqInstance"
	RocketmqInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RocketmqInstance_Kind}.String()
	RocketmqInstance_KindAPIVersion   = RocketmqInstance_Kind + "." + CRDGroupVersion.String()
	RocketmqInstance_GroupVersionKind = CRDGroupVersion.WithKind(RocketmqInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&RocketmqInstance{}, &RocketmqInstanceList{})
}