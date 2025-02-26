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

	msgStr := "Configuration unset successfully:\n"

	vVal, _, err := profiles.ViperValueFromOption(opt)
	if err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	if opt.Sensitive {
		msgStr += fmt.Sprintf("%s=%s", viperKey, profiles.MaskValue(vVal))
	} else {
		msgStr += fmt.Sprintf("%s=%s", viperKey, vVal)
	}

	output.Success(msgStr, nil)

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
