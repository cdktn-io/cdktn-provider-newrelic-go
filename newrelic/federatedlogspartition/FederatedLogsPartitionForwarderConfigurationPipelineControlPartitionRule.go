// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogspartition


type FederatedLogsPartitionForwarderConfigurationPipelineControlPartitionRule struct {
	// OTTL expression for routing logs to this partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/federated_logs_partition#expression FederatedLogsPartition#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

