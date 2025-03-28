// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"fmt"
	"io"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigDeleteProfile(args []string, rc io.ReadCloser) (err error) {
	var pName string
	if len(args) == 1 {
		pName = args[0]
	} else {
		pName, err = promptUserToDeleteProfile(rc)
		if err != nil {
			return fmt.Errorf("failed to delete profile: %w", err)
		}
	}

	if err = profiles.GetMainConfig().ValidateExistingProfileName(pName); err != nil {
		return fmt.Errorf("failed to delete profile: %w", err)
	}

	confirmed, err := promptUserToConfirmDelete(pName, rc)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %w", err)
	}

	if !confirmed {
		output.Message("Profile deletion cancelled.", nil)

		return nil
	}

	err = deleteProfile(pName)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %w", err)
	}

	return nil
}

func promptUserToDeleteProfile(rc io.ReadCloser) (pName string, err error) {
	pName, err = input.RunPromptSelect("Select profile to delete: ", profiles.GetMainConfig().ProfileNames(), rc)

	if err != nil {
		return pName, err
	}

	return pName, nil
}

func promptUserToConfirmDelete(pName string, rc io.ReadCloser) (confirmed bool, err error) {
	autoAccept := "false"
	if options.ConfigDeleteAutoAcceptOption.Flag.Changed {
		autoAccept, err = profiles.GetOptionValue(options.ConfigDeleteAutoAcceptOption)
		if err != nil {
			return false, err
		}
	}

	if autoAccept == "true" {
		return true, nil
	}

	return input.RunPromptConfirm(fmt.Sprintf("Are you sure you want to delete profile '%s'?", pName), rc)
}

func deleteProfile(pName string) (err error) {
	output.Message(fmt.Sprintf("Deleting profile '%s'...", pName), nil)

	if err = profiles.GetMainConfig().DeleteProfile(pName); err != nil {
		return err
	}

	output.Success(fmt.Sprintf("Profile '%s' deleted.", pName), nil)

	return nil
}
