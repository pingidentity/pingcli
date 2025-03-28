// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// Test RunInternalConfigViewProfile function
func Test_RunInternalConfigViewProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigViewProfile([]string{})
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigViewProfile function fails with invalid profile name
func Test_RunInternalConfigViewProfile_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^failed to view profile: invalid profile name: '.*' profile does not exist$`
	err := RunInternalConfigViewProfile([]string{"invalid"})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigViewProfile function with different profile
func Test_RunInternalConfigViewProfile_DifferentProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigViewProfile([]string{"production"})
	testutils.CheckExpectedError(t, err, nil)
}
