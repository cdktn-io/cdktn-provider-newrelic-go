// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsettings


type ApplicationSettingsTransactionTracerSql struct {
	// The level of SQL recording, either 'OBFUSCATED', 'OFF', or 'RAW'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.80.2/docs/resources/application_settings#record_sql ApplicationSettings#record_sql}
	RecordSql *string `field:"required" json:"recordSql" yaml:"recordSql"`
}

