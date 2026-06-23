// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogssetup


type FederatedLogsSetupDefaultPartitionStorage struct {
	// The URI location of the partition in object storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.2/docs/resources/federated_logs_setup#data_location_uri FederatedLogsSetup#data_location_uri}
	DataLocationUri *string `field:"required" json:"dataLocationUri" yaml:"dataLocationUri"`
	// The table name associated with the default partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.2/docs/resources/federated_logs_setup#table FederatedLogsSetup#table}
	Table *string `field:"required" json:"table" yaml:"table"`
}

