// Copyright © 2025 Ping Identity Corporation

package auth_test

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/cmd/auth"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

func TestLoginCommand_Creation(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	cmd := auth.NewLoginCommand()

	if cmd.Use != "login [flags]" {
		t.Errorf("Expected Use to be 'login [flags]', got %q", cmd.Use)
	}
	if cmd.Short != "Authenticate a supported provider" {
		t.Errorf("Expected Short to be 'Authenticate a supported provider', got %q", cmd.Short)
	}
	if !cmd.DisableFlagsInUseLine {
		t.Error("Expected DisableFlagsInUseLine to be true")
	}

	// Test that required flags are present
	deviceCodeFlag := cmd.Flags().Lookup("device-code")
	if deviceCodeFlag == nil {
		t.Error("device-code flag should be present")
	}

	authorizationCodeFlag := cmd.Flags().Lookup("authorization-code")
	if authorizationCodeFlag == nil {
		t.Error("authorization-code flag should be present")
	}

	clientCredentialsFlag := cmd.Flags().Lookup("client-credentials")
	if clientCredentialsFlag == nil {
		t.Error("client-credentials flag should be present")
	}

	// Test shorthand flags are mapped correctly
	if cmd.Flags().ShorthandLookup("d") == nil {
		t.Error("device-code shorthand -d should be present")
	}
	if cmd.Flags().ShorthandLookup("a") == nil {
		t.Error("auth-code shorthand -a should be present")
	}
	if cmd.Flags().ShorthandLookup("c") == nil {
		t.Error("client-credentials shorthand -c should be present")
	}
}

func TestLoginCommand_ShorthandFlags(t *testing.T) {
	// Test shorthand flags are properly recognized using ExecutePingcli approach
	// Focus on flag parsing validation rather than command execution

	// Test that shorthand flags work in argument validation context
	err := testutils_cobra.ExecutePingcli(t, "login", "-x")
	if err == nil {
		t.Fatal("Expected error for unknown shorthand flag")
	}
	if !strings.Contains(err.Error(), "unknown shorthand flag: 'x'") {
		t.Errorf("Expected unknown shorthand flag error, got: %v", err)
	}

	// Test that help works for shorthand
	err = testutils_cobra.ExecutePingcli(t, "login", "-h")
	if err != nil {
		t.Errorf("Shorthand help should work without error, got: %v", err)
	}
}

func TestLoginCommand_FlagValidationExecution(t *testing.T) {
	// Test basic flag validation using ExecutePingcli approach
	// This tests the complete command pipeline for argument validation

	// Test too many arguments
	err := testutils_cobra.ExecutePingcli(t, "login", "extra-arg")
	if err == nil {
		t.Fatal("Expected error when too many arguments are provided")
	}
	if !strings.Contains(err.Error(), "command accepts 0 arg(s), received 1") {
		t.Errorf("Expected argument validation error, got: %v", err)
	}

	// Test invalid flag
	err = testutils_cobra.ExecutePingcli(t, "login", "--invalid-flag")
	if err == nil {
		t.Fatal("Expected error when invalid flag is provided")
	}
	if !strings.Contains(err.Error(), "unknown flag: --invalid-flag") {
		t.Errorf("Expected unknown flag error, got: %v", err)
	}

	// Test help flag - should work without configuration issues
	err = testutils_cobra.ExecutePingcli(t, "login", "--help")
	if err != nil {
		t.Errorf("Help flag should work without error, got: %v", err)
	}

	// Test shorthand help flag
	err = testutils_cobra.ExecutePingcli(t, "login", "-h")
	if err != nil {
		t.Errorf("Shorthand help flag should work without error, got: %v", err)
	}
}

func TestLoginCommand_BooleanFlagBehavior(t *testing.T) {
	// Test flag behavior using ExecutePingcli approach
	// Focus on flag parsing and validation rather than command execution

	// Test help flag works
	err := testutils_cobra.ExecutePingcli(t, "login", "--help")
	if err != nil {
		t.Errorf("Help should work without error, got: %v", err)
	}

	// Test invalid flag combination (too many arguments)
	err = testutils_cobra.ExecutePingcli(t, "login", "extra", "arguments")
	if err == nil {
		t.Fatal("Expected error when too many arguments are provided")
	}
	if !strings.Contains(err.Error(), "command accepts 0 arg(s), received 2") {
		t.Errorf("Expected argument validation error, got: %v", err)
	}
}

func TestLoginCommand_DefaultAuthorizationCode(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Test that when no flags are provided, it defaults to auth_code
	// With valid credentials configured, may succeed; otherwise should fail
	err := testutils_cobra.ExecutePingcli(t, "login")
	if err == nil {
		// Success - valid auth_code credentials configured
		t.Skip("Login succeeded with configured auth_code credentials")
	}
	// Error expected when credentials not configured
	// The error might be "authorization code" (default) or "client credentials" (configured in testutils)
	if !strings.Contains(err.Error(), "authorization code") &&
		!strings.Contains(err.Error(), "client credentials") &&
		!strings.Contains(err.Error(), "failed to prompt for reconfiguration") &&
		!strings.Contains(err.Error(), "failed to get") {
		t.Errorf("Expected auth code or client credentials related error, got: %v", err)
	}
}

