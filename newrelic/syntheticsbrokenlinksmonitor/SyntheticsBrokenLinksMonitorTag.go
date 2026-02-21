// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticsbrokenlinksmonitor


type SyntheticsBrokenLinksMonitorTag struct {
	// Name of the tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/synthetics_broken_links_monitor#key SyntheticsBrokenLinksMonitor#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Values associated with the tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/synthetics_broken_links_monitor#values SyntheticsBrokenLinksMonitor#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

