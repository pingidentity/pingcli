// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var (
	unsetErrorPrefix = "failed to unset configuration"
)

func RunInternalConfigUnset(koanfKey string) (err error) {
	if err = configuration.ValidateKoanfKey(koanfKey); err != nil {
		return &errs.PingCLIError{Prefix: unsetErrorPrefix, Err: err}
	}

	pName, err := readConfigUnsetOptions()
	if err != nil {
		return &errs.PingCLIError{Prefix: unsetErrorPrefix, Err: err}
	}

	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &errs.PingCLIError{Prefix: unsetErrorPrefix, Err: err}
	}

	subKoanf, err := koanfConfig.GetProfileKoanf(pName)
	if err != nil {
		return &errs.PingCLIError{Prefix: unsetErrorPrefix, Err: err}
	}

	opt, err := configuration.OptionFromKoanfKey(koanfKey)
	if err != nil {
		return &errs.PingCLIError{Prefix: unsetErrorPrefix, Err: err}
	}

	err = subKoanf.Set(opt.KoanfKey, opt.DefaultValue)
	if err != nil {
		return &errs.PingCLIError{Prefix: unsetErrorPrefix, Err: err}
	}

	if err = koanfConfig.SaveProfile(pName, subKoanf); err != nil {
		return &errs.PingCLIError{Prefix: unsetErrorPrefix, Err: err}
	}

	msgStr := "Configuration unset successfully:\n"

	vVal, _, err := profiles.KoanfValueFromOption(opt, pName)
	if err != nil {
		return &errs.PingCLIError{Prefix: unsetErrorPrefix, Err: err}
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
		return pName, &errs.PingCLIError{Prefix: unsetErrorPrefix, Err: err}
	}

	if pName == "" {
		return pName, &errs.PingCLIError{Prefix: unsetErrorPrefix, Err: ErrUndeterminedProfile}
	}

	return pName, nil
}
