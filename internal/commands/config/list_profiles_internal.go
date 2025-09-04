// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

type ListProfilesError struct {
	Err error
}

func (e *ListProfilesError) Error() string {
	var err *ListProfilesError
	if errors.As(e.Err, &err) {
		return err.Error()
	}
	return fmt.Sprintf("failed to list profiles: %s", e.Err.Error())
}

func (e *ListProfilesError) Unwrap() error {
	var err *ListProfilesError
	if errors.As(e.Err, &err) {
		return err.Unwrap()
	}
	return e.Err
}

func RunInternalConfigListProfiles() (err error) {
	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &ListProfilesError{Err: err}
	}

	profileNames := koanfConfig.ProfileNames()
	activeProfileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		return &ListProfilesError{Err: err}
	}

	listStr := "Profiles:\n"

	// We need to enable/disable colorize before applying the color to the string below.
	output.SetColorize()
	activeFmt := color.New(color.Bold, color.FgGreen).SprintFunc()

	for i, profileName := range profileNames {
		if profileName == activeProfileName {
			listStr += "- " + profileName + activeFmt(" (active)") + " \n"
		} else {
			listStr += "- " + profileName + "\n"
		}

		description := koanfConfig.KoanfInstance().String(profileName + "." + options.ProfileDescriptionOption.KoanfKey)
		if description != "" {
			listStr += "    " + description
		}

		if i < len(profileNames)-1 {
			listStr += "\n"
		}
	}

	output.Message(listStr, nil)

	return nil
}
