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

	// Device code errors
	ErrDeviceCodeClientIDNotConfigured      = errors.New("device code client ID is not configured. Please run 'pingcli config set service.pingone.authentication.deviceCode.clientID=<your-client-id>'")
	ErrDeviceCodeEnvironmentIDNotConfigured = errors.New("device code environment ID is not configured. Please run 'pingcli config set service.pingone.authentication.deviceCode.environmentID=<your-env-id>'")

	// Auth code errors
	ErrAuthCodeClientIDNotConfigured      = errors.New("auth code client ID is not configured. Please run 'pingcli config set service.pingone.authentication.authCode.clientID=<your-client-id>'")
	ErrAuthCodeEnvironmentIDNotConfigured = errors.New("auth code environment ID is not configured. Please run 'pingcli config set service.pingone.authentication.authCode.environmentID=<your-env-id>'")
	ErrAuthCodeRedirectURINotConfigured   = errors.New("auth code redirect URI is not configured. Please run 'pingcli config set service.pingone.authentication.authCode.redirectURI=<your-redirect-uri>'")

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
)
