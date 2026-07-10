// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsconnection

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AwsConnectionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/aws_connection#credential AwsConnection#credential}
	Credential *AwsConnectionCredential `field:"required" json:"credential" yaml:"credential"`
	// The name of the AWS connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/aws_connection#name AwsConnection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account ID where the AWS connection will be created. Used when scope_type is ACCOUNT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/aws_connection#account_id AwsConnection#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// The description of the AWS connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/aws_connection#description AwsConnection#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Flag to indicate if the connection is enabled. True by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/aws_connection#enabled AwsConnection#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Optional field representing an identifier managed by the consumer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/aws_connection#external_id AwsConnection#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/aws_connection#id AwsConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Default region for this connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/aws_connection#region AwsConnection#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The scope ID (account ID or organization ID) for the AWS connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/aws_connection#scope_id AwsConnection#scope_id}
	ScopeId *string `field:"optional" json:"scopeId" yaml:"scopeId"`
	// The scope type for the AWS connection. Valid values are ACCOUNT and ORGANIZATION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/aws_connection#scope_type AwsConnection#scope_type}
	ScopeType *string `field:"optional" json:"scopeType" yaml:"scopeType"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/aws_connection#settings AwsConnection#settings}
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/aws_connection#tag AwsConnection#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

