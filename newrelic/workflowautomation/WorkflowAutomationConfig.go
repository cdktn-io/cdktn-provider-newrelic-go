// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workflowautomation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkflowAutomationConfig struct {
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
	// The YAML definition of the workflow automation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/workflow_automation#definition WorkflowAutomation#definition}
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// The name of the workflow automation. Must match the name in the YAML definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/workflow_automation#name WorkflowAutomation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The scope ID (account ID for ACCOUNT scope, organization ID for ORGANIZATION scope).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/workflow_automation#scope_id WorkflowAutomation#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// The scope type. Supported values are: ACCOUNT, ORGANIZATION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/workflow_automation#scope_type WorkflowAutomation#scope_type}
	ScopeType *string `field:"required" json:"scopeType" yaml:"scopeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.95.1/docs/resources/workflow_automation#id WorkflowAutomation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

