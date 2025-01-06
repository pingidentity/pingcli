package config_internal

import (
	"fmt"
	"strings"

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

	msgStr := fmt.Sprintf("Configuration values for profile '%s' and key '%s':\n", pName, viperKey)

	for _, opt := range options.Options() {
		if opt.ViperKey == "" || !strings.Contains(opt.ViperKey, viperKey) {
			continue
		}

		if opt.Sensitive {
			optVal, err := profiles.GetSensitiveOptionValue(opt, true)
			if err != nil {
				return fmt.Errorf("failed to get configuration: %v", err)
			}

			msgStr += fmt.Sprintf("%s=%s\n", opt.ViperKey, optVal)
		} else {
			optVal, err := profiles.GetOptionValue(opt)
			if err != nil {
				return fmt.Errorf("failed to get configuration: %v", err)
			}

			msgStr += fmt.Sprintf("%s=%s\n", opt.ViperKey, optVal)
		}
	}

	output.Message(msgStr, nil)

	return nil
}

func readConfigGetOptions() (pName string, err error) {
	if !options.RootProfileOption.Flag.Changed {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
	} else {
		pName, err = profiles.GetOptionValue(options.RootProfileOption)
	}

	if err != nil {
		return "", err
	}

	if pName == "" {
		return "", fmt.Errorf("unable to determine profile to get configuration from")
	}

	return pName, nil
}
