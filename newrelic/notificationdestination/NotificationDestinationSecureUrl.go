// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package notificationdestination


type NotificationDestinationSecureUrl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/notification_destination#prefix NotificationDestination#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/notification_destination#secure_suffix NotificationDestination#secure_suffix}.
	SecureSuffix *string `field:"required" json:"secureSuffix" yaml:"secureSuffix"`
}

