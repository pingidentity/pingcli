// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

const (
	ENUM_AUTH_PROVIDER_PINGONE string = "pingone"
)

var (
	authProviderErrorPrefix = "custom type auth provider error"
)

// AuthProvider represents a single supported authentication provider name (pingone)
type AuthProvider string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*AuthProvider)(nil)

// Set parses and sets a single authentication provider
func (ap *AuthProvider) Set(providerStr string) error {
	if ap == nil {
		return &errs.PingCLIError{Prefix: authProviderErrorPrefix, Err: ErrCustomTypeNil}
	}

	if providerStr == "" {
		return nil
	}

	// Create a map of valid provider values to check against user-provided provider
	validProviderMap := AuthProviderValidValuesMap()

	provider := strings.ToLower(strings.TrimSpace(providerStr))

	enumProvider, ok := validProviderMap[provider]
	if !ok {
		return &errs.PingCLIError{Prefix: authProviderErrorPrefix, Err: fmt.Errorf("%w '%s': must be %s", ErrUnrecognizedAuthProvider, provider, strings.Join(AuthProviderValidValues(), ", "))}
	}

	*ap = AuthProvider(enumProvider)

	return nil
}

// String returns the authentication provider as a string (implements pflag.Value)
func (ap *AuthProvider) String() string {
	if ap == nil {
		return ""
	}

	return string(*ap)
}

// Type returns the type string for this custom type (implements pflag.Value)
func (ap *AuthProvider) Type() string {
	return "string"
}

// ContainsPingOne checks if the PingOne provider is set
func (ap *AuthProvider) ContainsPingOne() bool {
	if ap == nil || len(*ap) == 0 {
		return false
	}

	return strings.EqualFold(string(*ap), ENUM_AUTH_PROVIDER_PINGONE)
}

// AuthProviderValidValues returns a sorted list of all valid authentication provider values
func AuthProviderValidValues() []string {
	allProvider := []string{
		ENUM_AUTH_PROVIDER_PINGONE,
	}

	slices.Sort(allProvider)

	return allProvider
}

// AuthProviderValidValuesMap returns a map of valid auth provider values with lowercase keys
func AuthProviderValidValuesMap() map[string]string {
	validProvider := AuthProviderValidValues()
	validProviderMap := make(map[string]string, len(validProvider))
	for _, s := range validProvider {
		validProviderMap[strings.ToLower(s)] = s
	}

	return validProviderMap
}
