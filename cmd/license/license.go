// Copyright © 2025 Ping Identity Corporation

package license

import (
	"fmt"

	"github.com/pingidentity/pingcli/cmd/common"
	license_internal "github.com/pingidentity/pingcli/internal/commands/license"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/spf13/cobra"
)

const (
	licenseCommandExamples = `  Request a new evaluation license for PingFederate 12.0.
    pingcli license request --product pingfederate --version 12.0
	
	  Request a new evaluation license for PingAccess 6.3.
	pingcli license request --product pingaccess --version 6.3`
)

func NewLicenseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               licenseCommandExamples,
		Long: `Request a new evaluation license for a specific product and version.

The new license request will be sent to the Ping Identity license server.`,
		RunE:  licenseRunE,
		Short: "Request a new evaluation license.",
		Use:   "license [flags]",
	}

	cmd.Flags().AddFlag(options.LicenseProductOption.Flag)
	cmd.Flags().AddFlag(options.LicenseVersionOption.Flag)
	cmd.Flags().AddFlag(options.LicenseDevopsUserOption.Flag)
	cmd.Flags().AddFlag(options.LicenseDevopsKeyOption.Flag)

	err := cmd.MarkFlagRequired(options.LicenseProductOption.CobraParamName)
	if err != nil {
		output.SystemError(fmt.Sprintf("Failed to mark flag '%s' as required: %v", options.LicenseProductOption.CobraParamName, err), nil)
	}
	err = cmd.MarkFlagRequired(options.LicenseVersionOption.CobraParamName)
	if err != nil {
		output.SystemError(fmt.Sprintf("Failed to mark flag '%s' as required: %v", options.LicenseVersionOption.CobraParamName, err), nil)
	}

	return cmd
}

func licenseRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("License Subcommand Called.")

	if err := license_internal.RunInternalLicense(); err != nil {
		return &errs.PingCLIError{Prefix: "", Err: err}
	}

	return nil
}
