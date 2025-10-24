// Copyright © 2025 Ping Identity Corporation

package auth_test

import (
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/cmd/auth"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// TestLoginCommand_DeviceCodeShorthandParsing_Integration tests that device-code shorthand -d is properly parsed
func TestLoginCommand_DeviceCodeShorthandParsing_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	// Set the args and parse flags
	args := []string{"-d"}
	cmd.SetArgs(args)
	err := cmd.ParseFlags(args)
	if err != nil {
		t.Fatalf("ParseFlags should not error: %v", err)
	}

	// Check that the expected flag was set to true
	flagValue, err := cmd.Flags().GetBool("device-code")
	if err != nil {
		t.Fatalf("GetBool should not error: %v", err)
	}
	if flagValue != true {
		t.Errorf("Flag device-code should be true, got %v", flagValue)
	}

	// Verify other flags are false
	allFlags := []string{"auth-code", "client-credentials"}
	for _, flagName := range allFlags {
		otherFlagValue, err := cmd.Flags().GetBool(flagName)
		if err != nil {
			t.Fatalf("GetBool should not error: %v", err)
		}
		if otherFlagValue {
			t.Errorf("Flag %s should be false when device-code is set", flagName)
		}
	}
}

// TestLoginCommand_AuthCodeShorthandParsing_Integration tests that auth-code shorthand -a is properly parsed
func TestLoginCommand_AuthCodeShorthandParsing_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	// Set the args and parse flags
	args := []string{"-a"}
	cmd.SetArgs(args)
	err := cmd.ParseFlags(args)
	if err != nil {
		t.Fatalf("ParseFlags should not error: %v", err)
	}

	// Check that the expected flag was set to true
	flagValue, err := cmd.Flags().GetBool("auth-code")
	if err != nil {
		t.Fatalf("GetBool should not error: %v", err)
	}
	if flagValue != true {
		t.Errorf("Flag auth-code should be true, got %v", flagValue)
	}

	// Verify other flags are false
	allFlags := []string{"device-code", "client-credentials"}
	for _, flagName := range allFlags {
		otherFlagValue, err := cmd.Flags().GetBool(flagName)
		if err != nil {
			t.Fatalf("GetBool should not error: %v", err)
		}
		if otherFlagValue {
			t.Errorf("Flag %s should be false when auth-code is set", flagName)
		}
	}
}

// TestLoginCommand_ClientCredentialsShorthandParsing_Integration tests that client-credentials shorthand -c is properly parsed
func TestLoginCommand_ClientCredentialsShorthandParsing_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	// Set the args and parse flags
	args := []string{"-c"}
	cmd.SetArgs(args)
	err := cmd.ParseFlags(args)
	if err != nil {
		t.Fatalf("ParseFlags should not error: %v", err)
	}

	// Check that the expected flag was set to true
	flagValue, err := cmd.Flags().GetBool("client-credentials")
	if err != nil {
		t.Fatalf("GetBool should not error: %v", err)
	}
	if flagValue != true {
		t.Errorf("Flag client-credentials should be true, got %v", flagValue)
	}

	// Verify other flags are false
	allFlags := []string{"device-code", "auth-code"}
	for _, flagName := range allFlags {
		otherFlagValue, err := cmd.Flags().GetBool(flagName)
		if err != nil {
			t.Fatalf("GetBool should not error: %v", err)
		}
		if otherFlagValue {
			t.Errorf("Flag %s should be false when client-credentials is set", flagName)
		}
	}
}

