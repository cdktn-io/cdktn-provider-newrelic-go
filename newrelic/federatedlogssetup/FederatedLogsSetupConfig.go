// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogssetup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FederatedLogsSetupConfig struct {
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
	// default_partition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/federated_logs_setup#default_partition FederatedLogsSetup#default_partition}
	DefaultPartition *FederatedLogsSetupDefaultPartition `field:"required" json:"defaultPartition" yaml:"defaultPartition"`
	// The name of the federated log setup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/federated_logs_setup#name FederatedLogsSetup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/federated_logs_setup#storage FederatedLogsSetup#storage}
	Storage *FederatedLogsSetupStorage `field:"required" json:"storage" yaml:"storage"`
	// The New Relic account ID where the federated logs setup will live.
	//
	// Defaults to the provider's account_id. Changing this after creation is rejected by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/federated_logs_setup#account_id FederatedLogsSetup#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Whether the setup is active. When false, log routing to this setup is turned off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/federated_logs_setup#active FederatedLogsSetup#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// The description of the federated log setup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/federated_logs_setup#description FederatedLogsSetup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// forwarder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/federated_logs_setup#forwarder FederatedLogsSetup#forwarder}
	Forwarder *FederatedLogsSetupForwarder `field:"optional" json:"forwarder" yaml:"forwarder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/federated_logs_setup#id FederatedLogsSetup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

