// Copyright © 2025 Ping Identity Corporation

package auth

import "errors"

var (
	// Configuration and validation errors
	ErrClientIDRequired      = errors.New("client ID is required")
	ErrClientSecretRequired  = errors.New("client secret is required")
	ErrEnvironmentIDRequired = errors.New("environment ID is required")
	ErrInvalidAuthType       = errors.New("invalid authentication type")
	ErrAuthConfigRequired    = errors.New("authentication configuration required. Please configure authentication using 'pingcli auth login' or 'pingcli config set'")
	ErrNoAuthTypeSpecified   = errors.New("no authentication type configured and no flag specified. Use --auth-code, --device-code, or --client-credentials to specify which credentials to clear")
	ErrNoAuthConfiguration   = errors.New("no configuration found. Nothing to logout from. Run 'pingcli login' to configure authentication")
)
