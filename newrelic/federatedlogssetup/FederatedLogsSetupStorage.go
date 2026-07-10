// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogssetup


type FederatedLogsSetupStorage struct {
	// cloud_provider_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/federated_logs_setup#cloud_provider_configuration FederatedLogsSetup#cloud_provider_configuration}
	CloudProviderConfiguration *FederatedLogsSetupStorageCloudProviderConfiguration `field:"required" json:"cloudProviderConfiguration" yaml:"cloudProviderConfiguration"`
	// The database name associated with the federated log setup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/federated_logs_setup#database FederatedLogsSetup#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// The connection manager entity GUID used for writing data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/federated_logs_setup#data_ingest_connection_id FederatedLogsSetup#data_ingest_connection_id}
	DataIngestConnectionId *string `field:"required" json:"dataIngestConnectionId" yaml:"dataIngestConnectionId"`
	// The object storage bucket where log data is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/federated_logs_setup#data_location_bucket FederatedLogsSetup#data_location_bucket}
	DataLocationBucket *string `field:"required" json:"dataLocationBucket" yaml:"dataLocationBucket"`
	// The connection manager entity GUID used by query workers for reading data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/federated_logs_setup#query_connection_id FederatedLogsSetup#query_connection_id}
	QueryConnectionId *string `field:"required" json:"queryConnectionId" yaml:"queryConnectionId"`
}

