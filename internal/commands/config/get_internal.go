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

func RunInternalConfigGet(viperKey string) (err error) {
	if err = configuration.ValidateParentViperKey(viperKey); err != nil {
		return fmt.Errorf("failed to get configuration: %w", err)
	}

	pName, err := readConfigGetOptions()
	if err != nil {
		return fmt.Errorf("failed to get configuration: %w", err)
	}

	msgStr := fmt.Sprintf("Configuration values for profile '%s' and key '%s':\n", strings.ToLower(pName), viperKey)

	for _, opt := range options.Options() {
		if opt.ViperKey == "" || !strings.Contains(opt.ViperKey, viperKey) {
			continue
		}

		vVal, _, err := profiles.ViperValueFromOption(opt)
		if err != nil {
			return fmt.Errorf("failed to get configuration: %w", err)
		}

		unmaskOptionVal, err := profiles.GetOptionValue(options.ConfigUnmaskSecretValueOption)
		if err != nil {
			unmaskOptionVal = "false"
		}

		if opt.Sensitive && strings.EqualFold(unmaskOptionVal, "false") {
			msgStr += fmt.Sprintf("%s=%s\n", opt.ViperKey, profiles.MaskValue(vVal))
		} else {
			msgStr += fmt.Sprintf("%s=%s\n", opt.ViperKey, vVal)
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
