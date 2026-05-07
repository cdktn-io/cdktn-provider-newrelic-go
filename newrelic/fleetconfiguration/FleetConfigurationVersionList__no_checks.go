// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package fleetconfiguration

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FleetConfigurationVersionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FleetConfigurationVersionList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FleetConfigurationVersionList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FleetConfigurationVersionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FleetConfigurationVersionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FleetConfigurationVersionList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FleetConfigurationVersionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFleetConfigurationVersionListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

