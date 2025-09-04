// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"errors"
	"fmt"
	"io"

	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

type SetActiveProfileError struct {
	Err error
}

func (e *SetActiveProfileError) Error() string {
	var err *SetActiveProfileError
	if errors.As(e.Err, &err) {
		return err.Error()
	}
	return fmt.Sprintf("failed to set active profile: %s", e.Err.Error())
}

func (e *SetActiveProfileError) Unwrap() error {
	var err *SetActiveProfileError
	if errors.As(e.Err, &err) {
		return err.Unwrap()
	}
	return e.Err
}

func RunInternalConfigSetActiveProfile(args []string, rc io.ReadCloser) (err error) {
	var pName string
	if len(args) == 1 {
		pName = args[0]
	} else {
		pName, err = promptUserToSelectActiveProfile(rc)
		if err != nil {
			return &SetActiveProfileError{Err: err}
		}
	}

	output.Message(fmt.Sprintf("Setting active profile to '%s'...", pName), nil)

	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &SetActiveProfileError{Err: err}
	}

	if err = koanfConfig.ChangeActiveProfile(pName); err != nil {
		return &SetActiveProfileError{Err: err}
	}

	output.Success(fmt.Sprintf("Active profile set to '%s'", pName), nil)

	return nil
}

func promptUserToSelectActiveProfile(rc io.ReadCloser) (pName string, err error) {
	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return "", &SetActiveProfileError{Err: err}
	}
	pName, err = input.RunPromptSelect("Select profile to set as active: ", koanfConfig.ProfileNames(), rc)

	if err != nil {
		return pName, &SetActiveProfileError{Err: err}
	}

	return pName, nil
}
