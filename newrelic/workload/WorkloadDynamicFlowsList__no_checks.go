// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workload

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkloadDynamicFlowsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkloadDynamicFlowsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkloadDynamicFlowsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkloadDynamicFlowsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkloadDynamicFlowsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkloadDynamicFlowsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkloadDynamicFlowsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkloadDynamicFlowsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

