// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package team


type TeamResources struct {
	// The URL or contact address for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#content Team#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// Resource category. Valid values are defined in TeamResourceTypeValues() in the go-client and stay in sync with the API automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#type Team#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A short label shown in the UI for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#title Team#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

