// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package maintenancewindow

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MaintenanceWindowConfig struct {
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
	// The list of entity IDs affected by the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.102.0/docs/resources/maintenance_window#affected_entities MaintenanceWindow#affected_entities}
	AffectedEntities *[]*string `field:"required" json:"affectedEntities" yaml:"affectedEntities"`
	// The type of the entities affected by the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.102.0/docs/resources/maintenance_window#affected_entity_type MaintenanceWindow#affected_entity_type}
	AffectedEntityType *string `field:"required" json:"affectedEntityType" yaml:"affectedEntityType"`
	// The duration of the maintenance window, in ISO 8601 duration format (e.g. PT1H for one hour). Capped at 24 hours when `rrule` is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.102.0/docs/resources/maintenance_window#duration MaintenanceWindow#duration}
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// The name of the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.102.0/docs/resources/maintenance_window#name MaintenanceWindow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.102.0/docs/resources/maintenance_window#scope MaintenanceWindow#scope}
	Scope *MaintenanceWindowScope `field:"required" json:"scope" yaml:"scope"`
	// The start time of the maintenance window, as an ISO 8601 date-time without a time zone (e.g. 2026-11-05T21:00:00).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.102.0/docs/resources/maintenance_window#start_time MaintenanceWindow#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// The timezone in which the maintenance window is scheduled (e.g. America/Los_Angeles).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.102.0/docs/resources/maintenance_window#timezone MaintenanceWindow#timezone}
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
	// The description of the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.102.0/docs/resources/maintenance_window#description MaintenanceWindow#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.102.0/docs/resources/maintenance_window#id MaintenanceWindow#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The recurrence rule of the maintenance window, in RRULE format, for recurring maintenance windows.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.102.0/docs/resources/maintenance_window#rrule MaintenanceWindow#rrule}
	Rrule *string `field:"optional" json:"rrule" yaml:"rrule"`
}

