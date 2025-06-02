// Copyright © 2025 Ping Identity Corporation

package plugin_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// Test RunInternalConfigGet function
func Test_RunInternalConfigGet(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := RunInternalPluginAdd("pingcli-feedback-plugin")
	if err != nil {
		t.Errorf("RunInternalConfigGet returned error: %v", err)
	}
}
