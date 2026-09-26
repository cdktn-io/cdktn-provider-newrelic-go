// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow


type PathpointFlowStagesStageKpis struct {
	// Display name of the KPI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#name PathpointFlow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#query PathpointFlow#query}
	Query *PathpointFlowStagesStageKpisQuery `field:"required" json:"query" yaml:"query"`
	// Account ID this KPI belongs to. Defaults to the flow's account_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#account_id PathpointFlow#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Optional category to group KPIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#category PathpointFlow#category}
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Optional description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#description PathpointFlow#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

