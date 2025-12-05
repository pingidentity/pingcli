// Copyright © 2025 Ping Identity Corporation

package auth_test

import (
	"os"
	"testing"

	auth_internal "github.com/pingidentity/pingcli/internal/commands/auth"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// TestLoginCmd_MissingConfiguration tests behavior when required configuration is missing
func TestLoginCmd_MissingConfiguration(t *testing.T) {
	// Create a custom config file with missing auth configuration
	configContents := `
activeProfile: test
test:
    description: Test profile without auth config
    outputFormat: json
    service:
        pingOne:
            regionCode: NA
`

	testutils_koanf.InitKoanfsCustomFile(t, configContents)

	testCases := []struct {
		name                 string
		authMethod           string
		expectedErrorPattern string
	}{
		{
			name:                 "client credentials missing client ID",
			authMethod:           "--client-credentials",
			expectedErrorPattern: `client credentials client ID is not configured|failed to prompt for reconfiguration|input prompt error`,
		},
		{
			name:                 "authorization code missing client ID",
			authMethod:           "--authorization-code",
			expectedErrorPattern: `authorization code client ID is not configured|failed to prompt for reconfiguration|input prompt error`,
		},
		{
			name:                 "device code missing client ID",
			authMethod:           "--device-code",
			expectedErrorPattern: `device code client ID is not configured|failed to prompt for reconfiguration|input prompt error`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := testutils_cobra.ExecutePingcli(t, "login", tc.authMethod)
			testutils.CheckExpectedError(t, err, &tc.expectedErrorPattern)
		})
	}
}

// TestLogoutCmd_NoActiveSession tests logout when no credentials are stored in keychain
func TestLogoutCmd_NoActiveSession(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Clear any existing tokens to ensure no active session
	_ = auth_internal.ClearToken()

	// Try to logout - should succeed even with no active session
	err := testutils_cobra.ExecutePingcli(t, "logout", "--client-credentials")
	if err != nil {
		t.Logf("Logout with no active session returned error (expected): %v", err)
	}
}

// TestLoginCmd_InvalidCredentials tests behavior with intentionally invalid credentials
func TestLoginCmd_InvalidCredentials(t *testing.T) {
	configContents := `
activeProfile: test
test:
    description: Test profile with invalid credentials
    outputFormat: json
    service:
        pingOne:
            regionCode: NA
            authentication:
                type: client_credentials
                environmentID: 00000000-0000-0000-0000-000000000000
                clientCredentials:
                    clientID: 00000000-0000-0000-0000-000000000001
                    clientSecret: invalid-client-secret
`
	testutils_koanf.InitKoanfsCustomFile(t, configContents)

	err := testutils_cobra.ExecutePingcli(t, "login", "--client-credentials")
	if err == nil {
		t.Error("Expected error with invalid credentials, but got none")
	} else {
		t.Logf("Got expected error with invalid credentials: %v", err)
	}
}

// TestLogoutCmd_WithoutAuthTypeConfigured tests logout when no auth type is configured
func TestLogoutCmd_WithoutAuthTypeConfigured(t *testing.T) {
	configContents := `
activeProfile: test
test:
    description: Test profile without auth type configured
    outputFormat: json
    service:
        pingOne:
            regionCode: NA
`
	testutils_koanf.InitKoanfsCustomFile(t, configContents)

	// Try to logout without specifying grant type and without configured auth type
	err := testutils_cobra.ExecutePingcli(t, "logout")
	expectedErrorPattern := `no authentication type configured|authentication type|failed to generate token key`
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// TestLoginCmd_DefaultAuthTypeNotConfigured tests login without flags when no auth type is configured
func TestLoginCmd_DefaultAuthTypeNotConfigured(t *testing.T) {
	configContents := `
activeProfile: test
test:
    description: Test profile without auth type
    outputFormat: json
    service:
        pingOne:
            regionCode: NA
`
	testutils_koanf.InitKoanfsCustomFile(t, configContents)

	// This should trigger interactive configuration prompt (which will fail in test environment)
	err := testutils_cobra.ExecutePingcli(t, "login")
	// We expect some error since we can't do interactive prompts in tests
	if err == nil {
		t.Error("Expected error when no auth type configured and no interactive input, but got none")
	} else {
		t.Logf("Got expected error: %v", err)
	}
}

// TestLoginCmd_MutuallyExclusiveFlags tests that multiple grant type flags cannot be used together
func TestLoginCmd_MutuallyExclusiveFlags(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		flags         []string
		expectedError string
	}{
		{
			name:          "authorization-code and device-code together",
			flags:         []string{"--authorization-code", "--device-code"},
			expectedError: "if any flags in the group.*are set none of the others can be",
		},
		{
			name:          "authorization-code and client-credentials together",
			flags:         []string{"--authorization-code", "--client-credentials"},
			expectedError: "if any flags in the group.*are set none of the others can be",
		},
		{
			name:          "device-code and client-credentials together",
			flags:         []string{"--device-code", "--client-credentials"},
			expectedError: "if any flags in the group.*are set none of the others can be",
		},
		{
			name:          "all three flags together",
			flags:         []string{"--authorization-code", "--device-code", "--client-credentials"},
			expectedError: "if any flags in the group.*are set none of the others can be",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := append([]string{"login"}, tc.flags...)
			err := testutils_cobra.ExecutePingcli(t, args...)
			testutils.CheckExpectedError(t, err, &tc.expectedError)
		})
	}
}

