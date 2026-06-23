// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogspartition


type FederatedLogsPartitionDataRetentionPolicy struct {
	// The duration value for retention.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.2/docs/resources/federated_logs_partition#duration FederatedLogsPartition#duration}
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// The time unit for the retention duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.2/docs/resources/federated_logs_partition#unit FederatedLogsPartition#unit}
	Unit *string `field:"required" json:"unit" yaml:"unit"`
}

