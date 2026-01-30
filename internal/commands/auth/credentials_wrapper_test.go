// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"golang.org/x/oauth2"
)

// mockTokenSource implements oauth2.TokenSource
type mockTokenSource struct {
	token *oauth2.Token
	err   error
}

func (m *mockTokenSource) Token() (*oauth2.Token, error) {
	return m.token, m.err
}

func TestFilePersistingTokenSource_Token(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping in CI environment due to filesystem writes")
	}

	testutils_koanf.InitKoanfs(t)

	// Explicitly set storage to file_system so SaveTokenForMethod writes to disk
	// Using environment variable as it takes precedence over Koanf defaults
	t.Setenv("PINGCLI_LOGIN_STORAGE_TYPE", "file_system")

	// Setup parameters
	authMethod := "test-persist-token-wrapper"
	expectedToken := &oauth2.Token{
		AccessToken:  "wrapper-access-token",
		TokenType:    "Bearer",
		RefreshToken: "wrapper-refresh-token",
		Expiry:       time.Now().Add(1 * time.Hour),
	}

	// Create wrapper with mock source
	mockSource := &mockTokenSource{
		token: expectedToken,
		err:   nil,
	}

	wrapper := &filePersistingTokenSource{
		source:     mockSource,
		authMethod: authMethod,
	}

	// Cleanup any potential existing file
	t.Cleanup(func() {
		_ = clearTokenFromFile(authMethod)
	})
	_ = clearTokenFromFile(authMethod)

	// Execute: Call Token() on the wrapper
	resultToken, err := wrapper.Token()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Verify 1: Token passed through correctly
	if resultToken.AccessToken != expectedToken.AccessToken {
		t.Errorf("Expected access token %s, got %s", expectedToken.AccessToken, resultToken.AccessToken)
	}

	// Verify 2: Token persisted to file
	loadedToken, err := loadTokenFromFile(authMethod)
	if err != nil {
		t.Fatalf("Failed to load token from file (persistence failed): %v", err)
	}

	if loadedToken.AccessToken != expectedToken.AccessToken {
		t.Errorf("Persisted token mismatch. Got %s, want %s", loadedToken.AccessToken, expectedToken.AccessToken)
	}
}

func TestFilePersistingTokenSource_ErrorPropagation(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErr := os.ErrPermission // Just a sample error
	mockSource := &mockTokenSource{
		token: nil,
		err:   expectedErr,
	}

	wrapper := &filePersistingTokenSource{
		source:     mockSource,
		authMethod: "test-method",
	}

	_, err := wrapper.Token()

	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected error %v, got %v", expectedErr, err)
	}
}
