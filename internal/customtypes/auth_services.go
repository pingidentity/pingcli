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
	ENUM_AUTH_SERVICE_PINGONE      string = "pingone"
	ENUM_AUTH_SERVICE_PINGFEDERATE string = "pingfederate"
)

var (
	authServicesErrorPrefix = "custom type auth services error"
)

// AuthServices represents a list of authentication service names (pingone, pingfederate)
type AuthServices []string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*AuthServices)(nil)

// GetServices returns the list of authentication services
func (as *AuthServices) GetServices() []string {
	if as == nil {
		return []string{}
	}

	return []string(*as)
}

// Set parses and adds authentication services from a comma-separated string
func (as *AuthServices) Set(servicesStr string) error {
	if as == nil {
		return &errs.PingCLIError{Prefix: authServicesErrorPrefix, Err: ErrCustomTypeNil}
	}

	if servicesStr == "" || servicesStr == "[]" {
		return nil
	}

	// Create a map of valid service values to check against user-provided services
	validServiceMap := AuthServicesValidValuesMap()

	// Create a map of existing services set in the AuthServices object
	existingServices := make(map[string]struct{}, len(*as))
	for _, s := range *as {
		existingServices[s] = struct{}{}
	}

	// Loop through user-provided services
	// check for valid value in map
	// check the service does not already exist
	for service := range strings.SplitSeq(servicesStr, ",") {
		service = strings.ToLower(strings.TrimSpace(service))

		enumService, ok := validServiceMap[service]
		if !ok {
			return &errs.PingCLIError{Prefix: authServicesErrorPrefix, Err: fmt.Errorf("%w '%s': must be one of %s", ErrUnrecognizedAuthService, service, strings.Join(AuthServicesValidValues(), ", "))}
		}

		if _, ok := existingServices[enumService]; ok {
			continue
		}

		*as = append(*as, enumService)
	}

	slices.Sort(*as)

	return nil
}

// ContainsPingOne checks if the PingOne service is in the list
func (as *AuthServices) ContainsPingOne() bool {
	if as == nil || len(*as) == 0 {
		return false
	}

	return slices.ContainsFunc(*as, func(s string) bool {
		return strings.EqualFold(s, ENUM_AUTH_SERVICE_PINGONE)
	})
}

// ContainsPingFederate checks if the PingFederate service is in the list
func (as *AuthServices) ContainsPingFederate() bool {
	if as == nil || len(*as) == 0 {
		return false
	}

	return slices.ContainsFunc(*as, func(s string) bool {
		return strings.EqualFold(s, ENUM_AUTH_SERVICE_PINGFEDERATE)
	})
}

// Type returns the type string for this custom type (implements pflag.Value)
func (as *AuthServices) Type() string {
	return "[]string"
}

// String returns a comma-separated string of the authentication services (implements pflag.Value)
func (as *AuthServices) String() string {
	if as == nil {
		return ""
	}

	slices.Sort(*as)

	return strings.Join(*as, ",")
}

// AuthServicesValidValues returns a sorted list of all valid authentication service values
func AuthServicesValidValues() []string {
	allServices := []string{
		ENUM_AUTH_SERVICE_PINGFEDERATE,
		ENUM_AUTH_SERVICE_PINGONE,
	}

	slices.Sort(allServices)

	return allServices
}

// AuthServicesValidValuesMap returns a map of valid auth service values with lowercase keys
func AuthServicesValidValuesMap() map[string]string {
	validServices := AuthServicesValidValues()
	validServiceMap := make(map[string]string, len(validServices))
	for _, s := range validServices {
		validServiceMap[strings.ToLower(s)] = s
	}

	return validServiceMap
}
