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
	err := testutils_cobra.ExecutePingcli(t, "config", "list-keys", options.RootColorOption.KoanfKey)
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
	// - export.pingOne.environmentID
	// - export.serviceGroup
	// - export.services
	// - noColor
	// - outputFormat
	// - request.accessToken
	// - request.accessTokenExpiry
	// - request.fail
	// - request.service
	// - service.pingFederate.adminAPIPath
	// - service.pingFederate.authentication.accessTokenAuth.accessToken
	// - service.pingFederate.authentication.basicAuth.password
	// - service.pingFederate.authentication.basicAuth.username
	// - service.pingFederate.authentication.clientCredentialsAuth.clientID
	// - service.pingFederate.authentication.clientCredentialsAuth.clientSecret
	// - service.pingFederate.authentication.clientCredentialsAuth.scopes
	// - service.pingFederate.authentication.clientCredentialsAuth.tokenURL
	// - service.pingFederate.authentication.type
	// - service.pingFederate.caCertificatePEMFiles
	// - service.pingFederate.httpsHost
	// - service.pingFederate.insecureTrustAllTLS
	// - service.pingFederate.xBypassExternalValidationHeader
	// - service.pingOne.authentication.type
	// - service.pingOne.authentication.worker.clientID
	// - service.pingOne.authentication.worker.clientSecret
	// - service.pingOne.authentication.worker.environmentID
	// - service.pingOne.regionCode
}
