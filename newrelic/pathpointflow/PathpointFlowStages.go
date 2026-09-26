// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow


type PathpointFlowStages struct {
	// Display name of the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#name PathpointFlow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Health rollup strategy: ALERT_CONDITIONS or AUTOMATIC_ROLL_UP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#health_rollup PathpointFlow#health_rollup}
	HealthRollup *string `field:"optional" json:"healthRollup" yaml:"healthRollup"`
	// When true, this stage is excluded from flow health calculation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#is_excluded PathpointFlow#is_excluded}
	IsExcluded interface{} `field:"optional" json:"isExcluded" yaml:"isExcluded"`
	// levels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#levels PathpointFlow#levels}
	Levels interface{} `field:"optional" json:"levels" yaml:"levels"`
	// Optional URL to an external resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#link PathpointFlow#link}
	Link *string `field:"optional" json:"link" yaml:"link"`
	// related block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#related PathpointFlow#related}
	Related *PathpointFlowStagesRelated `field:"optional" json:"related" yaml:"related"`
	// stage_kpis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#stage_kpis PathpointFlow#stage_kpis}
	StageKpis interface{} `field:"optional" json:"stageKpis" yaml:"stageKpis"`
}

