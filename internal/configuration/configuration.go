package configuration

import (
	"fmt"
	"slices"
	"strings"

	configuration_config "github.com/pingidentity/pingcli/internal/configuration/config"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	configuration_platform "github.com/pingidentity/pingcli/internal/configuration/platform"
	configuration_profiles "github.com/pingidentity/pingcli/internal/configuration/profiles"
	configuration_request "github.com/pingidentity/pingcli/internal/configuration/request"
	configuration_root "github.com/pingidentity/pingcli/internal/configuration/root"
	configuration_services "github.com/pingidentity/pingcli/internal/configuration/services"
)

func ViperKeys() (keys []string) {
	for _, opt := range options.Options() {
		if opt.ViperKey != "" {
			keys = append(keys, opt.ViperKey)
		}
	}

	slices.Sort(keys)
	return keys
}

func ValidateViperKey(viperKey string) error {
	validKeys := ViperKeys()
	for _, vKey := range validKeys {
		if strings.EqualFold(vKey, viperKey) {
			return nil
		}
	}

	return fmt.Errorf("key '%s' is not recognized as a valid configuration key.\nUse 'pingcli config list-keys' to view all available keys", viperKey)
}

// Return a list of all viper keys from Options
// Including all substrings of parent keys.
// For example, the option key export.environmentID adds the keys
// 'export' and 'export.environmentID' to the list.
func ExpandedViperKeys() (keys []string) {
	leafKeys := ViperKeys()
	for _, key := range leafKeys {
		keySplit := strings.Split(key, ".")
		for i := 0; i < len(keySplit); i++ {
			curKey := strings.Join(keySplit[:i+1], ".")
			if !slices.ContainsFunc(keys, func(v string) bool {
				return strings.EqualFold(v, curKey)
			}) {
				keys = append(keys, curKey)
			}
		}
	}

	slices.Sort(keys)
	return keys
}

func ValidateParentViperKey(viperKey string) error {
	validKeys := ExpandedViperKeys()
	for _, vKey := range validKeys {
		if strings.EqualFold(vKey, viperKey) {
			return nil
		}
	}

	return fmt.Errorf("key '%s' is not recognized as a valid configuration key.\nUse 'pingcli config list-keys' to view all available keys", viperKey)
}

func OptionFromViperKey(viperKey string) (opt options.Option, err error) {
	for _, opt := range options.Options() {
		if strings.EqualFold(opt.ViperKey, viperKey) {
			return opt, nil
		}
	}
	return opt, fmt.Errorf("failed to get option: no option found for viper key: %s", viperKey)
}

func InitAllOptions() {
	configuration_config.InitConfigAddProfileOptions()
	configuration_config.InitConfigDeleteProfileOptions()
	configuration_config.InitConfigListKeyOptions()

	configuration_platform.InitPlatformExportOptions()

	configuration_profiles.InitProfilesOptions()

	configuration_root.InitRootOptions()

	configuration_request.InitRequestOptions()

	configuration_services.InitPingFederateServiceOptions()
	configuration_services.InitPingOneServiceOptions()
}