// TestLoginCommand_DeviceCodeFullFlagParsing_Integration tests that device-code full flag is properly parsed
func TestLoginCommand_DeviceCodeFullFlagParsing_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	// Set the args and parse flags
	args := []string{"--device-code"}
	cmd.SetArgs(args)
	err := cmd.ParseFlags(args)
	if err != nil {
		t.Fatalf("ParseFlags should not error: %v", err)
	}

	// Check that the expected flag was set to true
	flagValue, err := cmd.Flags().GetBool("device-code")
	if err != nil {
		t.Fatalf("GetBool should not error: %v", err)
	}
	if flagValue != true {
		t.Errorf("Flag device-code should be true, got %v", flagValue)
	}

	// Verify other flags are false
	allFlags := []string{"auth-code", "client-credentials"}
	for _, flagName := range allFlags {
		otherFlagValue, err := cmd.Flags().GetBool(flagName)
		if err != nil {
			t.Fatalf("GetBool should not error: %v", err)
		}
		if otherFlagValue {
			t.Errorf("Flag %s should be false when device-code is set", flagName)
		}
	}
}

// TestLoginCommand_AuthCodeFullFlagParsing_Integration tests that auth-code full flag is properly parsed
func TestLoginCommand_AuthCodeFullFlagParsing_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	// Set the args and parse flags
	args := []string{"--auth-code"}
	cmd.SetArgs(args)
	err := cmd.ParseFlags(args)
	if err != nil {
		t.Fatalf("ParseFlags should not error: %v", err)
	}

	// Check that the expected flag was set to true
	flagValue, err := cmd.Flags().GetBool("auth-code")
	if err != nil {
		t.Fatalf("GetBool should not error: %v", err)
	}
	if flagValue != true {
		t.Errorf("Flag auth-code should be true, got %v", flagValue)
	}

	// Verify other flags are false
	allFlags := []string{"device-code", "client-credentials"}
	for _, flagName := range allFlags {
		otherFlagValue, err := cmd.Flags().GetBool(flagName)
		if err != nil {
			t.Fatalf("GetBool should not error: %v", err)
		}
		if otherFlagValue {
			t.Errorf("Flag %s should be false when auth-code is set", flagName)
		}
	}
}

// TestLoginCommand_ClientCredentialsFullFlagParsing_Integration tests that client-credentials full flag is properly parsed
func TestLoginCommand_ClientCredentialsFullFlagParsing_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	// Set the args and parse flags
	args := []string{"--client-credentials"}
	cmd.SetArgs(args)
	err := cmd.ParseFlags(args)
	if err != nil {
		t.Fatalf("ParseFlags should not error: %v", err)
	}

	// Check that the expected flag was set to true
	flagValue, err := cmd.Flags().GetBool("client-credentials")
	if err != nil {
		t.Fatalf("GetBool should not error: %v", err)
	}
	if flagValue != true {
		t.Errorf("Flag client-credentials should be true, got %v", flagValue)
	}

	// Verify other flags are false
	allFlags := []string{"device-code", "auth-code"}
	for _, flagName := range allFlags {
		otherFlagValue, err := cmd.Flags().GetBool(flagName)
		if err != nil {
			t.Fatalf("GetBool should not error: %v", err)
		}
		if otherFlagValue {
			t.Errorf("Flag %s should be false when client-credentials is set", flagName)
		}
	}
}

// TestLoginCommand_NoFlagsExecution_Integration tests that command defaults to auth_code when no flags are provided
func TestLoginCommand_NoFlagsExecution_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	cmd.SetArgs([]string{})
	err := cmd.Execute()

	if err == nil {
		t.Error("Expected error but got none")
	} else if !strings.Contains(err.Error(), "authorization code login failed") {
		t.Errorf("Error should contain 'authorization code login failed' (default), got: %v", err)
	}
}

// TestLoginCommand_MultipleFlagsDeviceCodeAndAuthCode_Integration tests that command fails with multiple flags -d -a
func TestLoginCommand_MultipleFlagsDeviceCodeAndAuthCode_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	cmd.SetArgs([]string{"-d", "-a"})
	err := cmd.Execute()

	if err == nil {
		t.Error("Expected error but got none")
	} else if !strings.Contains(err.Error(), "if any flags in the group") {
		t.Errorf("Expected mutually exclusive flags error, got: %v", err)
	}
}

