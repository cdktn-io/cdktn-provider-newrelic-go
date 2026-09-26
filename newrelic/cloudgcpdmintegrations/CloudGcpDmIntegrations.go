// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudgcpdmintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v16/cloudgcpdmintegrations/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/cloud_gcp_dm_integrations newrelic_cloud_gcp_dm_integrations}.
type CloudGcpDmIntegrations interface {
	cdktn.TerraformResource
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	AiPlatform() CloudGcpDmIntegrationsAiPlatformOutputReference
	AiPlatformInput() *CloudGcpDmIntegrationsAiPlatform
	AlloyDb() CloudGcpDmIntegrationsAlloyDbOutputReference
	AlloyDbInput() *CloudGcpDmIntegrationsAlloyDb
	ApiGateway() CloudGcpDmIntegrationsApiGatewayOutputReference
	ApiGatewayInput() *CloudGcpDmIntegrationsApiGateway
	AppEngine() CloudGcpDmIntegrationsAppEngineOutputReference
	AppEngineInput() *CloudGcpDmIntegrationsAppEngine
	BigQuery() CloudGcpDmIntegrationsBigQueryOutputReference
	BigQueryInput() *CloudGcpDmIntegrationsBigQuery
	BigTable() CloudGcpDmIntegrationsBigTableOutputReference
	BigTableInput() *CloudGcpDmIntegrationsBigTable
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Composer() CloudGcpDmIntegrationsComposerOutputReference
	ComposerInput() *CloudGcpDmIntegrationsComposer
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DataFlow() CloudGcpDmIntegrationsDataFlowOutputReference
	DataFlowInput() *CloudGcpDmIntegrationsDataFlow
	DataProc() CloudGcpDmIntegrationsDataProcOutputReference
	DataProcInput() *CloudGcpDmIntegrationsDataProc
	DataStore() CloudGcpDmIntegrationsDataStoreOutputReference
	DataStoreInput() *CloudGcpDmIntegrationsDataStore
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FirebaseAppHosting() CloudGcpDmIntegrationsFirebaseAppHostingOutputReference
	FirebaseAppHostingInput() *CloudGcpDmIntegrationsFirebaseAppHosting
	FirebaseAuth() CloudGcpDmIntegrationsFirebaseAuthOutputReference
	FirebaseAuthInput() *CloudGcpDmIntegrationsFirebaseAuth
	FirebaseDatabase() CloudGcpDmIntegrationsFirebaseDatabaseOutputReference
	FirebaseDatabaseInput() *CloudGcpDmIntegrationsFirebaseDatabase
	FirebaseHosting() CloudGcpDmIntegrationsFirebaseHostingOutputReference
	FirebaseHostingInput() *CloudGcpDmIntegrationsFirebaseHosting
	FirebaseStorage() CloudGcpDmIntegrationsFirebaseStorageOutputReference
	FirebaseStorageInput() *CloudGcpDmIntegrationsFirebaseStorage
	FirebaseVertexAi() CloudGcpDmIntegrationsFirebaseVertexAiOutputReference
	FirebaseVertexAiInput() *CloudGcpDmIntegrationsFirebaseVertexAi
	Firestore() CloudGcpDmIntegrationsFirestoreOutputReference
	FirestoreInput() *CloudGcpDmIntegrationsFirestore
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Functions() CloudGcpDmIntegrationsFunctionsOutputReference
	FunctionsInput() *CloudGcpDmIntegrationsFunctions
	Id() *string
	SetId(val *string)
	IdInput() *string
	Interconnect() CloudGcpDmIntegrationsInterconnectOutputReference
	InterconnectInput() *CloudGcpDmIntegrationsInterconnect
	Istio() CloudGcpDmIntegrationsIstioOutputReference
	IstioInput() *CloudGcpDmIntegrationsIstio
	Kubernetes() CloudGcpDmIntegrationsKubernetesOutputReference
	KubernetesInput() *CloudGcpDmIntegrationsKubernetes
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LinkedAccountId() *float64
	SetLinkedAccountId(val *float64)
	LinkedAccountIdInput() *float64
	LoadBalancing() CloudGcpDmIntegrationsLoadBalancingOutputReference
	LoadBalancingInput() *CloudGcpDmIntegrationsLoadBalancing
	ManagedKafka() CloudGcpDmIntegrationsManagedKafkaOutputReference
	ManagedKafkaInput() *CloudGcpDmIntegrationsManagedKafka
	MemCache() CloudGcpDmIntegrationsMemCacheOutputReference
	MemCacheInput() *CloudGcpDmIntegrationsMemCache
	MemoryStore() CloudGcpDmIntegrationsMemoryStoreOutputReference
	MemoryStoreInput() *CloudGcpDmIntegrationsMemoryStore
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PubSub() CloudGcpDmIntegrationsPubSubOutputReference
	PubSubInput() *CloudGcpDmIntegrationsPubSub
	// Experimental.
	RawOverrides() interface{}
	Redis() CloudGcpDmIntegrationsRedisOutputReference
	RedisInput() *CloudGcpDmIntegrationsRedis
	Router() CloudGcpDmIntegrationsRouterOutputReference
	RouterInput() *CloudGcpDmIntegrationsRouter
	Run() CloudGcpDmIntegrationsRunOutputReference
	RunInput() *CloudGcpDmIntegrationsRun
	Spanner() CloudGcpDmIntegrationsSpannerOutputReference
	SpannerInput() *CloudGcpDmIntegrationsSpanner
	Sql() CloudGcpDmIntegrationsSqlOutputReference
	SqlInput() *CloudGcpDmIntegrationsSql
	Storage() CloudGcpDmIntegrationsStorageOutputReference
	StorageInput() *CloudGcpDmIntegrationsStorage
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VirtualMachines() CloudGcpDmIntegrationsVirtualMachinesOutputReference
	VirtualMachinesInput() *CloudGcpDmIntegrationsVirtualMachines
	VpcAccess() CloudGcpDmIntegrationsVpcAccessOutputReference
	VpcAccessInput() *CloudGcpDmIntegrationsVpcAccess
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAiPlatform(value *CloudGcpDmIntegrationsAiPlatform)
	PutAlloyDb(value *CloudGcpDmIntegrationsAlloyDb)
	PutApiGateway(value *CloudGcpDmIntegrationsApiGateway)
	PutAppEngine(value *CloudGcpDmIntegrationsAppEngine)
	PutBigQuery(value *CloudGcpDmIntegrationsBigQuery)
	PutBigTable(value *CloudGcpDmIntegrationsBigTable)
	PutComposer(value *CloudGcpDmIntegrationsComposer)
	PutDataFlow(value *CloudGcpDmIntegrationsDataFlow)
	PutDataProc(value *CloudGcpDmIntegrationsDataProc)
	PutDataStore(value *CloudGcpDmIntegrationsDataStore)
	PutFirebaseAppHosting(value *CloudGcpDmIntegrationsFirebaseAppHosting)
	PutFirebaseAuth(value *CloudGcpDmIntegrationsFirebaseAuth)
	PutFirebaseDatabase(value *CloudGcpDmIntegrationsFirebaseDatabase)
	PutFirebaseHosting(value *CloudGcpDmIntegrationsFirebaseHosting)
	PutFirebaseStorage(value *CloudGcpDmIntegrationsFirebaseStorage)
	PutFirebaseVertexAi(value *CloudGcpDmIntegrationsFirebaseVertexAi)
	PutFirestore(value *CloudGcpDmIntegrationsFirestore)
	PutFunctions(value *CloudGcpDmIntegrationsFunctions)
	PutInterconnect(value *CloudGcpDmIntegrationsInterconnect)
	PutIstio(value *CloudGcpDmIntegrationsIstio)
	PutKubernetes(value *CloudGcpDmIntegrationsKubernetes)
	PutLoadBalancing(value *CloudGcpDmIntegrationsLoadBalancing)
	PutManagedKafka(value *CloudGcpDmIntegrationsManagedKafka)
	PutMemCache(value *CloudGcpDmIntegrationsMemCache)
	PutMemoryStore(value *CloudGcpDmIntegrationsMemoryStore)
	PutPubSub(value *CloudGcpDmIntegrationsPubSub)
	PutRedis(value *CloudGcpDmIntegrationsRedis)
	PutRouter(value *CloudGcpDmIntegrationsRouter)
	PutRun(value *CloudGcpDmIntegrationsRun)
	PutSpanner(value *CloudGcpDmIntegrationsSpanner)
	PutSql(value *CloudGcpDmIntegrationsSql)
	PutStorage(value *CloudGcpDmIntegrationsStorage)
	PutVirtualMachines(value *CloudGcpDmIntegrationsVirtualMachines)
	PutVpcAccess(value *CloudGcpDmIntegrationsVpcAccess)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetAccountId()
	ResetAiPlatform()
	ResetAlloyDb()
	ResetApiGateway()
	ResetAppEngine()
	ResetBigQuery()
	ResetBigTable()
	ResetComposer()
	ResetDataFlow()
	ResetDataProc()
	ResetDataStore()
	ResetFirebaseAppHosting()
	ResetFirebaseAuth()
	ResetFirebaseDatabase()
	ResetFirebaseHosting()
	ResetFirebaseStorage()
	ResetFirebaseVertexAi()
	ResetFirestore()
	ResetFunctions()
	ResetId()
	ResetInterconnect()
	ResetIstio()
	ResetKubernetes()
	ResetLoadBalancing()
	ResetManagedKafka()
	ResetMemCache()
	ResetMemoryStore()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPubSub()
	ResetRedis()
	ResetRouter()
	ResetRun()
	ResetSpanner()
	ResetSql()
	ResetStorage()
	ResetVirtualMachines()
	ResetVpcAccess()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CloudGcpDmIntegrations
type jsiiProxy_CloudGcpDmIntegrations struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_CloudGcpDmIntegrations) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) AiPlatform() CloudGcpDmIntegrationsAiPlatformOutputReference {
	var returns CloudGcpDmIntegrationsAiPlatformOutputReference
	_jsii_.Get(
		j,
		"aiPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) AiPlatformInput() *CloudGcpDmIntegrationsAiPlatform {
	var returns *CloudGcpDmIntegrationsAiPlatform
	_jsii_.Get(
		j,
		"aiPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) AlloyDb() CloudGcpDmIntegrationsAlloyDbOutputReference {
	var returns CloudGcpDmIntegrationsAlloyDbOutputReference
	_jsii_.Get(
		j,
		"alloyDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) AlloyDbInput() *CloudGcpDmIntegrationsAlloyDb {
	var returns *CloudGcpDmIntegrationsAlloyDb
	_jsii_.Get(
		j,
		"alloyDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) ApiGateway() CloudGcpDmIntegrationsApiGatewayOutputReference {
	var returns CloudGcpDmIntegrationsApiGatewayOutputReference
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) ApiGatewayInput() *CloudGcpDmIntegrationsApiGateway {
	var returns *CloudGcpDmIntegrationsApiGateway
	_jsii_.Get(
		j,
		"apiGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) AppEngine() CloudGcpDmIntegrationsAppEngineOutputReference {
	var returns CloudGcpDmIntegrationsAppEngineOutputReference
	_jsii_.Get(
		j,
		"appEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) AppEngineInput() *CloudGcpDmIntegrationsAppEngine {
	var returns *CloudGcpDmIntegrationsAppEngine
	_jsii_.Get(
		j,
		"appEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) BigQuery() CloudGcpDmIntegrationsBigQueryOutputReference {
	var returns CloudGcpDmIntegrationsBigQueryOutputReference
	_jsii_.Get(
		j,
		"bigQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) BigQueryInput() *CloudGcpDmIntegrationsBigQuery {
	var returns *CloudGcpDmIntegrationsBigQuery
	_jsii_.Get(
		j,
		"bigQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) BigTable() CloudGcpDmIntegrationsBigTableOutputReference {
	var returns CloudGcpDmIntegrationsBigTableOutputReference
	_jsii_.Get(
		j,
		"bigTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) BigTableInput() *CloudGcpDmIntegrationsBigTable {
	var returns *CloudGcpDmIntegrationsBigTable
	_jsii_.Get(
		j,
		"bigTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Composer() CloudGcpDmIntegrationsComposerOutputReference {
	var returns CloudGcpDmIntegrationsComposerOutputReference
	_jsii_.Get(
		j,
		"composer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) ComposerInput() *CloudGcpDmIntegrationsComposer {
	var returns *CloudGcpDmIntegrationsComposer
	_jsii_.Get(
		j,
		"composerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) DataFlow() CloudGcpDmIntegrationsDataFlowOutputReference {
	var returns CloudGcpDmIntegrationsDataFlowOutputReference
	_jsii_.Get(
		j,
		"dataFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) DataFlowInput() *CloudGcpDmIntegrationsDataFlow {
	var returns *CloudGcpDmIntegrationsDataFlow
	_jsii_.Get(
		j,
		"dataFlowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) DataProc() CloudGcpDmIntegrationsDataProcOutputReference {
	var returns CloudGcpDmIntegrationsDataProcOutputReference
	_jsii_.Get(
		j,
		"dataProc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) DataProcInput() *CloudGcpDmIntegrationsDataProc {
	var returns *CloudGcpDmIntegrationsDataProc
	_jsii_.Get(
		j,
		"dataProcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) DataStore() CloudGcpDmIntegrationsDataStoreOutputReference {
	var returns CloudGcpDmIntegrationsDataStoreOutputReference
	_jsii_.Get(
		j,
		"dataStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) DataStoreInput() *CloudGcpDmIntegrationsDataStore {
	var returns *CloudGcpDmIntegrationsDataStore
	_jsii_.Get(
		j,
		"dataStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FirebaseAppHosting() CloudGcpDmIntegrationsFirebaseAppHostingOutputReference {
	var returns CloudGcpDmIntegrationsFirebaseAppHostingOutputReference
	_jsii_.Get(
		j,
		"firebaseAppHosting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FirebaseAppHostingInput() *CloudGcpDmIntegrationsFirebaseAppHosting {
	var returns *CloudGcpDmIntegrationsFirebaseAppHosting
	_jsii_.Get(
		j,
		"firebaseAppHostingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FirebaseAuth() CloudGcpDmIntegrationsFirebaseAuthOutputReference {
	var returns CloudGcpDmIntegrationsFirebaseAuthOutputReference
	_jsii_.Get(
		j,
		"firebaseAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FirebaseAuthInput() *CloudGcpDmIntegrationsFirebaseAuth {
	var returns *CloudGcpDmIntegrationsFirebaseAuth
	_jsii_.Get(
		j,
		"firebaseAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FirebaseDatabase() CloudGcpDmIntegrationsFirebaseDatabaseOutputReference {
	var returns CloudGcpDmIntegrationsFirebaseDatabaseOutputReference
	_jsii_.Get(
		j,
		"firebaseDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FirebaseDatabaseInput() *CloudGcpDmIntegrationsFirebaseDatabase {
	var returns *CloudGcpDmIntegrationsFirebaseDatabase
	_jsii_.Get(
		j,
		"firebaseDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FirebaseHosting() CloudGcpDmIntegrationsFirebaseHostingOutputReference {
	var returns CloudGcpDmIntegrationsFirebaseHostingOutputReference
	_jsii_.Get(
		j,
		"firebaseHosting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FirebaseHostingInput() *CloudGcpDmIntegrationsFirebaseHosting {
	var returns *CloudGcpDmIntegrationsFirebaseHosting
	_jsii_.Get(
		j,
		"firebaseHostingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FirebaseStorage() CloudGcpDmIntegrationsFirebaseStorageOutputReference {
	var returns CloudGcpDmIntegrationsFirebaseStorageOutputReference
	_jsii_.Get(
		j,
		"firebaseStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FirebaseStorageInput() *CloudGcpDmIntegrationsFirebaseStorage {
	var returns *CloudGcpDmIntegrationsFirebaseStorage
	_jsii_.Get(
		j,
		"firebaseStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FirebaseVertexAi() CloudGcpDmIntegrationsFirebaseVertexAiOutputReference {
	var returns CloudGcpDmIntegrationsFirebaseVertexAiOutputReference
	_jsii_.Get(
		j,
		"firebaseVertexAi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FirebaseVertexAiInput() *CloudGcpDmIntegrationsFirebaseVertexAi {
	var returns *CloudGcpDmIntegrationsFirebaseVertexAi
	_jsii_.Get(
		j,
		"firebaseVertexAiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Firestore() CloudGcpDmIntegrationsFirestoreOutputReference {
	var returns CloudGcpDmIntegrationsFirestoreOutputReference
	_jsii_.Get(
		j,
		"firestore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FirestoreInput() *CloudGcpDmIntegrationsFirestore {
	var returns *CloudGcpDmIntegrationsFirestore
	_jsii_.Get(
		j,
		"firestoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Functions() CloudGcpDmIntegrationsFunctionsOutputReference {
	var returns CloudGcpDmIntegrationsFunctionsOutputReference
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) FunctionsInput() *CloudGcpDmIntegrationsFunctions {
	var returns *CloudGcpDmIntegrationsFunctions
	_jsii_.Get(
		j,
		"functionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Interconnect() CloudGcpDmIntegrationsInterconnectOutputReference {
	var returns CloudGcpDmIntegrationsInterconnectOutputReference
	_jsii_.Get(
		j,
		"interconnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) InterconnectInput() *CloudGcpDmIntegrationsInterconnect {
	var returns *CloudGcpDmIntegrationsInterconnect
	_jsii_.Get(
		j,
		"interconnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Istio() CloudGcpDmIntegrationsIstioOutputReference {
	var returns CloudGcpDmIntegrationsIstioOutputReference
	_jsii_.Get(
		j,
		"istio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) IstioInput() *CloudGcpDmIntegrationsIstio {
	var returns *CloudGcpDmIntegrationsIstio
	_jsii_.Get(
		j,
		"istioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Kubernetes() CloudGcpDmIntegrationsKubernetesOutputReference {
	var returns CloudGcpDmIntegrationsKubernetesOutputReference
	_jsii_.Get(
		j,
		"kubernetes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) KubernetesInput() *CloudGcpDmIntegrationsKubernetes {
	var returns *CloudGcpDmIntegrationsKubernetes
	_jsii_.Get(
		j,
		"kubernetesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) LinkedAccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"linkedAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) LinkedAccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"linkedAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) LoadBalancing() CloudGcpDmIntegrationsLoadBalancingOutputReference {
	var returns CloudGcpDmIntegrationsLoadBalancingOutputReference
	_jsii_.Get(
		j,
		"loadBalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) LoadBalancingInput() *CloudGcpDmIntegrationsLoadBalancing {
	var returns *CloudGcpDmIntegrationsLoadBalancing
	_jsii_.Get(
		j,
		"loadBalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) ManagedKafka() CloudGcpDmIntegrationsManagedKafkaOutputReference {
	var returns CloudGcpDmIntegrationsManagedKafkaOutputReference
	_jsii_.Get(
		j,
		"managedKafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) ManagedKafkaInput() *CloudGcpDmIntegrationsManagedKafka {
	var returns *CloudGcpDmIntegrationsManagedKafka
	_jsii_.Get(
		j,
		"managedKafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) MemCache() CloudGcpDmIntegrationsMemCacheOutputReference {
	var returns CloudGcpDmIntegrationsMemCacheOutputReference
	_jsii_.Get(
		j,
		"memCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) MemCacheInput() *CloudGcpDmIntegrationsMemCache {
	var returns *CloudGcpDmIntegrationsMemCache
	_jsii_.Get(
		j,
		"memCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) MemoryStore() CloudGcpDmIntegrationsMemoryStoreOutputReference {
	var returns CloudGcpDmIntegrationsMemoryStoreOutputReference
	_jsii_.Get(
		j,
		"memoryStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) MemoryStoreInput() *CloudGcpDmIntegrationsMemoryStore {
	var returns *CloudGcpDmIntegrationsMemoryStore
	_jsii_.Get(
		j,
		"memoryStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) PubSub() CloudGcpDmIntegrationsPubSubOutputReference {
	var returns CloudGcpDmIntegrationsPubSubOutputReference
	_jsii_.Get(
		j,
		"pubSub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) PubSubInput() *CloudGcpDmIntegrationsPubSub {
	var returns *CloudGcpDmIntegrationsPubSub
	_jsii_.Get(
		j,
		"pubSubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Redis() CloudGcpDmIntegrationsRedisOutputReference {
	var returns CloudGcpDmIntegrationsRedisOutputReference
	_jsii_.Get(
		j,
		"redis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) RedisInput() *CloudGcpDmIntegrationsRedis {
	var returns *CloudGcpDmIntegrationsRedis
	_jsii_.Get(
		j,
		"redisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Router() CloudGcpDmIntegrationsRouterOutputReference {
	var returns CloudGcpDmIntegrationsRouterOutputReference
	_jsii_.Get(
		j,
		"router",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) RouterInput() *CloudGcpDmIntegrationsRouter {
	var returns *CloudGcpDmIntegrationsRouter
	_jsii_.Get(
		j,
		"routerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Run() CloudGcpDmIntegrationsRunOutputReference {
	var returns CloudGcpDmIntegrationsRunOutputReference
	_jsii_.Get(
		j,
		"run",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) RunInput() *CloudGcpDmIntegrationsRun {
	var returns *CloudGcpDmIntegrationsRun
	_jsii_.Get(
		j,
		"runInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Spanner() CloudGcpDmIntegrationsSpannerOutputReference {
	var returns CloudGcpDmIntegrationsSpannerOutputReference
	_jsii_.Get(
		j,
		"spanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) SpannerInput() *CloudGcpDmIntegrationsSpanner {
	var returns *CloudGcpDmIntegrationsSpanner
	_jsii_.Get(
		j,
		"spannerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Sql() CloudGcpDmIntegrationsSqlOutputReference {
	var returns CloudGcpDmIntegrationsSqlOutputReference
	_jsii_.Get(
		j,
		"sql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) SqlInput() *CloudGcpDmIntegrationsSql {
	var returns *CloudGcpDmIntegrationsSql
	_jsii_.Get(
		j,
		"sqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) Storage() CloudGcpDmIntegrationsStorageOutputReference {
	var returns CloudGcpDmIntegrationsStorageOutputReference
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) StorageInput() *CloudGcpDmIntegrationsStorage {
	var returns *CloudGcpDmIntegrationsStorage
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) VirtualMachines() CloudGcpDmIntegrationsVirtualMachinesOutputReference {
	var returns CloudGcpDmIntegrationsVirtualMachinesOutputReference
	_jsii_.Get(
		j,
		"virtualMachines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) VirtualMachinesInput() *CloudGcpDmIntegrationsVirtualMachines {
	var returns *CloudGcpDmIntegrationsVirtualMachines
	_jsii_.Get(
		j,
		"virtualMachinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) VpcAccess() CloudGcpDmIntegrationsVpcAccessOutputReference {
	var returns CloudGcpDmIntegrationsVpcAccessOutputReference
	_jsii_.Get(
		j,
		"vpcAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpDmIntegrations) VpcAccessInput() *CloudGcpDmIntegrationsVpcAccess {
	var returns *CloudGcpDmIntegrationsVpcAccess
	_jsii_.Get(
		j,
		"vpcAccessInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/cloud_gcp_dm_integrations newrelic_cloud_gcp_dm_integrations} Resource.
func NewCloudGcpDmIntegrations(scope constructs.Construct, id *string, config *CloudGcpDmIntegrationsConfig) CloudGcpDmIntegrations {
	_init_.Initialize()

	if err := validateNewCloudGcpDmIntegrationsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudGcpDmIntegrations{}

	_jsii_.Create(
		"@cdktn/provider-newrelic.cloudGcpDmIntegrations.CloudGcpDmIntegrations",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/cloud_gcp_dm_integrations newrelic_cloud_gcp_dm_integrations} Resource.
func NewCloudGcpDmIntegrations_Override(c CloudGcpDmIntegrations, scope constructs.Construct, id *string, config *CloudGcpDmIntegrationsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-newrelic.cloudGcpDmIntegrations.CloudGcpDmIntegrations",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudGcpDmIntegrations)SetAccountId(val *float64) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_CloudGcpDmIntegrations)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudGcpDmIntegrations)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudGcpDmIntegrations)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudGcpDmIntegrations)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudGcpDmIntegrations)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudGcpDmIntegrations)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudGcpDmIntegrations)SetLinkedAccountId(val *float64) {
	if err := j.validateSetLinkedAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkedAccountId",
		val,
	)
}

func (j *jsiiProxy_CloudGcpDmIntegrations)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudGcpDmIntegrations)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a CloudGcpDmIntegrations resource upon running "cdktn plan <stack-name>".
func CloudGcpDmIntegrations_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateCloudGcpDmIntegrations_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-newrelic.cloudGcpDmIntegrations.CloudGcpDmIntegrations",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CloudGcpDmIntegrations_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudGcpDmIntegrations_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-newrelic.cloudGcpDmIntegrations.CloudGcpDmIntegrations",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudGcpDmIntegrations_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudGcpDmIntegrations_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-newrelic.cloudGcpDmIntegrations.CloudGcpDmIntegrations",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudGcpDmIntegrations_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudGcpDmIntegrations_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-newrelic.cloudGcpDmIntegrations.CloudGcpDmIntegrations",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudGcpDmIntegrations_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-newrelic.cloudGcpDmIntegrations.CloudGcpDmIntegrations",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := c.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutAiPlatform(value *CloudGcpDmIntegrationsAiPlatform) {
	if err := c.validatePutAiPlatformParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAiPlatform",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutAlloyDb(value *CloudGcpDmIntegrationsAlloyDb) {
	if err := c.validatePutAlloyDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAlloyDb",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutApiGateway(value *CloudGcpDmIntegrationsApiGateway) {
	if err := c.validatePutApiGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putApiGateway",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutAppEngine(value *CloudGcpDmIntegrationsAppEngine) {
	if err := c.validatePutAppEngineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAppEngine",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutBigQuery(value *CloudGcpDmIntegrationsBigQuery) {
	if err := c.validatePutBigQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBigQuery",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutBigTable(value *CloudGcpDmIntegrationsBigTable) {
	if err := c.validatePutBigTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBigTable",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutComposer(value *CloudGcpDmIntegrationsComposer) {
	if err := c.validatePutComposerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putComposer",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutDataFlow(value *CloudGcpDmIntegrationsDataFlow) {
	if err := c.validatePutDataFlowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataFlow",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutDataProc(value *CloudGcpDmIntegrationsDataProc) {
	if err := c.validatePutDataProcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataProc",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutDataStore(value *CloudGcpDmIntegrationsDataStore) {
	if err := c.validatePutDataStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataStore",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutFirebaseAppHosting(value *CloudGcpDmIntegrationsFirebaseAppHosting) {
	if err := c.validatePutFirebaseAppHostingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFirebaseAppHosting",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutFirebaseAuth(value *CloudGcpDmIntegrationsFirebaseAuth) {
	if err := c.validatePutFirebaseAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFirebaseAuth",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutFirebaseDatabase(value *CloudGcpDmIntegrationsFirebaseDatabase) {
	if err := c.validatePutFirebaseDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFirebaseDatabase",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutFirebaseHosting(value *CloudGcpDmIntegrationsFirebaseHosting) {
	if err := c.validatePutFirebaseHostingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFirebaseHosting",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutFirebaseStorage(value *CloudGcpDmIntegrationsFirebaseStorage) {
	if err := c.validatePutFirebaseStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFirebaseStorage",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutFirebaseVertexAi(value *CloudGcpDmIntegrationsFirebaseVertexAi) {
	if err := c.validatePutFirebaseVertexAiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFirebaseVertexAi",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutFirestore(value *CloudGcpDmIntegrationsFirestore) {
	if err := c.validatePutFirestoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFirestore",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutFunctions(value *CloudGcpDmIntegrationsFunctions) {
	if err := c.validatePutFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFunctions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutInterconnect(value *CloudGcpDmIntegrationsInterconnect) {
	if err := c.validatePutInterconnectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInterconnect",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutIstio(value *CloudGcpDmIntegrationsIstio) {
	if err := c.validatePutIstioParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIstio",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutKubernetes(value *CloudGcpDmIntegrationsKubernetes) {
	if err := c.validatePutKubernetesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKubernetes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutLoadBalancing(value *CloudGcpDmIntegrationsLoadBalancing) {
	if err := c.validatePutLoadBalancingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLoadBalancing",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutManagedKafka(value *CloudGcpDmIntegrationsManagedKafka) {
	if err := c.validatePutManagedKafkaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putManagedKafka",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutMemCache(value *CloudGcpDmIntegrationsMemCache) {
	if err := c.validatePutMemCacheParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMemCache",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutMemoryStore(value *CloudGcpDmIntegrationsMemoryStore) {
	if err := c.validatePutMemoryStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMemoryStore",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutPubSub(value *CloudGcpDmIntegrationsPubSub) {
	if err := c.validatePutPubSubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPubSub",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutRedis(value *CloudGcpDmIntegrationsRedis) {
	if err := c.validatePutRedisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRedis",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutRouter(value *CloudGcpDmIntegrationsRouter) {
	if err := c.validatePutRouterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRouter",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutRun(value *CloudGcpDmIntegrationsRun) {
	if err := c.validatePutRunParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRun",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutSpanner(value *CloudGcpDmIntegrationsSpanner) {
	if err := c.validatePutSpannerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSpanner",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutSql(value *CloudGcpDmIntegrationsSql) {
	if err := c.validatePutSqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSql",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutStorage(value *CloudGcpDmIntegrationsStorage) {
	if err := c.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStorage",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutVirtualMachines(value *CloudGcpDmIntegrationsVirtualMachines) {
	if err := c.validatePutVirtualMachinesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVirtualMachines",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) PutVpcAccess(value *CloudGcpDmIntegrationsVpcAccess) {
	if err := c.validatePutVpcAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVpcAccess",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := c.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetAccountId() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetAiPlatform() {
	_jsii_.InvokeVoid(
		c,
		"resetAiPlatform",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetAlloyDb() {
	_jsii_.InvokeVoid(
		c,
		"resetAlloyDb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetApiGateway() {
	_jsii_.InvokeVoid(
		c,
		"resetApiGateway",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetAppEngine() {
	_jsii_.InvokeVoid(
		c,
		"resetAppEngine",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetBigQuery() {
	_jsii_.InvokeVoid(
		c,
		"resetBigQuery",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetBigTable() {
	_jsii_.InvokeVoid(
		c,
		"resetBigTable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetComposer() {
	_jsii_.InvokeVoid(
		c,
		"resetComposer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetDataFlow() {
	_jsii_.InvokeVoid(
		c,
		"resetDataFlow",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetDataProc() {
	_jsii_.InvokeVoid(
		c,
		"resetDataProc",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetDataStore() {
	_jsii_.InvokeVoid(
		c,
		"resetDataStore",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetFirebaseAppHosting() {
	_jsii_.InvokeVoid(
		c,
		"resetFirebaseAppHosting",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetFirebaseAuth() {
	_jsii_.InvokeVoid(
		c,
		"resetFirebaseAuth",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetFirebaseDatabase() {
	_jsii_.InvokeVoid(
		c,
		"resetFirebaseDatabase",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetFirebaseHosting() {
	_jsii_.InvokeVoid(
		c,
		"resetFirebaseHosting",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetFirebaseStorage() {
	_jsii_.InvokeVoid(
		c,
		"resetFirebaseStorage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetFirebaseVertexAi() {
	_jsii_.InvokeVoid(
		c,
		"resetFirebaseVertexAi",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetFirestore() {
	_jsii_.InvokeVoid(
		c,
		"resetFirestore",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetFunctions() {
	_jsii_.InvokeVoid(
		c,
		"resetFunctions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetInterconnect() {
	_jsii_.InvokeVoid(
		c,
		"resetInterconnect",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetIstio() {
	_jsii_.InvokeVoid(
		c,
		"resetIstio",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetKubernetes() {
	_jsii_.InvokeVoid(
		c,
		"resetKubernetes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetLoadBalancing() {
	_jsii_.InvokeVoid(
		c,
		"resetLoadBalancing",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetManagedKafka() {
	_jsii_.InvokeVoid(
		c,
		"resetManagedKafka",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetMemCache() {
	_jsii_.InvokeVoid(
		c,
		"resetMemCache",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetMemoryStore() {
	_jsii_.InvokeVoid(
		c,
		"resetMemoryStore",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetPubSub() {
	_jsii_.InvokeVoid(
		c,
		"resetPubSub",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetRedis() {
	_jsii_.InvokeVoid(
		c,
		"resetRedis",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetRouter() {
	_jsii_.InvokeVoid(
		c,
		"resetRouter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetRun() {
	_jsii_.InvokeVoid(
		c,
		"resetRun",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetSpanner() {
	_jsii_.InvokeVoid(
		c,
		"resetSpanner",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetSql() {
	_jsii_.InvokeVoid(
		c,
		"resetSql",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetStorage() {
	_jsii_.InvokeVoid(
		c,
		"resetStorage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetVirtualMachines() {
	_jsii_.InvokeVoid(
		c,
		"resetVirtualMachines",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ResetVpcAccess() {
	_jsii_.InvokeVoid(
		c,
		"resetVpcAccess",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpDmIntegrations) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpDmIntegrations) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

