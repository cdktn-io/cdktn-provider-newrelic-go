// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogssetup


type FederatedLogsSetupForwarder struct {
	// The type of forwarder. Currently only PIPELINE_CONTROL is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/federated_logs_setup#type FederatedLogsSetup#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// pipeline_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/federated_logs_setup#pipeline_control FederatedLogsSetup#pipeline_control}
	PipelineControl *FederatedLogsSetupForwarderPipelineControl `field:"optional" json:"pipelineControl" yaml:"pipelineControl"`
}

