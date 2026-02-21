// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicelevel


type ServiceLevelObjectiveTimeWindow struct {
	// rolling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/service_level#rolling ServiceLevel#rolling}
	Rolling *ServiceLevelObjectiveTimeWindowRolling `field:"required" json:"rolling" yaml:"rolling"`
}

