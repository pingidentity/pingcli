package config_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// Test RunInternalConfigUnset function
func Test_RunInternalConfigUnset(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigUnset("color")
	if err != nil {
		t.Errorf("RunInternalConfigUnset returned error: %v", err)
	}
}

// Test RunInternalConfigUnset function fails with invalid key
func Test_RunInternalConfigUnset_InvalidKey(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^failed to unset configuration: key '.*' is not recognized as a valid configuration key. Valid keys: .*$`
	err := RunInternalConfigUnset("invalid-key")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigUnset function with different profile
func Test_RunInternalConfigUnset_DifferentProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("production")
	)

	options.ConfigUnsetProfileOption.Flag.Changed = true
	options.ConfigUnsetProfileOption.CobraParamValue = &profileName

	err := RunInternalConfigUnset("color")
	if err != nil {
		t.Errorf("RunInternalConfigUnset returned error: %v", err)
	}
}

// Test RunInternalConfigUnset function fails with invalid profile name
func Test_RunInternalConfigUnset_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("invalid")
	)

	options.ConfigUnsetProfileOption.Flag.Changed = true
	options.ConfigUnsetProfileOption.CobraParamValue = &profileName

	expectedErrorPattern := `^failed to unset configuration: invalid profile name: '.*' profile does not exist$`
	err := RunInternalConfigUnset("color")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
