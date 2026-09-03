// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudgcpdmintegrations

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudGcpDmIntegrationsConfig struct {
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
	// The ID of the GCP Dimensional Metrics linked account (from newrelic_cloud_gcp_link_account with use_workload_identity_federation = true).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#linked_account_id CloudGcpDmIntegrations#linked_account_id}
	LinkedAccountId *float64 `field:"required" json:"linkedAccountId" yaml:"linkedAccountId"`
	// The New Relic account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#account_id CloudGcpDmIntegrations#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// ai_platform block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#ai_platform CloudGcpDmIntegrations#ai_platform}
	AiPlatform *CloudGcpDmIntegrationsAiPlatform `field:"optional" json:"aiPlatform" yaml:"aiPlatform"`
	// alloy_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#alloy_db CloudGcpDmIntegrations#alloy_db}
	AlloyDb *CloudGcpDmIntegrationsAlloyDb `field:"optional" json:"alloyDb" yaml:"alloyDb"`
	// api_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#api_gateway CloudGcpDmIntegrations#api_gateway}
	ApiGateway *CloudGcpDmIntegrationsApiGateway `field:"optional" json:"apiGateway" yaml:"apiGateway"`
	// app_engine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#app_engine CloudGcpDmIntegrations#app_engine}
	AppEngine *CloudGcpDmIntegrationsAppEngine `field:"optional" json:"appEngine" yaml:"appEngine"`
	// big_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#big_query CloudGcpDmIntegrations#big_query}
	BigQuery *CloudGcpDmIntegrationsBigQuery `field:"optional" json:"bigQuery" yaml:"bigQuery"`
	// big_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#big_table CloudGcpDmIntegrations#big_table}
	BigTable *CloudGcpDmIntegrationsBigTable `field:"optional" json:"bigTable" yaml:"bigTable"`
	// composer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#composer CloudGcpDmIntegrations#composer}
	Composer *CloudGcpDmIntegrationsComposer `field:"optional" json:"composer" yaml:"composer"`
	// data_flow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#data_flow CloudGcpDmIntegrations#data_flow}
	DataFlow *CloudGcpDmIntegrationsDataFlow `field:"optional" json:"dataFlow" yaml:"dataFlow"`
	// data_proc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#data_proc CloudGcpDmIntegrations#data_proc}
	DataProc *CloudGcpDmIntegrationsDataProc `field:"optional" json:"dataProc" yaml:"dataProc"`
	// data_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#data_store CloudGcpDmIntegrations#data_store}
	DataStore *CloudGcpDmIntegrationsDataStore `field:"optional" json:"dataStore" yaml:"dataStore"`
	// firebase_app_hosting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#firebase_app_hosting CloudGcpDmIntegrations#firebase_app_hosting}
	FirebaseAppHosting *CloudGcpDmIntegrationsFirebaseAppHosting `field:"optional" json:"firebaseAppHosting" yaml:"firebaseAppHosting"`
	// firebase_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#firebase_auth CloudGcpDmIntegrations#firebase_auth}
	FirebaseAuth *CloudGcpDmIntegrationsFirebaseAuth `field:"optional" json:"firebaseAuth" yaml:"firebaseAuth"`
	// firebase_database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#firebase_database CloudGcpDmIntegrations#firebase_database}
	FirebaseDatabase *CloudGcpDmIntegrationsFirebaseDatabase `field:"optional" json:"firebaseDatabase" yaml:"firebaseDatabase"`
	// firebase_hosting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#firebase_hosting CloudGcpDmIntegrations#firebase_hosting}
	FirebaseHosting *CloudGcpDmIntegrationsFirebaseHosting `field:"optional" json:"firebaseHosting" yaml:"firebaseHosting"`
	// firebase_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#firebase_storage CloudGcpDmIntegrations#firebase_storage}
	FirebaseStorage *CloudGcpDmIntegrationsFirebaseStorage `field:"optional" json:"firebaseStorage" yaml:"firebaseStorage"`
	// firebase_vertex_ai block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#firebase_vertex_ai CloudGcpDmIntegrations#firebase_vertex_ai}
	FirebaseVertexAi *CloudGcpDmIntegrationsFirebaseVertexAi `field:"optional" json:"firebaseVertexAi" yaml:"firebaseVertexAi"`
	// firestore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#firestore CloudGcpDmIntegrations#firestore}
	Firestore *CloudGcpDmIntegrationsFirestore `field:"optional" json:"firestore" yaml:"firestore"`
	// functions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#functions CloudGcpDmIntegrations#functions}
	Functions *CloudGcpDmIntegrationsFunctions `field:"optional" json:"functions" yaml:"functions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#id CloudGcpDmIntegrations#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// interconnect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#interconnect CloudGcpDmIntegrations#interconnect}
	Interconnect *CloudGcpDmIntegrationsInterconnect `field:"optional" json:"interconnect" yaml:"interconnect"`
	// istio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#istio CloudGcpDmIntegrations#istio}
	Istio *CloudGcpDmIntegrationsIstio `field:"optional" json:"istio" yaml:"istio"`
	// kubernetes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#kubernetes CloudGcpDmIntegrations#kubernetes}
	Kubernetes *CloudGcpDmIntegrationsKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// load_balancing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#load_balancing CloudGcpDmIntegrations#load_balancing}
	LoadBalancing *CloudGcpDmIntegrationsLoadBalancing `field:"optional" json:"loadBalancing" yaml:"loadBalancing"`
	// managed_kafka block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#managed_kafka CloudGcpDmIntegrations#managed_kafka}
	ManagedKafka *CloudGcpDmIntegrationsManagedKafka `field:"optional" json:"managedKafka" yaml:"managedKafka"`
	// mem_cache block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#mem_cache CloudGcpDmIntegrations#mem_cache}
	MemCache *CloudGcpDmIntegrationsMemCache `field:"optional" json:"memCache" yaml:"memCache"`
	// memory_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#memory_store CloudGcpDmIntegrations#memory_store}
	MemoryStore *CloudGcpDmIntegrationsMemoryStore `field:"optional" json:"memoryStore" yaml:"memoryStore"`
	// pub_sub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#pub_sub CloudGcpDmIntegrations#pub_sub}
	PubSub *CloudGcpDmIntegrationsPubSub `field:"optional" json:"pubSub" yaml:"pubSub"`
	// redis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#redis CloudGcpDmIntegrations#redis}
	Redis *CloudGcpDmIntegrationsRedis `field:"optional" json:"redis" yaml:"redis"`
	// router block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#router CloudGcpDmIntegrations#router}
	Router *CloudGcpDmIntegrationsRouter `field:"optional" json:"router" yaml:"router"`
	// run block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#run CloudGcpDmIntegrations#run}
	Run *CloudGcpDmIntegrationsRun `field:"optional" json:"run" yaml:"run"`
	// spanner block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#spanner CloudGcpDmIntegrations#spanner}
	Spanner *CloudGcpDmIntegrationsSpanner `field:"optional" json:"spanner" yaml:"spanner"`
	// sql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#sql CloudGcpDmIntegrations#sql}
	Sql *CloudGcpDmIntegrationsSql `field:"optional" json:"sql" yaml:"sql"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#storage CloudGcpDmIntegrations#storage}
	Storage *CloudGcpDmIntegrationsStorage `field:"optional" json:"storage" yaml:"storage"`
	// virtual_machines block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#virtual_machines CloudGcpDmIntegrations#virtual_machines}
	VirtualMachines *CloudGcpDmIntegrationsVirtualMachines `field:"optional" json:"virtualMachines" yaml:"virtualMachines"`
	// vpc_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.97.2/docs/resources/cloud_gcp_dm_integrations#vpc_access CloudGcpDmIntegrations#vpc_access}
	VpcAccess *CloudGcpDmIntegrationsVpcAccess `field:"optional" json:"vpcAccess" yaml:"vpcAccess"`
}

