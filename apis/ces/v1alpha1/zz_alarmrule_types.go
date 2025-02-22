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

type AlarmActionsInitParameters struct {

	// specifies the list of objects to be notified if the alarm status changes, the
	// maximum length is 5. If type is set to notification, the value of notification_list cannot be empty. If type is
	// set to autoscaling, the value of notification_list must be []
	// and the value of namespace must be SYS.AS.
	NotificationList []*string `json:"notificationList,omitempty" tf:"notification_list,omitempty"`

	// Specifies the type of action triggered by an alarm. the
	// value can be notification or autoscaling.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AlarmActionsObservation struct {

	// specifies the list of objects to be notified if the alarm status changes, the
	// maximum length is 5. If type is set to notification, the value of notification_list cannot be empty. If type is
	// set to autoscaling, the value of notification_list must be []
	// and the value of namespace must be SYS.AS.
	NotificationList []*string `json:"notificationList,omitempty" tf:"notification_list,omitempty"`

	// Specifies the type of action triggered by an alarm. the
	// value can be notification or autoscaling.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AlarmActionsParameters struct {

	// specifies the list of objects to be notified if the alarm status changes, the
	// maximum length is 5. If type is set to notification, the value of notification_list cannot be empty. If type is
	// set to autoscaling, the value of notification_list must be []
	// and the value of namespace must be SYS.AS.
	// +kubebuilder:validation:Optional
	NotificationList []*string `json:"notificationList" tf:"notification_list,omitempty"`

	// Specifies the type of action triggered by an alarm. the
	// value can be notification or autoscaling.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type AlarmruleInitParameters struct {

	// Specifies whether to enable the action to be triggered by an alarm. The
	// default value is true.
	AlarmActionEnabled *bool `json:"alarmActionEnabled,omitempty" tf:"alarm_action_enabled,omitempty"`

	// Specifies the action triggered by an alarm. The structure is described
	// below.
	AlarmActions []AlarmActionsInitParameters `json:"alarmActions,omitempty" tf:"alarm_actions,omitempty"`

	// The value can be a string of 0 to 256 characters.
	AlarmDescription *string `json:"alarmDescription,omitempty" tf:"alarm_description,omitempty"`

	// Specifies whether to enable the alarm. The default value is true.
	AlarmEnabled *bool `json:"alarmEnabled,omitempty" tf:"alarm_enabled,omitempty"`

	// Specifies the alarm severity of the condition. The value can be 1, 2, 3 or 4,
	// which indicates critical, major, minor, and informational, respectively.
	// The default value is 2.
	// schema: Deprecated
	AlarmLevel *float64 `json:"alarmLevel,omitempty" tf:"alarm_level,omitempty"`

	// Specifies the name of an alarm rule. The value can be a string of 1 to 128
	// characters that can consist of English letters, Chinese characters, digits, underscores (_), hyphens (-).
	AlarmName *string `json:"alarmName,omitempty" tf:"alarm_name,omitempty"`

	// Specifies the alarm type. The value can be EVENT.SYS, EVENT.CUSTOM,
	// MULTI_INSTANCE and ALL_INSTANCE. Defaults to MULTI_INSTANCE.
	AlarmType *string `json:"alarmType,omitempty" tf:"alarm_type,omitempty"`

	// Specifies the alarm triggering condition. The structure is described below.
	Condition []ConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies the enterprise project ID of the alarm rule.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	InsufficientdataActions []InsufficientdataActionsInitParameters `json:"insufficientdataActions,omitempty" tf:"insufficientdata_actions,omitempty"`

	// Specifies the alarm metrics. The structure is described below. Changing this
	// creates a new resource.
	Metric []MetricInitParameters `json:"metric,omitempty" tf:"metric,omitempty"`

	// Specifies the alarm notification start time, for
	// example: 05:30. Changing this creates a new resource.
	NotificationBeginTime *string `json:"notificationBeginTime,omitempty" tf:"notification_begin_time,omitempty"`

	// Specifies the alarm notification stop time, for
	// example: 22:10. Changing this creates a new resource.
	NotificationEndTime *string `json:"notificationEndTime,omitempty" tf:"notification_end_time,omitempty"`

	// Specifies the action triggered by the clearing of an alarm. The structure is
	// described below.
	OkActions []OkActionsInitParameters `json:"okActions,omitempty" tf:"ok_actions,omitempty"`

	// The region in which to create the alarm rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates the alarm rule ID.
	// schema: Internal
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// Specifies the list of the resources to add into the alarm rule.
	// The structure is described below.
	Resources []ResourcesInitParameters `json:"resources,omitempty" tf:"resources,omitempty"`
}

