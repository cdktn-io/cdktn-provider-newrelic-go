// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package notebook


type NotebookTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/notebook#create Notebook#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/notebook#delete Notebook#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/notebook#update Notebook#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

