// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package federatedlogssetup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-newrelic-go/newrelic/v15/federatedlogssetup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FederatedLogsSetupHealthCheckEnd2EndDataFlowList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) FederatedLogsSetupHealthCheckEnd2EndDataFlowOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FederatedLogsSetupHealthCheckEnd2EndDataFlowList
type jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFederatedLogsSetupHealthCheckEnd2EndDataFlowList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FederatedLogsSetupHealthCheckEnd2EndDataFlowList {
	_init_.Initialize()

	if err := validateNewFederatedLogsSetupHealthCheckEnd2EndDataFlowListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList{}

	_jsii_.Create(
		"@cdktn/provider-newrelic.federatedLogsSetup.FederatedLogsSetupHealthCheckEnd2EndDataFlowList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFederatedLogsSetupHealthCheckEnd2EndDataFlowList_Override(f FederatedLogsSetupHealthCheckEnd2EndDataFlowList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-newrelic.federatedLogsSetup.FederatedLogsSetupHealthCheckEnd2EndDataFlowList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := f.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		f,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList) Get(index *float64) FederatedLogsSetupHealthCheckEnd2EndDataFlowOutputReference {
	if err := f.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns FederatedLogsSetupHealthCheckEnd2EndDataFlowOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (f *jsiiProxy_FederatedLogsSetupHealthCheckEnd2EndDataFlowList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

