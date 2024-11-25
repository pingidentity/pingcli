package config_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// Test RunInternalConfigGet function
func Test_RunInternalConfigGet(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigGet("service")
	if err != nil {
		t.Errorf("RunInternalConfigGet returned error: %v", err)
	}
}

// Test RunInternalConfigGet function fails with invalid key
func Test_RunInternalConfigGet_InvalidKey(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `(?s)^failed to get configuration: key '.*' is not recognized as a valid configuration key. Valid keys: .*$`
	err := RunInternalConfigGet("invalid-key")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigGet function with different profile
func Test_RunInternalConfigGet_DifferentProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("production")
	)

	options.RootProfileOption.Flag.Changed = true
	options.RootProfileOption.CobraParamValue = &profileName

	err := RunInternalConfigGet("service")
	if err != nil {
		t.Errorf("RunInternalConfigGet returned error: %v", err)
	}
}

// Test RunInternalConfigGet function fails with invalid profile name
func Test_RunInternalConfigGet_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("invalid")
	)

	options.RootProfileOption.Flag.Changed = true
	options.RootProfileOption.CobraParamValue = &profileName

	expectedErrorPattern := `^failed to get configuration: invalid profile name: '.*' profile does not exist$`
	err := RunInternalConfigGet("service")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
