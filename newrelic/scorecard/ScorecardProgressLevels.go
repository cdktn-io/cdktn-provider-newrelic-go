// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scorecard


type ScorecardProgressLevels struct {
	// A unique identifier for this level (e.g. 'red', 'green').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard#id Scorecard#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Display name of the level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard#name Scorecard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of what this level means.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard#description Scorecard#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Hex colour for this level, e.g. '#FF0000'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard#hex_color_code Scorecard#hex_color_code}
	HexColorCode *string `field:"optional" json:"hexColorCode" yaml:"hexColorCode"`
}

