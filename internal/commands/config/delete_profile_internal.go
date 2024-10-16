package config_internal

import (
	"fmt"
	"io"

	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigDeleteProfile(rc io.ReadCloser) (err error) {
	pName, err := promptUser(rc)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %v", err)
	}

	if pName == "" {
		output.Print(output.Opts{
			Message: "Profile deletion cancelled.",
			Result:  output.ENUM_RESULT_NIL,
		})
		return nil
	}

	err = deleteProfile(pName)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %v", err)
	}

	return nil
}

func promptUser(rc io.ReadCloser) (string, error) {
	pName, err := input.RunPromptSelect("Select profile to delete: ", profiles.GetMainConfig().ProfileNames(), rc)
	if err != nil {
		return "", err
	}

	confirmed, err := input.RunPromptConfirm(fmt.Sprintf("Confirm that you want to delete profile: '%s'", pName), rc)
	if err != nil {
		return "", fmt.Errorf("failed to delete profile: %v", err)
	}

	if confirmed {
		return pName, nil
	} else {
		return "", nil
	}
}

func deleteProfile(pName string) (err error) {
	output.Print(output.Opts{
		Message: fmt.Sprintf("Deleting profile '%s'...", pName),
		Result:  output.ENUM_RESULT_NIL,
	})

	if err = profiles.GetMainConfig().DeleteProfile(pName); err != nil {
		return err
	}

	output.Print(output.Opts{
		Message: fmt.Sprintf("Profile '%s' deleted.", pName),
		Result:  output.ENUM_RESULT_SUCCESS,
	})

	return nil
}
