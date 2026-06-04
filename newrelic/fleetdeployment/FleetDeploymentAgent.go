// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fleetdeployment


type FleetDeploymentAgent struct {
	// The agent type. Allowed values: NRInfra, NRDOT, FluentBit, NRPrometheusAgent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.91.0/docs/resources/fleet_deployment#agent_type FleetDeployment#agent_type}
	AgentType *string `field:"required" json:"agentType" yaml:"agentType"`
	// Configuration version entity GUID to associate with this agent in the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.91.0/docs/resources/fleet_deployment#configuration_version_id FleetDeployment#configuration_version_id}
	ConfigurationVersionId *string `field:"required" json:"configurationVersionId" yaml:"configurationVersionId"`
	// The agent version to deploy (e.g. "1.58.0").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.91.0/docs/resources/fleet_deployment#version FleetDeployment#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

