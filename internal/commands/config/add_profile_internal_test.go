package config_internal

import (
	"os"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// Test RunInternalConfigAddProfile function
func Test_RunInternalConfigAddProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("test-profile")
		description = customtypes.String("test-description")
		setActive   = customtypes.Bool(false)
	)

	options.ConfigAddProfileNameOption.Flag.Changed = true
	options.ConfigAddProfileNameOption.CobraParamValue = &profileName

	options.ConfigAddProfileDescriptionOption.Flag.Changed = true
	options.ConfigAddProfileDescriptionOption.CobraParamValue = &description

	options.ConfigAddProfileSetActiveOption.Flag.Changed = true
	options.ConfigAddProfileSetActiveOption.CobraParamValue = &setActive

	err := RunInternalConfigAddProfile(os.Stdin)
	if err != nil {
		t.Errorf("RunInternalConfigAddProfile returned error: %v", err)
	}
}

// Test RunInternalConfigAddProfile function fails when existing profile name is provided
func Test_RunInternalConfigAddProfile_ExistingProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("default")
		description = customtypes.String("test-description")
		setActive   = customtypes.Bool(false)
	)

	options.ConfigAddProfileNameOption.Flag.Changed = true
	options.ConfigAddProfileNameOption.CobraParamValue = &profileName

	options.ConfigAddProfileDescriptionOption.Flag.Changed = true
	options.ConfigAddProfileDescriptionOption.CobraParamValue = &description

	options.ConfigAddProfileSetActiveOption.Flag.Changed = true
	options.ConfigAddProfileSetActiveOption.CobraParamValue = &setActive

	expectedErrorPattern := `^failed to add profile: invalid profile name: '.*'. profile already exists$`
	err := RunInternalConfigAddProfile(os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigAddProfile function fails when profile name is not provided
func Test_RunInternalConfigAddProfile_NoProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("")
		description = customtypes.String("test-description")
		setActive   = customtypes.Bool(false)
	)

	options.ConfigAddProfileNameOption.Flag.Changed = true
	options.ConfigAddProfileNameOption.CobraParamValue = &profileName

	options.ConfigAddProfileDescriptionOption.Flag.Changed = true
	options.ConfigAddProfileDescriptionOption.CobraParamValue = &description

	options.ConfigAddProfileSetActiveOption.Flag.Changed = true
	options.ConfigAddProfileSetActiveOption.CobraParamValue = &setActive

	expectedErrorPattern := `^failed to add profile: unable to determine profile name$`
	err := RunInternalConfigAddProfile(os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigAddProfile function succeeds with set active flag set to true
func Test_RunInternalConfigAddProfile_SetActive(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("test-profile")
		description = customtypes.String("test-description")
		setActive   = customtypes.Bool(true)
	)

	options.ConfigAddProfileNameOption.Flag.Changed = true
	options.ConfigAddProfileNameOption.CobraParamValue = &profileName

	options.ConfigAddProfileDescriptionOption.Flag.Changed = true
	options.ConfigAddProfileDescriptionOption.CobraParamValue = &description

	options.ConfigAddProfileSetActiveOption.Flag.Changed = true
	options.ConfigAddProfileSetActiveOption.CobraParamValue = &setActive

	err := RunInternalConfigAddProfile(os.Stdin)
	if err != nil {
		t.Errorf("RunInternalConfigAddProfile returned error: %v", err)
	}
}

// Test RunInternalConfigAddProfile function fails with invalid set active flag
func Test_RunInternalConfigAddProfile_InvalidSetActive(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("test-profile")
		description = customtypes.String("test-description")
		setActive   = customtypes.String("invalid")
	)

	options.ConfigAddProfileNameOption.Flag.Changed = true
	options.ConfigAddProfileNameOption.CobraParamValue = &profileName

	options.ConfigAddProfileDescriptionOption.Flag.Changed = true
	options.ConfigAddProfileDescriptionOption.CobraParamValue = &description

	options.ConfigAddProfileSetActiveOption.Flag.Changed = true
	options.ConfigAddProfileSetActiveOption.CobraParamValue = &setActive

	expectedErrorPattern := `^failed to add profile: strconv.ParseBool: parsing ".*": invalid syntax$`
	err := RunInternalConfigAddProfile(os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
