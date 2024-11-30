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

type DependenciesInitParameters struct {

	// Installation mode. Only pip is supported.
	// Changing this parameter will create a new resource.
	// Installation mode. Only **pip** is supported.
	Installer *string `json:"installer,omitempty" tf:"installer,omitempty"`

	// Collection of dependency packages.
	// The package structure is documented below.
	// Changing this parameter will create a new resource.
	// Collection of dependency packages.
	Packages []PackagesInitParameters `json:"packages,omitempty" tf:"packages,omitempty"`
}

type DependenciesObservation struct {

	// Installation mode. Only pip is supported.
	// Changing this parameter will create a new resource.
	// Installation mode. Only **pip** is supported.
	Installer *string `json:"installer,omitempty" tf:"installer,omitempty"`

	// Collection of dependency packages.
	// The package structure is documented below.
	// Changing this parameter will create a new resource.
	// Collection of dependency packages.
	Packages []PackagesObservation `json:"packages,omitempty" tf:"packages,omitempty"`
}

type DependenciesParameters struct {

	// Installation mode. Only pip is supported.
	// Changing this parameter will create a new resource.
	// Installation mode. Only **pip** is supported.
	// +kubebuilder:validation:Optional
	Installer *string `json:"installer" tf:"installer,omitempty"`

	// Collection of dependency packages.
	// The package structure is documented below.
	// Changing this parameter will create a new resource.
	// Collection of dependency packages.
	// +kubebuilder:validation:Optional
	Packages []PackagesParameters `json:"packages" tf:"packages,omitempty"`
}

type ModelDocsInitParameters struct {

	// Document name, which must start with a letter. Enter 1 to 48 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	// Changing this parameter will create a new resource.
	// Document name, which must start with a letter. Enter 1 to 48 characters.
	DocName *string `json:"docName,omitempty" tf:"doc_name,omitempty"`

	// HTTP(S) link of the document.
	// Changing this parameter will create a new resource.
	// HTTP(S) link of the document.
	DocURL *string `json:"docUrl,omitempty" tf:"doc_url,omitempty"`
}

type ModelDocsObservation struct {

	// Document name, which must start with a letter. Enter 1 to 48 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	// Changing this parameter will create a new resource.
	// Document name, which must start with a letter. Enter 1 to 48 characters.
	DocName *string `json:"docName,omitempty" tf:"doc_name,omitempty"`

	// HTTP(S) link of the document.
	// Changing this parameter will create a new resource.
	// HTTP(S) link of the document.
	DocURL *string `json:"docUrl,omitempty" tf:"doc_url,omitempty"`
}

type ModelDocsParameters struct {

	// Document name, which must start with a letter. Enter 1 to 48 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	// Changing this parameter will create a new resource.
	// Document name, which must start with a letter. Enter 1 to 48 characters.
	// +kubebuilder:validation:Optional
	DocName *string `json:"docName,omitempty" tf:"doc_name,omitempty"`

	// HTTP(S) link of the document.
	// Changing this parameter will create a new resource.
	// HTTP(S) link of the document.
	// +kubebuilder:validation:Optional
	DocURL *string `json:"docUrl,omitempty" tf:"doc_url,omitempty"`
}

