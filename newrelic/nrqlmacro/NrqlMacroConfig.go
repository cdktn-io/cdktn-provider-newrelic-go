// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package nrqlmacro

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NrqlMacroConfig struct {
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
	// The NRQL fragment this macro expands to wherever `$name` appears in a query.
	//
	// Changes to this field take effect immediately across every dashboard widget, alert condition, and query that references the macro.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.10/docs/resources/nrql_macro#expression NrqlMacro#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// A unique name for the macro within your organization.
	//
	// Once created, the name cannot be changed — it is the identifier used to reference the macro as `$name` in NRQL queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.10/docs/resources/nrql_macro#name NrqlMacro#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.10/docs/resources/nrql_macro#id NrqlMacro#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The number of positional parameters the expression accepts, expressed as bare `$1`, `$2`, ... tokens. Required when the expression contains any positional parameters; defaults to `0` for expressions that do not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.10/docs/resources/nrql_macro#parameter_count NrqlMacro#parameter_count}
	ParameterCount *float64 `field:"optional" json:"parameterCount" yaml:"parameterCount"`
}

