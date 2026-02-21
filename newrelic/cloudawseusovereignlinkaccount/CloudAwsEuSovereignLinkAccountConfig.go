// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudawseusovereignlinkaccount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudAwsEuSovereignLinkAccountConfig struct {
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
	// The ARN of the IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/cloud_aws_eu_sovereign_link_account#arn CloudAwsEuSovereignLinkAccount#arn}
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The name of the AWS EU Sovereign account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/cloud_aws_eu_sovereign_link_account#name CloudAwsEuSovereignLinkAccount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/cloud_aws_eu_sovereign_link_account#account_id CloudAwsEuSovereignLinkAccount#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/cloud_aws_eu_sovereign_link_account#id CloudAwsEuSovereignLinkAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// How metrics are collected. PULL or PUSH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/cloud_aws_eu_sovereign_link_account#metric_collection_mode CloudAwsEuSovereignLinkAccount#metric_collection_mode}
	MetricCollectionMode *string `field:"optional" json:"metricCollectionMode" yaml:"metricCollectionMode"`
}

