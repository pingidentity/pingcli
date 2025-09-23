// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"github.com/fatih/color"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var (
	listProfilesErrorPrefix = "failed to list profiles"
)

func RunInternalConfigListProfiles() (err error) {
	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &errs.PingCLIError{Prefix: listProfilesErrorPrefix, Err: err}
	}

	profileNames := koanfConfig.ProfileNames()
	activeProfileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: listProfilesErrorPrefix, Err: err}
	}

	listStr := "Profiles:\n"

	// We need to enable/disable colorize before applying the color to the string below.
	output.SetColorize()
	activeFmt := color.New(color.Bold, color.FgGreen).SprintFunc()

	for i, profileName := range profileNames {
		if profileName == activeProfileName {
			listStr += "- " + profileName + activeFmt(" (active)") + " \n"
		} else {
			listStr += "- " + profileName + "\n"
		}

		description := koanfConfig.KoanfInstance().String(profileName + "." + options.ProfileDescriptionOption.KoanfKey)
		if description != "" {
			listStr += "    " + description
		}

		if i < len(profileNames)-1 {
			listStr += "\n"
		}
	}

	output.Message(listStr, nil)

	return nil
}
