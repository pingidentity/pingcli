package config_internal

import (
	"fmt"
	"io"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigDeleteProfile(rc io.ReadCloser) (err error) {
	pName, err := readConfigDeleteProfileOptions(rc)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %v", err)
	}

	output.Print(output.Opts{
		Message: fmt.Sprintf("Deleting profile '%s'...", pName),
		Result:  output.ENUM_RESULT_NIL,
	})

	if err := profiles.GetMainConfig().DeleteProfile(pName); err != nil {
		return fmt.Errorf("failed to delete profile: %v", err)
	}

	output.Print(output.Opts{
		Message: fmt.Sprintf("Profile '%s' deleted.", pName),
		Result:  output.ENUM_RESULT_SUCCESS,
	})

	return nil
}

func readConfigDeleteProfileOptions(rc io.ReadCloser) (pName string, err error) {
	if !options.ConfigDeleteProfileOption.Flag.Changed {
		pName, err = input.RunPromptSelect("Select profile to delete: ", profiles.GetMainConfig().ProfileNames(), rc)
	} else {
		pName, err = profiles.GetOptionValue(options.ConfigDeleteProfileOption)
	}

	if err != nil {
		return pName, err
	}

	if pName == "" {
		return pName, fmt.Errorf("unable to determine profile name to delete")
	}

	return pName, nil
}
