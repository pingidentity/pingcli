package config_internal

import (
	"fmt"
	"io"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigSetActiveProfile(rc io.ReadCloser) (err error) {
	pName, err := readConfigSetActiveProfileOptions(rc)
	if err != nil {
		return fmt.Errorf("failed to set active profile: %v", err)
	}

	output.Print(output.Opts{
		Message: fmt.Sprintf("Setting active profile to '%s'...", pName),
		Result:  output.ENUM_RESULT_NIL,
	})

	if err = profiles.GetMainConfig().ChangeActiveProfile(pName); err != nil {
		return fmt.Errorf("failed to set active profile: %v", err)
	}

	output.Print(output.Opts{
		Message: fmt.Sprintf("Active profile set to '%s'", pName),
		Result:  output.ENUM_RESULT_SUCCESS,
	})

	return nil
}

func readConfigSetActiveProfileOptions(rc io.ReadCloser) (pName string, err error) {
	if !options.ConfigSetActiveProfileOption.Flag.Changed {
		pName, err = input.RunPromptSelect("Select profile to set as active: ", profiles.GetMainConfig().ProfileNames(), rc)
	} else {
		pName, err = profiles.GetOptionValue(options.ConfigSetActiveProfileOption)
	}

	if err != nil {
		return pName, err
	}

	if pName == "" {
		return pName, fmt.Errorf("unable to determine profile name to set as active")
	}

	return pName, nil
}
