// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudawseusovereignintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v14/cloudawseusovereignintegrations/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/cloud_aws_eu_sovereign_integrations newrelic_cloud_aws_eu_sovereign_integrations}.
type CloudAwsEuSovereignIntegrations interface {
	cdktn.TerraformResource
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	Billing() CloudAwsEuSovereignIntegrationsBillingOutputReference
	BillingInput() *CloudAwsEuSovereignIntegrationsBilling
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Cloudtrail() CloudAwsEuSovereignIntegrationsCloudtrailOutputReference
	CloudtrailInput() *CloudAwsEuSovereignIntegrationsCloudtrail
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LinkedAccountId() *float64
	SetLinkedAccountId(val *float64)
	LinkedAccountIdInput() *float64
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
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	XRay() CloudAwsEuSovereignIntegrationsXRayOutputReference
	XRayInput() *CloudAwsEuSovereignIntegrationsXRay
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
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
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
	PutBilling(value *CloudAwsEuSovereignIntegrationsBilling)
	PutCloudtrail(value *CloudAwsEuSovereignIntegrationsCloudtrail)
	PutXRay(value *CloudAwsEuSovereignIntegrationsXRay)
	ResetAccountId()
	ResetBilling()
	ResetCloudtrail()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetXRay()
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
}

// The jsii proxy struct for CloudAwsEuSovereignIntegrations
type jsiiProxy_CloudAwsEuSovereignIntegrations struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) Billing() CloudAwsEuSovereignIntegrationsBillingOutputReference {
	var returns CloudAwsEuSovereignIntegrationsBillingOutputReference
	_jsii_.Get(
		j,
		"billing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) BillingInput() *CloudAwsEuSovereignIntegrationsBilling {
	var returns *CloudAwsEuSovereignIntegrationsBilling
	_jsii_.Get(
		j,
		"billingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) Cloudtrail() CloudAwsEuSovereignIntegrationsCloudtrailOutputReference {
	var returns CloudAwsEuSovereignIntegrationsCloudtrailOutputReference
	_jsii_.Get(
		j,
		"cloudtrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) CloudtrailInput() *CloudAwsEuSovereignIntegrationsCloudtrail {
	var returns *CloudAwsEuSovereignIntegrationsCloudtrail
	_jsii_.Get(
		j,
		"cloudtrailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) LinkedAccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"linkedAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) LinkedAccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"linkedAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) XRay() CloudAwsEuSovereignIntegrationsXRayOutputReference {
	var returns CloudAwsEuSovereignIntegrationsXRayOutputReference
	_jsii_.Get(
		j,
		"xRay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations) XRayInput() *CloudAwsEuSovereignIntegrationsXRay {
	var returns *CloudAwsEuSovereignIntegrationsXRay
	_jsii_.Get(
		j,
		"xRayInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/cloud_aws_eu_sovereign_integrations newrelic_cloud_aws_eu_sovereign_integrations} Resource.
func NewCloudAwsEuSovereignIntegrations(scope constructs.Construct, id *string, config *CloudAwsEuSovereignIntegrationsConfig) CloudAwsEuSovereignIntegrations {
	_init_.Initialize()

	if err := validateNewCloudAwsEuSovereignIntegrationsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudAwsEuSovereignIntegrations{}

	_jsii_.Create(
		"@cdktn/provider-newrelic.cloudAwsEuSovereignIntegrations.CloudAwsEuSovereignIntegrations",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/cloud_aws_eu_sovereign_integrations newrelic_cloud_aws_eu_sovereign_integrations} Resource.
func NewCloudAwsEuSovereignIntegrations_Override(c CloudAwsEuSovereignIntegrations, scope constructs.Construct, id *string, config *CloudAwsEuSovereignIntegrationsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-newrelic.cloudAwsEuSovereignIntegrations.CloudAwsEuSovereignIntegrations",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations)SetAccountId(val *float64) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations)SetLinkedAccountId(val *float64) {
	if err := j.validateSetLinkedAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkedAccountId",
		val,
	)
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudAwsEuSovereignIntegrations)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a CloudAwsEuSovereignIntegrations resource upon running "cdktn plan <stack-name>".
func CloudAwsEuSovereignIntegrations_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateCloudAwsEuSovereignIntegrations_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-newrelic.cloudAwsEuSovereignIntegrations.CloudAwsEuSovereignIntegrations",
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
func CloudAwsEuSovereignIntegrations_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAwsEuSovereignIntegrations_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-newrelic.cloudAwsEuSovereignIntegrations.CloudAwsEuSovereignIntegrations",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudAwsEuSovereignIntegrations_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAwsEuSovereignIntegrations_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-newrelic.cloudAwsEuSovereignIntegrations.CloudAwsEuSovereignIntegrations",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudAwsEuSovereignIntegrations_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAwsEuSovereignIntegrations_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-newrelic.cloudAwsEuSovereignIntegrations.CloudAwsEuSovereignIntegrations",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudAwsEuSovereignIntegrations_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-newrelic.cloudAwsEuSovereignIntegrations.CloudAwsEuSovereignIntegrations",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) PutBilling(value *CloudAwsEuSovereignIntegrationsBilling) {
	if err := c.validatePutBillingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBilling",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) PutCloudtrail(value *CloudAwsEuSovereignIntegrationsCloudtrail) {
	if err := c.validatePutCloudtrailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCloudtrail",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) PutXRay(value *CloudAwsEuSovereignIntegrationsXRay) {
	if err := c.validatePutXRayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putXRay",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) ResetAccountId() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) ResetBilling() {
	_jsii_.InvokeVoid(
		c,
		"resetBilling",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) ResetCloudtrail() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudtrail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) ResetXRay() {
	_jsii_.InvokeVoid(
		c,
		"resetXRay",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsEuSovereignIntegrations) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

