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
	var pName string
	if len(args) == 1 {
		pName = args[0]
	} else {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
		if err != nil {
			return fmt.Errorf("failed to view profile: %w", err)
		}
	}

	// Validate the profile name
	err = profiles.GetMainConfig().ValidateExistingProfileName(pName)
	if err != nil {
		return fmt.Errorf("failed to view profile: %w", err)
	}

	msgStr := fmt.Sprintf("Configuration for profile '%s':\n", pName)

	for _, opt := range options.Options() {
		if opt.ViperKey == "" {
			continue
		}

		vVal, _, err := profiles.ViperValueFromOption(opt)
		if err != nil {
			return fmt.Errorf("failed to view profile: %w", err)
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
