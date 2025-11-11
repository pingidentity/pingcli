// Copyright © 2025 Ping Identity Corporation

package auth_internal

import "errors"

var (
	// Token errors
	ErrNoTokenFound                   = errors.New("no token found for any authentication method")
	ErrNoCachedToken                  = errors.New("no cached token available")
	ErrUnsupportedAuthType            = errors.New("unsupported authentication type. Please run 'pingcli login' to authenticate")
	ErrAuthMethodNotConfigured        = errors.New("auth method is not configured")
	ErrUnsupportedAuthMethod          = errors.New("unsupported auth method")
	ErrTokenKeyGenerationRequirements = errors.New("environment ID and client ID are required for token key generation")
	ErrGrantTypeNotSet                = errors.New("configuration does not have grant type set")
	ErrRegionCodeRequired             = errors.New("region code is required and must be valid. Please run 'pingcli config set service.pingone.regionCode=<region>'")
	ErrEnvironmentIDNotConfigured     = errors.New("environment ID is not configured. Please run 'pingcli config set service.pingone.authentication.environmentID=<your-env-id>'")

	// Device code errors
	ErrDeviceCodeClientIDNotConfigured      = errors.New("device code client ID is not configured. Please run 'pingcli config set service.pingone.authentication.deviceCode.clientID=<your-client-id>'")
	ErrDeviceCodeEnvironmentIDNotConfigured = errors.New("device code environment ID is not configured. Please run 'pingcli config set service.pingone.authentication.deviceCode.environmentID=<your-env-id>'")

	// Auth code errors
	ErrAuthorizationCodeClientIDNotConfigured        = errors.New("authorization code client ID is not configured. Please run 'pingcli config set service.pingone.authentication.authorizationCode.clientID=<your-client-id>'")
	ErrAuthorizationCodeEnvironmentIDNotConfigured   = errors.New("authorization code environment ID is not configured. Please run 'pingcli config set service.pingone.authentication.authorizationCode.environmentID=<your-env-id>'")
	ErrAuthorizationCodeRedirectURINotConfigured     = errors.New("authorization code redirect URI is not configured. Please run 'pingcli config set service.pingone.authentication.authorizationCode.redirectURI=<your-redirect-uri>'")
	ErrAuthorizationCodeRedirectURIPathNotConfigured = errors.New("authorization code redirect URI path is not configured. Please run 'pingcli config set service.pingone.authentication.authorizationCode.redirectURIPath=<path>'")
	ErrAuthorizationCodeRedirectURIPortNotConfigured = errors.New("authorization code redirect URI port is not configured. Please run 'pingcli config set service.pingone.authentication.authorizationCode.redirectURIPort=<port>'")

	// Client credentials errors
	ErrClientCredentialsClientIDNotConfigured      = errors.New("client credentials client ID is not configured. Please run 'pingcli config set service.pingone.authentication.clientCredentials.clientID=<your-client-id>'")
	ErrClientCredentialsClientSecretNotConfigured  = errors.New("client credentials client secret is not configured. Please run 'pingcli config set service.pingone.authentication.clientCredentials.clientSecret=<your-client-secret>'")
	ErrClientCredentialsEnvironmentIDNotConfigured = errors.New("client credentials environment ID is not configured. Please run 'pingcli config set service.pingone.authentication.clientCredentials.environmentID=<your-env-id>'")

	// PingFederate errors
	ErrPingFederateContextNil  = errors.New("failed to initialize PingFederate services. context is nil")
	ErrPingFederateCACertParse = errors.New("failed to parse CA certificate PEM file to certificate pool")

	// PingOne errors
	ErrPingOneUnrecognizedAuthType = errors.New("unrecognized or unsupported PingOne authentication type")
	ErrPingOneClientConfigNil      = errors.New("PingOne client configuration is nil")

	// Configuration and validation errors
	ErrClientIDRequired      = errors.New("client ID is required")
	ErrClientSecretRequired  = errors.New("client secret is required")
	ErrEnvironmentIDRequired = errors.New("environment ID is required")
	ErrInvalidAuthType       = errors.New("invalid authentication type")
	ErrAuthConfigRequired    = errors.New("authentication configuration required. Please configure authentication using 'pingcli auth login' or 'pingcli config set'")
	ErrNoAuthTypeSpecified   = errors.New("no authentication type configured and no flag specified. Use --auth-code, --device-code, or --client-credentials to specify which credentials to clear")
	ErrNoAuthConfiguration   = errors.New("no configuration found. Nothing to logout from. Run 'pingcli login' to configure authentication")

	// Redirect URI validation errors
	ErrRedirectURIPathInvalid = errors.New("redirect URI path must start with '/'")
	ErrPortInvalid            = errors.New("port must be a number")
	ErrPortOutOfRange         = errors.New("port must be between 1 and 65535")
)
