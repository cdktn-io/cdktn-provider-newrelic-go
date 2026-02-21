// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package obfuscationrule

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObfuscationRuleActionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_ObfuscationRuleActionList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_ObfuscationRuleActionList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ObfuscationRuleActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ObfuscationRuleActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObfuscationRuleActionList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ObfuscationRuleActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewObfuscationRuleActionListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