type AlarmruleObservation struct {

	// Specifies whether to enable the action to be triggered by an alarm. The
	// default value is true.
	AlarmActionEnabled *bool `json:"alarmActionEnabled,omitempty" tf:"alarm_action_enabled,omitempty"`

	// Specifies the action triggered by an alarm. The structure is described
	// below.
	AlarmActions []AlarmActionsObservation `json:"alarmActions,omitempty" tf:"alarm_actions,omitempty"`

	// The value can be a string of 0 to 256 characters.
	AlarmDescription *string `json:"alarmDescription,omitempty" tf:"alarm_description,omitempty"`

	// Specifies whether to enable the alarm. The default value is true.
	AlarmEnabled *bool `json:"alarmEnabled,omitempty" tf:"alarm_enabled,omitempty"`

	// Specifies the alarm severity of the condition. The value can be 1, 2, 3 or 4,
	// which indicates critical, major, minor, and informational, respectively.
	// The default value is 2.
	// schema: Deprecated
	AlarmLevel *float64 `json:"alarmLevel,omitempty" tf:"alarm_level,omitempty"`

	// Specifies the name of an alarm rule. The value can be a string of 1 to 128
	// characters that can consist of English letters, Chinese characters, digits, underscores (_), hyphens (-).
	AlarmName *string `json:"alarmName,omitempty" tf:"alarm_name,omitempty"`

	// Indicates the alarm status. The value can be:
	AlarmState *string `json:"alarmState,omitempty" tf:"alarm_state,omitempty"`

	// Specifies the alarm type. The value can be EVENT.SYS, EVENT.CUSTOM,
	// MULTI_INSTANCE and ALL_INSTANCE. Defaults to MULTI_INSTANCE.
	AlarmType *string `json:"alarmType,omitempty" tf:"alarm_type,omitempty"`

	// Specifies the alarm triggering condition. The structure is described below.
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies the enterprise project ID of the alarm rule.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Indicates the alarm rule ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InsufficientdataActions []InsufficientdataActionsObservation `json:"insufficientdataActions,omitempty" tf:"insufficientdata_actions,omitempty"`

	// Specifies the alarm metrics. The structure is described below. Changing this
	// creates a new resource.
	Metric []MetricObservation `json:"metric,omitempty" tf:"metric,omitempty"`

	// Specifies the alarm notification start time, for
	// example: 05:30. Changing this creates a new resource.
	NotificationBeginTime *string `json:"notificationBeginTime,omitempty" tf:"notification_begin_time,omitempty"`

	// Specifies the alarm notification stop time, for
	// example: 22:10. Changing this creates a new resource.
	NotificationEndTime *string `json:"notificationEndTime,omitempty" tf:"notification_end_time,omitempty"`

	// Specifies the action triggered by the clearing of an alarm. The structure is
	// described below.
	OkActions []OkActionsObservation `json:"okActions,omitempty" tf:"ok_actions,omitempty"`

	// The region in which to create the alarm rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates the alarm rule ID.
	// schema: Internal
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// Specifies the list of the resources to add into the alarm rule.
	// The structure is described below.
	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// Indicates the time when the alarm status changed. The value is a UNIX timestamp and the unit is ms.
	UpdateTime *float64 `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type AlarmruleParameters struct {

	// Specifies whether to enable the action to be triggered by an alarm. The
	// default value is true.
	// +kubebuilder:validation:Optional
	AlarmActionEnabled *bool `json:"alarmActionEnabled,omitempty" tf:"alarm_action_enabled,omitempty"`

	// Specifies the action triggered by an alarm. The structure is described
	// below.
	// +kubebuilder:validation:Optional
	AlarmActions []AlarmActionsParameters `json:"alarmActions,omitempty" tf:"alarm_actions,omitempty"`

	// The value can be a string of 0 to 256 characters.
	// +kubebuilder:validation:Optional
	AlarmDescription *string `json:"alarmDescription,omitempty" tf:"alarm_description,omitempty"`

	// Specifies whether to enable the alarm. The default value is true.
	// +kubebuilder:validation:Optional
	AlarmEnabled *bool `json:"alarmEnabled,omitempty" tf:"alarm_enabled,omitempty"`

	// Specifies the alarm severity of the condition. The value can be 1, 2, 3 or 4,
	// which indicates critical, major, minor, and informational, respectively.
	// The default value is 2.
	// schema: Deprecated
	// +kubebuilder:validation:Optional
	AlarmLevel *float64 `json:"alarmLevel,omitempty" tf:"alarm_level,omitempty"`

	// Specifies the name of an alarm rule. The value can be a string of 1 to 128
	// characters that can consist of English letters, Chinese characters, digits, underscores (_), hyphens (-).
	// +kubebuilder:validation:Optional
	AlarmName *string `json:"alarmName,omitempty" tf:"alarm_name,omitempty"`

	// Specifies the alarm type. The value can be EVENT.SYS, EVENT.CUSTOM,
	// MULTI_INSTANCE and ALL_INSTANCE. Defaults to MULTI_INSTANCE.
	// +kubebuilder:validation:Optional
	AlarmType *string `json:"alarmType,omitempty" tf:"alarm_type,omitempty"`

	// Specifies the alarm triggering condition. The structure is described below.
	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies the enterprise project ID of the alarm rule.
	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// +kubebuilder:validation:Optional
	InsufficientdataActions []InsufficientdataActionsParameters `json:"insufficientdataActions,omitempty" tf:"insufficientdata_actions,omitempty"`

	// Specifies the alarm metrics. The structure is described below. Changing this
	// creates a new resource.
	// +kubebuilder:validation:Optional
	Metric []MetricParameters `json:"metric,omitempty" tf:"metric,omitempty"`

	// Specifies the alarm notification start time, for
	// example: 05:30. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	NotificationBeginTime *string `json:"notificationBeginTime,omitempty" tf:"notification_begin_time,omitempty"`

	// Specifies the alarm notification stop time, for
	// example: 22:10. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	NotificationEndTime *string `json:"notificationEndTime,omitempty" tf:"notification_end_time,omitempty"`

	// Specifies the action triggered by the clearing of an alarm. The structure is
	// described below.
	// +kubebuilder:validation:Optional
	OkActions []OkActionsParameters `json:"okActions,omitempty" tf:"ok_actions,omitempty"`

	// The region in which to create the alarm rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates the alarm rule ID.
	// schema: Internal
	// +kubebuilder:validation:Optional
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// Specifies the list of the resources to add into the alarm rule.
	// The structure is described below.
	// +kubebuilder:validation:Optional
	Resources []ResourcesParameters `json:"resources,omitempty" tf:"resources,omitempty"`
}

type ConditionInitParameters struct {

	// Specifies the alarm severity of the condition. The value can be 1, 2, 3 or 4,
	// which indicates critical, major, minor, and informational, respectively.
	// The default value is 2.
	AlarmLevel *float64 `json:"alarmLevel,omitempty" tf:"alarm_level,omitempty"`

	// Specifies the comparison condition of alarm thresholds. The value can be >,
	// =, <, >=, or <=.
	ComparisonOperator *string `json:"comparisonOperator,omitempty" tf:"comparison_operator,omitempty"`

	// Specifies the number of consecutive occurrence times. The value ranges from 1 to 5.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// Specifies the data rollup methods. The value can be max, min, average, sum, and variance.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// Specifies the metric name of the condition. The value can be a string of
	// 1 to 64 characters that must start with a letter and contain only letters, digits, and underscores (_).
	// For details, see Services Interconnected with Cloud Eye.
	// schema: Required
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Specifies the alarm checking period in seconds. The value can be 0, 1, 300, 1200, 3600, 14400,
	// and 86400.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the interval for triggering an alarm if the alarm persists.
	// Possible values are as follows:
	SuppressDuration *float64 `json:"suppressDuration,omitempty" tf:"suppress_duration,omitempty"`

	// Specifies the data unit.
	// For details, see Services Interconnected with Cloud Eye.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// Specifies the alarm threshold. The value ranges from 0 to Number of
	// 1.7976931348623157e+108.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type ConditionObservation struct {

	// Specifies the alarm severity of the condition. The value can be 1, 2, 3 or 4,
	// which indicates critical, major, minor, and informational, respectively.
	// The default value is 2.
	AlarmLevel *float64 `json:"alarmLevel,omitempty" tf:"alarm_level,omitempty"`

	// Specifies the comparison condition of alarm thresholds. The value can be >,
	// =, <, >=, or <=.
	ComparisonOperator *string `json:"comparisonOperator,omitempty" tf:"comparison_operator,omitempty"`

	// Specifies the number of consecutive occurrence times. The value ranges from 1 to 5.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// Specifies the data rollup methods. The value can be max, min, average, sum, and variance.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// Specifies the metric name of the condition. The value can be a string of
	// 1 to 64 characters that must start with a letter and contain only letters, digits, and underscores (_).
	// For details, see Services Interconnected with Cloud Eye.
	// schema: Required
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Specifies the alarm checking period in seconds. The value can be 0, 1, 300, 1200, 3600, 14400,
	// and 86400.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the interval for triggering an alarm if the alarm persists.
	// Possible values are as follows:
	SuppressDuration *float64 `json:"suppressDuration,omitempty" tf:"suppress_duration,omitempty"`

	// Specifies the data unit.
	// For details, see Services Interconnected with Cloud Eye.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// Specifies the alarm threshold. The value ranges from 0 to Number of
	// 1.7976931348623157e+108.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type ConditionParameters struct {

	// Specifies the alarm severity of the condition. The value can be 1, 2, 3 or 4,
	// which indicates critical, major, minor, and informational, respectively.
	// The default value is 2.
	// +kubebuilder:validation:Optional
	AlarmLevel *float64 `json:"alarmLevel,omitempty" tf:"alarm_level,omitempty"`

	// Specifies the comparison condition of alarm thresholds. The value can be >,
	// =, <, >=, or <=.
	// +kubebuilder:validation:Optional
	ComparisonOperator *string `json:"comparisonOperator" tf:"comparison_operator,omitempty"`

	// Specifies the number of consecutive occurrence times. The value ranges from 1 to 5.
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count" tf:"count,omitempty"`

	// Specifies the data rollup methods. The value can be max, min, average, sum, and variance.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter" tf:"filter,omitempty"`

	// Specifies the metric name of the condition. The value can be a string of
	// 1 to 64 characters that must start with a letter and contain only letters, digits, and underscores (_).
	// For details, see Services Interconnected with Cloud Eye.
	// schema: Required
	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Specifies the alarm checking period in seconds. The value can be 0, 1, 300, 1200, 3600, 14400,
	// and 86400.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period" tf:"period,omitempty"`

	// Specifies the interval for triggering an alarm if the alarm persists.
	// Possible values are as follows:
	// +kubebuilder:validation:Optional
	SuppressDuration *float64 `json:"suppressDuration,omitempty" tf:"suppress_duration,omitempty"`

	// Specifies the data unit.
	// For details, see Services Interconnected with Cloud Eye.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// Specifies the alarm threshold. The value ranges from 0 to Number of
	// 1.7976931348623157e+108.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type DimensionsInitParameters struct {

	// Specifies the dimension name. The value can be a string of 1 to 32 characters
	// that must start with a letter and contain only letters, digits, underscores (_), and hyphens (-).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the dimension value. The value can be a string of 1 to 64 characters
	// that must start with a letter or a number and contain only letters, digits, underscores (_), and hyphens (-).
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DimensionsObservation struct {

	// Specifies the dimension name. The value can be a string of 1 to 32 characters
	// that must start with a letter and contain only letters, digits, underscores (_), and hyphens (-).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the dimension value. The value can be a string of 1 to 64 characters
	// that must start with a letter or a number and contain only letters, digits, underscores (_), and hyphens (-).
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DimensionsParameters struct {

	// Specifies the dimension name. The value can be a string of 1 to 32 characters
	// that must start with a letter and contain only letters, digits, underscores (_), and hyphens (-).
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the dimension value. The value can be a string of 1 to 64 characters
	// that must start with a letter or a number and contain only letters, digits, underscores (_), and hyphens (-).
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type InsufficientdataActionsInitParameters struct {

	// specifies the list of objects to be notified if the alarm status changes, the
	// maximum length is 5.
	NotificationList []*string `json:"notificationList,omitempty" tf:"notification_list,omitempty"`

	// Specifies the type of action triggered by an alarm. the value is notification.
	// notification: indicates that a notification will be sent to the user.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InsufficientdataActionsObservation struct {

	// specifies the list of objects to be notified if the alarm status changes, the
	// maximum length is 5.
	NotificationList []*string `json:"notificationList,omitempty" tf:"notification_list,omitempty"`

	// Specifies the type of action triggered by an alarm. the value is notification.
	// notification: indicates that a notification will be sent to the user.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InsufficientdataActionsParameters struct {

	// specifies the list of objects to be notified if the alarm status changes, the
	// maximum length is 5.
	// +kubebuilder:validation:Optional
	NotificationList []*string `json:"notificationList" tf:"notification_list,omitempty"`

	// Specifies the type of action triggered by an alarm. the value is notification.
	// notification: indicates that a notification will be sent to the user.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type MetricInitParameters struct {

	// Specifies the list of metric dimensions. The structure is described below.
	// schema: Deprecated
	Dimensions []DimensionsInitParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Specifies the metric name of the condition. The value can be a string of
	// 1 to 64 characters that must start with a letter and contain only letters, digits, and underscores (_).
	// For details, see Services Interconnected with Cloud Eye.
	// schema: Deprecated
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Specifies the namespace in service.item format. service and item
	// each must be a string that starts with a letter and contains only letters, digits, and underscores (_).
	// Changing this creates a new resource.
	// For details, see Services Interconnected with Cloud Eye.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type MetricObservation struct {

	// Specifies the list of metric dimensions. The structure is described below.
	// schema: Deprecated
	Dimensions []DimensionsObservation `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Specifies the metric name of the condition. The value can be a string of
	// 1 to 64 characters that must start with a letter and contain only letters, digits, and underscores (_).
	// For details, see Services Interconnected with Cloud Eye.
	// schema: Deprecated
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Specifies the namespace in service.item format. service and item
	// each must be a string that starts with a letter and contains only letters, digits, and underscores (_).
	// Changing this creates a new resource.
	// For details, see Services Interconnected with Cloud Eye.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type MetricParameters struct {

	// Specifies the list of metric dimensions. The structure is described below.
	// schema: Deprecated
	// +kubebuilder:validation:Optional
	Dimensions []DimensionsParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Specifies the metric name of the condition. The value can be a string of
	// 1 to 64 characters that must start with a letter and contain only letters, digits, and underscores (_).
	// For details, see Services Interconnected with Cloud Eye.
	// schema: Deprecated
	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Specifies the namespace in service.item format. service and item
	// each must be a string that starts with a letter and contains only letters, digits, and underscores (_).
	// Changing this creates a new resource.
	// For details, see Services Interconnected with Cloud Eye.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

type OkActionsInitParameters struct {

	// specifies the list of objects to be notified if the alarm status changes, the
	// maximum length is 5.
	NotificationList []*string `json:"notificationList,omitempty" tf:"notification_list,omitempty"`

	// Specifies the type of action triggered by an alarm. the value is notification.
	// notification: indicates that a notification will be sent to the user.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OkActionsObservation struct {

	// specifies the list of objects to be notified if the alarm status changes, the
	// maximum length is 5.
	NotificationList []*string `json:"notificationList,omitempty" tf:"notification_list,omitempty"`

	// Specifies the type of action triggered by an alarm. the value is notification.
	// notification: indicates that a notification will be sent to the user.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OkActionsParameters struct {

	// specifies the list of objects to be notified if the alarm status changes, the
	// maximum length is 5.
	// +kubebuilder:validation:Optional
	NotificationList []*string `json:"notificationList" tf:"notification_list,omitempty"`

	// Specifies the type of action triggered by an alarm. the value is notification.
	// notification: indicates that a notification will be sent to the user.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ResourcesDimensionsInitParameters struct {

	// Specifies the dimension name. The value can be a string of 1 to 32 characters
	// that must start with a letter and contain only letters, digits, underscores (_), and hyphens (-).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the dimension value. The value can be a string of 1 to 64 characters
	// that must start with a letter or a number and contain only letters, digits, underscores (_), and hyphens (-).
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResourcesDimensionsObservation struct {

	// Specifies the dimension name. The value can be a string of 1 to 32 characters
	// that must start with a letter and contain only letters, digits, underscores (_), and hyphens (-).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the dimension value. The value can be a string of 1 to 64 characters
	// that must start with a letter or a number and contain only letters, digits, underscores (_), and hyphens (-).
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResourcesDimensionsParameters struct {

	// Specifies the dimension name. The value can be a string of 1 to 32 characters
	// that must start with a letter and contain only letters, digits, underscores (_), and hyphens (-).
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the dimension value. The value can be a string of 1 to 64 characters
	// that must start with a letter or a number and contain only letters, digits, underscores (_), and hyphens (-).
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResourcesInitParameters struct {

	// Specifies the list of metric dimensions. The structure is described below.
	Dimensions []ResourcesDimensionsInitParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`
}

type ResourcesObservation struct {

	// Specifies the list of metric dimensions. The structure is described below.
	Dimensions []ResourcesDimensionsObservation `json:"dimensions,omitempty" tf:"dimensions,omitempty"`
}

type ResourcesParameters struct {

	// Specifies the list of metric dimensions. The structure is described below.
	// +kubebuilder:validation:Optional
	Dimensions []ResourcesDimensionsParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`
}

// AlarmruleSpec defines the desired state of Alarmrule
type AlarmruleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlarmruleParameters `json:"forProvider"`
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
	InitProvider AlarmruleInitParameters `json:"initProvider,omitempty"`
}

// AlarmruleStatus defines the observed state of Alarmrule.
type AlarmruleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlarmruleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Alarmrule is the Schema for the Alarmrules API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type Alarmrule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.alarmName) || (has(self.initProvider) && has(self.initProvider.alarmName))",message="spec.forProvider.alarmName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.condition) || (has(self.initProvider) && has(self.initProvider.condition))",message="spec.forProvider.condition is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metric) || (has(self.initProvider) && has(self.initProvider.metric))",message="spec.forProvider.metric is a required parameter"
	Spec   AlarmruleSpec   `json:"spec"`
	Status AlarmruleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlarmruleList contains a list of Alarmrules
type AlarmruleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alarmrule `json:"items"`
}

// Repository type metadata.
var (
	Alarmrule_Kind             = "Alarmrule"
	Alarmrule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Alarmrule_Kind}.String()
	Alarmrule_KindAPIVersion   = Alarmrule_Kind + "." + CRDGroupVersion.String()
	Alarmrule_GroupVersionKind = CRDGroupVersion.WithKind(Alarmrule_Kind)
)

func init() {
	SchemeBuilder.Register(&Alarmrule{}, &AlarmruleList{})
}
