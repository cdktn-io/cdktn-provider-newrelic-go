// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datanewrelicgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataNewrelicGroupConfig struct {
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
	// The ID of the Authentication Domain the group being queried would belong to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/data-sources/group#authentication_domain_id DataNewrelicGroup#authentication_domain_id}
	AuthenticationDomainId *string `field:"required" json:"authenticationDomainId" yaml:"authenticationDomainId"`
	// The name of the group to be queried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/data-sources/group#name DataNewrelicGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

