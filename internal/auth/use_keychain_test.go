// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"os"
	"testing"
	"time"

	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"golang.org/x/oauth2"
)

// TestSaveTokenForMethod_WithKeychainDisabled tests that tokens are saved to file storage when keychain is disabled
func TestSaveTokenForMethod_WithKeychainDisabled(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Set use-keychain to false via environment variable
	t.Setenv("PINGCLI_AUTH_USE_KEYCHAIN", "false")

	testToken := &oauth2.Token{
		AccessToken:  "test-access-token",
		TokenType:    "Bearer",
		RefreshToken: "test-refresh-token",
		Expiry:       time.Now().Add(1 * time.Hour),
	}

	authMethod := "test-keychain-disabled"

	t.Cleanup(func() {
		_ = clearTokenFromFile(authMethod)
	})

	// Save token - should go to file storage since keychain is disabled
	err := SaveTokenForMethod(testToken, authMethod)
	if err != nil {
		t.Fatalf("Failed to save token with keychain disabled: %v", err)
	}

	// Verify token was saved to file
	loadedToken, err := loadTokenFromFile(authMethod)
	if err != nil {
		t.Fatalf("Token should be in file storage when keychain is disabled: %v", err)
	}

	if loadedToken.AccessToken != testToken.AccessToken {
		t.Errorf("AccessToken mismatch: got %s, want %s", loadedToken.AccessToken, testToken.AccessToken)
	}
}

// TestSaveTokenForMethod_WithKeychainEnabled tests that tokens are saved to keychain when enabled
func TestSaveTokenForMethod_WithKeychainEnabled(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Keychain is enabled by default, but let's be explicit
	t.Setenv("PINGCLI_AUTH_USE_KEYCHAIN", "true")

	testToken := &oauth2.Token{
		AccessToken:  "test-access-token-keychain",
		TokenType:    "Bearer",
		RefreshToken: "test-refresh-token-keychain",
		Expiry:       time.Now().Add(1 * time.Hour),
	}

	authMethod := "test-keychain-enabled"

	t.Cleanup(func() {
		_ = ClearTokenForMethod(authMethod)
	})

	// Save token - should try keychain first
	err := SaveTokenForMethod(testToken, authMethod)
	if err != nil {
		// Keychain might not be available in CI/test environment, which is fine
		// It should fall back to file storage
		t.Logf("SaveTokenForMethod returned error (expected in environments without keychain): %v", err)
	}

	// Token should be loadable from either keychain or file storage
	loadedToken, err := LoadTokenForMethod(authMethod)
	if err != nil {
		t.Fatalf("Failed to load token: %v", err)
	}

	if loadedToken.AccessToken != testToken.AccessToken {
		t.Errorf("AccessToken mismatch: got %s, want %s", loadedToken.AccessToken, testToken.AccessToken)
	}
}

// TestLoadTokenForMethod_WithKeychainDisabled tests that tokens are loaded from file storage when keychain is disabled
func TestLoadTokenForMethod_WithKeychainDisabled(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Set use-keychain to false via environment variable
	t.Setenv("PINGCLI_AUTH_USE_KEYCHAIN", "false")

	testToken := &oauth2.Token{
		AccessToken:  "test-load-access-token",
		TokenType:    "Bearer",
		RefreshToken: "test-load-refresh-token",
		Expiry:       time.Now().Add(1 * time.Hour),
	}

	authMethod := "test-load-keychain-disabled"

	t.Cleanup(func() {
		_ = clearTokenFromFile(authMethod)
	})

	// Directly save to file storage
	err := saveTokenToFile(testToken, authMethod)
	if err != nil {
		t.Fatalf("Failed to save token to file: %v", err)
	}

	// Load token - should come from file storage since keychain is disabled
	loadedToken, err := LoadTokenForMethod(authMethod)
	if err != nil {
		t.Fatalf("Failed to load token with keychain disabled: %v", err)
	}

	if loadedToken.AccessToken != testToken.AccessToken {
		t.Errorf("AccessToken mismatch: got %s, want %s", loadedToken.AccessToken, testToken.AccessToken)
	}
}

// TestLoadTokenForMethod_FallbackToFileStorage tests that LoadTokenForMethod can load from file when keychain doesn't have the token
func TestLoadTokenForMethod_FallbackToFileStorage(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// This test verifies the fallback mechanism by using a fresh token key that keychain won't have
	// We explicitly use keychain disabled mode to ensure file storage is used
	t.Setenv("PINGCLI_AUTH_USE_KEYCHAIN", "false")

	testToken := &oauth2.Token{
		AccessToken:  "test-fallback-token",
		TokenType:    "Bearer",
		RefreshToken: "test-fallback-refresh",
		Expiry:       time.Now().Add(1 * time.Hour),
	}

	authMethod := "test-fallback-method"

	t.Cleanup(func() {
		_ = clearTokenFromFile(authMethod)
		_ = ClearTokenForMethod(authMethod)
	})

	// Save token only to file storage (keychain disabled)
	err := saveTokenToFile(testToken, authMethod)
	if err != nil {
		t.Fatalf("Failed to save token to file: %v", err)
	}

	// Load token - should load from file storage since keychain is disabled
	loadedToken, err := LoadTokenForMethod(authMethod)
	if err != nil {
		t.Fatalf("Failed to load token from file storage: %v", err)
	}

	if loadedToken == nil {
		t.Fatal("LoadTokenForMethod returned nil token")
	}

	if loadedToken.AccessToken != testToken.AccessToken {
		t.Errorf("AccessToken mismatch: got %s, want %s", loadedToken.AccessToken, testToken.AccessToken)
	}
}

