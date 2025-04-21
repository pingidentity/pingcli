// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// Test RunInternalConfigListProfiles function
func Test_RunInternalConfigListProfiles(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := RunInternalConfigListProfiles()
	testutils.CheckExpectedError(t, err, nil)
}
