// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v16/pathpointflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PathpointFlowStagesOutputReference interface {
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
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsExcluded() interface{}
	SetIsExcluded(val interface{})
	IsExcludedInput() interface{}
	Levels() PathpointFlowStagesLevelsList
	LevelsInput() interface{}
	Link() *string
	SetLink(val *string)
	LinkInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Related() PathpointFlowStagesRelatedOutputReference
	RelatedInput() *PathpointFlowStagesRelated
	StageKpis() PathpointFlowStagesStageKpisList
	StageKpisInput() interface{}
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
	PutLevels(value interface{})
	PutRelated(value *PathpointFlowStagesRelated)
	PutStageKpis(value interface{})
	ResetHealthRollup()
	ResetIsExcluded()
	ResetLevels()
	ResetLink()
	ResetRelated()
	ResetStageKpis()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PathpointFlowStagesOutputReference
type jsiiProxy_PathpointFlowStagesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) HealthRollup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthRollup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) HealthRollupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthRollupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) IsExcluded() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isExcluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) IsExcludedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isExcludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) Levels() PathpointFlowStagesLevelsList {
	var returns PathpointFlowStagesLevelsList
	_jsii_.Get(
		j,
		"levels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) LevelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"levelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) Link() *string {
	var returns *string
	_jsii_.Get(
		j,
		"link",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) LinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) Related() PathpointFlowStagesRelatedOutputReference {
	var returns PathpointFlowStagesRelatedOutputReference
	_jsii_.Get(
		j,
		"related",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) RelatedInput() *PathpointFlowStagesRelated {
	var returns *PathpointFlowStagesRelated
	_jsii_.Get(
		j,
		"relatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) StageKpis() PathpointFlowStagesStageKpisList {
	var returns PathpointFlowStagesStageKpisList
	_jsii_.Get(
		j,
		"stageKpis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) StageKpisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stageKpisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPathpointFlowStagesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PathpointFlowStagesOutputReference {
	_init_.Initialize()

	if err := validateNewPathpointFlowStagesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PathpointFlowStagesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-newrelic.pathpointFlow.PathpointFlowStagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPathpointFlowStagesOutputReference_Override(p PathpointFlowStagesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-newrelic.pathpointFlow.PathpointFlowStagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference)SetHealthRollup(val *string) {
	if err := j.validateSetHealthRollupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthRollup",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference)SetIsExcluded(val interface{}) {
	if err := j.validateSetIsExcludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isExcluded",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference)SetLink(val *string) {
	if err := j.validateSetLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"link",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PathpointFlowStagesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PathpointFlowStagesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PathpointFlowStagesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PathpointFlowStagesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PathpointFlowStagesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PathpointFlowStagesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PathpointFlowStagesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PathpointFlowStagesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PathpointFlowStagesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PathpointFlowStagesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PathpointFlowStagesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PathpointFlowStagesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PathpointFlowStagesOutputReference) PutLevels(value interface{}) {
	if err := p.validatePutLevelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLevels",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PathpointFlowStagesOutputReference) PutRelated(value *PathpointFlowStagesRelated) {
	if err := p.validatePutRelatedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRelated",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PathpointFlowStagesOutputReference) PutStageKpis(value interface{}) {
	if err := p.validatePutStageKpisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putStageKpis",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PathpointFlowStagesOutputReference) ResetHealthRollup() {
	_jsii_.InvokeVoid(
		p,
		"resetHealthRollup",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PathpointFlowStagesOutputReference) ResetIsExcluded() {
	_jsii_.InvokeVoid(
		p,
		"resetIsExcluded",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PathpointFlowStagesOutputReference) ResetLevels() {
	_jsii_.InvokeVoid(
		p,
		"resetLevels",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PathpointFlowStagesOutputReference) ResetLink() {
	_jsii_.InvokeVoid(
		p,
		"resetLink",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PathpointFlowStagesOutputReference) ResetRelated() {
	_jsii_.InvokeVoid(
		p,
		"resetRelated",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PathpointFlowStagesOutputReference) ResetStageKpis() {
	_jsii_.InvokeVoid(
		p,
		"resetStageKpis",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PathpointFlowStagesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PathpointFlowStagesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

