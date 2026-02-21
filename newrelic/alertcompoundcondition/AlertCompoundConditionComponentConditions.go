// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package alertcompoundcondition


type AlertCompoundConditionComponentConditions struct {
	// The identifier that will be used in the compound alert condition's trigger_expression (e.g., 'A', 'B', 'C').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/alert_compound_condition#alias AlertCompoundCondition#alias}
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The ID of the existing alert condition to use as a component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/alert_compound_condition#id AlertCompoundCondition#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
}

