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

func TestLoginCommand_DefaultAuthCode(t *testing.T) {
	// Test that when no flags are provided, it defaults to auth_code
	expectedErrorPattern := `^authorization code login failed`
	err := testutils_cobra.ExecutePingcli(t, "login")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
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
			args: []string{"--device-code", "--auth-code"},
		},
		{
			name: "client-credentials and auth-code",
			args: []string{"--client-credentials", "--auth-code"},
		},
		{
			name: "all three flags",
			args: []string{"--device-code", "--client-credentials", "--auth-code"},
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
	testCases := []struct {
		name                 string
		flag                 string
		expectedErrorPattern string
	}{
		{
			name:                 "auth-code flag",
			flag:                 "--auth-code",
			expectedErrorPattern: `^authorization code login failed`,
		},
		{
			name:                 "auth-code shorthand",
			flag:                 "-a",
			expectedErrorPattern: `^authorization code login failed`,
		},
		{
			name:                 "device-code flag",
			flag:                 "--device-code",
			expectedErrorPattern: `^device code login failed`,
		},
		{
			name:                 "device-code shorthand",
			flag:                 "-d",
			expectedErrorPattern: `^device code login failed`,
		},
		{
			name:                 "client-credentials flag",
			flag:                 "--client-credentials",
			expectedErrorPattern: `^client credentials login failed`,
		},
		{
			name:                 "client-credentials shorthand",
			flag:                 "-c",
			expectedErrorPattern: `^client credentials login failed`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := testutils_cobra.ExecutePingcli(t, "login", tc.flag)
			testutils.CheckExpectedError(t, err, &tc.expectedErrorPattern)
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
