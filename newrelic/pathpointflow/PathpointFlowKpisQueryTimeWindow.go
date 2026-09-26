// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow


type PathpointFlowKpisQueryTimeWindow struct {
	// Raw NRQL time fragment, e.g. 'SINCE 3 days ago COMPARE WITH 1 day ago'. Mutually exclusive with relative_range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#custom_range PathpointFlow#custom_range}
	CustomRange *string `field:"optional" json:"customRange" yaml:"customRange"`
	// relative_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#relative_range PathpointFlow#relative_range}
	RelativeRange *PathpointFlowKpisQueryTimeWindowRelativeRange `field:"optional" json:"relativeRange" yaml:"relativeRange"`
}

