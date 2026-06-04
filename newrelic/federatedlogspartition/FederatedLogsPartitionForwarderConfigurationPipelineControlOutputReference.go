// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogspartition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v15/federatedlogspartition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference interface {
	cdktn.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *FederatedLogsPartitionForwarderConfigurationPipelineControl
	SetInternalValue(val *FederatedLogsPartitionForwarderConfigurationPipelineControl)
	PartitionRule() FederatedLogsPartitionForwarderConfigurationPipelineControlPartitionRuleOutputReference
	PartitionRuleInput() *FederatedLogsPartitionForwarderConfigurationPipelineControlPartitionRule
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
	PutPartitionRule(value *FederatedLogsPartitionForwarderConfigurationPipelineControlPartitionRule)
	ResetPartitionRule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference
type jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) InternalValue() *FederatedLogsPartitionForwarderConfigurationPipelineControl {
	var returns *FederatedLogsPartitionForwarderConfigurationPipelineControl
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) PartitionRule() FederatedLogsPartitionForwarderConfigurationPipelineControlPartitionRuleOutputReference {
	var returns FederatedLogsPartitionForwarderConfigurationPipelineControlPartitionRuleOutputReference
	_jsii_.Get(
		j,
		"partitionRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) PartitionRuleInput() *FederatedLogsPartitionForwarderConfigurationPipelineControlPartitionRule {
	var returns *FederatedLogsPartitionForwarderConfigurationPipelineControlPartitionRule
	_jsii_.Get(
		j,
		"partitionRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference {
	_init_.Initialize()

	if err := validateNewFederatedLogsPartitionForwarderConfigurationPipelineControlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-newrelic.federatedLogsPartition.FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference_Override(f FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-newrelic.federatedLogsPartition.FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference)SetInternalValue(val *FederatedLogsPartitionForwarderConfigurationPipelineControl) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) PutPartitionRule(value *FederatedLogsPartitionForwarderConfigurationPipelineControlPartitionRule) {
	if err := f.validatePutPartitionRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putPartitionRule",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) ResetPartitionRule() {
	_jsii_.InvokeVoid(
		f,
		"resetPartitionRule",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (f *jsiiProxy_FederatedLogsPartitionForwarderConfigurationPipelineControlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

