// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicelevel


type ServiceLevelEventsGoodEvents struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/service_level#from ServiceLevel#from}.
	From *string `field:"required" json:"from" yaml:"from"`
	// select block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/service_level#select ServiceLevel#select}
	Select *ServiceLevelEventsGoodEventsSelect `field:"optional" json:"select" yaml:"select"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/service_level#where ServiceLevel#where}.
	Where *string `field:"optional" json:"where" yaml:"where"`
}

