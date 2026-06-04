// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogspartition


type FederatedLogsPartitionForwarderConfigurationPipelineControl struct {
	// partition_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.91.0/docs/resources/federated_logs_partition#partition_rule FederatedLogsPartition#partition_rule}
	PartitionRule *FederatedLogsPartitionForwarderConfigurationPipelineControlPartitionRule `field:"optional" json:"partitionRule" yaml:"partitionRule"`
}

