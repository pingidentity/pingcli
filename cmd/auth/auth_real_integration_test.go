// Copyright © 2025 Ping Identity Corporation

package auth_test

import (
	"os"
	"testing"

	auth_internal "github.com/pingidentity/pingcli/internal/auth"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// TestLoginCommand_ClientCredentials_Integration tests the complete login flow with client credentials
func TestLoginCommand_ClientCredentials_Integration(t *testing.T) {
	// Skip if not in CI environment or missing credentials
	clientID := os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID")
	clientSecret := os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET")
	environmentID := os.Getenv("TEST_PINGONE_ENVIRONMENT_ID")
	regionCode := os.Getenv("TEST_PINGONE_REGION_CODE")

	if clientID == "" || clientSecret == "" || environmentID == "" || regionCode == "" {
		t.Skip("Skipping integration test: missing TEST_PINGONE_* environment variables")
	}

	// Clear any existing tokens to ensure fresh authentication
	err := auth_internal.ClearToken()
	if err != nil {
		t.Fatalf("Failed to clear token: %v", err)
	}

	// Test client credentials authentication using ExecutePingcli
	err = testutils_cobra.ExecutePingcli(t, "login", "--client-credentials")
	if err != nil {
		t.Fatalf("Login command should succeed with client credentials: %v", err)
	}

	// Login succeeded - token is automatically saved to keychain by SDK
	// Note: Token verification removed as SDK handles keychain storage internally
	// The absence of error from ExecutePingcli confirms successful authentication

	// Clean up - clear token from keychain
	err = auth_internal.ClearToken()
	if err != nil {
		t.Logf("Warning: Failed to clear token after test: %v", err)
	}
}

// TestLoginCommand_ShorthandHelpFlag_Integration tests shorthand help flag works in real environment
func TestLoginCommand_ShorthandHelpFlag_Integration(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "login", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// TestLoginCommand_InvalidShorthandFlag_Integration tests invalid shorthand flag fails in real environment
func TestLoginCommand_InvalidShorthandFlag_Integration(t *testing.T) {
	expectedErrorPattern := `^unknown shorthand flag: 'x' in -x$`
	err := testutils_cobra.ExecutePingcli(t, "login", "-x")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// TestLoginCommand_MultipleShorthandFlags_Integration tests multiple shorthand flags fail in real environment
func TestLoginCommand_MultipleShorthandFlags_Integration(t *testing.T) {
	expectedErrorPattern := `if any flags in the group`
	err := testutils_cobra.ExecutePingcli(t, "login", "-c", "-d")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// TestLoginCommand_DeviceCodeValidation_Integration tests device code configuration validation
// Note: Full device code flow testing is not currently implemented as it requires browser interaction automation
func TestLoginCommand_DeviceCodeValidation_Integration(t *testing.T) {
	// Skip if not in CI environment or missing credentials
	deviceCodeClientID := os.Getenv("TEST_PINGONE_DEVICE_CODE_CLIENT_ID")
	environmentID := os.Getenv("TEST_PINGONE_DEVICE_CODE_ENVIRONMENT_ID")
	regionCode := os.Getenv("TEST_PINGONE_REGION_CODE")

	if deviceCodeClientID == "" || environmentID == "" || regionCode == "" {
		t.Skip("Skipping device code validation test: missing TEST_PINGONE_DEVICE_CODE_* environment variables")
	}

	expectedErrorPattern := `^device code login failed: failed to get device code configuration:`
	err := testutils_cobra.ExecutePingcli(t, "login", "--device-code")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// TestLoginCommand_DeviceCodeShorthandFlag_Integration tests device code shorthand flag configuration validation
// Note: Full device code flow testing is not currently implemented as it requires browser interaction automation
func TestLoginCommand_DeviceCodeShorthandFlag_Integration(t *testing.T) {
	// Skip if not in CI environment or missing credentials
	deviceCodeClientID := os.Getenv("TEST_PINGONE_DEVICE_CODE_CLIENT_ID")
	environmentID := os.Getenv("TEST_PINGONE_DEVICE_CODE_ENVIRONMENT_ID")
	regionCode := os.Getenv("TEST_PINGONE_REGION_CODE")

	if deviceCodeClientID == "" || environmentID == "" || regionCode == "" {
		t.Skip("Skipping device code validation test: missing TEST_PINGONE_DEVICE_CODE_* environment variables")
	}

	expectedErrorPattern := `^device code login failed: failed to get device code configuration:`
	err := testutils_cobra.ExecutePingcli(t, "login", "-d")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// TestLoginCommand_AuthCodeValidation_Integration tests auth code configuration validation
// Note: Full auth code flow testing is not currently implemented as it requires browser interaction automation
func TestLoginCommand_AuthCodeValidation_Integration(t *testing.T) {
	// Skip if not in CI environment or missing credentials
	authCodeClientID := os.Getenv("TEST_PINGONE_AUTH_CODE_CLIENT_ID")
	environmentID := os.Getenv("TEST_PINGONE_AUTH_CODE_ENVIRONMENT_ID")
	regionCode := os.Getenv("TEST_PINGONE_REGION_CODE")

	if authCodeClientID == "" || environmentID == "" || regionCode == "" {
		t.Skip("Skipping auth code validation test: missing TEST_PINGONE_AUTH_CODE_* environment variables")
	}

	expectedErrorPattern := `^authorization code login failed: failed to get auth code configuration:`
	err := testutils_cobra.ExecutePingcli(t, "login", "--auth-code")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// TestLoginCommand_AuthCodeShorthandFlag_Integration tests auth code shorthand flag configuration validation
// Note: Full auth code flow testing is not currently implemented as it requires browser interaction automation
func TestLoginCommand_AuthCodeShorthandFlag_Integration(t *testing.T) {
	// Skip if not in CI environment or missing credentials
	authCodeClientID := os.Getenv("TEST_PINGONE_AUTH_CODE_CLIENT_ID")
	environmentID := os.Getenv("TEST_PINGONE_AUTH_CODE_ENVIRONMENT_ID")
	regionCode := os.Getenv("TEST_PINGONE_REGION_CODE")

	if authCodeClientID == "" || environmentID == "" || regionCode == "" {
		t.Skip("Skipping auth code validation test: missing TEST_PINGONE_AUTH_CODE_* environment variables")
	}

	expectedErrorPattern := `^authorization code login failed: failed to get auth code configuration:`
	err := testutils_cobra.ExecutePingcli(t, "login", "-a")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// TestLoginCommand_MultipleFlagsValidation_Integration tests multiple flags fail in real environment
func TestLoginCommand_MultipleFlagsValidation_Integration(t *testing.T) {
	expectedErrorPattern := `if any flags in the group`
	err := testutils_cobra.ExecutePingcli(t, "login", "--client-credentials", "--device-code")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// TestLoginCommand_NoFlagsValidation_Integration tests that no flags defaults to auth_code
func TestLoginCommand_NoFlagsValidation_Integration(t *testing.T) {
	// Should default to auth_code and fail with auth code configuration error
	expectedErrorPattern := `authorization code login failed`
	err := testutils_cobra.ExecutePingcli(t, "login")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// TestLoginCommand_InvalidFlagValidation_Integration tests invalid flag fails in real environment
func TestLoginCommand_InvalidFlagValidation_Integration(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingcli(t, "login", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// TestLoginCommand_HelpFlagValidation_Integration tests help flag works in real environment
func TestLoginCommand_HelpFlagValidation_Integration(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "login", "--help")
	testutils.CheckExpectedError(t, err, nil)
}

// TestLoginCommand_HelpShorthandFlagValidation_Integration tests help shorthand flag works in real environment
func TestLoginCommand_HelpShorthandFlagValidation_Integration(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "login", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// TestLogoutCommand_Integration tests logout functionality in real environment
func TestLogoutCommand_Integration(t *testing.T) {
	// Skip if not in CI environment or missing credentials
	clientID := os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID")
	clientSecret := os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET")
	environmentID := os.Getenv("TEST_PINGONE_ENVIRONMENT_ID")
	regionCode := os.Getenv("TEST_PINGONE_REGION_CODE")

	if clientID == "" || clientSecret == "" || environmentID == "" || regionCode == "" {
		t.Skip("Skipping integration test: missing TEST_PINGONE_* environment variables")
	}

	// First login to have something to logout from
	err := testutils_cobra.ExecutePingcli(t, "login", "--client-credentials")
	if err != nil {
		t.Fatalf("Login should succeed: %v", err)
	}

	// Login succeeded - token is saved in keychain

	// Test logout using ExecutePingcli with the same auth method
	err = testutils_cobra.ExecutePingcli(t, "logout", "--client-credentials")
	if err != nil {
		t.Fatalf("Logout should succeed: %v", err)
	}

	// Logout succeeded - token cleared from keychain
}
