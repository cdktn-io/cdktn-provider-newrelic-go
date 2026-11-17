// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package maintenancewindow


type MaintenanceWindowScope struct {
	// The unique identifier of the object within the given scope (e.g. an account ID).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.102.0/docs/resources/maintenance_window#id MaintenanceWindow#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The kind of scope `id` refers to. Valid values are `ACCOUNT` and `ORGANIZATION`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.102.0/docs/resources/maintenance_window#type MaintenanceWindow#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

