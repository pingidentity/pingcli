// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// Test RunInternalConfigUnset function
func Test_RunInternalConfigUnset(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := RunInternalConfigUnset("noColor")
	if err != nil {
		t.Errorf("RunInternalConfigUnset returned error: %v", err)
	}
}

// Test RunInternalConfigUnset function fails with invalid key
func Test_RunInternalConfigUnset_InvalidKey(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `^failed to unset configuration: key '.*' is not recognized as a valid configuration key.\s*Use 'pingcli config list-keys' to view all available keys`
	err := RunInternalConfigUnset("invalid-key")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigUnset function with different profile
func Test_RunInternalConfigUnset_DifferentProfile(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	var (
		profileName = customtypes.String("production")
	)

	options.RootProfileOption.Flag.Changed = true
	options.RootProfileOption.CobraParamValue = &profileName

	err := RunInternalConfigUnset("noColor")
	if err != nil {
		t.Errorf("RunInternalConfigUnset returned error: %v", err)
	}
}

// Test RunInternalConfigUnset function fails with invalid profile name
func Test_RunInternalConfigUnset_InvalidProfileName(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	var (
		profileName = customtypes.String("invalid")
	)

	options.RootProfileOption.Flag.Changed = true
	options.RootProfileOption.CobraParamValue = &profileName

	expectedErrorPattern := `^failed to unset configuration: invalid profile name: '.*' profile does not exist$`
	err := RunInternalConfigUnset("noColor")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigUnset function succeeds with case-insensitive key
func Test_RunInternalConfigUnset_CaseInsensitiveKeys(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := RunInternalConfigUnset("NoCoLoR")
	if err != nil {
		t.Errorf("RunInternalConfigUnset returned error: %v", err)
	}

	//Make sure the actual correct key was unset, not the case-insensitive one
	vVal, err := profiles.GetOptionValue(options.RootColorOption)
	if err != nil {
		t.Errorf("GetOptionValue returned error: %v", err)
	}

	if vVal != options.RootColorOption.DefaultValue.String() {
		t.Errorf("Expected %s to be %s, got %v", options.RootColorOption.KoanfKey, options.RootColorOption.DefaultValue.String(), vVal)
	}
}
