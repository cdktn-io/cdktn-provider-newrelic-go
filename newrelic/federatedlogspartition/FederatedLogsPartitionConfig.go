// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogspartition

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FederatedLogsPartitionConfig struct {
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
	// The name of the partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/federated_logs_partition#name FederatedLogsPartition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the federated log setup this partition belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/federated_logs_partition#setup_id FederatedLogsPartition#setup_id}
	SetupId *string `field:"required" json:"setupId" yaml:"setupId"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/federated_logs_partition#storage FederatedLogsPartition#storage}
	Storage *FederatedLogsPartitionStorage `field:"required" json:"storage" yaml:"storage"`
	// The New Relic account ID where the federated logs partition will live.
	//
	// Defaults to the provider's account_id. Changing this after creation is rejected by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/federated_logs_partition#account_id FederatedLogsPartition#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Whether the partition is active. When false, log routing to this partition is turned off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/federated_logs_partition#active FederatedLogsPartition#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// data_retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/federated_logs_partition#data_retention_policy FederatedLogsPartition#data_retention_policy}
	DataRetentionPolicy *FederatedLogsPartitionDataRetentionPolicy `field:"optional" json:"dataRetentionPolicy" yaml:"dataRetentionPolicy"`
	// The description of the partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/federated_logs_partition#description FederatedLogsPartition#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// forwarder_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/federated_logs_partition#forwarder_configuration FederatedLogsPartition#forwarder_configuration}
	ForwarderConfiguration *FederatedLogsPartitionForwarderConfiguration `field:"optional" json:"forwarderConfiguration" yaml:"forwarderConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.93.1/docs/resources/federated_logs_partition#id FederatedLogsPartition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

