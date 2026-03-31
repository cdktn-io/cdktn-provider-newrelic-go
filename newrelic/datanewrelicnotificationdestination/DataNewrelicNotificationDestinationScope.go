// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datanewrelicnotificationdestination


type DataNewrelicNotificationDestinationScope struct {
	// The ID of the scope (Organization UUID for ORGANIZATION scope, Account ID for ACCOUNT scope).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.84.0/docs/data-sources/notification_destination#id DataNewrelicNotificationDestination#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// (Required) The scope type of the destination. One of: (ACCOUNT, ORGANIZATION).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.84.0/docs/data-sources/notification_destination#type DataNewrelicNotificationDestination#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

