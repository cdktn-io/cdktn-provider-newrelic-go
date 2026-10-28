// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package teamsorganizationsettings


type TeamsOrganizationSettingsSyncGroupRulesConditions struct {
	// Match type: STARTS_WITH, ENDS_WITH, or CONTAINS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/teams_organization_settings#type TeamsOrganizationSettings#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The string to match against the group name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/teams_organization_settings#value TeamsOrganizationSettings#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

