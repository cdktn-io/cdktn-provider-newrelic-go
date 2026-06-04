// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package awsconnection

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsConnectionTagList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsConnectionTagList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsConnectionTagList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsConnectionTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsConnectionTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsConnectionTagList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsConnectionTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsConnectionTagListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

