// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package teamsorganizationsettings


type TeamsOrganizationSettingsSyncGroupRules struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.0/docs/resources/teams_organization_settings#conditions TeamsOrganizationSettings#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
}

