// Copyright © 2025 Ping Identity Corporation

package auth_test

import (
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/cmd/auth"
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
	if cmd.Short != "Login user to the CLI" {
		t.Errorf("Expected Short to be 'Login user to the CLI', got %q", cmd.Short)
	}
	if !cmd.DisableFlagsInUseLine {
		t.Error("Expected DisableFlagsInUseLine to be true")
	}

	// Test that required flags are present
	deviceCodeFlag := cmd.Flags().Lookup("device-code")
	if deviceCodeFlag == nil {
		t.Error("device-code flag should be present")
	}

	authCodeFlag := cmd.Flags().Lookup("auth-code")
	if authCodeFlag == nil {
		t.Error("auth-code flag should be present")
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

func TestAuthLoginRunE_NoFlagsSet(t *testing.T) {
	expectedErrorPattern := `^please specify an authentication method: --auth-code, --client-credentials, or --device-code$`
	err := testutils_cobra.ExecutePingcli(t, "login")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

func TestAuthLoginRunE_MultipleFlagsSet(t *testing.T) {
	expectedErrorPattern := `^please specify only one authentication method$`
	err := testutils_cobra.ExecutePingcli(t, "login", "--device-code", "--client-credentials")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

func TestAuthLoginRunE_DeviceCodeFlag(t *testing.T) {
	expectedErrorPattern := `^device code login failed`
	err := testutils_cobra.ExecutePingcli(t, "login", "--device-code")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

func TestAuthLoginRunE_DeviceCodeShorthandFlag(t *testing.T) {
	// Test that shorthand flag -d maps correctly and produces expected error
	expectedErrorPattern := `^device code login failed`
	err := testutils_cobra.ExecutePingcli(t, "login", "-d")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

func TestAuthLoginRunE_ClientCredentialsFlag(t *testing.T) {
	expectedErrorPattern := `^client credentials login failed`
	err := testutils_cobra.ExecutePingcli(t, "login", "--client-credentials")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

func TestAuthLoginRunE_ClientCredentialsShorthandFlag(t *testing.T) {
	// Test that shorthand flag -c maps correctly and produces expected error
	expectedErrorPattern := `^client credentials login failed`
	err := testutils_cobra.ExecutePingcli(t, "login", "-c")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

func TestAuthLoginRunE_AuthCodeFlag(t *testing.T) {
	expectedErrorPattern := `^authorization code login failed`
	err := testutils_cobra.ExecutePingcli(t, "login", "--auth-code")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

func TestAuthLoginRunE_AuthCodeShorthandFlag(t *testing.T) {
	// Test that shorthand flag -a maps correctly and produces expected error
	expectedErrorPattern := `^authorization code login failed`
	err := testutils_cobra.ExecutePingcli(t, "login", "-a")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

func TestAuthLoginRunE_DeviceCodeAndClientCredentials(t *testing.T) {
	expectedErrorPattern := `^please specify only one authentication method$`
	err := testutils_cobra.ExecutePingcli(t, "login", "--device-code", "--client-credentials")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

func TestAuthLoginRunE_DeviceCodeAndAuthCode(t *testing.T) {
	expectedErrorPattern := `^please specify only one authentication method$`
	err := testutils_cobra.ExecutePingcli(t, "login", "--device-code", "--auth-code")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

func TestAuthLoginRunE_ClientCredentialsAndAuthCode(t *testing.T) {
	expectedErrorPattern := `^please specify only one authentication method$`
	err := testutils_cobra.ExecutePingcli(t, "login", "--client-credentials", "--auth-code")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
