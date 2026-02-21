// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workflow


type WorkflowEnrichments struct {
	// nrql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/workflow#nrql Workflow#nrql}
	Nrql interface{} `field:"required" json:"nrql" yaml:"nrql"`
}

