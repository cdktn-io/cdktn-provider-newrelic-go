// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scorecard

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ScorecardConfig struct {
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
	// The name of the scorecard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard#name Scorecard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the scorecard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard#description Scorecard#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard#id Scorecard#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// progress_levels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard#progress_levels Scorecard#progress_levels}
	ProgressLevels interface{} `field:"optional" json:"progressLevels" yaml:"progressLevels"`
	// Set of newrelic_scorecard_rule GUIDs to attach to this scorecard.
	//
	// Rules must exist before being attached. Removing a GUID detaches the rule without deleting it — rules are standalone entities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard#rule_ids Scorecard#rule_ids}
	RuleIds *[]*string `field:"optional" json:"ruleIds" yaml:"ruleIds"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard#tags Scorecard#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard#timeouts Scorecard#timeouts}
	Timeouts *ScorecardTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

