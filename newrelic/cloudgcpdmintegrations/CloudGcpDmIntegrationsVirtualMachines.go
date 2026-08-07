// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudgcpdmintegrations


type CloudGcpDmIntegrationsVirtualMachines struct {
	// The data polling interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.96.1/docs/resources/cloud_gcp_dm_integrations#metrics_polling_interval CloudGcpDmIntegrations#metrics_polling_interval}
	MetricsPollingInterval *float64 `field:"optional" json:"metricsPollingInterval" yaml:"metricsPollingInterval"`
}

