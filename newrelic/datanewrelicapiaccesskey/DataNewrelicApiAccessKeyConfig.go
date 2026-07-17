// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datanewrelicapiaccesskey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataNewrelicApiAccessKeyConfig struct {
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
	// The type of the key, one of INGEST or USER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.2/docs/data-sources/api_access_key#key_type DataNewrelicApiAccessKey#key_type}
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// The New Relic account ID the key belongs to.
	//
	// Defaults to the account ID configured on the provider when not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.2/docs/data-sources/api_access_key#account_id DataNewrelicApiAccessKey#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.2/docs/data-sources/api_access_key#id DataNewrelicApiAccessKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The type of the ingest key, one of LICENSE or BROWSER. Only applies when `key_type` is INGEST.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.2/docs/data-sources/api_access_key#ingest_type DataNewrelicApiAccessKey#ingest_type}
	IngestType *string `field:"optional" json:"ingestType" yaml:"ingestType"`
	// The ID of the key.
	//
	// When specified, the key is fetched directly by its ID instead of searching by other attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.2/docs/data-sources/api_access_key#key_id DataNewrelicApiAccessKey#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// The name of the key. Used to narrow down the search when `key_id` is not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.2/docs/data-sources/api_access_key#name DataNewrelicApiAccessKey#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the user that owns the key. Only applies when `key_type` is USER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.2/docs/data-sources/api_access_key#user_id DataNewrelicApiAccessKey#user_id}
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
}

