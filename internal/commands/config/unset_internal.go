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

var ErrDetermineProfileUnset = errors.New("unable to determine profile to unset configuration from")

type UnsetError struct {
	Err error
}

func (e *UnsetError) Error() string {
	var err *UnsetError
	if errors.As(e.Err, &err) {
		return err.Error()
	}
	return fmt.Sprintf("failed to unset configuration: %s", e.Err.Error())
}

func (e *UnsetError) Unwrap() error {
	var err *UnsetError
	if errors.As(e.Err, &err) {
		return err.Unwrap()
	}
	return e.Err
}

func RunInternalConfigUnset(koanfKey string) (err error) {
	if err = configuration.ValidateKoanfKey(koanfKey); err != nil {
		return &UnsetError{Err: err}
	}

	pName, err := readConfigUnsetOptions()
	if err != nil {
		return &UnsetError{Err: err}
	}

	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &UnsetError{Err: err}
	}

	subKoanf, err := koanfConfig.GetProfileKoanf(pName)
	if err != nil {
		return &UnsetError{Err: err}
	}

	opt, err := configuration.OptionFromKoanfKey(koanfKey)
	if err != nil {
		return &UnsetError{Err: err}
	}

	err = subKoanf.Set(opt.KoanfKey, opt.DefaultValue)
	if err != nil {
		return &UnsetError{Err: err}
	}

	if err = koanfConfig.SaveProfile(pName, subKoanf); err != nil {
		return &UnsetError{Err: err}
	}

	msgStr := "Configuration unset successfully:\n"

	vVal, _, err := profiles.KoanfValueFromOption(opt, pName)
	if err != nil {
		return &UnsetError{Err: err}
	}

	unmaskOptionVal, err := profiles.GetOptionValue(options.ConfigUnmaskSecretValueOption)
	if err != nil {
		unmaskOptionVal = "false"
	}

	if opt.Sensitive && strings.EqualFold(unmaskOptionVal, "false") {
		msgStr += fmt.Sprintf("%s=%s", opt.KoanfKey, profiles.MaskValue(vVal))
	} else {
		msgStr += fmt.Sprintf("%s=%s", opt.KoanfKey, vVal)
	}

	output.Success(msgStr, nil)

	return nil
}

func readConfigUnsetOptions() (pName string, err error) {
	if !options.RootProfileOption.Flag.Changed {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
	} else {
		pName, err = profiles.GetOptionValue(options.RootProfileOption)
	}

	if err != nil {
		return pName, &UnsetError{Err: err}
	}

	if pName == "" {
		return pName, &UnsetError{Err: ErrDetermineProfileUnset}
	}

	return pName, nil
}
