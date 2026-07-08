// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsconnection


type AwsConnectionSettings struct {
	// The key or name of the setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.2/docs/resources/aws_connection#key AwsConnection#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.2/docs/resources/aws_connection#value AwsConnection#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

