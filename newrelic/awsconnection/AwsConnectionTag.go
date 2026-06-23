// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsconnection


type AwsConnectionTag struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.2/docs/resources/aws_connection#key AwsConnection#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.2/docs/resources/aws_connection#values AwsConnection#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

