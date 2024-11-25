package config_internal

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigUnset(viperKey string) (err error) {
	if err = configuration.ValidateViperKey(viperKey); err != nil {
		return fmt.Errorf("failed to unset configuration: %v", err)
	}

	pName, err := readConfigUnsetOptions()
	if err != nil {
		return fmt.Errorf("failed to unset configuration: %v", err)
	}

	subViper, err := profiles.GetMainConfig().GetProfileViper(pName)
	if err != nil {
		return fmt.Errorf("failed to unset configuration: %v", err)
	}

	opt, err := configuration.OptionFromViperKey(viperKey)
	if err != nil {
		return fmt.Errorf("failed to unset configuration: %v", err)
	}

	subViper.Set(viperKey, opt.DefaultValue)

	if err = profiles.GetMainConfig().SaveProfile(pName, subViper); err != nil {
		return fmt.Errorf("failed to unset configuration: %v", err)
	}

	yamlStr, err := profiles.GetMainConfig().ProfileToString(pName)
	if err != nil {
		return fmt.Errorf("failed to unset configuration: %v", err)
	}

	output.Success("Configuration unset successfully", map[string]interface{}{"Profile YAML": yamlStr})

	return nil
}

func readConfigUnsetOptions() (pName string, err error) {
	if !options.RootProfileOption.Flag.Changed {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
	} else {
		pName, err = profiles.GetOptionValue(options.RootProfileOption)
	}

	if err != nil {
		return pName, err
	}

	if pName == "" {
		return pName, fmt.Errorf("unable to determine profile to unset configuration from")
	}

	return pName, nil
}
