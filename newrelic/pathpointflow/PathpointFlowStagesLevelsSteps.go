// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow


type PathpointFlowStagesLevelsSteps struct {
	// Display name of the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#name PathpointFlow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#config PathpointFlow#config}
	Config *PathpointFlowStagesLevelsStepsConfig `field:"optional" json:"config" yaml:"config"`
	// entity_search_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#entity_search_query PathpointFlow#entity_search_query}
	EntitySearchQuery *PathpointFlowStagesLevelsStepsEntitySearchQuery `field:"optional" json:"entitySearchQuery" yaml:"entitySearchQuery"`
	// When true, this step is excluded from level health calculation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#is_excluded PathpointFlow#is_excluded}
	IsExcluded interface{} `field:"optional" json:"isExcluded" yaml:"isExcluded"`
	// Optional URL to an external resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#link PathpointFlow#link}
	Link *string `field:"optional" json:"link" yaml:"link"`
	// Account IDs whose data is scoped to this step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#scoped_accounts PathpointFlow#scoped_accounts}
	ScopedAccounts *[]*float64 `field:"optional" json:"scopedAccounts" yaml:"scopedAccounts"`
	// signals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#signals PathpointFlow#signals}
	Signals interface{} `field:"optional" json:"signals" yaml:"signals"`
}

