// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigUnset(koanfKey string) (err error) {
	if err = configuration.ValidateKoanfKey(koanfKey); err != nil {
		return fmt.Errorf("failed to unset configuration: %w", err)
	}

	pName, err := readConfigUnsetOptions()
	if err != nil {
		return fmt.Errorf("failed to unset configuration: %w", err)
	}

	subKoanf, err := profiles.GetKoanfConfig().GetProfileKoanf(pName)
	if err != nil {
		return fmt.Errorf("failed to unset configuration: %w", err)
	}

	opt, err := configuration.OptionFromKoanfKey(koanfKey)
	if err != nil {
		return fmt.Errorf("failed to unset configuration: %w", err)
	}

	err = subKoanf.Set(opt.KoanfKey, opt.DefaultValue)
	if err != nil {
		return fmt.Errorf("failed to unset configuration: %w", err)
	}

	if err = profiles.GetKoanfConfig().SaveProfile(pName, subKoanf); err != nil {
		return fmt.Errorf("failed to unset configuration: %w", err)
	}

	msgStr := "Configuration unset successfully:\n"

	vVal, _, err := profiles.KoanfValueFromOption(opt, pName)
	if err != nil {
		return fmt.Errorf("failed to unset configuration: %w", err)
	}

	unmaskOptionVal, err := profiles.GetOptionValue(options.ConfigUnmaskSecretValueOption)
	if err != nil {
		unmaskOptionVal = "false"
	}

	if opt.Sensitive && strings.EqualFold(unmaskOptionVal, "false") {
		msgStr += fmt.Sprintf("%s=%s", opt.KoanfKey, profiles.MaskValue(vVal))
	} else {
		msgStr += fmt.Sprintf("%s=%s", opt.KoanfKey, vVal)
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
