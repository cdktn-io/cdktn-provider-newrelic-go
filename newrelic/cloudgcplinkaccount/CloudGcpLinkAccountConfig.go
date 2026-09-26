// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudgcplinkaccount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudGcpLinkAccountConfig struct {
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
	// name of the linked account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/cloud_gcp_link_account#name CloudGcpLinkAccount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// project id of the Gcp account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/cloud_gcp_link_account#project_id CloudGcpLinkAccount#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// accountID of newrelic account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/cloud_gcp_link_account#account_id CloudGcpLinkAccount#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// The Workload Identity Federation pool provider audience URI, used for GCP Dimensional Metrics (keyless) linking.
	//
	// Format: //iam.googleapis.com/projects/{PROJECT_NUMBER}/locations/global/workloadIdentityPools/{POOL_ID}/providers/{PROVIDER_ID}. Required when use_workload_identity_federation = true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/cloud_gcp_link_account#audience CloudGcpLinkAccount#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/cloud_gcp_link_account#id CloudGcpLinkAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The GCP service account email New Relic impersonates to collect metrics when linking via Workload Identity Federation (GCP Dimensional Metrics).
	//
	// The service account must grant the WIF pool the roles/iam.workloadIdentityUser binding. Required when use_workload_identity_federation = true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/cloud_gcp_link_account#service_account_email CloudGcpLinkAccount#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// Set to true to link this GCP account for New Relic GCP Dimensional Metrics (v2) using keyless Workload Identity Federation (WIF).
	//
	// When true, audience and service_account_email are required. When false (the default), the account is linked using the legacy service-account-key flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/cloud_gcp_link_account#use_workload_identity_federation CloudGcpLinkAccount#use_workload_identity_federation}
	UseWorkloadIdentityFederation interface{} `field:"optional" json:"useWorkloadIdentityFederation" yaml:"useWorkloadIdentityFederation"`
}

