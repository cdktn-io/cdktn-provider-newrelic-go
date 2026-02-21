// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetJsonColors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/one_dashboard#color OneDashboard#color}.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// series_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/one_dashboard#series_overrides OneDashboard#series_overrides}
	SeriesOverrides interface{} `field:"optional" json:"seriesOverrides" yaml:"seriesOverrides"`
}

