package auth_internal_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	auth_internal "github.com/pingidentity/pingcli/internal/auth"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// createIntegrationTestConfig generates test configuration from environment variables
func createIntegrationTestConfig() string {
	return fmt.Sprintf(`activeProfile: integration
integration:
    description: "Integration test profile"
    noColor: true
    outputFormat: json
    service:
        pingOne:
            regionCode: %s
            authentication:
                type: clientCredentials
                clientCredentials:
                    clientID: %s
                    clientSecret: %s
                    environmentID: %s
                    scopes: ["%s"]
                deviceCode:
                    clientID: %s
                    environmentID: %s
                    scopes: ["%s"]
                authCode:
                    clientID: %s
                    environmentID: %s
                    redirectURI: "http://localhost:8080/callback"
                    scopes: ["%s"]
`,
		os.Getenv("TEST_PINGONE_REGION_CODE"),
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID"),
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET"),
		os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
		os.Getenv("TEST_PINGONE_SCOPES"),
		os.Getenv("TEST_PINGONE_DEVICE_CODE_CLIENT_ID"),
		os.Getenv("TEST_PINGONE_DEVICE_CODE_ENVIRONMENT_ID"),
		os.Getenv("TEST_PINGONE_DEVICE_CODE_SCOPES"),
		os.Getenv("TEST_PINGONE_AUTH_CODE_CLIENT_ID"),
		os.Getenv("TEST_PINGONE_AUTH_CODE_ENVIRONMENT_ID"),
		os.Getenv("TEST_PINGONE_AUTH_CODE_SCOPES"))
}

func TestClientCredentialsAuthentication_Integration(t *testing.T) {
	// Skip if running in CI environment without credentials
	if os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID") == "" ||
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET") == "" ||
		os.Getenv("TEST_PINGONE_ENVIRONMENT_ID") == "" ||
		os.Getenv("TEST_PINGONE_REGION_CODE") == "" {
		t.Skip("Skipping integration test - missing required environment variables")
	}

	// Initialize configuration with test config
	configuration.InitAllOptions()
	testutils_koanf.InitKoanfsCustomFile(t, createIntegrationTestConfig())

	// Clear any existing tokens to ensure fresh authentication
	err := auth_internal.ClearToken()
	if err != nil {
		t.Fatalf("Should be able to clear existing tokens: %v", err)
	}

	// Test performing fresh client credentials authentication
	token, newAuth, err := auth_internal.PerformClientCredentialsLogin(context.Background())
	if err != nil {
		t.Fatalf("Client credentials authentication should succeed: %v", err)
	}
	if token == nil {
		t.Fatal("Token should not be nil")
	}
	if token.AccessToken == "" {
		t.Error("Access token should not be empty")
	}
	if !token.Valid() {
		t.Error("Token should be valid")
	}
	if !newAuth {
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

	// Initialize configuration with test config
	configuration.InitAllOptions()
	testutils_koanf.InitKoanfsCustomFile(t, createIntegrationTestConfig())

	// Clear any existing tokens to ensure fresh authentication
	err := auth_internal.ClearToken()
	if err != nil {
		t.Fatalf("Should be able to clear existing tokens: %v", err)
	}

	// First authenticate to have a token
	token, _, err := auth_internal.PerformClientCredentialsLogin(context.Background())
	if err != nil {
		t.Fatalf("Client credentials authentication should succeed: %v", err)
	}
	if token == nil {
		t.Fatal("Token should not be nil")
	}

	// Save the token to keychain for the next test
	err = auth_internal.SaveToken(token)
	if err != nil {
		t.Fatalf("Should be able to save token to keychain: %v", err)
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
	// Skip if running in CI environment without credentials
	if os.Getenv("TEST_PINGONE_DEVICE_CODE_CLIENT_ID") == "" ||
		os.Getenv("TEST_PINGONE_DEVICE_CODE_ENVIRONMENT_ID") == "" ||
		os.Getenv("TEST_PINGONE_REGION_CODE") == "" {
		t.Skip("Skipping integration test - missing required environment variables")
	}

	// Initialize configuration with test config
	configuration.InitAllOptions()
	testutils_koanf.InitKoanfsCustomFile(t, createIntegrationTestConfig())

	// Test getting device code configuration - this validates the configuration is properly set
	_, err := auth_internal.GetDeviceCodeConfiguration()
	if err != nil {
		t.Fatalf("Should be able to get device code configuration: %v", err)
	}
}

func TestAuthCodeConfiguration_Integration(t *testing.T) {
	// Skip if running in CI environment without credentials
	if os.Getenv("TEST_PINGONE_AUTH_CODE_CLIENT_ID") == "" ||
		os.Getenv("TEST_PINGONE_AUTH_CODE_ENVIRONMENT_ID") == "" ||
		os.Getenv("TEST_PINGONE_REGION_CODE") == "" {
		t.Skip("Skipping integration test - missing required environment variables")
	}

	// Initialize configuration with test config
	configuration.InitAllOptions()
	testutils_koanf.InitKoanfsCustomFile(t, createIntegrationTestConfig())

	// Test getting auth code configuration - this validates the configuration is properly set
	_, err := auth_internal.GetAuthCodeConfiguration()
	if err != nil {
		t.Fatalf("Should be able to get auth code configuration: %v", err)
	}
}
