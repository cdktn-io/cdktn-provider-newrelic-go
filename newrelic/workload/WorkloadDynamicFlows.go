// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workload


type WorkloadDynamicFlows struct {
	// The unique entity identifier of the dynamic flow entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.2/docs/resources/workload#entity_guid Workload#entity_guid}
	EntityGuid *string `field:"required" json:"entityGuid" yaml:"entityGuid"`
	// The transaction name associated with the dynamic flow entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.2/docs/resources/workload#transaction_name Workload#transaction_name}
	TransactionName *string `field:"required" json:"transactionName" yaml:"transactionName"`
}

