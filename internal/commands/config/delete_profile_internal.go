// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"errors"
	"fmt"
	"io"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

type DeleteProfileError struct {
	Err error
}

func (e *DeleteProfileError) Error() string {
	var err *DeleteProfileError
	if errors.As(e.Err, &err) {
		return err.Error()
	}
	return fmt.Sprintf("failed to delete profile: %s", e.Err.Error())
}

func (e *DeleteProfileError) Unwrap() error {
	var err *DeleteProfileError
	if errors.As(e.Err, &err) {
		return err.Unwrap()
	}
	return e.Err
}

func RunInternalConfigDeleteProfile(args []string, rc io.ReadCloser) (err error) {
	var pName string
	if len(args) == 1 {
		pName = args[0]
	} else {
		pName, err = promptUserToDeleteProfile(rc)
		if err != nil {
			return &DeleteProfileError{Err: err}
		}
	}

	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &DeleteProfileError{Err: err}
	}

	if err = koanfConfig.ValidateExistingProfileName(pName); err != nil {
		return &DeleteProfileError{Err: err}
	}

	confirmed, err := promptUserToConfirmDelete(pName, rc)
	if err != nil {
		return &DeleteProfileError{Err: err}
	}

	if !confirmed {
		output.Message("Profile deletion cancelled.", nil)

		return nil
	}

	err = deleteProfile(pName)
	if err != nil {
		return &DeleteProfileError{Err: err}
	}

	return nil
}

func promptUserToDeleteProfile(rc io.ReadCloser) (pName string, err error) {
	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return pName, &DeleteProfileError{Err: err}
	}
	pName, err = input.RunPromptSelect("Select profile to delete", koanfConfig.ProfileNames(), rc)

	if err != nil {
		return pName, &DeleteProfileError{Err: err}
	}

	return pName, nil
}

func promptUserToConfirmDelete(pName string, rc io.ReadCloser) (confirmed bool, err error) {
	autoAccept := "false"
	if options.ConfigDeleteAutoAcceptOption.Flag.Changed {
		autoAccept, err = profiles.GetOptionValue(options.ConfigDeleteAutoAcceptOption)
		if err != nil {
			return false, &DeleteProfileError{Err: err}
		}
	}

	if autoAccept == "true" {
		return true, nil
	}

	confirmed, err = input.RunPromptConfirm(fmt.Sprintf("Are you sure you want to delete profile '%s'", pName), rc)
	if err != nil {
		return false, &DeleteProfileError{Err: err}
	}
	return confirmed, nil
}

func deleteProfile(pName string) (err error) {
	output.Message(fmt.Sprintf("Deleting profile '%s'...", pName), nil)

	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &DeleteProfileError{Err: err}
	}

	if err = koanfConfig.DeleteProfile(pName); err != nil {
		return &DeleteProfileError{Err: err}
	}

	output.Success(fmt.Sprintf("Profile '%s' deleted.", pName), nil)

	return nil
}
