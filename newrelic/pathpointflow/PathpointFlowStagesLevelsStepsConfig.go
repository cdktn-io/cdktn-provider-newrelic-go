// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pathpointflow


type PathpointFlowStagesLevelsStepsConfig struct {
	// How step health is rolled up: BEST_STATUS_WINS or WORST_STATUS_WINS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#health_rollup PathpointFlow#health_rollup}
	HealthRollup *string `field:"optional" json:"healthRollup" yaml:"healthRollup"`
	// Whether threshold is FIXED or PERCENTAGE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#threshold_type PathpointFlow#threshold_type}
	ThresholdType *string `field:"optional" json:"thresholdType" yaml:"thresholdType"`
	// Numeric threshold value for step health evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.99.3/docs/resources/pathpoint_flow#threshold_value PathpointFlow#threshold_value}
	ThresholdValue *float64 `field:"optional" json:"thresholdValue" yaml:"thresholdValue"`
}

