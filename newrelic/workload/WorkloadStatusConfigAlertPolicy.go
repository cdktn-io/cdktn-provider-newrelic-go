// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workload


type WorkloadStatusConfigAlertPolicy struct {
	// Whether the alert policy status configuration is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/workload#enabled Workload#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

