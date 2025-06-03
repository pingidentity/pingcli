// Copyright © 2025 Ping Identity Corporation

package plugin_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// Test RunInternalPluginList function
func Test_RunInternalPluginList(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := RunInternalPluginList()
	if err != nil {
		t.Errorf("RunInternalPluginList returned error: %v", err)
	}
}
