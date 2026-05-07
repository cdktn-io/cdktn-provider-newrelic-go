// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package fleetdeployment

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FleetDeploymentAgentList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FleetDeploymentAgentList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FleetDeploymentAgentList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FleetDeploymentAgentList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FleetDeploymentAgentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FleetDeploymentAgentList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FleetDeploymentAgentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFleetDeploymentAgentListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

