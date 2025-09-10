// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var ErrUndeterminedProfile = errors.New("unable to determine profile to get configuration from")

type GetError struct {
	Err error
}

func (e *GetError) Error() string {
	var err *GetError
	if errors.As(e.Err, &err) {
		return err.Error()
	}
	return fmt.Sprintf("failed to get configuration: %s", e.Err.Error())
}

func (e *GetError) Unwrap() error {
	var err *GetError
	if errors.As(e.Err, &err) {
		return err.Unwrap()
	}
	return e.Err
}

func RunInternalConfigGet(koanfKey string) (err error) {
	if err = configuration.ValidateParentKoanfKey(koanfKey); err != nil {
		return &GetError{Err: err}
	}

	pName, err := readConfigGetOptions()
	if err != nil {
		return &GetError{Err: err}
	}

	msgStr := fmt.Sprintf("Configuration values for profile '%s' and key '%s':\n", pName, koanfKey)

	for _, opt := range options.Options() {
		// We only want options that have a key in the configuration file
		if opt.KoanfKey == "" {
			continue
		}

		// Match the koanfKey (which can be a "parent key". E.g 'service.pingOne' would match all options like 'service.pingone.authentication.type') to all options.
		if !strings.HasPrefix(strings.ToLower(opt.KoanfKey), strings.ToLower(koanfKey)) {
			continue
		}

		vVal, _, err := profiles.KoanfValueFromOption(opt, pName)
		if err != nil {
			return &GetError{Err: err}
		}

		unmaskOptionVal, err := profiles.GetOptionValue(options.ConfigUnmaskSecretValueOption)
		if err != nil {
			unmaskOptionVal = "false"
		}

		if opt.Sensitive && strings.EqualFold(unmaskOptionVal, "false") {
			msgStr += fmt.Sprintf("%s=%s\n", opt.KoanfKey, profiles.MaskValue(vVal))
		} else {
			msgStr += fmt.Sprintf("%s=%s\n", opt.KoanfKey, vVal)
		}
	}

	output.Message(msgStr, nil)

	return nil
}

func readConfigGetOptions() (pName string, err error) {
	if !options.RootProfileOption.Flag.Changed {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
	} else {
		pName, err = profiles.GetOptionValue(options.RootProfileOption)
	}

	if err != nil {
		return pName, &GetError{Err: err}
	}

	if pName == "" {
		return pName, &GetError{Err: ErrUndeterminedProfile}
	}

	return pName, nil
}
