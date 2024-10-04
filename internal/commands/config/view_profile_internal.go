package config_internal

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigViewProfile() (err error) {
	pName, err := readConfigViewProfileOptions()
	if err != nil {
		return fmt.Errorf("failed to view profile: %v", err)
	}

	profileStr, err := profiles.GetMainConfig().ProfileToString(pName)
	if err != nil {
		return fmt.Errorf("failed to view profile: %v", err)
	}

	profileStr = fmt.Sprintf("Profile: %s\n\n%s", pName, profileStr)

	output.Print(output.Opts{
		Message: profileStr,
		Result:  output.ENUM_RESULT_NIL,
	})

	return nil
}

func readConfigViewProfileOptions() (pName string, err error) {
	if !options.ConfigViewProfileOption.Flag.Changed {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
	} else {
		pName, err = profiles.GetOptionValue(options.ConfigViewProfileOption)
	}

	if err != nil {
		return pName, err
	}

	if pName == "" {
		return pName, fmt.Errorf("unable to determine profile name to view")
	}

	return pName, nil
}
