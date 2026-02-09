// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
)

var (
	errKeychainUnavailable = errors.New("keychain unavailable")
)

type MockTokenStorage struct {
	mock.Mock
}

func (m *MockTokenStorage) SaveToken(token *oauth2.Token) error {
	args := m.Called(token)

	return args.Error(0)
}

func (m *MockTokenStorage) LoadToken() (*oauth2.Token, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	token, ok := args.Get(0).(*oauth2.Token)
	if !ok {
		return nil, args.Error(1)
	}

	return token, args.Error(1)
}

func (m *MockTokenStorage) ClearToken() error {
	args := m.Called()

	return args.Error(0)
}

func (m *MockTokenStorage) ClearAllTokens() error {
	args := m.Called()

	return args.Error(0)
}

func TestSaveTokenForMethod_FallsBackToFileWhenKeychainSaveFails(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	tmp := t.TempDir()
	t.Setenv("HOME", tmp)
	// Ensure keychain is the selected storage mode (default is secure_local)
	t.Setenv("PINGCLI_LOGIN_STORAGE_TYPE", "secure_local")

	mockStorage := new(MockTokenStorage)
	mockStorage.On("SaveToken", mock.Anything).Return(errKeychainUnavailable)

	old := newKeychainStorage
	t.Cleanup(func() { newKeychainStorage = old })
	newKeychainStorage = func(serviceName, username string) (tokenStorage, error) {
		return mockStorage, nil
	}

	authMethod := "test-auth-method"
	expectedToken := &oauth2.Token{
		AccessToken:  "access",
		TokenType:    "Bearer",
		RefreshToken: "refresh",
		Expiry:       time.Now().Add(1 * time.Hour),
	}

	location, err := SaveTokenForMethod(expectedToken, authMethod)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if location.Keychain {
		t.Fatalf("expected Keychain=false, got true")
	}
	if !location.File {
		t.Fatalf("expected File=true, got false")
	}

	// Verify it actually wrote the expected file under HOME
	filePath := filepath.Join(tmp, ".pingcli", "credentials", authMethod+".json")
	if _, statErr := os.Stat(filePath); statErr != nil {
		t.Fatalf("expected credentials file to exist at %s, got stat error: %v", filePath, statErr)
	}

	loaded, err := loadTokenFromFile(authMethod)
	if err != nil {
		t.Fatalf("expected to load token from file, got %v", err)
	}
	if loaded.AccessToken != expectedToken.AccessToken {
		t.Fatalf("expected access token %q, got %q", expectedToken.AccessToken, loaded.AccessToken)
	}

	mockStorage.AssertExpectations(t)
}

func TestSaveTokenForMethod_UsesKeychainWhenAvailable(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	tmp := t.TempDir()
	t.Setenv("HOME", tmp)
	t.Setenv("PINGCLI_LOGIN_STORAGE_TYPE", "secure_local")

	old := newKeychainStorage
	t.Cleanup(func() { newKeychainStorage = old })

	mockStorage := new(MockTokenStorage)
	mockStorage.On("SaveToken", mock.Anything).Return(nil)

	newKeychainStorage = func(serviceName, username string) (tokenStorage, error) {
		return mockStorage, nil
	}

	authMethod := "test-auth-method-keychain"
	token := &oauth2.Token{AccessToken: "access", TokenType: "Bearer", Expiry: time.Now().Add(1 * time.Hour)}

	location, err := SaveTokenForMethod(token, authMethod)
	if err == nil {
		// ok
	} else {
		t.Fatalf("expected no error, got %v", err)
	}

	mockStorage.AssertExpectations(t)

	if !location.Keychain {
		t.Fatalf("expected Keychain=true")
	}
	if location.File {
		t.Fatalf("expected File=false")
	}

	// File should not be written when keychain save succeeds
	filePath := filepath.Join(tmp, ".pingcli", "credentials", authMethod+".json")
	if _, statErr := os.Stat(filePath); statErr == nil {
		t.Fatalf("expected no credentials file at %s when keychain save succeeds", filePath)
	}
}

func TestClearAllTokens_UsesStorageClearAll(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Ensure keychain is the selected storage mode
	t.Setenv("PINGCLI_LOGIN_STORAGE_TYPE", "secure_local")

	old := newKeychainStorage
	t.Cleanup(func() { newKeychainStorage = old })

	mockStorage := new(MockTokenStorage)
	mockStorage.On("ClearAllTokens").Return(nil)

	newKeychainStorage = func(serviceName, username string) (tokenStorage, error) {
		return mockStorage, nil
	}

	err := ClearAllTokens()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	mockStorage.AssertExpectations(t)
}
