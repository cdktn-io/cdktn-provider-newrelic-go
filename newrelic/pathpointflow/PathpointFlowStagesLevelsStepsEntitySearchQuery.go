// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow


type PathpointFlowStagesLevelsStepsEntitySearchQuery struct {
	// Filter query for signals, e.g. domain='NR1' AND type='APPLICATION'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#query PathpointFlow#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// When true, this query is excluded from health calculation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#is_excluded PathpointFlow#is_excluded}
	IsExcluded interface{} `field:"optional" json:"isExcluded" yaml:"isExcluded"`
}

