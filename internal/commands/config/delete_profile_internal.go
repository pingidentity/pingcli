package config_internal

import (
	"fmt"
	"io"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigDeleteProfile(args []string, rc io.ReadCloser) (err error) {
	var pName string
	if len(args) == 1 {
		pName = args[0]
	} else {
		pName, err = promptUserToDeleteProfile(rc)
		if err != nil {
			return fmt.Errorf("failed to delete profile: %v", err)
		}
	}

	confirmed, err := promptUserToConfirmDelete(pName, rc)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %v", err)
	}

	if !confirmed {
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

func promptUserToDeleteProfile(rc io.ReadCloser) (pName string, err error) {
	pName, err = input.RunPromptSelect("Select profile to delete: ", profiles.GetMainConfig().ProfileNames(), rc)

	if err != nil {
		return pName, err
	}

	return pName, nil
}

func promptUserToConfirmDelete(pName string, rc io.ReadCloser) (confirmed bool, err error) {
	autoAccept := "false"
	if options.ConfigDeleteAutoAcceptOption.Flag.Changed {
		autoAccept, err = profiles.GetOptionValue(options.ConfigDeleteAutoAcceptOption)
		if err != nil {
			return false, err
		}
	}

	if autoAccept == "true" {
		return true, nil
	}

	return input.RunPromptConfirm(fmt.Sprintf("Are you sure you want to delete profile '%s'?", pName), rc)
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