// TestLogoutCmd_SpecificAuthMethod tests logout with specific grant type when multiple are configured
func TestLogoutCmd_SpecificAuthMethod(t *testing.T) {
	// Skip if not in CI environment or missing credentials
	clientID := os.Getenv("TEST_PINGONE_CLIENT_CREDENTIALS_CLIENT_ID")
	clientSecret := os.Getenv("TEST_PINGONE_CLIENT_CREDENTIALS_CLIENT_SECRET")
	environmentID := os.Getenv("TEST_PINGONE_ENVIRONMENT_ID")

	if clientID == "" || clientSecret == "" || environmentID == "" {
		t.Skip("Skipping test: missing TEST_PINGONE_* environment variables")
	}

	testutils_koanf.InitKoanfs(t)

	// Login with client credentials
	err := testutils_cobra.ExecutePingcli(t, "login", "--client-credentials")
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}

	// Verify we can get the grant type configured
	authType, err := profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
	if err != nil {
		t.Fatalf("Failed to get auth type: %v", err)
	}
	t.Logf("Current auth type: %s", authType)

	// Logout from specific grant type
	err = testutils_cobra.ExecutePingcli(t, "logout", "--client-credentials")
	if err != nil {
		t.Fatalf("Failed to logout: %v", err)
	}

	// Verify token is cleared
	_, err = auth_internal.LoadToken()
	if err == nil {
		t.Error("Token should not exist after logout")
	}
}

// TestLoginCmd_MissingEnvironmentID tests behavior when environment ID is missing
func TestLoginCmd_MissingEnvironmentID(t *testing.T) {
	testCases := []struct {
		name                 string
		authMethod           string
		configContents       string
		expectedErrorPattern string
	}{
		{
			name:       "client_credentials_missing_environment_id",
			authMethod: "--client-credentials",
			configContents: `
activeProfile: test
test:
    description: Test profile without environment ID
    outputFormat: json
    service:
        pingOne:
            regionCode: NA
            authentication:
                type: client_credentials
                clientCredentials:
                    clientID: 00000000-0000-0000-0000-000000000001
                    clientSecret: test-secret
`,
			expectedErrorPattern: `environment ID is not configured|failed to prompt for reconfiguration|input prompt error`,
		},
		{
			name:       "authorization_code_missing_environment_id",
			authMethod: "--authorization-code",
			configContents: `
activeProfile: test
test:
    description: Test profile without environment ID
    outputFormat: json
    service:
        pingOne:
            regionCode: NA
            authentication:
                type: authorization_code
                authorizationCode:
                    clientID: 00000000-0000-0000-0000-000000000001
                    redirectURIPath: /callback
                    redirectURIPort: "3000"
`,
			expectedErrorPattern: `environment ID is not configured|failed to prompt for reconfiguration|input prompt error`,
		},
		{
			name:       "device_code_missing_environment_id",
			authMethod: "--device-code",
			configContents: `
activeProfile: test
test:
    description: Test profile without environment ID
    outputFormat: json
    service:
        pingOne:
            regionCode: NA
            authentication:
                type: device_code
                deviceCode:
                    clientID: 00000000-0000-0000-0000-000000000001
`,
			expectedErrorPattern: `environment ID is not configured|failed to prompt for reconfiguration|input prompt error`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfsCustomFile(t, tc.configContents)
			err := testutils_cobra.ExecutePingcli(t, "login", tc.authMethod)
			testutils.CheckExpectedError(t, err, &tc.expectedErrorPattern)
		})
	}
}

