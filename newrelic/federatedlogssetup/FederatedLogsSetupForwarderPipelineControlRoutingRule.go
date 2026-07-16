// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogssetup


type FederatedLogsSetupForwarderPipelineControlRoutingRule struct {
	// OTTL expression for routing logs to this setup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/federated_logs_setup#expression FederatedLogsSetup#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

