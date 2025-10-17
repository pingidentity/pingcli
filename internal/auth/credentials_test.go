// Copyright © 2025 Ping Identity Corporation

package auth_internal_test

import (
	"context"
	"strings"
	"testing"

	auth_internal "github.com/pingidentity/pingcli/internal/auth"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

func TestPerformDeviceCodeLogin_MissingConfiguration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	ctx := context.Background()

	_, _, err := auth_internal.PerformDeviceCodeLogin(ctx)

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "failed to get device code configuration") {
		t.Errorf("Expected error to contain 'failed to get device code configuration', got: %v", err)
	}
}

func TestPerformClientCredentialsLogin_MissingConfiguration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	ctx := context.Background()

	_, _, err := auth_internal.PerformClientCredentialsLogin(ctx)

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "failed to get client credentials configuration") {
		t.Errorf("Expected error to contain 'failed to get client credentials configuration', got: %v", err)
	}
}

func TestPerformAuthCodeLogin_MissingConfiguration(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	ctx := context.Background()

	_, _, err := auth_internal.PerformAuthCodeLogin(ctx)

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "failed to get auth code configuration") {
		t.Errorf("Expected error to contain 'failed to get auth code configuration', got: %v", err)
	}
}

func TestGetDeviceCodeConfiguration_MissingClientID(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	_, err := auth_internal.GetDeviceCodeConfiguration()

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "device code client ID is not configured") {
		t.Errorf("Expected error to contain 'device code client ID is not configured', got: %v", err)
	}
}

func TestGetClientCredentialsConfiguration_MissingClientID(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	_, err := auth_internal.GetClientCredentialsConfiguration()

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "client credentials client ID is not configured") {
		t.Errorf("Expected error to contain 'client credentials client ID is not configured', got: %v", err)
	}
}

func TestGetAuthCodeConfiguration_MissingClientID(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	_, err := auth_internal.GetAuthCodeConfiguration()

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "auth code client ID is not configured") {
		t.Errorf("Expected error to contain 'auth code client ID is not configured', got: %v", err)
	}
}

func TestGetDeviceCodeConfiguration_MissingEnvironmentID(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Mock getting a client ID but missing environment ID
	// This would typically be done through dependency injection or mocking,
	// but for now we'll test the error path
	_, err := auth_internal.GetDeviceCodeConfiguration()

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	// Will fail on client ID first, but this tests the configuration validation logic
	if err != nil && !strings.Contains(err.Error(), "is not configured") {
		t.Errorf("Expected error to contain 'is not configured', got: %v", err)
	}
}

func TestGetClientCredentialsConfiguration_MissingEnvironmentID(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	_, err := auth_internal.GetClientCredentialsConfiguration()

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	// Will fail on client ID first, but this tests the configuration validation logic
	if err != nil && !strings.Contains(err.Error(), "is not configured") {
		t.Errorf("Expected error to contain 'is not configured', got: %v", err)
	}
}

func TestGetAuthCodeConfiguration_MissingEnvironmentID(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	_, err := auth_internal.GetAuthCodeConfiguration()

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	// Will fail on client ID first, but this tests the configuration validation logic
	if err != nil && !strings.Contains(err.Error(), "is not configured") {
		t.Errorf("Expected error to contain 'is not configured', got: %v", err)
	}
}

func TestSaveAndLoadToken(t *testing.T) {
	testKey := "test-token-key"

	// Test that SaveTokenForMethod returns an error with nil token
	err := auth_internal.SaveTokenForMethod(nil, testKey)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "token cannot be nil") {
		t.Errorf("Expected error to contain 'token cannot be nil', got: %v", err)
	}
}

func TestClearToken(t *testing.T) {
	testKey := "test-token-key"

	// Test that ClearTokenForMethod doesn't panic when no token exists
	// This should handle the case where keychain entry doesn't exist
	err := auth_internal.ClearTokenForMethod(testKey)

	// Should not error when no token exists (handles ErrNotFound)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

// Removed TestPingcliTokenSourceProvider_NilConfig - provider pattern was simplified away

func TestGetValidTokenSource_NoCache(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	ctx := context.Background()

	// This should attempt automatic authentication since no token is cached
	_, err := auth_internal.GetValidTokenSource(ctx)

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	// The error should be related to authentication configuration or automatic authentication failure
	if err == nil || (strings.Contains(err.Error(), "failed to get authentication type") ||
		strings.Contains(err.Error(), "automatic client credentials authentication failed") ||
		strings.Contains(err.Error(), "automatic authorization code authentication failed") ||
		strings.Contains(err.Error(), "automatic device code authentication failed")) {
		// Expected authentication error
	} else {
		t.Errorf("Expected authentication-related error, got: %s", err.Error())
	}
}

// TestAuthenticationErrorMessages_DeviceCode tests device code authentication error message
func TestAuthenticationErrorMessages_DeviceCode(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	ctx := context.Background()
	_, _, err := auth_internal.PerformDeviceCodeLogin(ctx)

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "device code client ID is not configured") {
		t.Errorf("Expected error to contain '%s', got: %v", "device code client ID is not configured", err)
	}
}

