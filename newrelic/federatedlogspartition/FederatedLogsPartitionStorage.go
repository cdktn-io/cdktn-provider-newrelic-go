// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogspartition


type FederatedLogsPartitionStorage struct {
	// The URI location of the partition in object storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.91.0/docs/resources/federated_logs_partition#data_location_uri FederatedLogsPartition#data_location_uri}
	DataLocationUri *string `field:"required" json:"dataLocationUri" yaml:"dataLocationUri"`
	// The table name associated with the partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.91.0/docs/resources/federated_logs_partition#table FederatedLogsPartition#table}
	Table *string `field:"required" json:"table" yaml:"table"`
}

