// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"github.com/fatih/color"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigListProfiles() (err error) {
	profileNames := profiles.GetKoanfConfig().ProfileNames()
	activeProfileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		return err
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

		description := profiles.GetKoanfConfig().KoanfInstance().String(profileName + "." + "description")
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