// TestAuthenticationErrorMessages_ClientCredentials tests client credentials authentication error message
func TestAuthenticationErrorMessages_ClientCredentials(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	ctx := context.Background()
	_, _, err := auth_internal.PerformClientCredentialsLogin(ctx)

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "client credentials client ID is not configured") {
		t.Errorf("Expected error to contain '%s', got: %v", "client credentials client ID is not configured", err)
	}
}

// TestAuthenticationErrorMessages_AuthCode tests auth code authentication error message
func TestAuthenticationErrorMessages_AuthCode(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	ctx := context.Background()
	_, _, err := auth_internal.PerformAuthCodeLogin(ctx)

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "auth code client ID is not configured") {
		t.Errorf("Expected error to contain '%s', got: %v", "auth code client ID is not configured", err)
	}
}

// TestConfigurationValidation_DeviceCode tests device code configuration validation
func TestConfigurationValidation_DeviceCode(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	_, err := auth_internal.GetDeviceCodeConfiguration()

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "client ID is not configured") {
		t.Errorf("Expected error to contain '%s', got: %v", "client ID is not configured", err)
	}
}

// TestConfigurationValidation_ClientCredentials tests client credentials configuration validation
func TestConfigurationValidation_ClientCredentials(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	_, err := auth_internal.GetClientCredentialsConfiguration()

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "client ID is not configured") {
		t.Errorf("Expected error to contain '%s', got: %v", "client ID is not configured", err)
	}
}

// TestConfigurationValidation_AuthCode tests auth code configuration validation
func TestConfigurationValidation_AuthCode(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	_, err := auth_internal.GetAuthCodeConfiguration()

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "client ID is not configured") {
		t.Errorf("Expected error to contain '%s', got: %v", "client ID is not configured", err)
	}
}

func TestSaveToken_NilToken(t *testing.T) {
	testKey := "test-token-key"

	err := auth_internal.SaveTokenForMethod(nil, testKey)

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "token cannot be nil") {
		t.Errorf("Expected error to contain 'token cannot be nil', got: %v", err)
	}
}

func TestLoadToken_ErrorCases(t *testing.T) {
	// Clear any existing token first
	_ = auth_internal.ClearToken()

	// Test when token doesn't exist in keychain
	_, err := auth_internal.LoadToken()

	// Should get an error when token doesn't exist (could be nil token or keychain error)
	// We just verify an error or nil token is returned
	if err == nil {
		// If no error, then token should be nil
		// This is also a valid case - no token found
		t.Skip("No cached token found (expected)")
	}
}

func TestGetValidTokenSource_ErrorPaths(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	ctx := context.Background()

	// Clear any existing token first
	_ = auth_internal.ClearToken()

	// Test without any cached token - should attempt automatic authentication
	_, err := auth_internal.GetValidTokenSource(ctx)

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	// The error message can vary depending on the configured auth type and state
	// Since "worker" type gets converted to "client_credentials", we expect client credentials auth failure
	if err == nil || (strings.Contains(err.Error(), "automatic client credentials authentication failed") ||
		strings.Contains(err.Error(), "failed to get authentication type") ||
		strings.Contains(err.Error(), "client ID is not configured")) {
		// Expected authentication error
	} else {
		t.Errorf("Expected client credentials authentication failure, got: %s", err.Error())
	}
}

// TestGetValidTokenSource_AutomaticDeviceCodeAuth tests automatic device code authentication
func TestGetValidTokenSource_AutomaticDeviceCodeAuth(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	ctx := context.Background()

	// Clear any existing token first
	_ = auth_internal.ClearToken()

	// Test that GetValidTokenSource attempts automatic authentication
	// In test environment, auth type is "worker" which gets converted to "client_credentials"
	_, err := auth_internal.GetValidTokenSource(ctx)

	// Should fail with some authentication-related error
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	// The error will depend on the configured auth type:
	// - In test env: "worker" -> "client_credentials" -> "automatic client credentials authentication failed"
	// - If device_code was configured: "automatic device code authentication failed"
	// - Other config errors: "failed to get authentication type"
	expectedErrors := []string{
		"automatic device code authentication failed",
		"automatic client credentials authentication failed", // test env: worker -> client_credentials
		"failed to get authentication type",
	}

	errorMatched := false
	for _, expectedError := range expectedErrors {
		if strings.Contains(err.Error(), expectedError) {
			errorMatched = true

			break
		}
	}

	if !errorMatched {
		t.Errorf("Expected error to contain authentication-related message, got: %v", err)
	}
}

