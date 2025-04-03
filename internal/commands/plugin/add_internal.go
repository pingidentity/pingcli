package plugin_internal

import (
	"fmt"
	"os/exec"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalPluginAdd(pluginExecutable string) error {
	if pluginExecutable == "" {
		return fmt.Errorf("plugin executable is required")
	}

	// Check if plugin executable is in PATH
	_, err := exec.LookPath(pluginExecutable)
	if err != nil {
		return fmt.Errorf("failed to add plugin: plugin executable not found in PATH: %w", err)
	}

	err = addPluginExecutable(pluginExecutable)
	if err != nil {
		return fmt.Errorf("failed to add plugin: %w", err)
	}

	output.Success(fmt.Sprintf("Plugin '%s' added.", pluginExecutable), nil)

	return nil
}

func addPluginExecutable(pluginExecutable string) error {
	pName, err := readPluginAddProfileName()
	if err != nil {
		return fmt.Errorf("failed to read profile name: %w", err)
	}

	subViper, err := profiles.GetMainConfig().GetProfileViper(pName)
	if err != nil {
		return fmt.Errorf("failed to get profile viper: %w", err)
	}

	existingPluginExectuables, _, err := profiles.ViperValueFromOption(options.PluginExecutablesOption)
	if err != nil {
		return fmt.Errorf("failed to get existing plugin configuration: %w", err)
	}

	strSlice := new(customtypes.StringSlice)
	if err = strSlice.Set(existingPluginExectuables); err != nil {
		return fmt.Errorf("failed to validate existing executables of key '%s': %w", options.PluginExecutablesOption.ViperKey, err)
	}
	if err = strSlice.Set(pluginExecutable); err != nil {
		return fmt.Errorf("failed to add new executable to key '%s': %w", options.PluginExecutablesOption.ViperKey, err)
	}

	subViper.Set(options.PluginExecutablesOption.ViperKey, strSlice)

	if err = profiles.GetMainConfig().SaveProfile(pName, subViper); err != nil {
		return err
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
		return pName, err
	}

	if pName == "" {
		return pName, fmt.Errorf("unable to determine profile to add plugin to")
	}

	return pName, nil
}
