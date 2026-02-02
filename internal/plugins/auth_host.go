// Copyright © 2025 Ping Identity Corporation

package plugins

import (
	"context"
	"fmt"

	auth "github.com/pingidentity/pingcli/internal/commands/auth"
	"github.com/pingidentity/pingcli/shared/grpc"
	"golang.org/x/oauth2"
)

var _ grpc.Authentication = (*PluginAuthenticator)(nil)

type PluginAuthenticator struct {
	getValidTokenSource func(context.Context) (oauth2.TokenSource, error)
}

func (a *PluginAuthenticator) GetToken() (string, error) {
	// Use the central auth manager which handles both Keychain and File Storage
	// It respects the active profile and auth method (Worker, Device Code, etc.)
	getter := a.getValidTokenSource
	if getter == nil {
		getter = auth.GetValidTokenSource
	}

	tokenSource, err := getter(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get authentication token source for plugin: %w", err)
	}

	token, err := tokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("failed to get authentication token for plugin: %w", err)
	}

	if token == nil || token.AccessToken == "" {
		return "", fmt.Errorf("no valid token found. Please run 'pingcli login'")
	}

	return token.AccessToken, nil
}
