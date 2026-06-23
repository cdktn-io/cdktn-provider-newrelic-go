// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsconnection


type AwsConnectionCredential struct {
	// assume_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.2/docs/resources/aws_connection#assume_role AwsConnection#assume_role}
	AssumeRole *AwsConnectionCredentialAssumeRole `field:"required" json:"assumeRole" yaml:"assumeRole"`
}

