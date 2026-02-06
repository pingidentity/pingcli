// Copyright © 2026 Ping Identity Corporation

package config_internal

import (
	"fmt"
	"io"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var (
	setActiveProfileErrorPrefix = "failed to set active profile"
)

func RunInternalConfigSetActiveProfile(args []string, rc io.ReadCloser) (err error) {
	var pName string
	if len(args) == 1 {
		pName = args[0]
	} else {
		pName, err = promptUserToSelectActiveProfile(rc)
		if err != nil {
			return &errs.PingCLIError{Prefix: setActiveProfileErrorPrefix, Err: err}
		}
	}

	output.Message(fmt.Sprintf("Setting active profile to '%s'...", pName), nil)

	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &errs.PingCLIError{Prefix: setActiveProfileErrorPrefix, Err: err}
	}

	if err = koanfConfig.ChangeActiveProfile(pName); err != nil {
		return &errs.PingCLIError{Prefix: setActiveProfileErrorPrefix, Err: err}
	}

	output.Success(fmt.Sprintf("Active profile set to '%s'", pName), nil)

	return nil
}

func promptUserToSelectActiveProfile(rc io.ReadCloser) (pName string, err error) {
	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return "", &errs.PingCLIError{Prefix: setActiveProfileErrorPrefix, Err: err}
	}
	pName, err = input.RunPromptSelect("Select profile to set as active: ", koanfConfig.ProfileNames(), rc)

	if err != nil {
		return pName, &errs.PingCLIError{Prefix: setActiveProfileErrorPrefix, Err: err}
	}

	return pName, nil
}
