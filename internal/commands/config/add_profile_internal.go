// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"fmt"
	"io"
	"strconv"

	"github.com/knadh/koanf/v2"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigAddProfile(rc io.ReadCloser) (err error) {
	newProfileName, newDescription, setActive, err := readConfigAddProfileOptions(rc)
	if err != nil {
		return fmt.Errorf("failed to add profile: %w", err)
	}

	err = profiles.GetKoanfConfig().ValidateNewProfileName(newProfileName)
	if err != nil {
		return fmt.Errorf("failed to add profile: %w", err)
	}

	output.Message(fmt.Sprintf("Adding new profile '%s'...", newProfileName), nil)

	subKoanf := koanf.New(".")
	err = subKoanf.Set(options.ProfileDescriptionOption.KoanfKey, newDescription)
	if err != nil {
		return fmt.Errorf("failed to add profile: %w", err)
	}

	if err = profiles.GetKoanfConfig().SaveProfile(newProfileName, subKoanf); err != nil {
		return fmt.Errorf("failed to add profile: %w", err)
	}

	output.Success(fmt.Sprintf("Profile created. Update additional profile attributes via 'pingcli config set' or directly within the config file at '%s'", profiles.GetKoanfConfig().GetKoanfConfigFile()), nil)

	if setActive {
		if err = profiles.GetKoanfConfig().ChangeActiveProfile(newProfileName); err != nil {
			return fmt.Errorf("failed to set active profile: %w", err)
		}

		output.Success(fmt.Sprintf("Profile '%s' set as active.", newProfileName), nil)
	}

	err = profiles.GetKoanfConfig().DefaultMissingKoanfKeys()
	if err != nil {
		return fmt.Errorf("failed to add profile: %w", err)
	}

	err = profiles.GetKoanfConfig().SaveProfile(newProfileName, subKoanf)
	if err != nil {
		return fmt.Errorf("failed to save added profile: %w", err)
	}

	return nil
}

func readConfigAddProfileOptions(rc io.ReadCloser) (newProfileName, newDescription string, setActive bool, err error) {
	if newProfileName, err = readConfigAddProfileNameOption(rc); err != nil {
		return newProfileName, newDescription, setActive, err
	}

	if newDescription, err = readConfigAddProfileDescriptionOption(rc); err != nil {
		return newProfileName, newDescription, setActive, err
	}

	if setActive, err = readConfigAddProfileSetActiveOption(rc); err != nil {
		return newProfileName, newDescription, setActive, err
	}

	return newProfileName, newDescription, setActive, nil
}

func readConfigAddProfileNameOption(rc io.ReadCloser) (newProfileName string, err error) {
	if !options.ConfigAddProfileNameOption.Flag.Changed {
		newProfileName, err = input.RunPrompt("New profile name", profiles.GetKoanfConfig().ValidateNewProfileName, rc)
		if err != nil {
			return newProfileName, err
		}

		if newProfileName == "" {
			return newProfileName, fmt.Errorf("unable to determine profile name")
		}
	} else {
		newProfileName, err = profiles.GetOptionValue(options.ConfigAddProfileNameOption)
		if err != nil {
			return newProfileName, err
		}

		if newProfileName == "" {
			return newProfileName, fmt.Errorf("unable to determine profile name")
		}
	}

	return newProfileName, nil
}

func readConfigAddProfileDescriptionOption(rc io.ReadCloser) (newDescription string, err error) {
	if !options.ConfigAddProfileDescriptionOption.Flag.Changed {
		return input.RunPrompt("New profile description: ", nil, rc)
	} else {
		return profiles.GetOptionValue(options.ConfigAddProfileDescriptionOption)
	}
}

func readConfigAddProfileSetActiveOption(rc io.ReadCloser) (setActive bool, err error) {
	if !options.ConfigAddProfileSetActiveOption.Flag.Changed {
		return input.RunPromptConfirm("Set new profile as active: ", rc)
	} else {
		boolStr, err := profiles.GetOptionValue(options.ConfigAddProfileSetActiveOption)
		if err != nil {
			return setActive, err
		}

		return strconv.ParseBool(boolStr)
	}
}