// TestLoginCommand_MultipleFlagsClientCredAndDeviceCode_Integration tests that command fails with multiple flags -c -d
func TestLoginCommand_MultipleFlagsClientCredAndDeviceCode_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	cmd.SetArgs([]string{"-c", "-d"})
	err := cmd.Execute()

	if err == nil {
		t.Error("Expected error but got none")
	} else if !strings.Contains(err.Error(), "if any flags in the group") {
		t.Errorf("Expected mutually exclusive flags error, got: %v", err)
	}
}

// TestLoginCommand_MultipleFlagsAuthCodeAndClientCred_Integration tests that command fails with multiple flags --auth-code --client-credentials
func TestLoginCommand_MultipleFlagsAuthCodeAndClientCred_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	cmd.SetArgs([]string{"--auth-code", "--client-credentials"})
	err := cmd.Execute()

	if err == nil {
		t.Error("Expected error but got none")
	} else if !strings.Contains(err.Error(), "if any flags in the group") {
		t.Errorf("Expected mutually exclusive flags error, got: %v", err)
	}
}

// TestLoginCommand_AllThreeFlagsExecution_Integration tests that command fails with all three flags -d -a -c
func TestLoginCommand_AllThreeFlagsExecution_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	cmd.SetArgs([]string{"-d", "-a", "-c"})
	err := cmd.Execute()

	if err == nil {
		t.Error("Expected error but got none")
	} else if !strings.Contains(err.Error(), "if any flags in the group") {
		t.Errorf("Expected mutually exclusive flags error, got: %v", err)
	}
}

// TestLoginCommand_DeviceCodeOnlyExecution_Integration tests that device-code flag only validates properly
func TestLoginCommand_DeviceCodeOnlyExecution_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	cmd.SetArgs([]string{"-d"})
	err := cmd.Execute()

	if err == nil {
		t.Error("Expected error but got none")
	} else if !strings.Contains(err.Error(), "device code login failed") {
		t.Errorf("Error should contain %q, got: %v", "device code login failed", err)
	}
}

// TestLoginCommand_AuthCodeOnlyExecution_Integration tests that auth-code flag only validates properly
func TestLoginCommand_AuthCodeOnlyExecution_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	cmd.SetArgs([]string{"-a"})
	err := cmd.Execute()

	if err == nil {
		t.Error("Expected error but got none")
	} else if !strings.Contains(err.Error(), "authorization code login failed") {
		t.Errorf("Error should contain %q, got: %v", "authorization code login failed", err)
	}
}

// TestLoginCommand_ClientCredentialsOnlyExecution_Integration tests that client-credentials flag only validates properly
func TestLoginCommand_ClientCredentialsOnlyExecution_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	cmd.SetArgs([]string{"-c"})
	err := cmd.Execute()

	if err == nil {
		t.Error("Expected error but got none")
	} else if !strings.Contains(err.Error(), "client credentials login failed") {
		t.Errorf("Error should contain %q, got: %v", "client credentials login failed", err)
	}
}

// TestLoginCommand_DeviceCodeBooleanFlagBehavior_Integration tests that device-code flag can be set without values
func TestLoginCommand_DeviceCodeBooleanFlagBehavior_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	args := []string{"--device-code"}
	cmd.SetArgs(args)
	err := cmd.ParseFlags(args)
	if err != nil {
		t.Fatalf("ParseFlags should not error: %v", err)
	}

	flagValue, err := cmd.Flags().GetBool("device-code")
	if err != nil {
		t.Fatalf("GetBool should not error: %v", err)
	}
	if !flagValue {
		t.Errorf("Flag device-code should be true when set without value")
	}
}

