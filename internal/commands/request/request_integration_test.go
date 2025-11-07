// Copyright © 2025 Ping Identity Corporation

package request_internal_test

import (
	"context"
	"os"
	"strings"
	"testing"

	auth_internal "github.com/pingidentity/pingcli/internal/commands/auth"
	request_internal "github.com/pingidentity/pingcli/internal/commands/request"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// TestRequestPingOne_RealAuth tests the complete request flow with real authentication
func TestRequestPingOne_RealAuth(t *testing.T) {
	// Skip if not in CI environment or missing credentials
	clientID := os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID")
	clientSecret := os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET")
	environmentID := os.Getenv("TEST_PINGONE_ENVIRONMENT_ID")
	regionCode := os.Getenv("TEST_PINGONE_REGION_CODE")

	if clientID == "" || clientSecret == "" || environmentID == "" || regionCode == "" {
		t.Skip("Skipping integration test: missing TEST_PINGONE_* environment variables")
	}

	// Initialize test configuration using existing pattern
	testutils_koanf.InitKoanfs(t)

	// Set service to pingone
	t.Setenv("PINGCLI_REQUEST_SERVICE", "pingone")

	// Clear any existing tokens to ensure fresh authentication
	err := auth_internal.ClearToken()
	if err != nil {
		t.Fatalf("Failed to clear token: %v", err)
	}

	// First authenticate
	_, err = auth_internal.PerformClientCredentialsLogin(context.Background())
	if err != nil {
		t.Fatalf("Authentication should succeed: %v", err)
	}

	// Test simple environment API request - this should succeed if auth is working
	err = request_internal.RunInternalRequest("environments")
	if err != nil {
		t.Fatalf("PingOne environments request should succeed with valid auth: %v", err)
	}

	// Clean up
	err = auth_internal.ClearToken()
	if err != nil {
		t.Fatalf("Failed to clear token after test: %v", err)
	}
}

// TestRequestPingOne_NoAuth tests that request command properly handles missing authentication
// Note: This test is skipped because GetValidTokenSource now automatically authenticates
// with client_credentials when properly configured, which is the desired behavior.
func TestRequestPingOne_NoAuth(t *testing.T) {
	t.Skip("Skipping: GetValidTokenSource now automatically handles authentication when configured")
}

// TestGetAPIURLForRegion_EnvironmentsEndpoint_Integration tests URL building for environments endpoint
func TestGetAPIURLForRegion_EnvironmentsEndpoint_Integration(t *testing.T) {
	// Skip if not in CI environment or missing region code
	regionCode := os.Getenv("TEST_PINGONE_REGION_CODE")
	if regionCode == "" {
		t.Skip("Skipping integration test: missing TEST_PINGONE_REGION_CODE environment variable")
	}

	// Initialize test configuration
	testutils_koanf.InitKoanfs(t)

	uri := "environments"
	url, err := request_internal.GetAPIURLForRegion(uri)
	if err != nil {
		t.Fatalf("Should be able to build API URL: %v", err)
	}
	if url == "" {
		t.Error("URL should not be empty")
	}

	// Verify URL contains the URI
	if !strings.Contains(url, uri) {
		t.Errorf("URL should contain the original URI %q, got: %q", uri, url)
	}
}

// TestGetAPIURLForRegion_NestedEndpoint_Integration tests URL building for nested endpoint
func TestGetAPIURLForRegion_NestedEndpoint_Integration(t *testing.T) {
	// Skip if not in CI environment or missing region code
	regionCode := os.Getenv("TEST_PINGONE_REGION_CODE")
	if regionCode == "" {
		t.Skip("Skipping integration test: missing TEST_PINGONE_REGION_CODE environment variable")
	}

	// Initialize test configuration
	testutils_koanf.InitKoanfs(t)

	uri := "environments/123/users"
	url, err := request_internal.GetAPIURLForRegion(uri)
	if err != nil {
		t.Fatalf("Should be able to build API URL: %v", err)
	}
	if url == "" {
		t.Error("URL should not be empty")
	}

	// Verify URL contains the URI
	if !strings.Contains(url, uri) {
		t.Errorf("URL should contain the original URI %q, got: %q", uri, url)
	}
}

// TestRequestDataFunctions_GetDataRaw_Integration tests getDataRaw function
func TestRequestDataFunctions_GetDataRaw_Integration(t *testing.T) {
	// Initialize test configuration
	testutils_koanf.InitKoanfs(t)

	data, err := request_internal.GetDataRaw()
	if err != nil {
		t.Fatalf("Should be able to get raw data: %v", err)
	}
	// Raw data should be empty by default in test environment
	if data != "" {
		t.Errorf("Raw data should be empty by default, got: %q", data)
	}
}

// TestRequestDataFunctions_GetDataFile_Integration tests getDataFile function
func TestRequestDataFunctions_GetDataFile_Integration(t *testing.T) {
	// Initialize test configuration
	testutils_koanf.InitKoanfs(t)

	data, err := request_internal.GetDataFile()
	if err != nil {
		t.Fatalf("Should be able to get file data: %v", err)
	}
	// File data should be empty by default in test environment
	if data != "" {
		t.Errorf("File data should be empty by default, got: %q", data)
	}
}
