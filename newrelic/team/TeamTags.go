// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package team


type TeamTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.0/docs/resources/team#key Team#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.0/docs/resources/team#values Team#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

