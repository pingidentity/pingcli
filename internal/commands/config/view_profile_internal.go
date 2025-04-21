// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigViewProfile(args []string) (err error) {
	var msgStr, pName string
	if len(args) == 1 {
		pName = args[0]
	} else {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
		if err != nil {
			return fmt.Errorf("failed to view profile: %w", err)
		}
	}

	// Validate the profile name
	err = profiles.GetKoanfConfig().ValidateExistingProfileName(pName)
	if err != nil {
		return fmt.Errorf("failed to view profile: %w", err)
	}

	// Get the Koanf configuration for the specified profile
	koanfProfile, err := profiles.GetKoanfConfig().GetProfileKoanf(pName)
	if err != nil {
		return fmt.Errorf("failed to get config from profile: %w", err)
	}

	// Iterate over the options in profile and print them
	for _, opt := range options.Options() {
		if !koanfProfile.Exists(opt.KoanfKey) {
			continue
		}

		vVal, ok, err := profiles.KoanfValueFromOption(opt, pName)
		if !ok {
			continue
		}

		if err != nil {
			return fmt.Errorf("failed to get koanf value from option: %w", err)
		}

		unmaskOptionVal, err := profiles.GetOptionValue(options.ConfigUnmaskSecretValueOption)
		if err != nil {
			unmaskOptionVal = "false"
		}

		if opt.Sensitive && strings.EqualFold(unmaskOptionVal, "false") {
			msgStr += fmt.Sprintf("%s=%s\n", opt.KoanfKey, profiles.MaskValue(vVal))
		} else {
			msgStr += fmt.Sprintf("%s=%s\n", opt.KoanfKey, vVal)
		}
	}

	output.Message(fmt.Sprintf("Configuration for profile '%s':\n", pName)+msgStr, nil)

	return nil
}
