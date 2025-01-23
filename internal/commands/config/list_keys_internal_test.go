package config_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// Test RunInternalConfigListKeys function
func Test_RunInternalConfigListKeys(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigListKeys()
	testutils.CheckExpectedError(t, err, nil)
}
