// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cardinalitymanagement

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CardinalityManagementMetricList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CardinalityManagementMetricList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CardinalityManagementMetricList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CardinalityManagementMetricList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CardinalityManagementMetricList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CardinalityManagementMetricList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CardinalityManagementMetricList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCardinalityManagementMetricListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

