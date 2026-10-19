// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package teamsorganizationsettings


type TeamsOrganizationSettingsHierarchyLevels struct {
	// GUID of the hierarchy level entity. Obtain from the newrelic_teams_hierarchy_levels data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.1/docs/resources/teams_organization_settings#id TeamsOrganizationSettings#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Display name of this hierarchy level (e.g. 'Division', 'Squad'). Renaming here updates the entity in the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.1/docs/resources/teams_organization_settings#name TeamsOrganizationSettings#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

