// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogssetup


type FederatedLogsSetupDefaultPartition struct {
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.2/docs/resources/federated_logs_setup#storage FederatedLogsSetup#storage}
	Storage *FederatedLogsSetupDefaultPartitionStorage `field:"required" json:"storage" yaml:"storage"`
	// data_retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.2/docs/resources/federated_logs_setup#data_retention_policy FederatedLogsSetup#data_retention_policy}
	DataRetentionPolicy *FederatedLogsSetupDefaultPartitionDataRetentionPolicy `field:"optional" json:"dataRetentionPolicy" yaml:"dataRetentionPolicy"`
}

