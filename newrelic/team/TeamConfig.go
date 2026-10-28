// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package team

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TeamConfig struct {
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
	// The name of the team. Must be unique within the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#name Team#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Additional searchable names for the team. Each alias must be unique within the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#aliases Team#aliases}
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// A free-text description of the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#description Team#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Set of entity GUIDs that this team owns. Added to the team's auto-created ownership collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#entities Team#entities}
	Entities *[]*string `field:"optional" json:"entities" yaml:"entities"`
	// Controls how Terraform manages this team's entity ownership collection.
	//
	// • `managed` (default): Terraform tracks the entities declared in the entities block and ensures     they remain in the collection. Drift surfaces only when a declared entity goes missing.     External additions are silently ignored.
	//   • `unmanaged`: Terraform does not control entity ownership. The ownership collection is left     entirely to tag-based discovery and/or manual UI management. No drift or warnings     are shown for entities, and the `entities` attribute cannot be set in this mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#entity_management_mode Team#entity_management_mode}
	EntityManagementMode *string `field:"optional" json:"entityManagementMode" yaml:"entityManagementMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#id Team#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of New Relic user IDs (integers) to designate as team managers. Each ID must also appear in members.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#managers Team#managers}
	Managers *[]*float64 `field:"optional" json:"managers" yaml:"managers"`
	// Set of New Relic user IDs (integers) to add as members of this team.
	//
	// Managers must be a subset of members.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#members Team#members}
	Members *[]*float64 `field:"optional" json:"members" yaml:"members"`
	// Entity management GUID of the parent team. Cannot be set to the team's own GUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#parent_id Team#parent_id}
	ParentId *string `field:"optional" json:"parentId" yaml:"parentId"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#resources Team#resources}
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#tags Team#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/team#timeouts Team#timeouts}
	Timeouts *TeamTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

