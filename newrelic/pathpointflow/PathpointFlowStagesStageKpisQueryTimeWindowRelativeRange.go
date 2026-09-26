// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow


type PathpointFlowStagesStageKpisQueryTimeWindowRelativeRange struct {
	// How far back the KPI is evaluated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#since PathpointFlow#since}
	Since *string `field:"required" json:"since" yaml:"since"`
	// The earlier window to compare against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#compare_against PathpointFlow#compare_against}
	CompareAgainst *string `field:"optional" json:"compareAgainst" yaml:"compareAgainst"`
}

