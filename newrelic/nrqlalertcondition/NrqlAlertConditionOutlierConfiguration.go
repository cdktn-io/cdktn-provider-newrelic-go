// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package nrqlalertcondition


type NrqlAlertConditionOutlierConfiguration struct {
	// dbscan block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/nrql_alert_condition#dbscan NrqlAlertCondition#dbscan}
	Dbscan *NrqlAlertConditionOutlierConfigurationDbscan `field:"required" json:"dbscan" yaml:"dbscan"`
}

