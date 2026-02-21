// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticsstepmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v14/syntheticsstepmonitor/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/synthetics_step_monitor newrelic_synthetics_step_monitor}.
type SyntheticsStepMonitor interface {
	cdktn.TerraformResource
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	Browsers() *[]*string
	SetBrowsers(val *[]*string)
	BrowsersInput() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	Devices() *[]*string
	SetDevices(val *[]*string)
	DevicesInput() *[]*string
	EnableScreenshotOnFailureAndScript() interface{}
	SetEnableScreenshotOnFailureAndScript(val interface{})
	EnableScreenshotOnFailureAndScriptInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Guid() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LocationPrivate() SyntheticsStepMonitorLocationPrivateList
	LocationPrivateInput() interface{}
	LocationsPublic() *[]*string
	SetLocationsPublic(val *[]*string)
	LocationsPublicInput() *[]*string
	MonitorId() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Period() *string
	SetPeriod(val *string)
	PeriodInMinutes() *float64
	PeriodInput() *string
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
	RuntimeType() *string
	SetRuntimeType(val *string)
	RuntimeTypeInput() *string
	RuntimeTypeVersion() *string
	SetRuntimeTypeVersion(val *string)
	RuntimeTypeVersionInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Steps() SyntheticsStepMonitorStepsList
	StepsInput() interface{}
	Tag() SyntheticsStepMonitorTagList
	TagInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UseUnsupportedLegacyRuntime() interface{}
	SetUseUnsupportedLegacyRuntime(val interface{})
	UseUnsupportedLegacyRuntimeInput() interface{}
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
	PutLocationPrivate(value interface{})
	PutSteps(value interface{})
	PutTag(value interface{})
	ResetAccountId()
	ResetBrowsers()
	ResetDevices()
	ResetEnableScreenshotOnFailureAndScript()
	ResetId()
	ResetLocationPrivate()
	ResetLocationsPublic()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRuntimeType()
	ResetRuntimeTypeVersion()
	ResetTag()
	ResetUseUnsupportedLegacyRuntime()
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

// The jsii proxy struct for SyntheticsStepMonitor
type jsiiProxy_SyntheticsStepMonitor struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SyntheticsStepMonitor) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Browsers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"browsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) BrowsersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"browsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Devices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"devices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) DevicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"devicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) EnableScreenshotOnFailureAndScript() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableScreenshotOnFailureAndScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) EnableScreenshotOnFailureAndScriptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableScreenshotOnFailureAndScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Guid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) LocationPrivate() SyntheticsStepMonitorLocationPrivateList {
	var returns SyntheticsStepMonitorLocationPrivateList
	_jsii_.Get(
		j,
		"locationPrivate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) LocationPrivateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locationPrivateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) LocationsPublic() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsPublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) LocationsPublicInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsPublicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) MonitorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Period() *string {
	var returns *string
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) PeriodInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) PeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) RuntimeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) RuntimeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) RuntimeTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) RuntimeTypeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeTypeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Steps() SyntheticsStepMonitorStepsList {
	var returns SyntheticsStepMonitorStepsList
	_jsii_.Get(
		j,
		"steps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) StepsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) Tag() SyntheticsStepMonitorTagList {
	var returns SyntheticsStepMonitorTagList
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) TagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) UseUnsupportedLegacyRuntime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useUnsupportedLegacyRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsStepMonitor) UseUnsupportedLegacyRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useUnsupportedLegacyRuntimeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/synthetics_step_monitor newrelic_synthetics_step_monitor} Resource.
func NewSyntheticsStepMonitor(scope constructs.Construct, id *string, config *SyntheticsStepMonitorConfig) SyntheticsStepMonitor {
	_init_.Initialize()

	if err := validateNewSyntheticsStepMonitorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsStepMonitor{}

	_jsii_.Create(
		"@cdktn/provider-newrelic.syntheticsStepMonitor.SyntheticsStepMonitor",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/synthetics_step_monitor newrelic_synthetics_step_monitor} Resource.
func NewSyntheticsStepMonitor_Override(s SyntheticsStepMonitor, scope constructs.Construct, id *string, config *SyntheticsStepMonitorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-newrelic.syntheticsStepMonitor.SyntheticsStepMonitor",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetAccountId(val *float64) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetBrowsers(val *[]*string) {
	if err := j.validateSetBrowsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"browsers",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetDevices(val *[]*string) {
	if err := j.validateSetDevicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"devices",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetEnableScreenshotOnFailureAndScript(val interface{}) {
	if err := j.validateSetEnableScreenshotOnFailureAndScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableScreenshotOnFailureAndScript",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetLocationsPublic(val *[]*string) {
	if err := j.validateSetLocationsPublicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationsPublic",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetPeriod(val *string) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetRuntimeType(val *string) {
	if err := j.validateSetRuntimeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeType",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetRuntimeTypeVersion(val *string) {
	if err := j.validateSetRuntimeTypeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeTypeVersion",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_SyntheticsStepMonitor)SetUseUnsupportedLegacyRuntime(val interface{}) {
	if err := j.validateSetUseUnsupportedLegacyRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useUnsupportedLegacyRuntime",
		val,
	)
}

// Generates CDKTN code for importing a SyntheticsStepMonitor resource upon running "cdktn plan <stack-name>".
func SyntheticsStepMonitor_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSyntheticsStepMonitor_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-newrelic.syntheticsStepMonitor.SyntheticsStepMonitor",
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
func SyntheticsStepMonitor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsStepMonitor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-newrelic.syntheticsStepMonitor.SyntheticsStepMonitor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsStepMonitor_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsStepMonitor_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-newrelic.syntheticsStepMonitor.SyntheticsStepMonitor",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsStepMonitor_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsStepMonitor_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-newrelic.syntheticsStepMonitor.SyntheticsStepMonitor",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SyntheticsStepMonitor_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-newrelic.syntheticsStepMonitor.SyntheticsStepMonitor",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) PutLocationPrivate(value interface{}) {
	if err := s.validatePutLocationPrivateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLocationPrivate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) PutSteps(value interface{}) {
	if err := s.validatePutStepsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSteps",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) PutTag(value interface{}) {
	if err := s.validatePutTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTag",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) ResetAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) ResetBrowsers() {
	_jsii_.InvokeVoid(
		s,
		"resetBrowsers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) ResetDevices() {
	_jsii_.InvokeVoid(
		s,
		"resetDevices",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) ResetEnableScreenshotOnFailureAndScript() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableScreenshotOnFailureAndScript",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) ResetLocationPrivate() {
	_jsii_.InvokeVoid(
		s,
		"resetLocationPrivate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) ResetLocationsPublic() {
	_jsii_.InvokeVoid(
		s,
		"resetLocationsPublic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) ResetRuntimeType() {
	_jsii_.InvokeVoid(
		s,
		"resetRuntimeType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) ResetRuntimeTypeVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetRuntimeTypeVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) ResetTag() {
	_jsii_.InvokeVoid(
		s,
		"resetTag",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) ResetUseUnsupportedLegacyRuntime() {
	_jsii_.InvokeVoid(
		s,
		"resetUseUnsupportedLegacyRuntime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsStepMonitor) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsStepMonitor) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

