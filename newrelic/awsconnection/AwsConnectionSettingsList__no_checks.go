// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package awsconnection

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsConnectionSettingsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsConnectionSettingsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsConnectionSettingsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsConnectionSettingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsConnectionSettingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsConnectionSettingsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsConnectionSettingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsConnectionSettingsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

