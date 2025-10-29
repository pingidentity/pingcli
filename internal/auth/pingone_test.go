package auth_internal_test

import (
	"strings"
	"testing"

	auth_internal "github.com/pingidentity/pingcli/internal/auth"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// Test GetPingOneAccessToken function with missing configuration
func TestGetPingOneAccessToken_MissingConfiguration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	token, err := auth_internal.GetPingOneAccessToken()

	// In test environment, worker credentials may be configured
	if err == nil {
		// Success - valid credentials were configured
		if token == "" {
			t.Error("Expected token when no error, but got empty string")
		}
		t.Skip("Authentication succeeded (valid credentials configured)")
	}
	// Should fail because no client ID is configured for worker authentication
	if !strings.Contains(err.Error(), "client ID is required") &&
		!strings.Contains(err.Error(), "client ID is not configured") &&
		!strings.Contains(err.Error(), "failed to get") {
		t.Errorf("Expected configuration error, got: %v", err)
	}
}

// Test ClearPingOneClientCache function
func TestClearPingOneClientCache(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// This should not panic or error
	auth_internal.ClearPingOneClientCache()

	// Function should complete without issue - if we get here, test passes
}
