// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scorecardrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ScorecardRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Whether this rule is active. Set to false to pause evaluation without deleting the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#enabled ScorecardRule#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the scorecard rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#name ScorecardRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// nrql_engine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#nrql_engine ScorecardRule#nrql_engine}
	NrqlEngine *ScorecardRuleNrqlEngine `field:"required" json:"nrqlEngine" yaml:"nrqlEngine"`
	// A description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#description ScorecardRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#id ScorecardRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Relative weight of this rule's contribution to the overall scorecard score.
	//
	// Must be a non-negative integer. Omit for equal weighting across all rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#impact_weight ScorecardRule#impact_weight}
	ImpactWeight *float64 `field:"optional" json:"impactWeight" yaml:"impactWeight"`
	// The progress level ID from the parent scorecard this rule maps to (e.g. 'red'). Must match a progress_levels.id defined on the scorecard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#progress_level ScorecardRule#progress_level}
	ProgressLevel *string `field:"optional" json:"progressLevel" yaml:"progressLevel"`
	// Evaluation frequency in minutes.
	//
	// Must be one of: 60 (hourly), 360 (6h), 720 (12h), 1440 (daily). Omit to use the API default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#run_interval ScorecardRule#run_interval}
	RunInterval *float64 `field:"optional" json:"runInterval" yaml:"runInterval"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#tags ScorecardRule#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#timeouts ScorecardRule#timeouts}
	Timeouts *ScorecardRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

