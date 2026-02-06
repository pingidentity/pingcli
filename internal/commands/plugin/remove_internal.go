// Copyright © 2026 Ping Identity Corporation

package plugin_internal

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var (
	removeErrorPrefix = "failed to remove plugin"
)

func RunInternalPluginRemove(pluginExecutable string) error {
	if pluginExecutable == "" {
		return &errs.PingCLIError{Prefix: removeErrorPrefix, Err: ErrPluginNameEmpty}
	}

	ok, err := removePluginExecutable(pluginExecutable)
	if err != nil {
		return &errs.PingCLIError{Prefix: removeErrorPrefix, Err: err}
	}

	if ok {
		output.Success(fmt.Sprintf("Plugin '%s' removed.", pluginExecutable), nil)
	} else {
		output.Warn(fmt.Sprintf("Plugin '%s' not found in configuration and was not removed.", pluginExecutable), nil)
	}

	return nil
}

func removePluginExecutable(pluginExecutable string) (bool, error) {
	pName, err := readPluginRemoveProfileName()
	if err != nil {
		return false, &errs.PingCLIError{Prefix: removeErrorPrefix, Err: err}
	}

	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return false, &errs.PingCLIError{Prefix: removeErrorPrefix, Err: err}
	}

	subKoanf, err := koanfConfig.GetProfileKoanf(pName)
	if err != nil {
		return false, &errs.PingCLIError{Prefix: removeErrorPrefix, Err: err}
	}

	existingPluginExectuables, _, err := profiles.KoanfValueFromOption(options.PluginExecutablesOption, pName)
	if err != nil {
		return false, &errs.PingCLIError{Prefix: removeErrorPrefix, Err: fmt.Errorf("%w: %w", ErrReadPluginNamesConfig, err)}
	}

	strSlice := new(customtypes.StringSlice)
	if err = strSlice.Set(existingPluginExectuables); err != nil {
		return false, &errs.PingCLIError{Prefix: removeErrorPrefix, Err: err}
	}
	removed, err := strSlice.Remove(pluginExecutable)
	if err != nil {
		return false, &errs.PingCLIError{Prefix: removeErrorPrefix, Err: err}
	}

	if !removed {
		return false, nil
	}

	err = subKoanf.Set(options.PluginExecutablesOption.KoanfKey, strSlice)
	if err != nil {
		return false, &errs.PingCLIError{Prefix: removeErrorPrefix, Err: err}
	}

	if err = koanfConfig.SaveProfile(pName, subKoanf); err != nil {
		return false, &errs.PingCLIError{Prefix: removeErrorPrefix, Err: err}
	}

	return true, nil
}

func readPluginRemoveProfileName() (pName string, err error) {
	if !options.RootProfileOption.Flag.Changed {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
	} else {
		pName, err = profiles.GetOptionValue(options.RootProfileOption)
	}

	if err != nil {
		return pName, &errs.PingCLIError{Prefix: removeErrorPrefix, Err: err}
	}

	if pName == "" {
		return pName, &errs.PingCLIError{Prefix: removeErrorPrefix, Err: ErrUndeterminedProfile}
	}

	return pName, nil
}
