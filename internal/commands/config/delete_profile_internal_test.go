package config_internal

import (
	"os"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// Test RunInternalConfigDeleteProfile function
func Test_RunInternalConfigDeleteProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("production")
	)

	options.ConfigDeleteProfileOption.Flag.Changed = true
	options.ConfigDeleteProfileOption.CobraParamValue = &profileName

	err := RunInternalConfigDeleteProfile(os.Stdin)
	if err != nil {
		t.Errorf("RunInternalConfigDeleteProfile returned error: %v", err)
	}
}

// Test RunInternalConfigDeleteProfile function fails with active profile
func Test_RunInternalConfigDeleteProfile_ActiveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("default")
	)

	options.ConfigDeleteProfileOption.Flag.Changed = true
	options.ConfigDeleteProfileOption.CobraParamValue = &profileName

	expectedErrorPattern := `^failed to delete profile: '.*' is the active profile and cannot be deleted$`
	err := RunInternalConfigDeleteProfile(os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigDeleteProfile function fails with invalid profile name
func Test_RunInternalConfigDeleteProfile_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("(*#&)")
	)

	options.ConfigDeleteProfileOption.Flag.Changed = true
	options.ConfigDeleteProfileOption.CobraParamValue = &profileName

	expectedErrorPattern := `^failed to delete profile: invalid profile name: '.*'\. name must contain only alphanumeric characters, underscores, and dashes$`
	err := RunInternalConfigDeleteProfile(os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigDeleteProfile function fails with empty profile name
func Test_RunInternalConfigDeleteProfile_EmptyProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("")
	)

	options.ConfigDeleteProfileOption.Flag.Changed = true
	options.ConfigDeleteProfileOption.CobraParamValue = &profileName

	expectedErrorPattern := `^failed to delete profile: unable to determine profile name to delete$`
	err := RunInternalConfigDeleteProfile(os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigDeleteProfile function fails with non-existent profile name
func Test_RunInternalConfigDeleteProfile_NonExistentProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("non-existent")
	)

	options.ConfigDeleteProfileOption.Flag.Changed = true
	options.ConfigDeleteProfileOption.CobraParamValue = &profileName

	expectedErrorPattern := `^failed to delete profile: invalid profile name: '.*' profile does not exist$`
	err := RunInternalConfigDeleteProfile(os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
