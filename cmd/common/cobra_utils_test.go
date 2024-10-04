package common_test

import (
	"testing"

	"github.com/pingidentity/pingcli/cmd/common"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/spf13/cobra"
)

// Test ExactArgs returns no error when the number of arguments matches the expected number
func TestExactArgs_Matches(t *testing.T) {
	posArgsFunc := common.ExactArgs(2)
	err := posArgsFunc(nil, []string{"arg1", "arg2"})
	testutils.CheckExpectedError(t, err, nil)
}

// Test ExactArgs returns an error when the number of arguments does not match the expected number
func TestExactArgs_DoesNotMatch(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'test': command accepts 2 arg\(s\), received 3$`
	posArgsFunc := common.ExactArgs(2)
	err := posArgsFunc(&cobra.Command{Use: "test"}, []string{"arg1", "arg2", "arg3"})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RangeArgs returns no error when the number of arguments is within the expected range
func TestRangeArgs_Matches(t *testing.T) {
	posArgsFunc := common.RangeArgs(2, 4)
	err := posArgsFunc(nil, []string{"arg1", "arg2", "arg3"})
	testutils.CheckExpectedError(t, err, nil)
}

// Test RangeArgs returns an error when the number of arguments is below the expected range
func TestRangeArgs_BelowRange(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'test': command accepts 2 to 4 arg\(s\), received 1$`
	posArgsFunc := common.RangeArgs(2, 4)
	err := posArgsFunc(&cobra.Command{Use: "test"}, []string{"arg1"})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RangeArgs returns an error when the number of arguments is above the expected range
func TestRangeArgs_AboveRange(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'test': command accepts 2 to 4 arg\(s\), received 5$`
	posArgsFunc := common.RangeArgs(2, 4)
	err := posArgsFunc(&cobra.Command{Use: "test"}, []string{"arg1", "arg2", "arg3", "arg4", "arg5"})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
