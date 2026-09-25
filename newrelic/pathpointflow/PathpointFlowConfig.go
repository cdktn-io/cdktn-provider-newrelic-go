// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PathpointFlowConfig struct {
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
	// Display name of the Pathpoint flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#name PathpointFlow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The New Relic account ID that owns this Pathpoint flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#account_id PathpointFlow#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Optional category used to group flows (e.g. Marketing, Checkout).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#category PathpointFlow#category}
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Optional description of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#description PathpointFlow#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Health rollup strategy: ALERT_CONDITIONS or AUTOMATIC_ROLL_UP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#health_rollup PathpointFlow#health_rollup}
	HealthRollup *string `field:"optional" json:"healthRollup" yaml:"healthRollup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#id PathpointFlow#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kpis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#kpis PathpointFlow#kpis}
	Kpis interface{} `field:"optional" json:"kpis" yaml:"kpis"`
	// How often health statuses refresh: ONE_MINUTE, FIVE_MINUTES, TEN_MINUTES, FIFTEEN_MINUTES, THIRTY_MINUTES.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#refresh_interval PathpointFlow#refresh_interval}
	RefreshInterval *string `field:"optional" json:"refreshInterval" yaml:"refreshInterval"`
	// stages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.1/docs/resources/pathpoint_flow#stages PathpointFlow#stages}
	Stages interface{} `field:"optional" json:"stages" yaml:"stages"`
}

