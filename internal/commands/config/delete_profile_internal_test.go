// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// Test deleteProfile function
func Test_deleteProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := deleteProfile("production")
	testutils.CheckExpectedError(t, err, nil)
}

// Test deleteProfile function fails with active profile
func Test_deleteProfile_ActiveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^'.*' is the active profile and cannot be deleted$`
	err := deleteProfile("default")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test deleteProfile function fails with invalid profile name
func Test_deleteProfile_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^invalid profile name: '.*' profile does not exist$`
	err := deleteProfile("(*#&)")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test deleteProfile function fails with empty profile name
func Test_deleteProfile_EmptyProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^invalid profile name: profile name cannot be empty$`
	err := deleteProfile("")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test deleteProfile function fails with non-existent profile name
func Test_deleteProfile_NonExistentProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^invalid profile name: '.*' profile does not exist$`
	err := deleteProfile("non-existent")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
