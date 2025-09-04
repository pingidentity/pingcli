// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/knadh/koanf/v2"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var (
	ErrProfileNameNotProvided = errors.New("unable to determine profile name")
	ErrSetActiveInvalid       = errors.New("invalid value for set-active flag. must be 'true' or 'false'")
)

type AddProfileError struct {
	Err error
}

func (e *AddProfileError) Error() string {
	var err *AddProfileError
	if errors.As(e.Err, &err) {
		return err.Error()
	}
	return fmt.Sprintf("failed to add profile: %s", e.Err.Error())
}

func (e *AddProfileError) Unwrap() error {
	var err *AddProfileError
	if errors.As(e.Err, &err) {
		return err.Unwrap()
	}
	return e.Err
}

func RunInternalConfigAddProfile(rc io.ReadCloser) (err error) {
	newProfileName, newDescription, setActive, err := readConfigAddProfileOptions(rc)
	if err != nil {
		return &AddProfileError{Err: err}
	}

	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &AddProfileError{Err: err}
	}

	err = koanfConfig.ValidateNewProfileName(newProfileName)
	if err != nil {
		return &AddProfileError{Err: err}
	}

	output.Message(fmt.Sprintf("Adding new profile '%s'...", newProfileName), nil)

	subKoanf := koanf.New(".")
	err = subKoanf.Set(options.ProfileDescriptionOption.KoanfKey, newDescription)
	if err != nil {
		return &AddProfileError{Err: err}
	}

	if err = koanfConfig.SaveProfile(newProfileName, subKoanf); err != nil {
		return &AddProfileError{Err: err}
	}

	output.Success(fmt.Sprintf("Profile created. Update additional profile attributes via 'pingcli config set' or directly within the config file at '%s'", koanfConfig.GetKoanfConfigFile()), nil)

	if setActive {
		if err = koanfConfig.ChangeActiveProfile(newProfileName); err != nil {
			return &AddProfileError{Err: err}
		}

		output.Success(fmt.Sprintf("Profile '%s' set as active.", newProfileName), nil)
	}

	err = koanfConfig.DefaultMissingKoanfKeys()
	if err != nil {
		return &AddProfileError{Err: err}
	}

	return nil
}

func readConfigAddProfileOptions(rc io.ReadCloser) (newProfileName, newDescription string, setActive bool, err error) {
	if newProfileName, err = readConfigAddProfileNameOption(rc); err != nil {
		return newProfileName, newDescription, setActive, &AddProfileError{Err: err}
	}

	if newDescription, err = readConfigAddProfileDescriptionOption(rc); err != nil {
		return newProfileName, newDescription, setActive, &AddProfileError{Err: err}
	}

	if setActive, err = readConfigAddProfileSetActiveOption(rc); err != nil {
		return newProfileName, newDescription, setActive, &AddProfileError{Err: err}
	}

	return newProfileName, newDescription, setActive, nil
}

func readConfigAddProfileNameOption(rc io.ReadCloser) (newProfileName string, err error) {
	if !options.ConfigAddProfileNameOption.Flag.Changed {
		koanfConfig, err := profiles.GetKoanfConfig()
		if err != nil {
			return newProfileName, &AddProfileError{Err: err}
		}

		newProfileName, err = input.RunPrompt("New profile name", koanfConfig.ValidateNewProfileName, rc)
		if err != nil {
			return newProfileName, &AddProfileError{Err: err}
		}

		if newProfileName == "" {
			return newProfileName, &AddProfileError{Err: ErrProfileNameNotProvided}
		}
	} else {
		newProfileName, err = profiles.GetOptionValue(options.ConfigAddProfileNameOption)
		if err != nil {
			return newProfileName, &AddProfileError{Err: err}
		}

		if newProfileName == "" {
			return newProfileName, &AddProfileError{Err: ErrProfileNameNotProvided}
		}
	}

	return newProfileName, nil
}

func readConfigAddProfileDescriptionOption(rc io.ReadCloser) (newDescription string, err error) {
	if !options.ConfigAddProfileDescriptionOption.Flag.Changed {
		newDescription, err = input.RunPrompt("New profile description: ", nil, rc)
		if err != nil {
			return newDescription, &AddProfileError{Err: err}
		}
	} else {
		newDescription, err = profiles.GetOptionValue(options.ConfigAddProfileDescriptionOption)
		if err != nil {
			return newDescription, &AddProfileError{Err: err}
		}
	}

	return newDescription, nil
}

func readConfigAddProfileSetActiveOption(rc io.ReadCloser) (setActive bool, err error) {
	if !options.ConfigAddProfileSetActiveOption.Flag.Changed {
		setActive, err = input.RunPromptConfirm("Set new profile as active: ", rc)
		if err != nil {
			return setActive, &AddProfileError{Err: err}
		}
	} else {
		boolStr, err := profiles.GetOptionValue(options.ConfigAddProfileSetActiveOption)
		if err != nil {
			return setActive, &AddProfileError{Err: err}
		}

		setActive, err = strconv.ParseBool(boolStr)
		if err != nil {
			return setActive, &AddProfileError{Err: ErrSetActiveInvalid}
		}
	}

	return setActive, nil
}
