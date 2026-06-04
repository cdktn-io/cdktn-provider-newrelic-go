// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogssetup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v15/federatedlogssetup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FederatedLogsSetupStorageOutputReference interface {
	cdktn.ComplexObject
	CloudProviderConfiguration() FederatedLogsSetupStorageCloudProviderConfigurationOutputReference
	CloudProviderConfigurationInput() *FederatedLogsSetupStorageCloudProviderConfiguration
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	DataIngestConnectionId() *string
	SetDataIngestConnectionId(val *string)
	DataIngestConnectionIdInput() *string
	DataLocationBucket() *string
	SetDataLocationBucket(val *string)
	DataLocationBucketInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *FederatedLogsSetupStorage
	SetInternalValue(val *FederatedLogsSetupStorage)
	QueryConnectionId() *string
	SetQueryConnectionId(val *string)
	QueryConnectionIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutCloudProviderConfiguration(value *FederatedLogsSetupStorageCloudProviderConfiguration)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FederatedLogsSetupStorageOutputReference
type jsiiProxy_FederatedLogsSetupStorageOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) CloudProviderConfiguration() FederatedLogsSetupStorageCloudProviderConfigurationOutputReference {
	var returns FederatedLogsSetupStorageCloudProviderConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudProviderConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) CloudProviderConfigurationInput() *FederatedLogsSetupStorageCloudProviderConfiguration {
	var returns *FederatedLogsSetupStorageCloudProviderConfiguration
	_jsii_.Get(
		j,
		"cloudProviderConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) DataIngestConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataIngestConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) DataIngestConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataIngestConnectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) DataLocationBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLocationBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) DataLocationBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLocationBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) InternalValue() *FederatedLogsSetupStorage {
	var returns *FederatedLogsSetupStorage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) QueryConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) QueryConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryConnectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFederatedLogsSetupStorageOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FederatedLogsSetupStorageOutputReference {
	_init_.Initialize()

	if err := validateNewFederatedLogsSetupStorageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FederatedLogsSetupStorageOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-newrelic.federatedLogsSetup.FederatedLogsSetupStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFederatedLogsSetupStorageOutputReference_Override(f FederatedLogsSetupStorageOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-newrelic.federatedLogsSetup.FederatedLogsSetupStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference)SetDataIngestConnectionId(val *string) {
	if err := j.validateSetDataIngestConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataIngestConnectionId",
		val,
	)
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference)SetDataLocationBucket(val *string) {
	if err := j.validateSetDataLocationBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataLocationBucket",
		val,
	)
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference)SetInternalValue(val *FederatedLogsSetupStorage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference)SetQueryConnectionId(val *string) {
	if err := j.validateSetQueryConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryConnectionId",
		val,
	)
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FederatedLogsSetupStorageOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) PutCloudProviderConfiguration(value *FederatedLogsSetupStorageCloudProviderConfiguration) {
	if err := f.validatePutCloudProviderConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putCloudProviderConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupStorageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

