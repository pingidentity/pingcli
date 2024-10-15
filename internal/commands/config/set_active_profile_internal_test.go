package config_internal

import (
	"os"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// Test RunInternalConfigSetActiveProfile function
func Test_RunInternalConfigSetActiveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("production")
	)
	options.ConfigSetActiveProfileOption.Flag.Changed = true
	options.ConfigSetActiveProfileOption.CobraParamValue = &profileName

	err := RunInternalConfigSetActiveProfile(os.Stdin)
	if err != nil {
		t.Errorf("RunInternalConfigSetActiveProfile returned error: %v", err)
	}
}

// Test RunInternalConfigSetActiveProfile function fails with invalid profile name
func Test_RunInternalConfigSetActiveProfile_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("(*#&)")
	)
	options.ConfigSetActiveProfileOption.Flag.Changed = true
	options.ConfigSetActiveProfileOption.CobraParamValue = &profileName

	expectedErrorPattern := `^failed to set active profile: invalid profile name: '.*'\. name must contain only alphanumeric characters, underscores, and dashes$`
	err := RunInternalConfigSetActiveProfile(os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSetActiveProfile function fails with non-existent profile
func Test_RunInternalConfigSetActiveProfile_NonExistentProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("non-existent")
	)
	options.ConfigSetActiveProfileOption.Flag.Changed = true
	options.ConfigSetActiveProfileOption.CobraParamValue = &profileName

	expectedErrorPattern := `^failed to set active profile: invalid profile name: '.*' profile does not exist$`
	err := RunInternalConfigSetActiveProfile(os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
