// Copyright © 2025 Ping Identity Corporation

package plugins

import (
	"fmt"

	auth "github.com/pingidentity/pingcli/internal/commands/auth"
	"github.com/pingidentity/pingcli/shared/grpc"
)

var _ grpc.Authentication = (*PluginAuthenticator)(nil)

type PluginAuthenticator struct{}

func (a *PluginAuthenticator) GetToken() (string, error) {
	// Use the central auth manager which handles both Keychain and File Storage
	// It respects the active profile and auth method (Worker, Device Code, etc.)
	token, err := auth.LoadToken()
	if err != nil {
		return "", fmt.Errorf("failed to load authentication token for plugin: %w", err)
	}

	if token == nil || token.AccessToken == "" {
		return "", fmt.Errorf("no valid token found. Please run 'pingcli login'")
	}

	return token.AccessToken, nil
}
