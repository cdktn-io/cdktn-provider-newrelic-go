// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scorecardrule


type ScorecardRuleNrqlEngine struct {
	// Account IDs where this rule runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#accounts ScorecardRule#accounts}
	Accounts *[]*float64 `field:"required" json:"accounts" yaml:"accounts"`
	// The NRQL query. Must be FACET-ed and alias the result as 'score'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#query ScorecardRule#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// Additional account IDs to join with the query accounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.101.6/docs/resources/scorecard_rule#join_accounts ScorecardRule#join_accounts}
	JoinAccounts *[]*float64 `field:"optional" json:"joinAccounts" yaml:"joinAccounts"`
}

