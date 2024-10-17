package config_internal

import (
	"fmt"
	"io"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfig(rc io.ReadCloser) (err error) {
	profileName, newProfileName, newDescription, err := readConfigOptions(rc)
	if err != nil {
		return fmt.Errorf("failed to update profile. %v", err)
	}

	output.Print(output.Opts{
		Message: fmt.Sprintf("Updating profile '%s'...", profileName),
		Result:  output.ENUM_RESULT_NIL,
	})

	if err = profiles.GetMainConfig().ChangeProfileName(profileName, newProfileName); err != nil {
		return fmt.Errorf("failed to update profile '%s' name to: %s. %v", profileName, newProfileName, err)
	}

	if err = profiles.GetMainConfig().ChangeProfileDescription(newProfileName, newDescription); err != nil {
		return fmt.Errorf("failed to update profile '%s' description to: %s. %v", newProfileName, newDescription, err)
	}

	output.Print(output.Opts{
		Message: fmt.Sprintf("Profile updated. Update additional profile attributes via 'pingcli config set' or directly within the config file at '%s'", profiles.GetMainConfig().ViperInstance().ConfigFileUsed()),
		Result:  output.ENUM_RESULT_SUCCESS,
	})

	return nil
}

func readConfigOptions(rc io.ReadCloser) (profileName, newName, description string, err error) {
	if profileName, err = readConfigProfileNameOption(); err != nil {
		return profileName, newName, description, err
	}

	if newName, err = readConfigNameOption(rc); err != nil {
		return profileName, newName, description, err
	}

	if description, err = readConfigDescriptionOption(rc); err != nil {
		return profileName, newName, description, err
	}

	return profileName, newName, description, nil
}

func readConfigProfileNameOption() (pName string, err error) {
	if !options.ConfigProfileOption.Flag.Changed {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
	} else {
		pName, err = profiles.GetOptionValue(options.ConfigProfileOption)
	}

	if err != nil {
		return pName, err
	}

	if pName == "" {
		return pName, fmt.Errorf("unable to determine profile name to update")
	}

	return pName, nil
}

func readConfigNameOption(rc io.ReadCloser) (newName string, err error) {
	if !options.ConfigNameOption.Flag.Changed {
		newName, err = input.RunPrompt("New profile name: ", validateChangeProfileName, rc)
	} else {
		newName, err = profiles.GetOptionValue(options.ConfigNameOption)
	}

	if err != nil {
		return newName, err
	}

	if newName == "" {
		return newName, fmt.Errorf("unable to determine new profile name")
	}

	return newName, nil
}

func readConfigDescriptionOption(rc io.ReadCloser) (description string, err error) {
	if !options.ConfigDescriptionOption.Flag.Changed {
		return input.RunPrompt("New profile description: ", nil, rc)
	} else {
		return profiles.GetOptionValue(options.ConfigDescriptionOption)
	}
}

func validateChangeProfileName(newName string) (err error) {
	oldName, err := readConfigProfileNameOption()
	if err != nil {
		return err
	}

	if err = profiles.GetMainConfig().ValidateUpdateExistingProfileName(oldName, newName); err != nil {
		return err
	}

	return nil
}
