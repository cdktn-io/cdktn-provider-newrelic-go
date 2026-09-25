// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v16/pathpointflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PathpointFlowStagesStageKpisQueryTimeWindowOutputReference interface {
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
	CustomRange() *string
	SetCustomRange(val *string)
	CustomRangeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *PathpointFlowStagesStageKpisQueryTimeWindow
	SetInternalValue(val *PathpointFlowStagesStageKpisQueryTimeWindow)
	RelativeRange() PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference
	RelativeRangeInput() *PathpointFlowStagesStageKpisQueryTimeWindowRelativeRange
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
	PutRelativeRange(value *PathpointFlowStagesStageKpisQueryTimeWindowRelativeRange)
	ResetCustomRange()
	ResetRelativeRange()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PathpointFlowStagesStageKpisQueryTimeWindowOutputReference
type jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) CustomRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) CustomRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) InternalValue() *PathpointFlowStagesStageKpisQueryTimeWindow {
	var returns *PathpointFlowStagesStageKpisQueryTimeWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) RelativeRange() PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference {
	var returns PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference
	_jsii_.Get(
		j,
		"relativeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) RelativeRangeInput() *PathpointFlowStagesStageKpisQueryTimeWindowRelativeRange {
	var returns *PathpointFlowStagesStageKpisQueryTimeWindowRelativeRange
	_jsii_.Get(
		j,
		"relativeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPathpointFlowStagesStageKpisQueryTimeWindowOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PathpointFlowStagesStageKpisQueryTimeWindowOutputReference {
	_init_.Initialize()

	if err := validateNewPathpointFlowStagesStageKpisQueryTimeWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-newrelic.pathpointFlow.PathpointFlowStagesStageKpisQueryTimeWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPathpointFlowStagesStageKpisQueryTimeWindowOutputReference_Override(p PathpointFlowStagesStageKpisQueryTimeWindowOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-newrelic.pathpointFlow.PathpointFlowStagesStageKpisQueryTimeWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference)SetCustomRange(val *string) {
	if err := j.validateSetCustomRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customRange",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference)SetInternalValue(val *PathpointFlowStagesStageKpisQueryTimeWindow) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) PutRelativeRange(value *PathpointFlowStagesStageKpisQueryTimeWindowRelativeRange) {
	if err := p.validatePutRelativeRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRelativeRange",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) ResetCustomRange() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomRange",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) ResetRelativeRange() {
	_jsii_.InvokeVoid(
		p,
		"resetRelativeRange",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

