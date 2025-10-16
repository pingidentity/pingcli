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

	_, err := auth_internal.GetPingOneAccessToken()

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	// Should fail because no client ID is configured for worker authentication
	if err != nil && !strings.Contains(err.Error(), "client ID is required") {
		t.Errorf("Expected error to contain 'client ID is required', got: %v", err)
	}
}

// Test ClearPingOneClientCache function
func TestClearPingOneClientCache(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// This should not panic or error
	auth_internal.ClearPingOneClientCache()

	// Function should complete without issue - if we get here, test passes
}
