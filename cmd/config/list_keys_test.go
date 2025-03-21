// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Config List Keys Command Executes without issue
func TestConfigListKeysCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "list-keys")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config List Keys YAML Command --yaml, -y flag
func TestConfigListKeysCmd_YAMLFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "list-keys", "--yaml")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "config", "list-keys", "-y")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config List Keys Command --help, -h flag
func TestConfigListKeysCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "list-keys", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "config", "list-keys", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config List Keys Command fails when provided too many arguments
func TestConfigListKeysCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli config list-keys': command accepts 0 arg\(s\), received 1$`
	err := testutils_cobra.ExecutePingcli(t, "config", "list-keys", options.RootColorOption.ViperKey)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_listKeysValue() {
	t := testing.T{}
	_ = testutils_cobra.ExecutePingcli(&t, "config", "list-keys")

	// Output:
	// Valid Keys:
	// - activeProfile
	// - description
	// - detailedExitCode
	// - export.format
	// - export.outputDirectory
	// - export.overwrite
	// - export.pingone.environmentID
	// - export.serviceGroup
	// - export.services
	// - noColor
	// - outputFormat
	// - request.accessToken
	// - request.accessTokenExpiry
	// - request.fail
	// - request.service
	// - service.pingfederate.adminAPIPath
	// - service.pingfederate.authentication.accessTokenAuth.accessToken
	// - service.pingfederate.authentication.basicAuth.password
	// - service.pingfederate.authentication.basicAuth.username
	// - service.pingfederate.authentication.clientCredentialsAuth.clientID
	// - service.pingfederate.authentication.clientCredentialsAuth.clientSecret
	// - service.pingfederate.authentication.clientCredentialsAuth.scopes
	// - service.pingfederate.authentication.clientCredentialsAuth.tokenURL
	// - service.pingfederate.authentication.type
	// - service.pingfederate.caCertificatePemFiles
	// - service.pingfederate.httpsHost
	// - service.pingfederate.insecureTrustAllTLS
	// - service.pingfederate.xBypassExternalValidationHeader
	// - service.pingone.authentication.type
	// - service.pingone.authentication.worker.clientID
	// - service.pingone.authentication.worker.clientSecret
	// - service.pingone.authentication.worker.environmentID
	// - service.pingone.regionCode
}
