// Copyright © 2025 Ping Identity Corporation

package plugin_internal

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalPluginRemove(pluginExecutable string) error {
	if pluginExecutable == "" {
		return fmt.Errorf("plugin executable is required")
	}

	ok, err := removePluginExecutable(pluginExecutable)
	if err != nil {
		return fmt.Errorf("failed to remove plugin: %w", err)
	}

	if ok {
		output.Success(fmt.Sprintf("Plugin '%s' removed.", pluginExecutable), nil)
	}

	return nil
}

func removePluginExecutable(pluginExecutable string) (bool, error) {
	pName, err := readPluginRemoveProfileName()
	if err != nil {
		return false, fmt.Errorf("failed to read profile name: %w", err)
	}

	subKoanf, err := profiles.GetKoanfConfig().GetProfileKoanf(pName)
	if err != nil {
		return false, fmt.Errorf("failed to get profile: %w", err)
	}

	existingPluginExectuables, _, err := profiles.KoanfValueFromOption(options.PluginExecutablesOption, pName)
	if err != nil {
		return false, fmt.Errorf("failed to get existing plugin configuration from profile '%s': %w", pName, err)
	}

	strSlice := new(customtypes.StringSlice)
	if err = strSlice.Set(existingPluginExectuables); err != nil {
		return false, err
	}
	removed, err := strSlice.Remove(pluginExecutable)
	if err != nil {
		return false, err
	}

	if !removed {
		output.Warn(fmt.Sprintf("plugin executable '%s' not found in profile '%s' plugins", pluginExecutable, pName), nil)

		return false, nil
	}

	err = subKoanf.Set(options.PluginExecutablesOption.KoanfKey, strSlice)
	if err != nil {
		return false, err
	}

	if err = profiles.GetKoanfConfig().SaveProfile(pName, subKoanf); err != nil {
		return false, err
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
		return pName, err
	}

	if pName == "" {
		return pName, fmt.Errorf("unable to determine active profile")
	}

	return pName, nil
}