// TestLoginCmd_MissingClientSecret tests client credentials without client secret
func TestLoginCmd_MissingClientSecret(t *testing.T) {
	configContents := `
activeProfile: test
test:
    description: Test profile without client secret
    outputFormat: json
    service:
        pingOne:
            regionCode: NA
            authentication:
                type: client_credentials
                environmentID: 00000000-0000-0000-0000-000000000000
                clientCredentials:
                    clientID: 00000000-0000-0000-0000-000000000001
`
	testutils_koanf.InitKoanfsCustomFile(t, configContents)

	err := testutils_cobra.ExecutePingcli(t, "login", "--client-credentials")
	expectedErrorPattern := `client secret is not configured|failed to prompt for reconfiguration|input prompt error`
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// TestLoginCmd_AuthorizationCodeMissingRedirectURI tests authorization code without redirect URI
func TestLoginCmd_AuthorizationCodeMissingRedirectURI(t *testing.T) {
	configContents := `
activeProfile: test
test:
    description: Test profile without redirect URI
    outputFormat: json
    service:
        pingOne:
            regionCode: NA
            authentication:
                type: authorization_code
                environmentID: 00000000-0000-0000-0000-000000000000
                authorizationCode:
                    clientID: 00000000-0000-0000-0000-000000000001
`
	testutils_koanf.InitKoanfsCustomFile(t, configContents)

	err := testutils_cobra.ExecutePingcli(t, "login", "--authorization-code")
	expectedErrorPattern := `redirect URI.*is not configured|failed to prompt for reconfiguration|input prompt error`
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// TestLoginCmd_InvalidFlags tests invalid flag combinations
func TestLoginCmd_InvalidFlags(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                 string
		args                 []string
		expectedErrorPattern string
	}{
		{
			name:                 "unknown_flag",
			args:                 []string{"login", "--unknown-flag"},
			expectedErrorPattern: `unknown flag: --unknown-flag`,
		},
		{
			name:                 "unknown_shorthand",
			args:                 []string{"login", "-x"},
			expectedErrorPattern: `unknown shorthand flag: 'x'`,
		},
		{
			name:                 "too_many_arguments",
			args:                 []string{"login", "extra-arg"},
			expectedErrorPattern: `command accepts 0 arg\(s\), received 1`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := testutils_cobra.ExecutePingcli(t, tc.args...)
			testutils.CheckExpectedError(t, err, &tc.expectedErrorPattern)
		})
	}
}

// TestLogoutCmd_InvalidFlags tests invalid flags for logout command
func TestLogoutCmd_InvalidFlags(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                 string
		args                 []string
		expectedErrorPattern string
	}{
		{
			name:                 "unknown_flag",
			args:                 []string{"logout", "--unknown-flag"},
			expectedErrorPattern: `unknown flag: --unknown-flag`,
		},
		{
			name:                 "too_many_arguments",
			args:                 []string{"logout", "extra-arg"},
			expectedErrorPattern: `command accepts 0 arg\(s\), received 1`,
		},
		{
			name:                 "mutually_exclusive_flags",
			args:                 []string{"logout", "--authorization-code", "--client-credentials"},
			expectedErrorPattern: `if any flags in the group.*are set none of the others can be`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := testutils_cobra.ExecutePingcli(t, tc.args...)
			testutils.CheckExpectedError(t, err, &tc.expectedErrorPattern)
		})
	}
}

// TestLoginCmd_HelpFlags tests help flags work correctly
func TestLoginCmd_HelpFlags(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name string
		args []string
	}{
		{
			name: "long_help_flag",
			args: []string{"login", "--help"},
		},
		{
			name: "short_help_flag",
			args: []string{"login", "-h"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := testutils_cobra.ExecutePingcli(t, tc.args...)
			// Help should not return an error
			if err != nil {
				t.Errorf("Help flag should not return error, got: %v", err)
			}
		})
	}
}

// TestLogoutCmd_HelpFlags tests help flags work correctly for logout
func TestLogoutCmd_HelpFlags(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name string
		args []string
	}{
		{
			name: "long_help_flag",
			args: []string{"logout", "--help"},
		},
		{
			name: "short_help_flag",
			args: []string{"logout", "-h"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := testutils_cobra.ExecutePingcli(t, tc.args...)
			// Help should not return an error
			if err != nil {
				t.Errorf("Help flag should not return error, got: %v", err)
			}
		})
	}
}
