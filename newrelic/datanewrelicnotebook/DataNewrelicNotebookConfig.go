// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datanewrelicnotebook

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataNewrelicNotebookConfig struct {
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
	// The unique entity identifier (GUID) of the notebook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.98.0/docs/data-sources/notebook#guid DataNewrelicNotebook#guid}
	Guid *string `field:"required" json:"guid" yaml:"guid"`
	// When true, the full notebook body is fetched from the Blob Storage API and stored in the `content` attribute.
	//
	// When false (default), only NerdGraph metadata (title, organization_id, blob_id) is retrieved, which is faster and avoids an extra API call.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.98.0/docs/data-sources/notebook#fetch_content DataNewrelicNotebook#fetch_content}
	FetchContent interface{} `field:"optional" json:"fetchContent" yaml:"fetchContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.98.0/docs/data-sources/notebook#id DataNewrelicNotebook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

