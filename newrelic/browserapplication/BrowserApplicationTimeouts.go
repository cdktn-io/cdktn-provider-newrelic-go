// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package browserapplication


type BrowserApplicationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/browser_application#create BrowserApplication#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/browser_application#read BrowserApplication#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