func TestLoginCommand_MutuallyExclusiveFlags(t *testing.T) {
	testCases := []struct {
		name string
		args []string
	}{
		{
			name: "device-code and client-credentials",
			args: []string{"--device-code", "--client-credentials"},
		},
		{
			name: "device-code and auth-code",
			args: []string{"--device-code", "--authorization-code"},
		},
		{
			name: "client-credentials and auth-code",
			args: []string{"--client-credentials", "--authorization-code"},
		},
		{
			name: "all three flags",
			args: []string{"--device-code", "--client-credentials", "--authorization-code"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := append([]string{"login"}, tc.args...)
			err := testutils_cobra.ExecutePingcli(t, args...)
			if err == nil {
				t.Fatal("Expected error for mutually exclusive flags, got nil")
			}
			// Check that error mentions mutual exclusivity
			if !strings.Contains(err.Error(), "if any flags in the group") {
				t.Errorf("Expected mutually exclusive flags error, got: %v", err)
			}
		})
	}
}

func TestLoginCommand_SpecificAuthMethod(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                 string
		args                 []string
		expectedErrorPattern string
		expectSuccess        bool
		allowBoth            bool // Allow either success or specific error
	}{
		{
			name:                 "auth-code flag",
			args:                 []string{"--authorization-code"},
			expectedErrorPattern: `authorization code`,
			allowBoth:            true, // May succeed with valid config
		},
		{
			name:                 "auth-code shorthand",
			args:                 []string{"-a"},
			expectedErrorPattern: `authorization code`,
			allowBoth:            true, // May succeed with valid config
		},
		{
			name:                 "device-code flag",
			args:                 []string{"--device-code"},
			expectedErrorPattern: `device (code|auth)`,
			allowBoth:            true, // May succeed with valid config
		},
		{
			name:                 "device-code shorthand",
			args:                 []string{"-d"},
			expectedErrorPattern: `device (code|auth)`,
			allowBoth:            true, // May succeed with valid config
		},
		{
			name:          "client-credentials flag",
			args:          []string{"--client-credentials"},
			expectSuccess: true, // With valid config, login succeeds
		},
		{
			name:          "client-credentials shorthand",
			args:          []string{"-c"},
			expectSuccess: true, // With valid config, login succeeds
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Skip client credentials tests if environment variables are not set
			isClientCreds := false
			for _, arg := range tc.args {
				if strings.Contains(arg, "client-credentials") || arg == "-c" {
					isClientCreds = true

					break
				}
			}

			if tc.expectSuccess && isClientCreds {
				if val, _ := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsClientIDOption); val == "" && os.Getenv("TEST_PINGONE_CLIENT_ID") == "" && os.Getenv("PINGONE_CLIENT_ID") == "" {
					t.Skip("Skipping test: Client Credentials Client ID not configured")
				}
			}

			// Pre-configure environment variables for the test execution
			// The login command relies on global configuration/env vars
			if isClientCreds {
				clientID := os.Getenv("TEST_PINGONE_CLIENT_ID")
				if clientID != "" {
					t.Setenv("PINGONE_CLIENT_ID", clientID)
					t.Setenv("PINGONE_CLIENT_CREDENTIALS_CLIENT_ID", clientID)
					// Set Koanf-style env var overrides
					t.Setenv("PINGCLI_SERVICE_PINGONE_AUTHENTICATION_CLIENTCREDENTIALS_CLIENTID", clientID)
				}

				clientSecret := os.Getenv("TEST_PINGONE_CLIENT_SECRET")
				if clientSecret != "" {
					t.Setenv("PINGONE_CLIENT_SECRET", clientSecret)
					t.Setenv("PINGONE_CLIENT_CREDENTIALS_CLIENT_SECRET", clientSecret)
					t.Setenv("PINGCLI_SERVICE_PINGONE_AUTHENTICATION_CLIENTCREDENTIALS_CLIENTSECRET", clientSecret)
				}

				region := os.Getenv("TEST_PINGONE_REGION_CODE")
				if region != "" {
					t.Setenv("PINGONE_REGION_CODE", region)
					t.Setenv("PINGCLI_SERVICE_PINGONE_REGION", region)
				}
			}

			cmdArgs := append([]string{"login"}, tc.args...)

			err := testutils_cobra.ExecutePingcli(t, cmdArgs...)
			switch {
			case tc.expectSuccess:
				if err != nil {
					t.Errorf("Expected success but got error: %v", err)
				}
			case tc.allowBoth:
				// Either success or expected error is acceptable
				if err != nil {
					// Check error matches expected pattern
					matched, _ := regexp.MatchString(tc.expectedErrorPattern, err.Error())
					if !matched && !strings.Contains(err.Error(), "failed to prompt") &&
						!strings.Contains(err.Error(), "failed to configure authentication") &&
						!strings.Contains(err.Error(), "input prompt error") &&
						!strings.Contains(err.Error(), "failed to get") &&
						!strings.Contains(err.Error(), "context deadline exceeded") &&
						!strings.Contains(err.Error(), "address already in use") {
						t.Errorf("Error did not match expected pattern '%s', got: %v", tc.expectedErrorPattern, err)
					}
				}
				// Success is also acceptable
			default:
				testutils.CheckExpectedError(t, err, &tc.expectedErrorPattern)
			}
		})
	}
}

func TestLoginCommandValidation(t *testing.T) {
	// Test invalid flag combination (too many arguments)
	err := testutils_cobra.ExecutePingcli(t, "login", "extra", "arguments")
	if err == nil {
		t.Fatal("Expected error when too many arguments are provided")
	}
	if !strings.Contains(err.Error(), "command accepts 0 arg(s), received 2") {
		t.Errorf("Expected argument validation error, got: %v", err)
	}
}
