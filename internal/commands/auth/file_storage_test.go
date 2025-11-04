// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

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
	expectedDir := filepath.Join(homeDir, ".pingcli", "credentials")

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
