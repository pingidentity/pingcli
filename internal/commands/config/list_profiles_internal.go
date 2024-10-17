package config_internal

import (
	"slices"

	"github.com/fatih/color"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigListProfiles() {
	profileNames := profiles.GetMainConfig().ProfileNames()
	activeProfile := profiles.GetMainConfig().ActiveProfile().Name()

	listStr := "Profiles:\n"

	slices.Sort(profileNames)

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
			continue
		}

		listStr += "    " + description + "\n"
	}

	output.Print(output.Opts{
		Message: listStr,
		Result:  output.ENUM_RESULT_NIL,
	})
}
