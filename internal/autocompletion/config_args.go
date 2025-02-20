package autocompletion

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/spf13/cobra"
)

func ConfigViewProfileFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return profiles.GetMainConfig().ProfileNames(), cobra.ShellCompDirectiveNoFileComp
}

func ConfigReturnNonActiveProfilesFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	profileNames := profiles.GetMainConfig().ProfileNames()
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	activeProfileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		output.SystemError(fmt.Sprintf("Unable to get active profile: %v", err), nil)
	}

	nonActiveProfiles := []string{}
	for _, p := range profileNames {
		if p != activeProfileName {
			nonActiveProfiles = append(nonActiveProfiles, p)
		}
	}
	return nonActiveProfiles, cobra.ShellCompDirectiveNoFileComp
}
