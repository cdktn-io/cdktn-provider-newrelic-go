// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cardinalitymanagement

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CardinalityManagementConfig struct {
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
	// The override mode.
	//
	// Use `DEFAULT` to set a single account-wide limit that applies to all metrics, or `PER_METRIC` to set individual limits for one or more named metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/cardinality_management#mode CardinalityManagement#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The account-wide cardinality limit — the maximum number of unique dimension-value combinations allowed per metric per day.
	//
	// Required when `mode` is `DEFAULT`; must not be set when `mode` is `PER_METRIC`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/cardinality_management#cardinality_limit CardinalityManagement#cardinality_limit}
	CardinalityLimit *float64 `field:"optional" json:"cardinalityLimit" yaml:"cardinalityLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/cardinality_management#id CardinalityManagement#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.94.3/docs/resources/cardinality_management#metric CardinalityManagement#metric}
	Metric interface{} `field:"optional" json:"metric" yaml:"metric"`
}

