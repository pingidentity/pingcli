package config_internal

import (
	"github.com/fatih/color"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigListProfiles() {
	l := logger.Get()

	profileNames := profiles.GetMainConfig().ProfileNames()
	activeProfile := profiles.GetMainConfig().ActiveProfile().Name()

	listStr := "Profiles:\n"

	// We need to enable/disable colorize before applying the color to the string below.
	output.SetColorize()
	activeFmt := color.New(color.Bold, color.FgGreen).SprintFunc()

	for _, profileName := range profileNames {
		if profileName == activeProfile {
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
			listStr += "    " + description + "\n"
		}
	}

	output.Print(output.Opts{
		Message: listStr,
		Result:  output.ENUM_RESULT_NIL,
	})
}
