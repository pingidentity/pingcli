// Copyright © 2025 Ping Identity Corporation

package completion_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Completion Command Executes without issue
func TestCompletionCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t)
	testutils.CheckExpectedError(t, err, nil)
}
