// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/pingidentity/pingcli/internal/constants"
	"golang.org/x/oauth2"
)

func TestSaveAndLoadTokenFromFile(t *testing.T) {
	testToken := &oauth2.Token{
		AccessToken:  "test-access-token",
		TokenType:    "Bearer",
		RefreshToken: "test-refresh-token",
		Expiry:       time.Now().Add(1 * time.Hour),
	}

	authMethod := "test-method"

	t.Cleanup(func() {
		_ = clearTokenFromFile(authMethod)
	})

	err := saveTokenToFile(testToken, authMethod)
	if err != nil {
		t.Fatalf("Failed to save token to file: %v", err)
	}

	loadedToken, err := loadTokenFromFile(authMethod)
	if err != nil {
		t.Fatalf("Failed to load token from file: %v", err)
	}

	if loadedToken.AccessToken != testToken.AccessToken {
		t.Errorf("AccessToken mismatch: got %s, want %s", loadedToken.AccessToken, testToken.AccessToken)
	}
	if loadedToken.TokenType != testToken.TokenType {
		t.Errorf("TokenType mismatch: got %s, want %s", loadedToken.TokenType, testToken.TokenType)
	}
	if loadedToken.RefreshToken != testToken.RefreshToken {
		t.Errorf("RefreshToken mismatch: got %s, want %s", loadedToken.RefreshToken, testToken.RefreshToken)
	}
	if loadedToken.Expiry.Sub(testToken.Expiry).Abs() > time.Second {
		t.Errorf("Expiry mismatch: got %v, want %v", loadedToken.Expiry, testToken.Expiry)
	}
}

func TestClearTokenFromFile(t *testing.T) {
	testToken := &oauth2.Token{
		AccessToken: "test-access-token",
		TokenType:   "Bearer",
	}

	authMethod := "test-clear-method"

	err := saveTokenToFile(testToken, authMethod)
	if err != nil {
		t.Fatalf("Failed to save token: %v", err)
	}

	err = clearTokenFromFile(authMethod)
	if err != nil {
		t.Fatalf("Failed to clear token: %v", err)
	}

	filePath, _ := getCredentialsFilePath(authMethod)
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		t.Errorf("Token file should not exist after clearing")
	}
}

func TestLoadTokenFromFile_NotExists(t *testing.T) {
	authMethod := "non-existent-method"

	_, err := loadTokenFromFile(authMethod)
	if err == nil {
		t.Error("Expected error when loading non-existent token")
	}
}

func TestSaveTokenToFile_NilToken(t *testing.T) {
	authMethod := "nil-token-test"

	err := saveTokenToFile(nil, authMethod)
	if err == nil {
		t.Error("Expected error when saving nil token")
	}
}

func TestGetCredentialsFilePath(t *testing.T) {
	authMethod := "test-path-method"

	filePath, err := getCredentialsFilePath(authMethod)
	if err != nil {
		t.Fatalf("Failed to get credentials file path: %v", err)
	}

	homeDir, _ := os.UserHomeDir()
	expectedDir := filepath.Join(homeDir, constants.PingCliDirName, constants.CredentialsDirName)

	if !strings.HasPrefix(filePath, expectedDir) {
		t.Errorf("File path %s does not start with expected directory %s", filePath, expectedDir)
	}

	if filepath.Base(filePath) != "test-path-method.json" {
		t.Errorf("File name should be test-path-method.json, got %s", filepath.Base(filePath))
	}
}

func TestClearTokenFromFile_NotExists(t *testing.T) {
	authMethod := "non-existent-clear"

	err := clearTokenFromFile(authMethod)
	if err != nil {
		t.Errorf("Expected no error when clearing non-existent file, got: %v", err)
	}
}

func TestClearAllTokenFilesForGrantType(t *testing.T) {
	// Create test tokens for different profiles and grant types
	homeDir, _ := os.UserHomeDir()
	credentialsDir := filepath.Join(homeDir, constants.PingCliDirName, constants.CredentialsDirName)
	_ = os.MkdirAll(credentialsDir, 0700)

	testFiles := []string{
		"token-abc12345_pingone_device_code_production.json",
		"token-def67890_pingone_device_code_production.json",        // Another device_code token for production
		"token-abc12345_pingone_device_code_staging.json",           // Same hash, different profile
		"token-ghi11111_pingone_authorization_code_production.json", // Different grant type, same profile
		"token-jkl22222_pingone_client_credentials_production.json",
	}

	// Create test files
	for _, filename := range testFiles {
		filePath := filepath.Join(credentialsDir, filename)
		if err := os.WriteFile(filePath, []byte("test"), 0600); err != nil {
			t.Fatalf("Failed to create test file %s: %v", filename, err)
		}
	}

	t.Cleanup(func() {
		// Clean up all test files
		for _, filename := range testFiles {
			_ = os.Remove(filepath.Join(credentialsDir, filename))
		}
	})

	// Clear device_code tokens for production profile only
	err := clearAllTokenFilesForGrantType("pingone", "device_code", "production")
	if err != nil {
		t.Fatalf("Failed to clear token files: %v", err)
	}

	// Verify device_code production files are gone
	for _, filename := range []string{
		"token-abc12345_pingone_device_code_production.json",
		"token-def67890_pingone_device_code_production.json",
	} {
		filePath := filepath.Join(credentialsDir, filename)
		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			t.Errorf("File %s should have been deleted", filename)
		}
	}

	// Verify other files still exist
	for _, filename := range []string{
		"token-abc12345_pingone_device_code_staging.json",
		"token-ghi11111_pingone_authorization_code_production.json",
		"token-jkl22222_pingone_client_credentials_production.json",
	} {
		filePath := filepath.Join(credentialsDir, filename)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("File %s should still exist", filename)
		}
	}
}

func TestClearAllTokenFilesForGrantType_NoFiles(t *testing.T) {
	// Should not error when no matching files exist
	err := clearAllTokenFilesForGrantType("pingone", "device_code", "nonexistent-profile")
	if err != nil {
		t.Errorf("Expected no error when no files match, got: %v", err)
	}
}

func TestClearAllTokenFilesForGrantType_DefaultProfile(t *testing.T) {
	homeDir, _ := os.UserHomeDir()
	credentialsDir := filepath.Join(homeDir, constants.PingCliDirName, constants.CredentialsDirName)
	_ = os.MkdirAll(credentialsDir, 0700)

	testFile := "token-abc12345_pingone_device_code_default.json"
	filePath := filepath.Join(credentialsDir, testFile)
	if err := os.WriteFile(filePath, []byte("test"), 0600); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	t.Cleanup(func() {
		_ = os.Remove(filePath)
	})

	// Clear with empty profile name (should default to "default")
	err := clearAllTokenFilesForGrantType("pingone", "device_code", "")
	if err != nil {
		t.Fatalf("Failed to clear token files: %v", err)
	}

	// Verify file is gone
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		t.Errorf("File should have been deleted with default profile")
	}
}
