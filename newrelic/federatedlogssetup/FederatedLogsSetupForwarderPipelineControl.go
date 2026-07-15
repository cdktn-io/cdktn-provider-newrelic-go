// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogssetup


type FederatedLogsSetupForwarderPipelineControl struct {
	// The fleet entity GUID used for deploying the pipeline configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.0/docs/resources/federated_logs_setup#fleet_id FederatedLogsSetup#fleet_id}
	FleetId *string `field:"required" json:"fleetId" yaml:"fleetId"`
	// routing_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.0/docs/resources/federated_logs_setup#routing_rule FederatedLogsSetup#routing_rule}
	RoutingRule *FederatedLogsSetupForwarderPipelineControlRoutingRule `field:"optional" json:"routingRule" yaml:"routingRule"`
}

