// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package scorecardrule

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScorecardRuleTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ScorecardRuleTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ScorecardRuleTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ScorecardRuleTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ScorecardRuleTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ScorecardRuleTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ScorecardRuleTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewScorecardRuleTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

