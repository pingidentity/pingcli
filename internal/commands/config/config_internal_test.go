package config_internal

import (
	"os"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// Test RunInternalConfig function
func Test_RunInternalConfig(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		oldProfile  = customtypes.String("production")
		profileName = customtypes.String("test-profile")
		description = customtypes.String("test-description")
	)

	options.ConfigProfileOption.Flag.Changed = true
	options.ConfigProfileOption.CobraParamValue = &oldProfile

	options.ConfigNameOption.Flag.Changed = true
	options.ConfigNameOption.CobraParamValue = &profileName

	options.ConfigDescriptionOption.Flag.Changed = true
	options.ConfigDescriptionOption.CobraParamValue = &description

	err := RunInternalConfig(os.Stdin)
	if err != nil {
		t.Errorf("RunInternalConfig returned error: %v", err)
	}
}

// Test RunInternalConfig function fails when existing profile name is provided
func Test_RunInternalConfig_ExistingProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		oldProfile  = customtypes.String("production")
		profileName = customtypes.String("default")
		description = customtypes.String("test-description")
	)

	options.ConfigProfileOption.Flag.Changed = true
	options.ConfigProfileOption.CobraParamValue = &oldProfile

	options.ConfigNameOption.Flag.Changed = true
	options.ConfigNameOption.CobraParamValue = &profileName

	options.ConfigDescriptionOption.Flag.Changed = true
	options.ConfigDescriptionOption.CobraParamValue = &description

	expectedErrorPattern := `^failed to update profile '.*' name to: .*\. invalid profile name: '.*'\. profile already exists$`
	err := RunInternalConfig(os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfig function fails when invalid profile name is provided
func Test_RunInternalConfig_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		oldProfile  = customtypes.String("production")
		profileName = customtypes.String("test-profile!")
		description = customtypes.String("test-description")
	)

	options.ConfigProfileOption.Flag.Changed = true
	options.ConfigProfileOption.CobraParamValue = &oldProfile

	options.ConfigNameOption.Flag.Changed = true
	options.ConfigNameOption.CobraParamValue = &profileName

	options.ConfigDescriptionOption.Flag.Changed = true
	options.ConfigDescriptionOption.CobraParamValue = &description

	expectedErrorPattern := `^failed to update profile '.*' name to: .*\. invalid profile name: '.*'\. name must contain only alphanumeric characters, underscores, and dashes$`
	err := RunInternalConfig(os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfig function fails when profile name is not provided
func Test_RunInternalConfig_NoProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		oldProfile  = customtypes.String("production")
		profileName = customtypes.String("")
		description = customtypes.String("test-description")
	)

	options.ConfigProfileOption.Flag.Changed = true
	options.ConfigProfileOption.CobraParamValue = &oldProfile

	options.ConfigNameOption.Flag.Changed = true
	options.ConfigNameOption.CobraParamValue = &profileName

	options.ConfigDescriptionOption.Flag.Changed = true
	options.ConfigDescriptionOption.CobraParamValue = &description

	expectedErrorPattern := `^failed to update profile\. unable to determine new profile name$`
	err := RunInternalConfig(os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfig function fails when provided active profile name
func Test_RunInternalConfig_ActiveProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		oldProfile  = customtypes.String("default")
		profileName = customtypes.String("test-profile")
		description = customtypes.String("test-description")
	)

	options.ConfigProfileOption.Flag.Changed = true
	options.ConfigProfileOption.CobraParamValue = &oldProfile

	options.ConfigNameOption.Flag.Changed = true
	options.ConfigNameOption.CobraParamValue = &profileName

	options.ConfigDescriptionOption.Flag.Changed = true
	options.ConfigDescriptionOption.CobraParamValue = &description

	expectedErrorPattern := `^failed to update profile '.*' name to: .*\. '.*' is the active profile and cannot be deleted$`
	err := RunInternalConfig(os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