// TestLoginCommand_AuthCodeBooleanFlagBehavior_Integration tests that auth-code flag can be set without values
func TestLoginCommand_AuthCodeBooleanFlagBehavior_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	args := []string{"--auth-code"}
	cmd.SetArgs(args)
	err := cmd.ParseFlags(args)
	if err != nil {
		t.Fatalf("ParseFlags should not error: %v", err)
	}

	flagValue, err := cmd.Flags().GetBool("auth-code")
	if err != nil {
		t.Fatalf("GetBool should not error: %v", err)
	}
	if !flagValue {
		t.Errorf("Flag auth-code should be true when set without value")
	}
}

// TestLoginCommand_ClientCredentialsBooleanFlagBehavior_Integration tests that client-credentials flag can be set without values
func TestLoginCommand_ClientCredentialsBooleanFlagBehavior_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	args := []string{"--client-credentials"}
	cmd.SetArgs(args)
	err := cmd.ParseFlags(args)
	if err != nil {
		t.Fatalf("ParseFlags should not error: %v", err)
	}

	flagValue, err := cmd.Flags().GetBool("client-credentials")
	if err != nil {
		t.Fatalf("GetBool should not error: %v", err)
	}
	if !flagValue {
		t.Errorf("Flag client-credentials should be true when set without value")
	}
}

// TestLoginCommand_DeviceCodeShorthandExecution_Integration tests end-to-end execution with device-code shorthand flag
func TestLoginCommand_DeviceCodeShorthandExecution_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	cmd.SetArgs([]string{"-d"})
	err := cmd.Execute()

	// Should get an authentication error (not a flag parsing error)
	if err == nil {
		t.Error("Expected error but got none")
	} else {
		if !strings.Contains(err.Error(), "device code login failed") {
			t.Errorf("Error should contain %q, got: %v", "device code login failed", err)
		}
		// Ensure it's NOT a flag parsing error
		if strings.Contains(err.Error(), "unknown shorthand flag") {
			t.Errorf("Should not be a flag parsing error with 'unknown shorthand flag': %v", err)
		}
		if strings.Contains(err.Error(), "flag provided but not defined") {
			t.Errorf("Should not be a flag parsing error with 'flag provided but not defined': %v", err)
		}
	}
}

// TestLoginCommand_AuthCodeShorthandExecution_Integration tests end-to-end execution with auth-code shorthand flag
func TestLoginCommand_AuthCodeShorthandExecution_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	cmd.SetArgs([]string{"-a"})
	err := cmd.Execute()

	// Should get an authentication error (not a flag parsing error)
	if err == nil {
		t.Error("Expected error but got none")
	} else {
		if !strings.Contains(err.Error(), "authorization code login failed") {
			t.Errorf("Error should contain %q, got: %v", "authorization code login failed", err)
		}
		// Ensure it's NOT a flag parsing error
		if strings.Contains(err.Error(), "unknown shorthand flag") {
			t.Errorf("Should not be a flag parsing error with 'unknown shorthand flag': %v", err)
		}
		if strings.Contains(err.Error(), "flag provided but not defined") {
			t.Errorf("Should not be a flag parsing error with 'flag provided but not defined': %v", err)
		}
	}
}

// TestLoginCommand_ClientCredentialsShorthandExecution_Integration tests end-to-end execution with client-credentials shorthand flag
func TestLoginCommand_ClientCredentialsShorthandExecution_Integration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLoginCommand()

	cmd.SetArgs([]string{"-c"})
	err := cmd.Execute()

	// Should get an authentication error (not a flag parsing error)
	if err == nil {
		t.Error("Expected error but got none")
	} else {
		if !strings.Contains(err.Error(), "client credentials login failed") {
			t.Errorf("Error should contain %q, got: %v", "client credentials login failed", err)
		}
		// Ensure it's NOT a flag parsing error
		if strings.Contains(err.Error(), "unknown shorthand flag") {
			t.Errorf("Should not be a flag parsing error with 'unknown shorthand flag': %v", err)
		}
		if strings.Contains(err.Error(), "flag provided but not defined") {
			t.Errorf("Should not be a flag parsing error with 'flag provided but not defined': %v", err)
		}
	}
}
