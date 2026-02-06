// Copyright © 2026 Ping Identity Corporation

package config_internal

import (
	"fmt"
	"io"
	"strconv"

	"github.com/knadh/koanf/v2"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var (
	addProfileErrorPrefix = "failed to add profile"
)

func RunInternalConfigAddProfile(rc io.ReadCloser, koanfConfig *profiles.KoanfConfig) (err error) {
	if koanfConfig == nil {
		return &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: ErrKoanfNotInitialized}
	}

	newProfileName, newDescription, setActive, err := readConfigAddProfileOptions(rc)
	if err != nil {
		return &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
	}

	err = koanfConfig.ValidateNewProfileName(newProfileName)
	if err != nil {
		return &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
	}

	output.Message(fmt.Sprintf("Adding new profile '%s'...", newProfileName), nil)

	subKoanf := koanf.New(".")
	err = subKoanf.Set(options.ProfileDescriptionOption.KoanfKey, newDescription)
	if err != nil {
		return &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
	}

	if err = koanfConfig.SaveProfile(newProfileName, subKoanf); err != nil {
		return &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
	}

	output.Success(fmt.Sprintf("Profile created. Update additional profile attributes via 'pingcli config set' or directly within the config file at '%s'", koanfConfig.GetKoanfConfigFile()), nil)

	if setActive {
		if err = koanfConfig.ChangeActiveProfile(newProfileName); err != nil {
			return &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
		}

		output.Success(fmt.Sprintf("Profile '%s' set as active.", newProfileName), nil)
	}

	err = koanfConfig.DefaultMissingKoanfKeys()
	if err != nil {
		return &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
	}

	return nil
}

func readConfigAddProfileOptions(rc io.ReadCloser) (newProfileName, newDescription string, setActive bool, err error) {
	if newProfileName, err = readConfigAddProfileNameOption(rc); err != nil {
		return newProfileName, newDescription, setActive, &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
	}

	if newDescription, err = readConfigAddProfileDescriptionOption(rc); err != nil {
		return newProfileName, newDescription, setActive, &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
	}

	if setActive, err = readConfigAddProfileSetActiveOption(rc); err != nil {
		return newProfileName, newDescription, setActive, &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
	}

	return newProfileName, newDescription, setActive, nil
}

func readConfigAddProfileNameOption(rc io.ReadCloser) (newProfileName string, err error) {
	if !options.ConfigAddProfileNameOption.Flag.Changed {
		koanfConfig, err := profiles.GetKoanfConfig()
		if err != nil {
			return newProfileName, &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
		}

		newProfileName, err = input.RunPrompt("New profile name", koanfConfig.ValidateNewProfileName, rc)
		if err != nil {
			return newProfileName, &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
		}

		if newProfileName == "" {
			return newProfileName, &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: ErrNoProfileProvided}
		}
	} else {
		newProfileName, err = profiles.GetOptionValue(options.ConfigAddProfileNameOption)
		if err != nil {
			return newProfileName, &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
		}

		if newProfileName == "" {
			return newProfileName, &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: ErrNoProfileProvided}
		}
	}

	return newProfileName, nil
}

func readConfigAddProfileDescriptionOption(rc io.ReadCloser) (newDescription string, err error) {
	if !options.ConfigAddProfileDescriptionOption.Flag.Changed {
		newDescription, err = input.RunPrompt("New profile description: ", nil, rc)
		if err != nil {
			return newDescription, &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
		}
	} else {
		newDescription, err = profiles.GetOptionValue(options.ConfigAddProfileDescriptionOption)
		if err != nil {
			return newDescription, &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
		}
	}

	return newDescription, nil
}

func readConfigAddProfileSetActiveOption(rc io.ReadCloser) (setActive bool, err error) {
	if !options.ConfigAddProfileSetActiveOption.Flag.Changed {
		setActive, err = input.RunPromptConfirm("Set new profile as active: ", rc)
		if err != nil {
			return setActive, &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
		}
	} else {
		boolStr, err := profiles.GetOptionValue(options.ConfigAddProfileSetActiveOption)
		if err != nil {
			return setActive, &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: err}
		}

		setActive, err = strconv.ParseBool(boolStr)
		if err != nil {
			return setActive, &errs.PingCLIError{Prefix: addProfileErrorPrefix, Err: ErrSetActiveFlagInvalid}
		}
	}

	return setActive, nil
}
