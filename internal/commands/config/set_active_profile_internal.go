package config_internal

import (
	"fmt"
	"io"

	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigSetActiveProfile(args []string, rc io.ReadCloser) (err error) {
	var pName string
	if len(args) == 1 {
		pName = args[0]
	} else {
		pName, err = promptUserToSelectActiveProfile(rc)
		if err != nil {
			return fmt.Errorf("failed to set active profile: %v", err)
		}
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

func promptUserToSelectActiveProfile(rc io.ReadCloser) (pName string, err error) {
	pName, err = input.RunPromptSelect("Select profile to set as active: ", profiles.GetMainConfig().ProfileNames(), rc)

	if err != nil {
		return pName, err
	}

	return pName, nil
}
