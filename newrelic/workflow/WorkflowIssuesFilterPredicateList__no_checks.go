// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workflow

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkflowIssuesFilterPredicateList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkflowIssuesFilterPredicateList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkflowIssuesFilterPredicateList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkflowIssuesFilterPredicateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkflowIssuesFilterPredicateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkflowIssuesFilterPredicateList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkflowIssuesFilterPredicateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkflowIssuesFilterPredicateListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

