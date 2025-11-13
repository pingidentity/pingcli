package auth_internal_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	auth_internal "github.com/pingidentity/pingcli/internal/commands/auth"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// createIntegrationTestConfig generates test configuration with dummy values
// This is only used for configuration validation tests, not actual authentication
func createIntegrationTestConfig() string {
	return `activeProfile: integration
integration:
    description: "Integration test profile"
    noColor: true
    outputFormat: json
    service:
        pingOne:
            regionCode: NA
            authentication:
                type: clientCredentials
                environmentID: 00000000-0000-0000-0000-000000000000
                clientCredentials:
                    clientID: 00000000-0000-0000-0000-000000000001
                    clientSecret: dummy-secret-for-config-test
                    scopes: ["openid"]
                deviceCode:
                    clientID: ""
                    scopes: []
                authorizationCode:
                    clientID: ""
                    redirectURIPath: ""
                    redirectURIPort: ""
                    scopes: []
`
}

func TestClientCredentialsAuthentication_Integration(t *testing.T) {
	// Skip if running in CI environment without credentials
	if os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID") == "" ||
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET") == "" ||
		os.Getenv("TEST_PINGONE_ENVIRONMENT_ID") == "" ||
		os.Getenv("TEST_PINGONE_REGION_CODE") == "" {
		t.Skip("Skipping integration test - missing required environment variables")
	}

	// Default scopes if not provided
	scopes := os.Getenv("TEST_PINGONE_SCOPES")
	if scopes == "" {
		scopes = "openid"
	}

	// Initialize configuration with test config
	configuration.InitAllOptions()
	testConfig := fmt.Sprintf(`activeProfile: integration
integration:
    description: "Integration test profile"
    noColor: true
    outputFormat: json
    service:
        pingOne:
            regionCode: %s
            authentication:
                type: clientCredentials
                environmentID: %s
                clientCredentials:
                    clientID: %s
                    clientSecret: %s
                    scopes: ["%s"]
`,
		os.Getenv("TEST_PINGONE_REGION_CODE"),
		os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID"),
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET"),
		scopes)
	testutils_koanf.InitKoanfsCustomFile(t, testConfig)

	// Clear any existing tokens to ensure fresh authentication
	err := auth_internal.ClearToken()
	if err != nil {
		t.Fatalf("Should be able to clear existing tokens: %v", err)
	}

	// Test performing fresh client credentials authentication
	result, err := auth_internal.PerformClientCredentialsLogin(context.Background())
	if err != nil {
		t.Fatalf("Client credentials authentication should succeed: %v", err)
	}
	if result.Token == nil {
		t.Fatal("Token should not be nil")
	}
	if result.Token.AccessToken == "" {
		t.Error("Access token should not be empty")
	}
	if !result.Token.Valid() {
		t.Error("Token should be valid")
	}
	if !result.NewAuth {
		t.Error("Should be a new authentication since we cleared tokens")
	}
}

func TestValidTokenSource_Integration(t *testing.T) {
	// Skip if running in CI environment without credentials
	if os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID") == "" ||
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET") == "" ||
		os.Getenv("TEST_PINGONE_ENVIRONMENT_ID") == "" ||
		os.Getenv("TEST_PINGONE_REGION_CODE") == "" {
		t.Skip("Skipping integration test - missing required environment variables")
	}

	// Default scopes if not provided
	scopes := os.Getenv("TEST_PINGONE_SCOPES")
	if scopes == "" {
		scopes = "openid"
	}

	// Initialize configuration with test config
	configuration.InitAllOptions()
	testConfig := fmt.Sprintf(`activeProfile: integration
integration:
    description: "Integration test profile"
    noColor: true
    outputFormat: json
    service:
        pingOne:
            regionCode: %s
            authentication:
                type: clientCredentials
                environmentID: %s
                clientCredentials:
                    clientID: %s
                    clientSecret: %s
                    scopes: ["%s"]
`,
		os.Getenv("TEST_PINGONE_REGION_CODE"),
		os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID"),
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET"),
		scopes)
	testutils_koanf.InitKoanfsCustomFile(t, testConfig)

	// Clear any existing tokens to ensure fresh authentication
	err := auth_internal.ClearToken()
	if err != nil {
		t.Fatalf("Should be able to clear existing tokens: %v", err)
	}

	// First authenticate to have a token
	result, err := auth_internal.PerformClientCredentialsLogin(context.Background())
	if err != nil {
		t.Fatalf("Client credentials authentication should succeed: %v", err)
	}
	if result.Token == nil {
		t.Fatal("Token should not be nil")
	}

	// Now test getting valid token source from cached token
	tokenSource, err := auth_internal.GetValidTokenSource(context.Background())
	if err != nil {
		t.Fatalf("Should be able to get valid token source after authentication: %v", err)
	}
	if tokenSource == nil {
		t.Fatal("Valid token source should not be nil")
	}

	// Test getting token from source
	retrievedToken, err := tokenSource.Token()
	if err != nil {
		t.Fatalf("Should be able to get token from valid token source: %v", err)
	}
	if retrievedToken.AccessToken == "" {
		t.Error("Retrieved access token should not be empty")
	}
}

func TestDeviceCodeConfiguration_Integration(t *testing.T) {
	// Initialize configuration with test config
	configuration.InitAllOptions()
	testutils_koanf.InitKoanfsCustomFile(t, createIntegrationTestConfig())

	// Test getting device code configuration - with empty values, this should fail validation
	// This test verifies that empty device code configuration is properly validated
	_, err := auth_internal.GetDeviceCodeConfiguration()
	if err == nil {
		t.Fatal("Should get validation error with empty device code configuration")
	}
	// Verify we get the expected configuration error
	if !strings.Contains(err.Error(), "client ID is not configured") {
		t.Errorf("Expected client ID configuration error, got: %v", err)
	}
}

func TestAuthorizationCodeConfiguration_Integration(t *testing.T) {
	// Initialize configuration with test config
	configuration.InitAllOptions()
	testutils_koanf.InitKoanfsCustomFile(t, createIntegrationTestConfig())

	// Test getting auth code configuration - with empty values, this should fail validation
	// This test verifies that empty auth code configuration is properly validated
	_, err := auth_internal.GetAuthorizationCodeConfiguration()
	if err == nil {
		t.Fatal("Should get validation error with empty auth code configuration")
	}
	// Verify we get the expected configuration error
	if !strings.Contains(err.Error(), "client ID is not configured") {
		t.Errorf("Expected client ID configuration error, got: %v", err)
	}
}
