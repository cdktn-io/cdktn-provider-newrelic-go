// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetStackedBarChartStyles struct {
	// gradient block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.81.0/docs/resources/one_dashboard#gradient OneDashboard#gradient}
	Gradient *OneDashboardPageWidgetStackedBarChartStylesGradient `field:"optional" json:"gradient" yaml:"gradient"`
	// Line interpolation style. Valid values: 'linear', 'smooth', 'stepBefore', 'stepAfter'. Applicable to line widgets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.81.0/docs/resources/one_dashboard#line_interpolation OneDashboard#line_interpolation}
	LineInterpolation *string `field:"optional" json:"lineInterpolation" yaml:"lineInterpolation"`
}

