// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// Test RunInternalConfigListKeys function
func Test_RunInternalConfigListKeys(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := RunInternalConfigListKeys()
	testutils.CheckExpectedError(t, err, nil)
}
