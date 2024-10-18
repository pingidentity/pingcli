package config_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// Test RunInternalConfigSet function
func Test_RunInternalConfigSet(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigSet("noColor=true")
	if err != nil {
		t.Errorf("RunInternalConfigSet returned error: %v", err)
	}
}

// Test RunInternalConfigSet function fails with invalid key
func Test_RunInternalConfigSet_InvalidKey(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^failed to set configuration: key '.*' is not recognized as a valid configuration key. Valid keys: .*$`
	err := RunInternalConfigSet("invalid-key=false")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSet function fails with invalid value
func Test_RunInternalConfigSet_InvalidValue(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^failed to set configuration: value for key '.*' must be a boolean. Allowed .*: strconv.ParseBool: parsing ".*": invalid syntax$`
	err := RunInternalConfigSet("noColor=invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSet function fails with non-existent profile name
func Test_RunInternalConfigSet_NonExistentProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("non-existent")
	)

	options.ConfigSetProfileOption.Flag.Changed = true
	options.ConfigSetProfileOption.CobraParamValue = &profileName

	expectedErrorPattern := `^failed to set configuration: invalid profile name: '.*' profile does not exist$`
	err := RunInternalConfigSet("noColor=true")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSet function with different profile
func Test_RunInternalConfigSet_DifferentProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("production")
	)

	options.ConfigSetProfileOption.Flag.Changed = true
	options.ConfigSetProfileOption.CobraParamValue = &profileName

	err := RunInternalConfigSet("noColor=true")
	if err != nil {
		t.Errorf("RunInternalConfigSet returned error: %v", err)
	}
}

// Test RunInternalConfigSet function fails with invalid profile name
func Test_RunInternalConfigSet_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	var (
		profileName = customtypes.String("*&%*&")
	)

	options.ConfigSetProfileOption.Flag.Changed = true
	options.ConfigSetProfileOption.CobraParamValue = &profileName

	expectedErrorPattern := `^failed to set configuration: invalid profile name: '.*'\. name must contain only alphanumeric characters, underscores, and dashes$`
	err := RunInternalConfigSet("noColor=true")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSet function fails with no value provided
func Test_RunInternalConfigSet_NoValue(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^failed to set configuration: value for key '.*' is empty. Use 'pingcli config unset .*' to unset the key$`
	err := RunInternalConfigSet("noColor=")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSet function fails with no keyValue provided
func Test_RunInternalConfigSet_NoKeyValue(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^failed to set configuration: invalid assignment format ''\. Expect 'key=value' format$`
	err := RunInternalConfigSet("")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
