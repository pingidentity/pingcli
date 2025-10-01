// Copyright © 2025 Ping Identity Corporation

package plugin_internal

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var (
	addErrorPrefix         = "failed to add plugin"
	ErrPluginAlreadyExists = errors.New("plugin executable already exists in configuration")
	ErrPluginNotFound      = errors.New("plugin executable not found in system PATH")
)

func RunInternalPluginAdd(pluginExecutable string) error {
	if pluginExecutable == "" {
		return &errs.PingCLIError{Prefix: addErrorPrefix, Err: ErrPluginNameEmpty}
	}

	_, err := exec.LookPath(pluginExecutable)
	if err != nil {
		return &errs.PingCLIError{Prefix: addErrorPrefix, Err: fmt.Errorf("%w: %w", ErrPluginNotFound, err)}
	}

	err = addPluginExecutable(pluginExecutable)
	if err != nil {
		return &errs.PingCLIError{Prefix: addErrorPrefix, Err: err}
	}

	output.Success(fmt.Sprintf("Plugin '%s' added.", pluginExecutable), nil)

	return nil
}

func addPluginExecutable(pluginExecutable string) error {
	pName, err := readPluginAddProfileName()
	if err != nil {
		return &errs.PingCLIError{Prefix: addErrorPrefix, Err: err}
	}

	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &errs.PingCLIError{Prefix: addErrorPrefix, Err: err}
	}

	subKoanf, err := koanfConfig.GetProfileKoanf(pName)
	if err != nil {
		return &errs.PingCLIError{Prefix: addErrorPrefix, Err: err}
	}

	existingPluginExectuables, ok, err := profiles.KoanfValueFromOption(options.PluginExecutablesOption, pName)
	if err != nil {
		return &errs.PingCLIError{Prefix: addErrorPrefix, Err: fmt.Errorf("%w: %w", ErrReadPluginNamesConfig, err)}
	}
	if !ok {
		existingPluginExectuables = ""
	}

	strSlice := new(customtypes.StringSlice)
	if err = strSlice.Set(existingPluginExectuables); err != nil {
		return &errs.PingCLIError{Prefix: addErrorPrefix, Err: err}
	}

	// Check if the plugin is already added
	for _, existingPlugin := range strSlice.StringSlice() {
		if strings.EqualFold(existingPlugin, pluginExecutable) {
			return &errs.PingCLIError{Prefix: addErrorPrefix, Err: fmt.Errorf("%w: '%s'", ErrPluginAlreadyExists, pluginExecutable)}
		}
	}

	if err = strSlice.Set(pluginExecutable); err != nil {
		return &errs.PingCLIError{Prefix: addErrorPrefix, Err: err}
	}

	err = subKoanf.Set(options.PluginExecutablesOption.KoanfKey, strSlice)
	if err != nil {
		return &errs.PingCLIError{Prefix: addErrorPrefix, Err: err}
	}

	if err = koanfConfig.SaveProfile(pName, subKoanf); err != nil {
		return &errs.PingCLIError{Prefix: addErrorPrefix, Err: err}
	}

	return nil
}

func readPluginAddProfileName() (pName string, err error) {
	if !options.RootProfileOption.Flag.Changed {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
	} else {
		pName, err = profiles.GetOptionValue(options.RootProfileOption)
	}

	if err != nil {
		return pName, &errs.PingCLIError{Prefix: addErrorPrefix, Err: err}
	}

	if pName == "" {
		return pName, &errs.PingCLIError{Prefix: addErrorPrefix, Err: ErrUndeterminedProfile}
	}

	return pName, nil
}
