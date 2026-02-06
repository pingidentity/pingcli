// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"golang.org/x/oauth2"
)

var (
	errNotImplemented      = errors.New("not implemented")
	errKeychainUnavailable = errors.New("keychain unavailable")
)

type mockTokenStorage struct {
	saveErr error
}

func (m *mockTokenStorage) SaveToken(token *oauth2.Token) error { return m.saveErr }
func (m *mockTokenStorage) LoadToken() (*oauth2.Token, error) {
	return nil, errNotImplemented
}
func (m *mockTokenStorage) ClearToken() error     { return nil }
func (m *mockTokenStorage) ClearAllTokens() error { return nil }

type funcTokenStorage struct {
	saveFn     func(*oauth2.Token) error
	loadFn     func() (*oauth2.Token, error)
	clearFn    func() error
	clearAllFn func() error
}

func (s *funcTokenStorage) SaveToken(token *oauth2.Token) error {
	if s.saveFn == nil {
		return nil
	}

	return s.saveFn(token)
}

func (s *funcTokenStorage) LoadToken() (*oauth2.Token, error) {
	if s.loadFn == nil {
		return nil, errNotImplemented
	}

	return s.loadFn()
}

func (s *funcTokenStorage) ClearToken() error {
	if s.clearFn == nil {
		return nil
	}

	return s.clearFn()
}

func (s *funcTokenStorage) ClearAllTokens() error {
	if s.clearAllFn == nil {
		return nil
	}

	return s.clearAllFn()
}

func TestSaveTokenForMethod_FallsBackToFileWhenKeychainSaveFails(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	tmp := t.TempDir()
	t.Setenv("HOME", tmp)
	// Ensure keychain is the selected storage mode (default is secure_local)
	t.Setenv("PINGCLI_LOGIN_STORAGE_TYPE", "secure_local")

	old := newKeychainStorage
	t.Cleanup(func() { newKeychainStorage = old })
	newKeychainStorage = func(serviceName, username string) (tokenStorage, error) {
		return &mockTokenStorage{saveErr: errKeychainUnavailable}, nil
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
}

func TestSaveTokenForMethod_UsesKeychainWhenAvailable(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	tmp := t.TempDir()
	t.Setenv("HOME", tmp)
	t.Setenv("PINGCLI_LOGIN_STORAGE_TYPE", "secure_local")

	old := newKeychainStorage
	t.Cleanup(func() { newKeychainStorage = old })

	sawSave := false
	newKeychainStorage = func(serviceName, username string) (tokenStorage, error) {
		return &funcTokenStorage{
			saveFn: func(*oauth2.Token) error {
				sawSave = true

				return nil
			},
		}, nil
	}

	authMethod := "test-auth-method-keychain"
	token := &oauth2.Token{AccessToken: "access", TokenType: "Bearer", Expiry: time.Now().Add(1 * time.Hour)}

	location, err := SaveTokenForMethod(token, authMethod)
	if err == nil {
		// ok
	} else {
		t.Fatalf("expected no error, got %v", err)
	}
	if !sawSave {
		t.Fatalf("expected keychain SaveToken to be attempted")
	}
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

func TestClearToken_UsesClearAllTokens(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Ensure keychain is the selected storage mode
	t.Setenv("PINGCLI_LOGIN_STORAGE_TYPE", "secure_local")

	old := newKeychainStorage
	t.Cleanup(func() { newKeychainStorage = old })

	sawClearAll := false
	newKeychainStorage = func(serviceName, username string) (tokenStorage, error) {
		return &funcTokenStorage{
			clearAllFn: func() error {
				sawClearAll = true

				return nil
			},
		}, nil
	}

	// Calling ClearToken should trigger ClearAllTokens on the storage
	// because ClearToken() in credentials.go calls newKeychainStorage("pingcli", "clearAllTokens")
	// and then invokes ClearAllTokens()
	err := ClearToken()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !sawClearAll {
		t.Error("Expected ClearAllTokens to be called on storage backend")
	}
}
