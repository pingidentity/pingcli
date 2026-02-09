// Copyright © 2026 Ping Identity Corporation

package autocompletion

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/spf13/cobra"
)

func ConfigViewProfileFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		wrappedErr := fmt.Errorf("%w: %w", ErrGetConfiguration, err)
		output.SystemError((&errs.PingCLIError{Prefix: autocompletionErrorPrefix, Err: wrappedErr}).Error(), nil)
	}

	return koanfConfig.ProfileNames(), cobra.ShellCompDirectiveNoFileComp
}

func ConfigReturnNonActiveProfilesFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		wrappedErr := fmt.Errorf("%w: %w", ErrGetConfiguration, err)
		output.SystemError((&errs.PingCLIError{Prefix: autocompletionErrorPrefix, Err: wrappedErr}).Error(), nil)
	}

	profileNames := koanfConfig.ProfileNames()
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	activeProfileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		wrappedErr := fmt.Errorf("%w: %w", ErrGetActiveProfile, err)
		output.SystemError((&errs.PingCLIError{Prefix: autocompletionErrorPrefix, Err: wrappedErr}).Error(), nil)
	}

	nonActiveProfiles := []string{}
	for _, p := range profileNames {
		if p != activeProfileName {
			nonActiveProfiles = append(nonActiveProfiles, p)
		}
	}

	return nonActiveProfiles, cobra.ShellCompDirectiveNoFileComp
}
