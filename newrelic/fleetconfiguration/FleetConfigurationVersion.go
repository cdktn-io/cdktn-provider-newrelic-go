// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fleetconfiguration


type FleetConfigurationVersion struct {
	// Configuration content for this version (YAML or JSON).
	//
	// Content must be unique across version blocks. Use file() to load from a file: file("${path.module}/config.yaml").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.87.3/docs/resources/fleet_configuration#configuration_content FleetConfiguration#configuration_content}
	ConfigurationContent *string `field:"required" json:"configurationContent" yaml:"configurationContent"`
}

