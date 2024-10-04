package feedback_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Feedback Command Executes without issue
func TestFeedbackCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "feedback")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Feedback Command fails when provided too many arguments
func TestFeedbackCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli feedback': command accepts 0 arg\(s\), received 1$`
	err := testutils_cobra.ExecutePingcli(t, "feedback", "extra-arg")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Feedback Command help flag
func TestFeedbackCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "feedback", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "feedback", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Feedback Command fails with invalid flag
func TestFeedbackCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingcli(t, "feedback", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
