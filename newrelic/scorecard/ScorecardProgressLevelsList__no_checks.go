// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package scorecard

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScorecardProgressLevelsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ScorecardProgressLevelsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ScorecardProgressLevelsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ScorecardProgressLevelsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ScorecardProgressLevelsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ScorecardProgressLevelsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ScorecardProgressLevelsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewScorecardProgressLevelsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

