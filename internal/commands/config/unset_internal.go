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

	if err = profiles.GetMainConfig().ValidateExistingProfileName(pName); err != nil {
		return fmt.Errorf("failed to unset configuration: %v", err)
	}

	subViper := profiles.GetMainConfig().ViperInstance().Sub(pName)

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

	output.Print(output.Opts{
		Message: "Configuration unset successfully",
		Result:  output.ENUM_RESULT_SUCCESS,
	})

	output.Print(output.Opts{
		Message: yamlStr,
		Result:  output.ENUM_RESULT_NIL,
	})

	return nil
}

func readConfigUnsetOptions() (pName string, err error) {
	if !options.ConfigUnsetProfileOption.Flag.Changed {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
	} else {
		pName, err = profiles.GetOptionValue(options.ConfigUnsetProfileOption)
	}

	if err != nil {
		return pName, err
	}

	if pName == "" {
		return pName, fmt.Errorf("unable to determine profile to unset configuration from")
	}

	return pName, nil
}
