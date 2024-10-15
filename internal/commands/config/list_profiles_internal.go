package config_internal

import (
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigListProfiles() {
	profileNames := profiles.GetMainConfig().ProfileNames()
	activeProfile := profiles.GetMainConfig().ActiveProfile().Name()

	listStr := "Profiles:\n"
	for _, profileName := range profileNames {
		if profileName == activeProfile {
			listStr += "- " + profileName + " (active)\n"
		} else {
			listStr += "- " + profileName + "\n"
		}
	}

	output.Print(output.Opts{
		Message: listStr,
		Result:  output.ENUM_RESULT_NIL,
	})
}
