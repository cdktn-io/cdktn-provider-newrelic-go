// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsconnection


type AwsConnectionCredentialAssumeRole struct {
	// ARN of the IAM role New Relic should assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/aws_connection#role_arn AwsConnection#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// External ID supplied by New Relic during AssumeRole.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/aws_connection#external_id AwsConnection#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
}

