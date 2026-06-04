// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package federatedlogssetup

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FederatedLogsSetupLifecycleStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FederatedLogsSetupLifecycleStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FederatedLogsSetupLifecycleStatusList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FederatedLogsSetupLifecycleStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FederatedLogsSetupLifecycleStatusList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FederatedLogsSetupLifecycleStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFederatedLogsSetupLifecycleStatusListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

