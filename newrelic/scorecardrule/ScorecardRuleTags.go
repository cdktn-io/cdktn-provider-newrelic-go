// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scorecardrule


type ScorecardRuleTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#key ScorecardRule#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#values ScorecardRule#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

