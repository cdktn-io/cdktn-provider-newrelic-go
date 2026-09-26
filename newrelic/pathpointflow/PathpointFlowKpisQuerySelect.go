// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow


type PathpointFlowKpisQuerySelect struct {
	// Aggregation function: AVERAGE, COUNT, HISTOGRAM, MAX, MIN, PERCENTILE, SUM, UNIQUE_COUNT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#aggregation_type PathpointFlow#aggregation_type}
	AggregationType *string `field:"required" json:"aggregationType" yaml:"aggregationType"`
	// Optional alias for the aggregated value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#alias PathpointFlow#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Attribute name to aggregate. Required for all functions except COUNT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#attribute PathpointFlow#attribute}
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// Threshold used in the selected function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#threshold PathpointFlow#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
}

