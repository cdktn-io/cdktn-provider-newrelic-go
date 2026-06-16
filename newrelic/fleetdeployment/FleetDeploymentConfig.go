// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fleetdeployment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FleetDeploymentConfig struct {
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
	// The GUID of the fleet this deployment belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/fleet_deployment#fleet_id FleetDeployment#fleet_id}
	FleetId *string `field:"required" json:"fleetId" yaml:"fleetId"`
	// agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/fleet_deployment#agent FleetDeployment#agent}
	Agent interface{} `field:"optional" json:"agent" yaml:"agent"`
	// A description of the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/fleet_deployment#description FleetDeployment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/fleet_deployment#id FleetDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/fleet_deployment#name FleetDeployment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The organization ID. Auto-fetched from the account if not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/fleet_deployment#organization_id FleetDeployment#organization_id}
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
	// Tags for the deployment in format 'key:value1,value2'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/fleet_deployment#tags FleetDeployment#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

