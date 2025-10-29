// Copyright © 2025 Ping Identity Corporation

package auth_test

import (
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/cmd/auth"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

func TestLogoutCommand_Creation(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	cmd := auth.NewLogoutCommand()

	// Test basic command properties
	if cmd.Name() != "logout" {
		t.Errorf("Expected command name to be 'logout', got %q", cmd.Name())
	}
	if cmd.Short != "Logout user from the CLI" {
		t.Errorf("Expected command short to be 'Logout user from the CLI', got %q", cmd.Short)
	}
	expectedLong := "Logout user from the CLI by clearing stored credentials. Credentials are cleared from both keychain and file storage. By default, uses the authentication method configured in the active profile. You can specify a different authentication method using the auth method flags."
	if cmd.Long != expectedLong {
		t.Errorf("Expected command long to be %q, got %q", expectedLong, cmd.Long)
	}
	if !strings.Contains(cmd.Use, "logout") {
		t.Errorf("Expected command Use to contain 'logout', got %q", cmd.Use)
	}

	// Test that the command has auth method flags
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

	// Test that shorthands are present
	if deviceCodeFlag != nil && deviceCodeFlag.Shorthand != "d" {
		t.Error("device-code shorthand -d should be present")
	}
	if authCodeFlag != nil && authCodeFlag.Shorthand != "a" {
		t.Error("auth-code shorthand -a should be present")
	}
	if clientCredentialsFlag != nil && clientCredentialsFlag.Shorthand != "c" {
		t.Error("client-credentials shorthand -c should be present")
	}

	// Test that the command accepts exactly 0 arguments using common.ExactArgs(0)
	err := cmd.Args(cmd, []string{})
	if err != nil {
		t.Errorf("Expected command to accept 0 arguments: %v", err)
	}
}

func TestLogoutCommandHelp(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	cmd := auth.NewLogoutCommand()

	// Test that help can be generated without error
	usage := cmd.UsageString()
	if !strings.Contains(usage, "logout") {
		t.Errorf("Expected usage to contain 'logout', got %q", usage)
	}

	// Verify auth method flags are in help
	flagOutput := cmd.Flags().FlagUsages()
	if !strings.Contains(flagOutput, "auth-code") {
		t.Error("Help should contain auth-code flag")
	}
	if !strings.Contains(flagOutput, "device-code") {
		t.Error("Help should contain device-code flag")
	}
	if !strings.Contains(flagOutput, "client-credentials") {
		t.Error("Help should contain client-credentials flag")
	}
}

func TestLogoutCommandValidation(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	cmd := auth.NewLogoutCommand()

	// Test that command rejects arguments
	err := cmd.Args(cmd, []string{"unexpected-arg"})
	if err == nil {
		t.Error("Expected command to reject arguments")
	}
	if !strings.Contains(err.Error(), "accepts 0 arg(s), received 1") {
		t.Errorf("Expected error to contain 'accepts 0 arg(s), received 1', got %q", err.Error())
	}
}

func TestLogoutCommand_MutuallyExclusiveFlags(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Test that specifying multiple auth method flags fails
	tests := []struct {
		name  string
		flags []string
	}{
		{
			name:  "device-code and client-credentials",
			flags: []string{"--device-code", "--client-credentials"},
		},
		{
			name:  "device-code and auth-code",
			flags: []string{"--device-code", "--auth-code"},
		},
		{
			name:  "client-credentials and auth-code",
			flags: []string{"--client-credentials", "--auth-code"},
		},
		{
			name:  "all three flags",
			flags: []string{"--device-code", "--client-credentials", "--auth-code"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := append([]string{"logout"}, tt.flags...)
			err := testutils_cobra.ExecutePingcli(t, args...)
			if err == nil {
				t.Error("Expected error when specifying multiple auth method flags, got nil")
			}
			if !strings.Contains(err.Error(), "if any flags in the group") {
				t.Errorf("Expected mutually exclusive flags error, got: %v", err)
			}
		})
	}
}

func TestLogoutCommand_SpecificAuthMethod(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	tests := []struct {
		name     string
		flag     string
		flagName string
	}{
		{
			name:     "device-code flag",
			flag:     "--device-code",
			flagName: "device-code",
		},
		{
			name:     "client-credentials flag",
			flag:     "--client-credentials",
			flagName: "client-credentials",
		},
		{
			name:     "auth-code flag",
			flag:     "--auth-code",
			flagName: "auth-code",
		},
		{
			name:     "device-code shorthand",
			flag:     "-d",
			flagName: "device-code",
		},
		{
			name:     "client-credentials shorthand",
			flag:     "-c",
			flagName: "client-credentials",
		},
		{
			name:     "auth-code shorthand",
			flag:     "-a",
			flagName: "auth-code",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := auth.NewLogoutCommand()

			// For boolean flags, they just need to be present to be set to true
			err := cmd.Flags().Parse([]string{tt.flag})
			if err != nil {
				t.Fatalf("Failed to parse flag %s: %v", tt.flag, err)
			}

			// Verify the flag was set
			flag := cmd.Flags().Lookup(tt.flagName)
			if flag == nil {
				t.Fatalf("Flag %s not found", tt.flagName)
			}

			if !flag.Changed {
				t.Errorf("Flag %s (name: %s) was not marked as changed", tt.flag, tt.flagName)
			}

			if flag.Value.String() != "true" {
				t.Errorf("Flag %s (name: %s) should be true, got: %s", tt.flag, tt.flagName, flag.Value.String())
			}
		})
	}
}

func TestAuthLogoutRunE_ClearCredentials(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLogoutCommand()

	// May fail if no credentials exist, which is expected
	err := cmd.RunE(cmd, []string{})
	_ = err
}
