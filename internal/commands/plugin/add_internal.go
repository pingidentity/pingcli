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
		// exec error contains executable name and $PATH error message
		return fmt.Errorf("failed to add plugin: %w", err)
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

	subKoanf, err := profiles.GetKoanfConfig().GetProfileKoanf(pName)
	if err != nil {
		return fmt.Errorf("failed to get profile: %w", err)
	}

	existingPluginExectuables, _, err := profiles.KoanfValueFromOption(options.PluginExecutablesOption, pName)
	if err != nil {
		return fmt.Errorf("failed to get existing plugin configuration from profile '%s': %w", pName, err)
	}

	strSlice := new(customtypes.StringSlice)
	if err = strSlice.Set(existingPluginExectuables); err != nil {
		return err
	}
	if err = strSlice.Set(pluginExecutable); err != nil {
		return err
	}

	err = subKoanf.Set(options.PluginExecutablesOption.KoanfKey, strSlice)
	if err != nil {
		return err
	}

	if err = profiles.GetKoanfConfig().SaveProfile(pName, subKoanf); err != nil {
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
		return pName, fmt.Errorf("unable to determine active profile")
	}

	return pName, nil
}
