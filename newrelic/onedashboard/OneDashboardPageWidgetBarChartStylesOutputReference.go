// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package onedashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v14/jsii"

	"github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v14/onedashboard/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OneDashboardPageWidgetBarChartStylesOutputReference interface {
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
	Gradient() OneDashboardPageWidgetBarChartStylesGradientOutputReference
	GradientInput() *OneDashboardPageWidgetBarChartStylesGradient
	InternalValue() *OneDashboardPageWidgetBarChartStyles
	SetInternalValue(val *OneDashboardPageWidgetBarChartStyles)
	LineInterpolation() *string
	SetLineInterpolation(val *string)
	LineInterpolationInput() *string
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
	PutGradient(value *OneDashboardPageWidgetBarChartStylesGradient)
	ResetGradient()
	ResetLineInterpolation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OneDashboardPageWidgetBarChartStylesOutputReference
type jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) Gradient() OneDashboardPageWidgetBarChartStylesGradientOutputReference {
	var returns OneDashboardPageWidgetBarChartStylesGradientOutputReference
	_jsii_.Get(
		j,
		"gradient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) GradientInput() *OneDashboardPageWidgetBarChartStylesGradient {
	var returns *OneDashboardPageWidgetBarChartStylesGradient
	_jsii_.Get(
		j,
		"gradientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) InternalValue() *OneDashboardPageWidgetBarChartStyles {
	var returns *OneDashboardPageWidgetBarChartStyles
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) LineInterpolation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lineInterpolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) LineInterpolationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lineInterpolationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOneDashboardPageWidgetBarChartStylesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OneDashboardPageWidgetBarChartStylesOutputReference {
	_init_.Initialize()

	if err := validateNewOneDashboardPageWidgetBarChartStylesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-newrelic.oneDashboard.OneDashboardPageWidgetBarChartStylesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOneDashboardPageWidgetBarChartStylesOutputReference_Override(o OneDashboardPageWidgetBarChartStylesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-newrelic.oneDashboard.OneDashboardPageWidgetBarChartStylesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference)SetInternalValue(val *OneDashboardPageWidgetBarChartStyles) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference)SetLineInterpolation(val *string) {
	if err := j.validateSetLineInterpolationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lineInterpolation",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) PutGradient(value *OneDashboardPageWidgetBarChartStylesGradient) {
	if err := o.validatePutGradientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGradient",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) ResetGradient() {
	_jsii_.InvokeVoid(
		o,
		"resetGradient",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) ResetLineInterpolation() {
	_jsii_.InvokeVoid(
		o,
		"resetLineInterpolation",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarChartStylesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

