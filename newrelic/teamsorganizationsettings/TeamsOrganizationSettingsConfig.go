// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package teamsorganizationsettings

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TeamsOrganizationSettingsConfig struct {
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
	// Whether tag-based entity discovery is enabled. When true, entities with a matching tag are automatically assigned to teams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.1/docs/resources/teams_organization_settings#discovery_enabled TeamsOrganizationSettings#discovery_enabled}
	DiscoveryEnabled interface{} `field:"optional" json:"discoveryEnabled" yaml:"discoveryEnabled"`
	// Tag keys used for entity discovery (e.g. ["team"]). Entities with these tags are auto-assigned to the team whose name or alias matches the tag value. Must contain at least one key when declared.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.1/docs/resources/teams_organization_settings#discovery_tag_keys TeamsOrganizationSettings#discovery_tag_keys}
	DiscoveryTagKeys *[]*string `field:"optional" json:"discoveryTagKeys" yaml:"discoveryTagKeys"`
	// hierarchy_levels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.1/docs/resources/teams_organization_settings#hierarchy_levels TeamsOrganizationSettings#hierarchy_levels}
	HierarchyLevels interface{} `field:"optional" json:"hierarchyLevels" yaml:"hierarchyLevels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.1/docs/resources/teams_organization_settings#id TeamsOrganizationSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// sync_group_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.1/docs/resources/teams_organization_settings#sync_group_rules TeamsOrganizationSettings#sync_group_rules}
	SyncGroupRules *TeamsOrganizationSettingsSyncGroupRules `field:"optional" json:"syncGroupRules" yaml:"syncGroupRules"`
	// Whether automatic team creation from IdP groups is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.1/docs/resources/teams_organization_settings#sync_groups_enabled TeamsOrganizationSettings#sync_groups_enabled}
	SyncGroupsEnabled interface{} `field:"optional" json:"syncGroupsEnabled" yaml:"syncGroupsEnabled"`
}

