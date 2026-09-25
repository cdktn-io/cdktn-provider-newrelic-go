// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow


type PathpointFlowStagesLevelsStepsSignals struct {
	// Entity GUID of the signal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#guid PathpointFlow#guid}
	Guid *string `field:"required" json:"guid" yaml:"guid"`
	// When true, this signal is excluded from step health calculation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#is_excluded PathpointFlow#is_excluded}
	IsExcluded interface{} `field:"optional" json:"isExcluded" yaml:"isExcluded"`
	// Display name of the signal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#name PathpointFlow#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether this GUID belongs to an entity or an alert condition: ENTITY or ALERT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#type PathpointFlow#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

