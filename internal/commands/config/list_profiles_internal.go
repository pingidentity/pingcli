// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"strings"

	"github.com/fatih/color"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigListProfiles() (err error) {
	l := logger.Get()

	profileNames := profiles.GetMainConfig().ProfileNames()
	activeProfileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		return err
	}

	listStr := "Profiles:\n"

	// We need to enable/disable colorize before applying the color to the string below.
	output.SetColorize()
	activeFmt := color.New(color.Bold, color.FgGreen).SprintFunc()

	for _, profileName := range profileNames {
		if strings.EqualFold(profileName, activeProfileName) {
			listStr += "- " + profileName + activeFmt(" (active)") + " \n"
		} else {
			listStr += "- " + profileName + "\n"
		}

		description, err := profiles.GetMainConfig().ProfileViperValue(profileName, "description")
		if err != nil {
			l.Warn().Msgf("Cannot retrieve profile description for profile %s: %v", profileName, err)
			continue
		}

		if description != "" {
			listStr += "    " + description
		}
	}

	output.Message(listStr, nil)

	return nil
}
