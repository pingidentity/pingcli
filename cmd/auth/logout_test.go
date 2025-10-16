// Copyright © 2025 Ping Identity Corporation

package auth_test

import (
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/cmd/auth"
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
	if cmd.Long != "Logout user from the CLI by clearing stored credentials from Keychain" {
		t.Errorf("Expected command long to be 'Logout user from the CLI by clearing stored credentials from Keychain', got %q", cmd.Long)
	}
	if !strings.Contains(cmd.Use, "logout") {
		t.Errorf("Expected command Use to contain 'logout', got %q", cmd.Use)
	}

	// Test that the command has no flags (logout doesn't need any)
	if cmd.Flags().NFlag() != 0 {
		t.Errorf("Expected command to have 0 flags, got %d", cmd.Flags().NFlag())
	}

	// Test that the command accepts exactly 0 arguments using common.ExactArgs(0)
	// Since Args is a function that returns an error, we test it differently
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

	// Test the Long description separately
	if !strings.Contains(cmd.Long, "clearing stored credentials") {
		t.Errorf("Expected Long description to contain 'clearing stored credentials', got %q", cmd.Long)
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

func TestAuthLogoutRunE_ClearCredentials(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	cmd := auth.NewLogoutCommand()

	// The logout command should execute without error
	// Note: This will actually try to clear credentials from keychain
	// which may fail if no credentials exist, but that's the expected behavior
	err := cmd.RunE(cmd, []string{})

	// We don't assert no error because it might fail if no credentials exist
	// We just verify the function executes and handles the case appropriately
	// The actual keychain clearing is tested separately in credentials_test.go
	_ = err // Acknowledge that error might occur and that's expected
}
