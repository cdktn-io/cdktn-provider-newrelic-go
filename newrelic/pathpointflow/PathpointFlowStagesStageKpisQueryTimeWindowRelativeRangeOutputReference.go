// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v16/pathpointflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference interface {
	cdktn.ComplexObject
	CompareAgainst() *string
	SetCompareAgainst(val *string)
	CompareAgainstInput() *string
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
	InternalValue() *PathpointFlowStagesStageKpisQueryTimeWindowRelativeRange
	SetInternalValue(val *PathpointFlowStagesStageKpisQueryTimeWindowRelativeRange)
	Since() *string
	SetSince(val *string)
	SinceInput() *string
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
	ResetCompareAgainst()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference
type jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) CompareAgainst() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareAgainst",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) CompareAgainstInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareAgainstInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) InternalValue() *PathpointFlowStagesStageKpisQueryTimeWindowRelativeRange {
	var returns *PathpointFlowStagesStageKpisQueryTimeWindowRelativeRange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) Since() *string {
	var returns *string
	_jsii_.Get(
		j,
		"since",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) SinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference {
	_init_.Initialize()

	if err := validateNewPathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-newrelic.pathpointFlow.PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference_Override(p PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-newrelic.pathpointFlow.PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference)SetCompareAgainst(val *string) {
	if err := j.validateSetCompareAgainstParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compareAgainst",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference)SetInternalValue(val *PathpointFlowStagesStageKpisQueryTimeWindowRelativeRange) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference)SetSince(val *string) {
	if err := j.validateSetSinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"since",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) ResetCompareAgainst() {
	_jsii_.InvokeVoid(
		p,
		"resetCompareAgainst",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryTimeWindowRelativeRangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

