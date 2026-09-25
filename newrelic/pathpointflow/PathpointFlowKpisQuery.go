// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow


type PathpointFlowKpisQuery struct {
	// Data source to query from (e.g., Transaction, Metric, Log).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#from PathpointFlow#from}
	From *string `field:"required" json:"from" yaml:"from"`
	// select block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#select PathpointFlow#select}
	Select *PathpointFlowKpisQuerySelect `field:"required" json:"select" yaml:"select"`
	// time_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#time_window PathpointFlow#time_window}
	TimeWindow *PathpointFlowKpisQueryTimeWindow `field:"optional" json:"timeWindow" yaml:"timeWindow"`
	// Optional WHERE clause to filter data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#where PathpointFlow#where}
	Where *string `field:"optional" json:"where" yaml:"where"`
}

