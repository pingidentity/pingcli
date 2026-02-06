// Copyright © 2026 Ping Identity Corporation

package config_internal

import (
	"fmt"
	"io"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var (
	deleteProfileErrorPrefix = "failed to delete profile"
)

func RunInternalConfigDeleteProfile(args []string, rc io.ReadCloser) (err error) {
	var pName string
	if len(args) == 1 {
		pName = args[0]
	} else {
		pName, err = promptUserToDeleteProfile(rc)
		if err != nil {
			return &errs.PingCLIError{Prefix: deleteProfileErrorPrefix, Err: err}
		}
	}

	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &errs.PingCLIError{Prefix: deleteProfileErrorPrefix, Err: err}
	}

	if err = koanfConfig.ValidateExistingProfileName(pName); err != nil {
		return &errs.PingCLIError{Prefix: deleteProfileErrorPrefix, Err: err}
	}

	confirmed, err := promptUserToConfirmDelete(pName, rc)
	if err != nil {
		return &errs.PingCLIError{Prefix: deleteProfileErrorPrefix, Err: err}
	}

	if !confirmed {
		output.Message("Profile deletion cancelled.", nil)

		return nil
	}

	err = deleteProfile(pName)
	if err != nil {
		return &errs.PingCLIError{Prefix: deleteProfileErrorPrefix, Err: err}
	}

	return nil
}

func promptUserToDeleteProfile(rc io.ReadCloser) (pName string, err error) {
	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return pName, &errs.PingCLIError{Prefix: deleteProfileErrorPrefix, Err: err}
	}
	pName, err = input.RunPromptSelect("Select profile to delete", koanfConfig.ProfileNames(), rc)

	if err != nil {
		return pName, &errs.PingCLIError{Prefix: deleteProfileErrorPrefix, Err: err}
	}

	return pName, nil
}

func promptUserToConfirmDelete(pName string, rc io.ReadCloser) (confirmed bool, err error) {
	autoAccept := "false"
	if options.ConfigDeleteAutoAcceptOption.Flag.Changed {
		autoAccept, err = profiles.GetOptionValue(options.ConfigDeleteAutoAcceptOption)
		if err != nil {
			return false, &errs.PingCLIError{Prefix: deleteProfileErrorPrefix, Err: err}
		}
	}

	if autoAccept == "true" {
		return true, nil
	}

	confirmed, err = input.RunPromptConfirm(fmt.Sprintf("Are you sure you want to delete profile '%s'", pName), rc)
	if err != nil {
		return false, &errs.PingCLIError{Prefix: deleteProfileErrorPrefix, Err: err}
	}

	return confirmed, nil
}

func deleteProfile(pName string) (err error) {
	output.Message(fmt.Sprintf("Deleting profile '%s'...", pName), nil)

	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &errs.PingCLIError{Prefix: deleteProfileErrorPrefix, Err: err}
	}

	if err = koanfConfig.DeleteProfile(pName); err != nil {
		return &errs.PingCLIError{Prefix: deleteProfileErrorPrefix, Err: err}
	}

	output.Success(fmt.Sprintf("Profile '%s' deleted.", pName), nil)

	return nil
}
