// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v16/pathpointflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PathpointFlowStagesStageKpisQueryOutputReference interface {
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
	From() *string
	SetFrom(val *string)
	FromInput() *string
	InternalValue() *PathpointFlowStagesStageKpisQuery
	SetInternalValue(val *PathpointFlowStagesStageKpisQuery)
	Select() PathpointFlowStagesStageKpisQuerySelectOutputReference
	SelectInput() *PathpointFlowStagesStageKpisQuerySelect
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeWindow() PathpointFlowStagesStageKpisQueryTimeWindowOutputReference
	TimeWindowInput() *PathpointFlowStagesStageKpisQueryTimeWindow
	Where() *string
	SetWhere(val *string)
	WhereInput() *string
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
	PutSelect(value *PathpointFlowStagesStageKpisQuerySelect)
	PutTimeWindow(value *PathpointFlowStagesStageKpisQueryTimeWindow)
	ResetTimeWindow()
	ResetWhere()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PathpointFlowStagesStageKpisQueryOutputReference
type jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) From() *string {
	var returns *string
	_jsii_.Get(
		j,
		"from",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) FromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) InternalValue() *PathpointFlowStagesStageKpisQuery {
	var returns *PathpointFlowStagesStageKpisQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) Select() PathpointFlowStagesStageKpisQuerySelectOutputReference {
	var returns PathpointFlowStagesStageKpisQuerySelectOutputReference
	_jsii_.Get(
		j,
		"select",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) SelectInput() *PathpointFlowStagesStageKpisQuerySelect {
	var returns *PathpointFlowStagesStageKpisQuerySelect
	_jsii_.Get(
		j,
		"selectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) TimeWindow() PathpointFlowStagesStageKpisQueryTimeWindowOutputReference {
	var returns PathpointFlowStagesStageKpisQueryTimeWindowOutputReference
	_jsii_.Get(
		j,
		"timeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) TimeWindowInput() *PathpointFlowStagesStageKpisQueryTimeWindow {
	var returns *PathpointFlowStagesStageKpisQueryTimeWindow
	_jsii_.Get(
		j,
		"timeWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) Where() *string {
	var returns *string
	_jsii_.Get(
		j,
		"where",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) WhereInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whereInput",
		&returns,
	)
	return returns
}


func NewPathpointFlowStagesStageKpisQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PathpointFlowStagesStageKpisQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPathpointFlowStagesStageKpisQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-newrelic.pathpointFlow.PathpointFlowStagesStageKpisQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPathpointFlowStagesStageKpisQueryOutputReference_Override(p PathpointFlowStagesStageKpisQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-newrelic.pathpointFlow.PathpointFlowStagesStageKpisQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference)SetFrom(val *string) {
	if err := j.validateSetFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"from",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference)SetInternalValue(val *PathpointFlowStagesStageKpisQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference)SetWhere(val *string) {
	if err := j.validateSetWhereParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"where",
		val,
	)
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) PutSelect(value *PathpointFlowStagesStageKpisQuerySelect) {
	if err := p.validatePutSelectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSelect",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) PutTimeWindow(value *PathpointFlowStagesStageKpisQueryTimeWindow) {
	if err := p.validatePutTimeWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeWindow",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) ResetTimeWindow() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeWindow",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) ResetWhere() {
	_jsii_.InvokeVoid(
		p,
		"resetWhere",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PathpointFlowStagesStageKpisQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

