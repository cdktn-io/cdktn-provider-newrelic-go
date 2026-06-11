// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cardinalitymanagement


type CardinalityManagementMetric struct {
	// The maximum number of unique dimension-value combinations allowed per day for this metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.0/docs/resources/cardinality_management#cardinality_limit CardinalityManagement#cardinality_limit}
	CardinalityLimit *float64 `field:"required" json:"cardinalityLimit" yaml:"cardinalityLimit"`
	// The full name of the metric (e.g. `http.server.duration`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.0/docs/resources/cardinality_management#name CardinalityManagement#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

