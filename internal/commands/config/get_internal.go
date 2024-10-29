package config_internal

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigGet(viperKey string) (err error) {
	if err = configuration.ValidateParentViperKey(viperKey); err != nil {
		return fmt.Errorf("failed to get configuration: %v", err)
	}

	pName, err := readConfigGetOptions()
	if err != nil {
		return fmt.Errorf("failed to get configuration: %v", err)
	}

	yamlStr, err := profiles.GetMainConfig().ProfileViperValue(pName, viperKey)
	if err != nil {
		return fmt.Errorf("failed to get configuration: %v", err)
	}

	output.Message(yamlStr, nil)

	return nil
}

func readConfigGetOptions() (pName string, err error) {
	if !options.ConfigGetProfileOption.Flag.Changed {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
	} else {
		pName, err = profiles.GetOptionValue(options.ConfigGetProfileOption)
	}

	if err != nil {
		return pName, err
	}

	if pName == "" {
		return pName, fmt.Errorf("unable to determine profile to get configuration from")
	}

	return pName, nil
}