type ModelInitParameters struct {

	// Package required for inference code and model.
	// If the package is read from the configuration file, this parameter can be left blank.
	// The Dependency structure is documented below.
	// Package required for inference code and model.
	Dependencies []DependenciesInitParameters `json:"dependencies,omitempty" tf:"dependencies,omitempty"`

	// Model description that consists of 1 to 100 characters.
	// The following special characters cannot be contained: &!'"<>=.
	// Model description that consists of 1 to 100 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// OBS path for storing the execution code.
	// The name of the execution code file is consistently to be customize_service.py.
	// The inference code file must be stored in the model directory.
	// This parameter can be left blank. Then, the system will automatically identify the inference
	// code in the model directory.
	// OBS path for storing the execution code.
	ExecutionCode *string `json:"executionCode,omitempty" tf:"execution_code,omitempty"`

	// The model configuration file describes the model usage,
	// computing framework, precision, inference code dependency package, and model API.
	// The fields such as model_algorithm, model_type, runtime, swr_location, metrics, apis,
	// dependencies, and health in the configuration file config.json.
	// For details, see Specifications for Writing the Model Configuration File
	// The model configuration file.
	InitialConfig *string `json:"initialConfig,omitempty" tf:"initial_config,omitempty"`

	// Deployment type. Only lowercase letters are supported.
	// The value can be real-time, edge, or batch. Default value: [real-time, edge, batch].
	// Deployment type. Only lowercase letters are supported.
	InstallType []*string `json:"installType,omitempty" tf:"install_type,omitempty"`

	// Model precision.
	// If the value is read from the configuration file, this parameter can be left blank.
	// Model precision.
	Metrics *string `json:"metrics,omitempty" tf:"metrics,omitempty"`

	// Model algorithm.
	// If the algorithm is read from the configuration file, this parameter can be left blank.
	// The value can be predict_analysis, object_detection, image_classification, or unknown_algorithm.
	// Model algorithm.
	ModelAlgorithm *string `json:"modelAlgorithm,omitempty" tf:"model_algorithm,omitempty"`

	// List of model description documents. A maximum of three documents are supported.
	// List of model description documents. A maximum of three documents are supported.
	ModelDocs []ModelDocsInitParameters `json:"modelDocs,omitempty" tf:"model_docs,omitempty"`

	// Model type.
	// It can be TensorFlow, MXNet, Caffe, Spark_MLlib, Scikit_Learn,
	// XGBoost, Image, PyTorch, or Template read from the configuration file.
	// Model type.
	ModelType *string `json:"modelType,omitempty" tf:"model_type,omitempty"`

	// Model name, which consists of 1 to 64 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	// Model name, which consists of 1 to 64 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Model runtime environment.
	// Its possible values are determined based on model_type.
	// For details, see Supported AI engines and their runtime
	// Model runtime environment.
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Whether to enable image replication.
	// This parameter is valid only when model_type is set to Image.
	// Value options are as follows:
	// Whether to enable image replication.
	SourceCopy *string `json:"sourceCopy,omitempty" tf:"source_copy,omitempty"`

	// ID of the source training job.
	// If the model is generated from a training job, input this parameter for source tracing.
	// If the model is imported from a third-party meta model, leave this parameter blank.
	// ID of the source training job.
	SourceJobID *string `json:"sourceJobId,omitempty" tf:"source_job_id,omitempty"`

	// Version of the source training job.
	// If the model is generated from a training job, input this parameter for source tracing.
	// If the model is imported from a third-party meta model, leave this parameter blank.
	// Version of the source training job.
	SourceJobVersion *string `json:"sourceJobVersion,omitempty" tf:"source_job_version,omitempty"`

	// OBS path where the model is located or the SWR image location.
	// OBS path where the model is located or the SWR image location.
	SourceLocation *string `json:"sourceLocation,omitempty" tf:"source_location,omitempty"`

	// Model source type, which can only be auto,
	// indicating an ExeML model (model download is not allowed).
	// If the model is obtained from a training job, leave this parameter blank.
	// Model source type
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Configuration items in a template.
	// This parameter is mandatory when model_type is set to Template.
	Template []TemplateInitParameters `json:"template,omitempty" tf:"template,omitempty"`

	// Model version in the format of Digit.Digit.Digit.
	// Each digit is a one-digit or two-digit positive integer, but cannot start with 0.
	// For example, 01.01.01 is not allowed.
	// Model version in the format of Digit.Digit.Digit.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Workspace ID, which defaults to 0.
	// Workspace ID, which defaults to 0.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type ModelObservation struct {

	// Package required for inference code and model.
	// If the package is read from the configuration file, this parameter can be left blank.
	// The Dependency structure is documented below.
	// Package required for inference code and model.
	Dependencies []DependenciesObservation `json:"dependencies,omitempty" tf:"dependencies,omitempty"`

	// Model description that consists of 1 to 100 characters.
	// The following special characters cannot be contained: &!'"<>=.
	// Model description that consists of 1 to 100 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// OBS path for storing the execution code.
	// The name of the execution code file is consistently to be customize_service.py.
	// The inference code file must be stored in the model directory.
	// This parameter can be left blank. Then, the system will automatically identify the inference
	// code in the model directory.
	// OBS path for storing the execution code.
	ExecutionCode *string `json:"executionCode,omitempty" tf:"execution_code,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Image path generated after model packaging.
	// Image path generated after model packaging.
	ImageAddress *string `json:"imageAddress,omitempty" tf:"image_address,omitempty"`

	// The model configuration file describes the model usage,
	// computing framework, precision, inference code dependency package, and model API.
	// The fields such as model_algorithm, model_type, runtime, swr_location, metrics, apis,
	// dependencies, and health in the configuration file config.json.
	// For details, see Specifications for Writing the Model Configuration File
	// The model configuration file.
	InitialConfig *string `json:"initialConfig,omitempty" tf:"initial_config,omitempty"`

	// Deployment type. Only lowercase letters are supported.
	// The value can be real-time, edge, or batch. Default value: [real-time, edge, batch].
	// Deployment type. Only lowercase letters are supported.
	InstallType []*string `json:"installType,omitempty" tf:"install_type,omitempty"`

	// Whether a model is subscribed from AI Gallery.
	// Whether a model is subscribed from AI Gallery.
	MarketFlag *bool `json:"marketFlag,omitempty" tf:"market_flag,omitempty"`

	// Model precision.
	// If the value is read from the configuration file, this parameter can be left blank.
	// Model precision.
	Metrics *string `json:"metrics,omitempty" tf:"metrics,omitempty"`

	// Model algorithm.
	// If the algorithm is read from the configuration file, this parameter can be left blank.
	// The value can be predict_analysis, object_detection, image_classification, or unknown_algorithm.
	// Model algorithm.
	ModelAlgorithm *string `json:"modelAlgorithm,omitempty" tf:"model_algorithm,omitempty"`

	// List of model description documents. A maximum of three documents are supported.
	// List of model description documents. A maximum of three documents are supported.
	ModelDocs []ModelDocsObservation `json:"modelDocs,omitempty" tf:"model_docs,omitempty"`

	// Model size, in bytes.
	// Model size, in bytes.
	ModelSize *float64 `json:"modelSize,omitempty" tf:"model_size,omitempty"`

	// Model source.
	// Value options are as follows:
	// Model source.
	ModelSource *string `json:"modelSource,omitempty" tf:"model_source,omitempty"`

	// Model type.
	// It can be TensorFlow, MXNet, Caffe, Spark_MLlib, Scikit_Learn,
	// XGBoost, Image, PyTorch, or Template read from the configuration file.
	// Model type.
	ModelType *string `json:"modelType,omitempty" tf:"model_type,omitempty"`

	// Model name, which consists of 1 to 64 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	// Model name, which consists of 1 to 64 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether a model is subscribed from AI Gallery.
	// Whether a model is subscribed from AI Gallery.
	PublishableFlag *bool `json:"publishableFlag,omitempty" tf:"publishable_flag,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Model runtime environment.
	// Its possible values are determined based on model_type.
	// For details, see Supported AI engines and their runtime
	// Model runtime environment.
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Download address of the model schema file.
	// Download address of the model schema file.
	SchemaDoc *string `json:"schemaDoc,omitempty" tf:"schema_doc,omitempty"`

	// Whether to enable image replication.
	// This parameter is valid only when model_type is set to Image.
	// Value options are as follows:
	// Whether to enable image replication.
	SourceCopy *string `json:"sourceCopy,omitempty" tf:"source_copy,omitempty"`

	// ID of the source training job.
	// If the model is generated from a training job, input this parameter for source tracing.
	// If the model is imported from a third-party meta model, leave this parameter blank.
	// ID of the source training job.
	SourceJobID *string `json:"sourceJobId,omitempty" tf:"source_job_id,omitempty"`

	// Version of the source training job.
	// If the model is generated from a training job, input this parameter for source tracing.
	// If the model is imported from a third-party meta model, leave this parameter blank.
	// Version of the source training job.
	SourceJobVersion *string `json:"sourceJobVersion,omitempty" tf:"source_job_version,omitempty"`

	// OBS path where the model is located or the SWR image location.
	// OBS path where the model is located or the SWR image location.
	SourceLocation *string `json:"sourceLocation,omitempty" tf:"source_location,omitempty"`

	// Model source type, which can only be auto,
	// indicating an ExeML model (model download is not allowed).
	// If the model is obtained from a training job, leave this parameter blank.
	// Model source type
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Model status.
	// Model status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Configuration items in a template.
	// This parameter is mandatory when model_type is set to Template.
	Template []TemplateObservation `json:"template,omitempty" tf:"template,omitempty"`

	// Whether a model can be tuned.
	// Whether a model can be tuned.
	Tunable *bool `json:"tunable,omitempty" tf:"tunable,omitempty"`

	// Model version in the format of Digit.Digit.Digit.
	// Each digit is a one-digit or two-digit positive integer, but cannot start with 0.
	// For example, 01.01.01 is not allowed.
	// Model version in the format of Digit.Digit.Digit.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Workspace ID, which defaults to 0.
	// Workspace ID, which defaults to 0.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type ModelParameters struct {

	// Package required for inference code and model.
	// If the package is read from the configuration file, this parameter can be left blank.
	// The Dependency structure is documented below.
	// Package required for inference code and model.
	// +kubebuilder:validation:Optional
	Dependencies []DependenciesParameters `json:"dependencies,omitempty" tf:"dependencies,omitempty"`

	// Model description that consists of 1 to 100 characters.
	// The following special characters cannot be contained: &!'"<>=.
	// Model description that consists of 1 to 100 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// OBS path for storing the execution code.
	// The name of the execution code file is consistently to be customize_service.py.
	// The inference code file must be stored in the model directory.
	// This parameter can be left blank. Then, the system will automatically identify the inference
	// code in the model directory.
	// OBS path for storing the execution code.
	// +kubebuilder:validation:Optional
	ExecutionCode *string `json:"executionCode,omitempty" tf:"execution_code,omitempty"`

	// The model configuration file describes the model usage,
	// computing framework, precision, inference code dependency package, and model API.
	// The fields such as model_algorithm, model_type, runtime, swr_location, metrics, apis,
	// dependencies, and health in the configuration file config.json.
	// For details, see Specifications for Writing the Model Configuration File
	// The model configuration file.
	// +kubebuilder:validation:Optional
	InitialConfig *string `json:"initialConfig,omitempty" tf:"initial_config,omitempty"`

	// Deployment type. Only lowercase letters are supported.
	// The value can be real-time, edge, or batch. Default value: [real-time, edge, batch].
	// Deployment type. Only lowercase letters are supported.
	// +kubebuilder:validation:Optional
	InstallType []*string `json:"installType,omitempty" tf:"install_type,omitempty"`

	// Model precision.
	// If the value is read from the configuration file, this parameter can be left blank.
	// Model precision.
	// +kubebuilder:validation:Optional
	Metrics *string `json:"metrics,omitempty" tf:"metrics,omitempty"`

	// Model algorithm.
	// If the algorithm is read from the configuration file, this parameter can be left blank.
	// The value can be predict_analysis, object_detection, image_classification, or unknown_algorithm.
	// Model algorithm.
	// +kubebuilder:validation:Optional
	ModelAlgorithm *string `json:"modelAlgorithm,omitempty" tf:"model_algorithm,omitempty"`

	// List of model description documents. A maximum of three documents are supported.
	// List of model description documents. A maximum of three documents are supported.
	// +kubebuilder:validation:Optional
	ModelDocs []ModelDocsParameters `json:"modelDocs,omitempty" tf:"model_docs,omitempty"`

	// Model type.
	// It can be TensorFlow, MXNet, Caffe, Spark_MLlib, Scikit_Learn,
	// XGBoost, Image, PyTorch, or Template read from the configuration file.
	// Model type.
	// +kubebuilder:validation:Optional
	ModelType *string `json:"modelType,omitempty" tf:"model_type,omitempty"`

	// Model name, which consists of 1 to 64 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	// Model name, which consists of 1 to 64 characters.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Model runtime environment.
	// Its possible values are determined based on model_type.
	// For details, see Supported AI engines and their runtime
	// Model runtime environment.
	// +kubebuilder:validation:Optional
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Whether to enable image replication.
	// This parameter is valid only when model_type is set to Image.
	// Value options are as follows:
	// Whether to enable image replication.
	// +kubebuilder:validation:Optional
	SourceCopy *string `json:"sourceCopy,omitempty" tf:"source_copy,omitempty"`

	// ID of the source training job.
	// If the model is generated from a training job, input this parameter for source tracing.
	// If the model is imported from a third-party meta model, leave this parameter blank.
	// ID of the source training job.
	// +kubebuilder:validation:Optional
	SourceJobID *string `json:"sourceJobId,omitempty" tf:"source_job_id,omitempty"`

	// Version of the source training job.
	// If the model is generated from a training job, input this parameter for source tracing.
	// If the model is imported from a third-party meta model, leave this parameter blank.
	// Version of the source training job.
	// +kubebuilder:validation:Optional
	SourceJobVersion *string `json:"sourceJobVersion,omitempty" tf:"source_job_version,omitempty"`

	// OBS path where the model is located or the SWR image location.
	// OBS path where the model is located or the SWR image location.
	// +kubebuilder:validation:Optional
	SourceLocation *string `json:"sourceLocation,omitempty" tf:"source_location,omitempty"`

	// Model source type, which can only be auto,
	// indicating an ExeML model (model download is not allowed).
	// If the model is obtained from a training job, leave this parameter blank.
	// Model source type
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Configuration items in a template.
	// This parameter is mandatory when model_type is set to Template.
	// +kubebuilder:validation:Optional
	Template []TemplateParameters `json:"template,omitempty" tf:"template,omitempty"`

	// Model version in the format of Digit.Digit.Digit.
	// Each digit is a one-digit or two-digit positive integer, but cannot start with 0.
	// For example, 01.01.01 is not allowed.
	// Model version in the format of Digit.Digit.Digit.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Workspace ID, which defaults to 0.
	// Workspace ID, which defaults to 0.
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type PackagesInitParameters struct {

	// Name of a dependency package.
	// Ensure that the package name is correct and available.
	// Chinese characters and special characters (&!'"<>=) are not allowed.
	// Changing this parameter will create a new resource.
	// Name of a dependency package.
	PackageName *string `json:"packageName,omitempty" tf:"package_name,omitempty"`

	// Version of a dependency package.
	// If this parameter is left blank, the latest version is installed by default.
	// Chinese characters and special characters (&!'"<>=) are not allowed.
	// Changing this parameter will create a new resource.
	// Version of a dependency package.
	PackageVersion *string `json:"packageVersion,omitempty" tf:"package_version,omitempty"`

	// Version restriction, which can be EXACT, ATLEAST, or ATMOST.
	// This parameter is mandatory only when package_version is available.
	// Changing this parameter will create a new resource.
	// Version restriction, which can be **EXACT**, **ATLEAST**, or **ATMOST**.
	Restraint *string `json:"restraint,omitempty" tf:"restraint,omitempty"`
}

type PackagesObservation struct {

	// Name of a dependency package.
	// Ensure that the package name is correct and available.
	// Chinese characters and special characters (&!'"<>=) are not allowed.
	// Changing this parameter will create a new resource.
	// Name of a dependency package.
	PackageName *string `json:"packageName,omitempty" tf:"package_name,omitempty"`

	// Version of a dependency package.
	// If this parameter is left blank, the latest version is installed by default.
	// Chinese characters and special characters (&!'"<>=) are not allowed.
	// Changing this parameter will create a new resource.
	// Version of a dependency package.
	PackageVersion *string `json:"packageVersion,omitempty" tf:"package_version,omitempty"`

	// Version restriction, which can be EXACT, ATLEAST, or ATMOST.
	// This parameter is mandatory only when package_version is available.
	// Changing this parameter will create a new resource.
	// Version restriction, which can be **EXACT**, **ATLEAST**, or **ATMOST**.
	Restraint *string `json:"restraint,omitempty" tf:"restraint,omitempty"`
}

type PackagesParameters struct {

	// Name of a dependency package.
	// Ensure that the package name is correct and available.
	// Chinese characters and special characters (&!'"<>=) are not allowed.
	// Changing this parameter will create a new resource.
	// Name of a dependency package.
	// +kubebuilder:validation:Optional
	PackageName *string `json:"packageName" tf:"package_name,omitempty"`

	// Version of a dependency package.
	// If this parameter is left blank, the latest version is installed by default.
	// Chinese characters and special characters (&!'"<>=) are not allowed.
	// Changing this parameter will create a new resource.
	// Version of a dependency package.
	// +kubebuilder:validation:Optional
	PackageVersion *string `json:"packageVersion,omitempty" tf:"package_version,omitempty"`

	// Version restriction, which can be EXACT, ATLEAST, or ATMOST.
	// This parameter is mandatory only when package_version is available.
	// Changing this parameter will create a new resource.
	// Version restriction, which can be **EXACT**, **ATLEAST**, or **ATMOST**.
	// +kubebuilder:validation:Optional
	Restraint *string `json:"restraint,omitempty" tf:"restraint,omitempty"`
}

type TemplateInitParameters struct {

	// ID of the input and output mode.
	// When this parameter is used, the input and output mode built in the template does not take effect.
	// Changing this parameter will create a new resource.
	// ID of the input and output mode.
	InferFormat *string `json:"inferFormat,omitempty" tf:"infer_format,omitempty"`

	// ID of the used template.
	// The template has a built-in input and output mode.
	// Changing this parameter will create a new resource.
	// ID of the used template.
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`

	// Template input configuration,
	// specifying the source path for configuring a model.
	// The TemplateInput structure is documented below.
	// Changing this parameter will create a new resource.
	// Template input configuration, specifying the source path for configuring a model.
	TemplateInputs []TemplateInputsInitParameters `json:"templateInputs,omitempty" tf:"template_inputs,omitempty"`
}

type TemplateInputsInitParameters struct {

	// Template input path, which can be a path to an OBS file or directory.
	// When you use a template with multiple input items to create a model,
	// if the target paths input_properties specified in the template are the same,
	// the OBS directory or OBS file name entered here must be unique to prevent files from being overwritten.
	// Changing this parameter will create a new resource.
	// Template input path, which can be a path to an OBS file or directory.
	Input *string `json:"input,omitempty" tf:"input,omitempty"`

	// Input item ID, which is obtained from template details.
	// Changing this parameter will create a new resource.
	// Input item ID, which is obtained from template details.
	InputID *string `json:"inputId,omitempty" tf:"input_id,omitempty"`
}

type TemplateInputsObservation struct {

	// Template input path, which can be a path to an OBS file or directory.
	// When you use a template with multiple input items to create a model,
	// if the target paths input_properties specified in the template are the same,
	// the OBS directory or OBS file name entered here must be unique to prevent files from being overwritten.
	// Changing this parameter will create a new resource.
	// Template input path, which can be a path to an OBS file or directory.
	Input *string `json:"input,omitempty" tf:"input,omitempty"`

	// Input item ID, which is obtained from template details.
	// Changing this parameter will create a new resource.
	// Input item ID, which is obtained from template details.
	InputID *string `json:"inputId,omitempty" tf:"input_id,omitempty"`
}

type TemplateInputsParameters struct {

	// Template input path, which can be a path to an OBS file or directory.
	// When you use a template with multiple input items to create a model,
	// if the target paths input_properties specified in the template are the same,
	// the OBS directory or OBS file name entered here must be unique to prevent files from being overwritten.
	// Changing this parameter will create a new resource.
	// Template input path, which can be a path to an OBS file or directory.
	// +kubebuilder:validation:Optional
	Input *string `json:"input" tf:"input,omitempty"`

	// Input item ID, which is obtained from template details.
	// Changing this parameter will create a new resource.
	// Input item ID, which is obtained from template details.
	// +kubebuilder:validation:Optional
	InputID *string `json:"inputId" tf:"input_id,omitempty"`
}

type TemplateObservation struct {

	// ID of the input and output mode.
	// When this parameter is used, the input and output mode built in the template does not take effect.
	// Changing this parameter will create a new resource.
	// ID of the input and output mode.
	InferFormat *string `json:"inferFormat,omitempty" tf:"infer_format,omitempty"`

	// ID of the used template.
	// The template has a built-in input and output mode.
	// Changing this parameter will create a new resource.
	// ID of the used template.
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`

	// Template input configuration,
	// specifying the source path for configuring a model.
	// The TemplateInput structure is documented below.
	// Changing this parameter will create a new resource.
	// Template input configuration, specifying the source path for configuring a model.
	TemplateInputs []TemplateInputsObservation `json:"templateInputs,omitempty" tf:"template_inputs,omitempty"`
}

type TemplateParameters struct {

	// ID of the input and output mode.
	// When this parameter is used, the input and output mode built in the template does not take effect.
	// Changing this parameter will create a new resource.
	// ID of the input and output mode.
	// +kubebuilder:validation:Optional
	InferFormat *string `json:"inferFormat,omitempty" tf:"infer_format,omitempty"`

	// ID of the used template.
	// The template has a built-in input and output mode.
	// Changing this parameter will create a new resource.
	// ID of the used template.
	// +kubebuilder:validation:Optional
	TemplateID *string `json:"templateId" tf:"template_id,omitempty"`

	// Template input configuration,
	// specifying the source path for configuring a model.
	// The TemplateInput structure is documented below.
	// Changing this parameter will create a new resource.
	// Template input configuration, specifying the source path for configuring a model.
	// +kubebuilder:validation:Optional
	TemplateInputs []TemplateInputsParameters `json:"templateInputs" tf:"template_inputs,omitempty"`
}

// ModelSpec defines the desired state of Model
type ModelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ModelParameters `json:"forProvider"`
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
	InitProvider ModelInitParameters `json:"initProvider,omitempty"`
}

// ModelStatus defines the observed state of Model.
type ModelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ModelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Model is the Schema for the Models API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type Model struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.modelType) || (has(self.initProvider) && has(self.initProvider.modelType))",message="spec.forProvider.modelType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sourceLocation) || (has(self.initProvider) && has(self.initProvider.sourceLocation))",message="spec.forProvider.sourceLocation is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version) || (has(self.initProvider) && has(self.initProvider.version))",message="spec.forProvider.version is a required parameter"
	Spec   ModelSpec   `json:"spec"`
	Status ModelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ModelList contains a list of Models
type ModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Model `json:"items"`
}

// Repository type metadata.
var (
	Model_Kind             = "Model"
	Model_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Model_Kind}.String()
	Model_KindAPIVersion   = Model_Kind + "." + CRDGroupVersion.String()
	Model_GroupVersionKind = CRDGroupVersion.WithKind(Model_Kind)
)

func init() {
	SchemeBuilder.Register(&Model{}, &ModelList{})
}