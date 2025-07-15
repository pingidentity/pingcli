// Copyright © 2025 Ping Identity Corporation

package configuration

import (
	"fmt"
	"slices"
	"strings"

	configuration_config "github.com/pingidentity/pingcli/internal/configuration/config"
	configuration_license "github.com/pingidentity/pingcli/internal/configuration/license"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	configuration_platform "github.com/pingidentity/pingcli/internal/configuration/platform"
	configuration_plugin "github.com/pingidentity/pingcli/internal/configuration/plugin"
	configuration_profiles "github.com/pingidentity/pingcli/internal/configuration/profiles"
	configuration_request "github.com/pingidentity/pingcli/internal/configuration/request"
	configuration_root "github.com/pingidentity/pingcli/internal/configuration/root"
	configuration_services "github.com/pingidentity/pingcli/internal/configuration/services"
)

func KoanfKeys() (keys []string) {
	for _, opt := range options.Options() {
		if opt.KoanfKey != "" {
			keys = append(keys, opt.KoanfKey)
		}
	}

	slices.Sort(keys)

	return keys
}

func ValidateKoanfKey(koanfKey string) error {
	validKeys := KoanfKeys()
	for _, vKey := range validKeys {
		if vKey == koanfKey {
			return nil
		}
	}

	return fmt.Errorf("key '%s' is not recognized as a valid configuration key.\nUse 'pingcli config list-keys' to view all available keys", koanfKey)
}

// Return a list of all koanf keys from Options
// Including all substrings of parent keys.
// For example, the option key export.environmentID adds the keys
// 'export' and 'export.environmentID' to the list.
func ExpandedKoanfKeys() (keys []string) {
	leafKeys := KoanfKeys()
	for _, key := range leafKeys {
		keySplit := strings.Split(key, ".")
		for i := range keySplit {
			curKey := strings.Join(keySplit[:i+1], ".")
			if !slices.ContainsFunc(keys, func(v string) bool {
				return v == curKey
			}) {
				keys = append(keys, curKey)
			}
		}
	}

	slices.Sort(keys)

	return keys
}

func ValidateParentKoanfKey(koanfKey string) error {
	validKeys := ExpandedKoanfKeys()
	for _, vKey := range validKeys {
		if vKey == koanfKey {
			return nil
		}
	}

	return fmt.Errorf("key '%s' is not recognized as a valid configuration key.\nUse 'pingcli config list-keys' to view all available keys", koanfKey)
}

func OptionFromKoanfKey(koanfKey string) (opt options.Option, err error) {
	for _, opt := range options.Options() {
		if opt.KoanfKey == koanfKey {
			return opt, nil
		}
	}

	return opt, fmt.Errorf("failed to get option: no option found for koanf key: %s", koanfKey)
}

func InitAllOptions() {
	configuration_config.InitConfigAddProfileOptions()
	configuration_config.InitConfigDeleteProfileOptions()
	configuration_config.InitConfigListKeyOptions()

	configuration_platform.InitPlatformExportOptions()

	configuration_plugin.InitPluginOptions()

	configuration_profiles.InitProfilesOptions()

	configuration_root.InitRootOptions()

	configuration_request.InitRequestOptions()

	configuration_services.InitPingFederateServiceOptions()
	configuration_services.InitPingOneServiceOptions()

	configuration_license.InitLicenseOptions()
}
