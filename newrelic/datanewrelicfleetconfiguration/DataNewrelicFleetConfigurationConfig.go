// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datanewrelicfleetconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataNewrelicFleetConfigurationConfig struct {
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
	// The GUID of the fleet configuration entity.
	//
	// Returns the content of the latest version. Populated automatically when looking up by version_entity_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/data-sources/fleet_configuration#configuration_id DataNewrelicFleetConfiguration#configuration_id}
	ConfigurationId *string `field:"optional" json:"configurationId" yaml:"configurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/data-sources/fleet_configuration#id DataNewrelicFleetConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the fleet configuration. The first matching configuration is returned. Returns the content of its latest version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/data-sources/fleet_configuration#name DataNewrelicFleetConfiguration#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The organization ID. Resolved automatically from the provider when omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/data-sources/fleet_configuration#organization_id DataNewrelicFleetConfiguration#organization_id}
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
	// The GUID of a specific configuration version entity. Returns the content of that exact version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/data-sources/fleet_configuration#version_entity_id DataNewrelicFleetConfiguration#version_entity_id}
	VersionEntityId *string `field:"optional" json:"versionEntityId" yaml:"versionEntityId"`
}

