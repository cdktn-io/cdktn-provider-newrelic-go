// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow


type PathpointFlowStagesRelated struct {
	// When true, this stage acts as a source to other stages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#source PathpointFlow#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// When true, this stage acts as a target to other stages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#target PathpointFlow#target}
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}

