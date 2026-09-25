// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v16/pathpointflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PathpointFlowStagesLevelsStepsConfigOutputReference interface {
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
	HealthRollup() *string
	SetHealthRollup(val *string)
	HealthRollupInput() *string
	InternalValue() *PathpointFlowStagesLevelsStepsConfig
	SetInternalValue(val *PathpointFlowStagesLevelsStepsConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThresholdType() *string
	SetThresholdType(val *string)
	ThresholdTypeInput() *string
	ThresholdValue() *float64
	SetThresholdValue(val *float64)
	ThresholdValueInput() *float64
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
	ResetHealthRollup()
	ResetThresholdType()
	ResetThresholdValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PathpointFlowStagesLevelsStepsConfigOutputReference
type jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) HealthRollup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthRollup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) HealthRollupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthRollupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) InternalValue() *PathpointFlowStagesLevelsStepsConfig {
	var returns *PathpointFlowStagesLevelsStepsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) ThresholdType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) ThresholdTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) ThresholdValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) ThresholdValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdValueInput",
		&returns,
	)
	return returns
}


func NewPathpointFlowStagesLevelsStepsConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PathpointFlowStagesLevelsStepsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewPathpointFlowStagesLevelsStepsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-newrelic.pathpointFlow.PathpointFlowStagesLevelsStepsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPathpointFlowStagesLevelsStepsConfigOutputReference_Override(p PathpointFlowStagesLevelsStepsConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-newrelic.pathpointFlow.PathpointFlowStagesLevelsStepsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference)SetHealthRollup(val *string) {
	if err := j.validateSetHealthRollupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthRollup",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference)SetInternalValue(val *PathpointFlowStagesLevelsStepsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference)SetThresholdType(val *string) {
	if err := j.validateSetThresholdTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdType",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference)SetThresholdValue(val *float64) {
	if err := j.validateSetThresholdValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdValue",
		val,
	)
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) ResetHealthRollup() {
	_jsii_.InvokeVoid(
		p,
		"resetHealthRollup",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) ResetThresholdType() {
	_jsii_.InvokeVoid(
		p,
		"resetThresholdType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) ResetThresholdValue() {
	_jsii_.InvokeVoid(
		p,
		"resetThresholdValue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesLevelsStepsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