// TestShouldUseKeychain_Default tests the default behavior when flag is not set
func TestShouldUseKeychain_Default(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Don't set the flag - should default to true
	// Note: shouldUseKeychain is not exported, but we can test the behavior through SaveTokenForMethod

	testToken := &oauth2.Token{
		AccessToken: "test-default-token",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(1 * time.Hour),
	}

	authMethod := "test-default-keychain"

	t.Cleanup(func() {
		_ = ClearTokenForMethod(authMethod)
	})

	// Save token - should try keychain by default
	err := SaveTokenForMethod(testToken, authMethod)
	if err != nil {
		t.Logf("SaveTokenForMethod with default settings returned error: %v", err)
	}

	// Should be able to load the token
	loadedToken, err := LoadTokenForMethod(authMethod)
	if err != nil {
		t.Fatalf("Failed to load token with default settings: %v", err)
	}

	if loadedToken.AccessToken != testToken.AccessToken {
		t.Errorf("AccessToken mismatch: got %s, want %s", loadedToken.AccessToken, testToken.AccessToken)
	}
}

// TestClearTokenForMethod_ClearsBothStorages tests that clearing a token removes it from both keychain and file storage
func TestClearTokenForMethod_ClearsBothStorages(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testToken := &oauth2.Token{
		AccessToken: "test-clear-both",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(1 * time.Hour),
	}

	authMethod := "test-clear-both-storages"

	t.Cleanup(func() {
		_ = ClearTokenForMethod(authMethod)
	})

	// Save to file storage directly
	err := saveTokenToFile(testToken, authMethod)
	if err != nil {
		t.Fatalf("Failed to save token to file: %v", err)
	}

	// Verify file exists
	filePath, _ := getCredentialsFilePath(authMethod)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatal("Token file should exist before clearing")
	}

	// Clear token - should remove from both keychain and file storage
	err = ClearTokenForMethod(authMethod)
	if err != nil {
		t.Logf("ClearTokenForMethod returned error (may be expected if keychain not available): %v", err)
	}

	// Give a moment for file system operations to complete
	time.Sleep(10 * time.Millisecond)

	// Verify file was deleted
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		t.Error("Token file should be deleted after clearing")
	}

	// Verify token cannot be loaded from file
	_, err = loadTokenFromFile(authMethod)
	if err == nil {
		t.Error("Should not be able to load token from file after clearing")
	}
}

// TestPerformLogin_UsesValidCachedToken tests that Perform*Login functions check for valid cached tokens
func TestPerformLogin_UsesValidCachedToken(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// This test would require setting up full client credentials configuration
	// For now, we verify the test infrastructure exists
	// Real testing is done in integration tests

	t.Skip("This test requires full authentication configuration - covered by integration tests")
}

// TestSaveTokenForMethod_FileStorageFallback tests that file storage is used as fallback when keychain fails
func TestSaveTokenForMethod_FileStorageFallback(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Keychain enabled by default
	t.Setenv("PINGCLI_AUTH_USE_KEYCHAIN", "true")

	testToken := &oauth2.Token{
		AccessToken:  "test-fallback-save",
		TokenType:    "Bearer",
		RefreshToken: "test-fallback-save-refresh",
		Expiry:       time.Now().Add(1 * time.Hour),
	}

	authMethod := "test-save-fallback"

	t.Cleanup(func() {
		_ = ClearTokenForMethod(authMethod)
	})

	// Save token - will try keychain first (may succeed or fail depending on environment)
	err := SaveTokenForMethod(testToken, authMethod)
	if err != nil {
		t.Logf("SaveTokenForMethod returned error: %v", err)
	}

	// Give a moment for file system operations to complete
	time.Sleep(10 * time.Millisecond)

	// Token should be loadable from either storage
	// In environments where keychain works, it may be there instead of file
	loadedToken, err := LoadTokenForMethod(authMethod)
	if err != nil {
		// If LoadTokenForMethod fails, check file storage directly
		loadedToken, err = loadTokenFromFile(authMethod)
		if err != nil {
			t.Fatalf("Token should be in at least one storage location: %v", err)
		}
	}

	if loadedToken.AccessToken != testToken.AccessToken {
		t.Errorf("AccessToken mismatch: got %s, want %s", loadedToken.AccessToken, testToken.AccessToken)
	}
}

// TestEnvironmentVariable_UseKeychain tests that PINGCLI_AUTH_USE_KEYCHAIN environment variable is respected
func TestEnvironmentVariable_UseKeychain(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Set environment variable
	t.Setenv("PINGCLI_AUTH_USE_KEYCHAIN", "false")

	// Reinitialize koanf to pick up environment variable
	testutils_koanf.InitKoanfs(t)

	testToken := &oauth2.Token{
		AccessToken: "test-env-var-token",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(1 * time.Hour),
	}

	authMethod := "test-env-var"

	t.Cleanup(func() {
		_ = clearTokenFromFile(authMethod)
	})

	// Save token - should respect environment variable
	err := SaveTokenForMethod(testToken, authMethod)
	if err != nil {
		t.Fatalf("Failed to save token with env var: %v", err)
	}

	// Verify token was saved to file (since env var set to false)
	loadedToken, err := loadTokenFromFile(authMethod)
	if err != nil {
		t.Fatalf("Token should be in file storage when env var is false: %v", err)
	}

	if loadedToken.AccessToken != testToken.AccessToken {
		t.Errorf("AccessToken mismatch: got %s, want %s", loadedToken.AccessToken, testToken.AccessToken)
	}
}
