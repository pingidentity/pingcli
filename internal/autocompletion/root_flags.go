// Copyright © 2026 Ping Identity Corporation

package autocompletion

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/spf13/cobra"
)

func RootProfileFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		wrappedErr := fmt.Errorf("%w: %w", ErrGetConfiguration, err)
		output.SystemError((&errs.PingCLIError{Prefix: autocompletionErrorPrefix, Err: wrappedErr}).Error(), nil)
	}

	return koanfConfig.ProfileNames(), cobra.ShellCompDirectiveNoFileComp
}

func RootOutputFormatFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return customtypes.OutputFormatValidValues(), cobra.ShellCompDirectiveNoFileComp
}
