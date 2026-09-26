// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow


type PathpointFlowStagesLevels struct {
	// steps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#steps PathpointFlow#steps}
	Steps interface{} `field:"optional" json:"steps" yaml:"steps"`
}

