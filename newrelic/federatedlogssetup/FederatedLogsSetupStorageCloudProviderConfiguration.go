// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogssetup


type FederatedLogsSetupStorageCloudProviderConfiguration struct {
	// The cloud provider. Currently only AWS is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/federated_logs_setup#provider FederatedLogsSetup#provider}
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// The cloud provider region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/federated_logs_setup#region FederatedLogsSetup#region}
	Region *string `field:"required" json:"region" yaml:"region"`
}

