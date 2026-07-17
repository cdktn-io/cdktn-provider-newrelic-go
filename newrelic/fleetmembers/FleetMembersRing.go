// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fleetmembers


type FleetMembersRing struct {
	// Ordered list of entity GUIDs to assign to this ring.
	//
	// Only the entities listed here are tracked by Terraform; any other entities already in the ring through other means are not affected. Removing a GUID from this list will remove that entity from the fleet ring on the next apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.2/docs/resources/fleet_members#entity_ids FleetMembers#entity_ids}
	EntityIds *[]*string `field:"required" json:"entityIds" yaml:"entityIds"`
	// The name of the ring as configured on the fleet (e.g. "default", "canary").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.2/docs/resources/fleet_members#name FleetMembers#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

