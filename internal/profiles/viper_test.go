// Copyright © 2025 Ping Identity Corporation

package profiles_test

import (
	"slices"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
	"github.com/spf13/viper"
)

// Test ChangeActiveProfile function
func Test_ChangeActiveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	mainConfig := profiles.GetMainConfig()

	err := mainConfig.ChangeActiveProfile("production")
	if err != nil {
		t.Errorf("ChangeActiveProfile returned error: %v", err)
	}

	activeProfile, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		t.Errorf("GetOptionValue returned error: %v", err)
	}

	if activeProfile != "production" {
		t.Errorf("activeProfile is %s, expected %s", activeProfile, "production")
	}
}

// Test ChangeActiveProfile function with invalid profile
func Test_ChangeActiveProfile_InvalidProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	mainConfig := profiles.GetMainConfig()

	expectedErrorPattern := `^invalid profile name: '.*' profile does not exist$`
	err := mainConfig.ChangeActiveProfile("invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ChangeProfileName function
func Test_ChangeProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	mainConfig := profiles.GetMainConfig()

	err := mainConfig.ChangeProfileName("production", "new")
	if err != nil {
		t.Errorf("ChangeProfileName returned error: %v", err)
	}

	profiles := mainConfig.ProfileNames()
	if len(profiles) != 2 {
		t.Errorf("profiles length is %d, expected %d", len(profiles), 2)
	}

	if slices.Contains(profiles, "production") {
		t.Errorf("profiles contains production, expected it to be removed")
	}

	if !slices.Contains(profiles, "new") {
		t.Errorf("profiles does not contain new, expected it to be added")
	}
}

// Test ChangeProfileName function with same profile name
func Test_ChangeProfileName_SameProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	mainConfig := profiles.GetMainConfig()

	err := mainConfig.ChangeProfileName("production", "production")
	if err != nil {
		t.Errorf("ChangeProfileName returned error: %v", err)
	}
}

// Test ChangeProfileName function with invalid old profile name
func Test_ChangeProfileName_InvalidOldProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	mainConfig := profiles.GetMainConfig()

	expectedErrorPattern := `^invalid profile name: '.*' profile does not exist$`
	err := mainConfig.ChangeProfileName("invalid", "new")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ChangeProfileName function with invalid new profile name
func Test_ChangeProfileName_InvalidNewProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	mainConfig := profiles.GetMainConfig()

	expectedErrorPattern := `^invalid profile name: '.*'. profile already exists$`
	err := mainConfig.ChangeProfileName("production", "default")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test DeleteProfile function
func Test_DeleteProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	mainConfig := profiles.GetMainConfig()

	err := mainConfig.DeleteProfile("production")
	if err != nil {
		t.Errorf("DeleteProfile returned error: %v", err)
	}

	profiles := mainConfig.ProfileNames()
	if len(profiles) != 1 {
		t.Errorf("profiles length is %d, expected %d", len(profiles), 0)
	}

	if slices.Contains(profiles, "production") {
		t.Errorf("profiles contains production, expected it to be removed")
	}
}

// Test DeleteProfile function with invalid profile name
func Test_DeleteProfile_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	mainConfig := profiles.GetMainConfig()

	expectedErrorPattern := `^invalid profile name: '.*' profile does not exist$`
	err := mainConfig.DeleteProfile("invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test DeleteProfile function with active profile
func Test_DeleteProfile_ActiveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	mainConfig := profiles.GetMainConfig()

	expectedErrorPattern := `^'.*' is the active profile and cannot be deleted$`
	err := mainConfig.DeleteProfile("default")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test SaveProfile function
func Test_SaveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	mainConfig := profiles.GetMainConfig()

	subViper := viper.New()
	subViper.Set("description", "new description")

	err := mainConfig.SaveProfile("new", subViper)
	if err != nil {
		t.Errorf("SaveProfile returned error: %v", err)
	}

	profiles := mainConfig.ProfileNames()
	if len(profiles) != 3 {
		t.Errorf("profiles length is %d, expected %d", len(profiles), 3)
	}

	if !slices.Contains(profiles, "new") {
		t.Errorf("profiles does not contain new, expected it to be added")
	}
}

// Test SaveProfile function with existing profile name
func Test_SaveProfile_ExistingProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	mainConfig := profiles.GetMainConfig()

	subViper := viper.New()
	subViper.Set("description", "new description")

	err := mainConfig.SaveProfile("production", subViper)
	testutils.CheckExpectedError(t, err, nil)

	profiles := mainConfig.ProfileNames()
	if len(profiles) != 2 {
		t.Errorf("profiles length is %d, expected %d", len(profiles), 2)
	}

	if !slices.Contains(profiles, "production") {
		t.Errorf("profiles does not contain production, expected it to be added")
	}

	actual := mainConfig.ViperInstance().Get("production.description")
	expected := "new description"

	if actual != expected {
		t.Errorf("description is %s, expected %s", actual, expected)
	}
}