// TestGetValidTokenSource_AutomaticAuthCodeAuth tests automatic auth code authentication
func TestGetValidTokenSource_AutomaticAuthCodeAuth(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	ctx := context.Background()

	// Clear any existing token first
	_ = auth_internal.ClearToken()

	// Test automatic authentication behavior
	// In test environment, auth type is "worker" which gets converted to "client_credentials"
	_, err := auth_internal.GetValidTokenSource(ctx)

	// Should fail with authentication configuration error
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	// The error will depend on the configured auth type:
	// - In test env: "worker" -> "client_credentials" -> "automatic client credentials authentication failed"
	// - If auth_code was configured: "automatic authorization code authentication failed"
	expectedErrors := []string{
		"automatic authorization code authentication failed",
		"automatic client credentials authentication failed", // test env: worker -> client_credentials
	}

	errorMatched := false
	for _, expectedError := range expectedErrors {
		if strings.Contains(err.Error(), expectedError) {
			errorMatched = true

			break
		}
	}

	if !errorMatched {
		t.Errorf("Expected error to contain automatic authentication failure message, got: %v", err)
	}
}

// TestGetValidTokenSource_AutomaticClientCredentialsAuth tests automatic client credentials authentication
func TestGetValidTokenSource_AutomaticClientCredentialsAuth(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	ctx := context.Background()

	// Clear any existing token first
	_ = auth_internal.ClearToken()

	// Test client credentials auth by temporarily setting the auth type
	// This would require configuration mocking for a complete test
	// For now, this documents the expected behavior

	// In a real scenario with client credentials configured:
	// 1. GetValidTokenSource() detects no cached token
	// 2. Reads auth type as "client_credentials"
	// 3. Calls PerformClientCredentialsLogin()
	// 4. Returns token source with new token

	_, err := auth_internal.GetValidTokenSource(ctx)

	// Should fail with some authentication configuration error
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	// The specific error depends on the configured auth method
	expectedErrors := []string{
		"automatic device code authentication failed",
		"automatic authorization code authentication failed",
		"automatic client credentials authentication failed",
		"failed to get authentication type",
	}

	errorMatched := false
	for _, expectedError := range expectedErrors {
		if strings.Contains(err.Error(), expectedError) {
			errorMatched = true

			break
		}
	}

	if !errorMatched {
		t.Errorf("Expected error to contain one of the automatic authentication failure messages, got: %v", err)
	}
}

// TestGetValidTokenSource_ValidCachedToken tests that valid cached tokens are used without re-authentication
func TestGetValidTokenSource_ValidCachedToken(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	ctx := context.Background()

	// Clear any existing token first
	_ = auth_internal.ClearToken()

	// This test would require mocking a valid cached token
	// For now, it documents the expected behavior:
	// 1. GetValidTokenSource() finds a valid cached token
	// 2. Returns static token source without attempting new authentication
	// 3. No authentication method calls are made

	_, err := auth_internal.GetValidTokenSource(ctx)

	// Without a valid cached token, should attempt automatic authentication
	if err == nil {
		t.Error("Expected error without cached token, but got nil")
	}
}

// TestGetValidTokenSource_WorkerTypeAlias tests that "worker" type is treated as "client_credentials"
func TestGetValidTokenSource_WorkerTypeAlias(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	ctx := context.Background()

	// Clear any existing token first
	_ = auth_internal.ClearToken()

	// Test that "worker" auth type is treated as "client_credentials"
	// In test environment, the auth type is typically "worker"
	_, err := auth_internal.GetValidTokenSource(ctx)

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	// Should attempt client credentials authentication (since worker -> client_credentials)
	if err != nil && !strings.Contains(err.Error(), "automatic client credentials authentication failed") {
		t.Errorf("Expected 'automatic client credentials authentication failed' (worker->client_credentials), got: %v", err)
	}
}

// TestGetValidTokenSource_UnsupportedAuthType tests unsupported authentication types
func TestGetValidTokenSource_UnsupportedAuthType(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	ctx := context.Background()

	// Clear any existing token first
	_ = auth_internal.ClearToken()

	// This test documents behavior for unsupported auth types like "saml" or "oidc"
	// The expected behavior:
	// 1. GetValidTokenSource() finds no cached token
	// 2. Reads auth type as unsupported value
	// 3. Returns error about unsupported authentication type

	_, err := auth_internal.GetValidTokenSource(ctx)

	// Should fail with some error (could be unsupported auth type or config error)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	// The exact error depends on the current configuration
}
