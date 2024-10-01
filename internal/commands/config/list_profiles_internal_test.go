package config_internal

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test RunInternalConfigListProfiles function
func Test_RunInternalConfigListProfiles(t *testing.T) {
	testutils_viper.InitVipers(t)

	RunInternalConfigListProfiles()
}
