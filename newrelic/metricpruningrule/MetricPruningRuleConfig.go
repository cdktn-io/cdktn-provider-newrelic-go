// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package metricpruningrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MetricPruningRuleConfig struct {
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
	// The NRQL query that identifies the metric attributes to prune.
	//
	// Must select specific attributes from Metric (e.g. `SELECT collector.name FROM Metric WHERE metricName = 'my.metric'`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.92.0/docs/resources/metric_pruning_rule#nrql MetricPruningRule#nrql}
	Nrql *string `field:"required" json:"nrql" yaml:"nrql"`
	// The account ID in which the pruning rule is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.92.0/docs/resources/metric_pruning_rule#account_id MetricPruningRule#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// A human-readable description of the pruning rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.92.0/docs/resources/metric_pruning_rule#description MetricPruningRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.92.0/docs/resources/metric_pruning_rule#id MetricPruningRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

