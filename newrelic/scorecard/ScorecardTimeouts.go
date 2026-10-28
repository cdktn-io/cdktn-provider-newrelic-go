// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scorecard


type ScorecardTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard#create Scorecard#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

