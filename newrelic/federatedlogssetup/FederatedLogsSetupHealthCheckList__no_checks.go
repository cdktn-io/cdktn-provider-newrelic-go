// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package federatedlogssetup

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FederatedLogsSetupHealthCheckList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FederatedLogsSetupHealthCheckList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FederatedLogsSetupHealthCheckList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FederatedLogsSetupHealthCheckList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FederatedLogsSetupHealthCheckList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FederatedLogsSetupHealthCheckList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFederatedLogsSetupHealthCheckListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

