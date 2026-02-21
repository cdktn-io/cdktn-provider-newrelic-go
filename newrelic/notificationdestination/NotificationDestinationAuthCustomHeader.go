// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package notificationdestination


type NotificationDestinationAuthCustomHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/notification_destination#key NotificationDestination#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/notification_destination#value NotificationDestination#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

