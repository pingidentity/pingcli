package config_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// Test RunInternalConfigViewProfile function
func Test_RunInternalConfigViewProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigViewProfile()
	if err != nil {
		t.Errorf("RunInternalConfigViewProfile returned error: %v", err)
	}
}

// Test RunInternalConfigViewProfile function fails with invalid profile name
func Test_RunInternalConfigViewProfile_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("invalid")
	)

	options.ConfigViewProfileOption.Flag.Changed = true
	options.ConfigViewProfileOption.CobraParamValue = &profileName

	expectedErrorPattern := `^failed to view profile: invalid profile name: '.*' profile does not exist$`
	err := RunInternalConfigViewProfile()
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigViewProfile function with different profile
func Test_RunInternalConfigViewProfile_DifferentProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("production")
	)

	options.ConfigViewProfileOption.Flag.Changed = true
	options.ConfigViewProfileOption.CobraParamValue = &profileName

	err := RunInternalConfigViewProfile()
	if err != nil {
		t.Errorf("RunInternalConfigViewProfile returned error: %v", err)
	}
}
