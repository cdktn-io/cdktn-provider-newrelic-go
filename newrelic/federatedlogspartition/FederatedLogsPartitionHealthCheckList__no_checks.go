// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package federatedlogspartition

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FederatedLogsPartitionHealthCheckList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FederatedLogsPartitionHealthCheckList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FederatedLogsPartitionHealthCheckList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FederatedLogsPartitionHealthCheckList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FederatedLogsPartitionHealthCheckList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FederatedLogsPartitionHealthCheckList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFederatedLogsPartitionHealthCheckListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

