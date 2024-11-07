package config_internal

import (
	"os"
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// Test RunInternalConfigSetActiveProfile function
func Test_RunInternalConfigSetActiveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigSetActiveProfile([]string{"production"}, os.Stdin)
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigSetActiveProfile function fails with invalid profile name
func Test_RunInternalConfigSetActiveProfile_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^failed to set active profile: invalid profile name: '.*'\. name must contain only alphanumeric characters, underscores, and dashes$`
	err := RunInternalConfigSetActiveProfile([]string{"(*#&)"}, os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSetActiveProfile function fails with non-existent profile
func Test_RunInternalConfigSetActiveProfile_NonExistentProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^failed to set active profile: invalid profile name: '.*' profile does not exist$`
	err := RunInternalConfigSetActiveProfile([]string{"non-existent"}, os.Stdin)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
