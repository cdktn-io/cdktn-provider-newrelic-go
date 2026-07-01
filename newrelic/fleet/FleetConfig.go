// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fleet

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FleetConfig struct {
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
	// The type of entities this fleet will manage. Allowed values: HOST, KUBERNETESCLUSTER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.0/docs/resources/fleet#managed_entity_type Fleet#managed_entity_type}
	ManagedEntityType *string `field:"required" json:"managedEntityType" yaml:"managedEntityType"`
	// The name of the fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.0/docs/resources/fleet#name Fleet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.0/docs/resources/fleet#description Fleet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.0/docs/resources/fleet#id Fleet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The operating system type. Required for HOST fleets. Allowed values: LINUX, WINDOWS. Must not be set for KUBERNETESCLUSTER fleets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.0/docs/resources/fleet#operating_system Fleet#operating_system}
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// The organization ID. If not provided, it will be auto-fetched from the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.0/docs/resources/fleet#organization_id Fleet#organization_id}
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
	// Tags for the fleet in format 'key:value1,value2'. Each tag can have multiple values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.0/docs/resources/fleet#tags Fleet#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

