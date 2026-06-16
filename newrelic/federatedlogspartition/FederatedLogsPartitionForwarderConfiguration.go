// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogspartition


type FederatedLogsPartitionForwarderConfiguration struct {
	// The type of forwarder. Must match the parent setup's forwarder type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/federated_logs_partition#type FederatedLogsPartition#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// pipeline_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/federated_logs_partition#pipeline_control FederatedLogsPartition#pipeline_control}
	PipelineControl *FederatedLogsPartitionForwarderConfigurationPipelineControl `field:"optional" json:"pipelineControl" yaml:"pipelineControl"`
}

