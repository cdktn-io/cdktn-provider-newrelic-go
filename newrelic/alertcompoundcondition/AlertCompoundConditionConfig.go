// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package alertcompoundcondition

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AlertCompoundConditionConfig struct {
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
	// component_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/alert_compound_condition#component_conditions AlertCompoundCondition#component_conditions}
	ComponentConditions interface{} `field:"required" json:"componentConditions" yaml:"componentConditions"`
	// Whether or not to enable the alert condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/alert_compound_condition#enabled AlertCompoundCondition#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The title of the compound alert condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/alert_compound_condition#name AlertCompoundCondition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the policy where this condition should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/alert_compound_condition#policy_id AlertCompoundCondition#policy_id}
	PolicyId *float64 `field:"required" json:"policyId" yaml:"policyId"`
	// Expression that defines how component condition evaluations are combined.
	//
	// Valid operators are 'AND', 'OR', 'NOT'. For more complex expressions, use parentheses. Simple example: 'A AND B'. Complex example: 'A AND (B OR C) AND NOT D'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/alert_compound_condition#trigger_expression AlertCompoundCondition#trigger_expression}
	TriggerExpression *string `field:"required" json:"triggerExpression" yaml:"triggerExpression"`
	// The New Relic account ID for managing your compound alert conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/alert_compound_condition#account_id AlertCompoundCondition#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// How the compound condition will take into account the component conditions' facets during evaluation.
	//
	// Valid values: 'FACETS_IGNORED' (default) - facets are not taken into consideration when determining when the compound alert condition activates; 'FACETS_MATCH' - the compound alert condition will activate only when shared facets have matching values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/alert_compound_condition#facet_matching_behavior AlertCompoundCondition#facet_matching_behavior}
	FacetMatchingBehavior *string `field:"optional" json:"facetMatchingBehavior" yaml:"facetMatchingBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/alert_compound_condition#id AlertCompoundCondition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Runbook URL to display in notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/alert_compound_condition#runbook_url AlertCompoundCondition#runbook_url}
	RunbookUrl *string `field:"optional" json:"runbookUrl" yaml:"runbookUrl"`
	// The duration, in seconds, that the trigger expression must be true before the compound alert condition will activate.
	//
	// Between 30-86400 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/alert_compound_condition#threshold_duration AlertCompoundCondition#threshold_duration}
	ThresholdDuration *float64 `field:"optional" json:"thresholdDuration" yaml:"thresholdDuration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/alert_compound_condition#timeouts AlertCompoundCondition#timeouts}
	Timeouts *AlertCompoundConditionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

