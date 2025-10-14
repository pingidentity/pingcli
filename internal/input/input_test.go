// Copyright © 2025 Ping Identity Corporation

package input_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/manifoldco/promptui"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/stretchr/testify/require"
)

var (
	errInvalidInput = errors.New("invalid input")
)

func mockValidateFunc(input string) error {
	if input == "invalid" {
		return errInvalidInput
	}

	return nil
}

func TestRunPrompt(t *testing.T) {
	testInput := "test-input"
	reader := testutils.WriteStringToPipe(t, fmt.Sprintf("%s\n", testInput))

	parsedInput, err := input.RunPrompt("test", nil, reader)
	require.NoError(t, err)
	require.Equal(t, testInput, parsedInput)
}

func TestRunPromptWithValidation(t *testing.T) {
	testInput := "test-input"
	reader := testutils.WriteStringToPipe(t, fmt.Sprintf("%s\n", testInput))

	parsedInput, err := input.RunPrompt("test", mockValidateFunc, reader)
	require.NoError(t, err)
	require.Equal(t, testInput, parsedInput)
}

func TestRunPromptWithValidationError(t *testing.T) {
	testInput := "invalid"
	reader := testutils.WriteStringToPipe(t, fmt.Sprintf("%s\n", testInput))

	_, err := input.RunPrompt("test", mockValidateFunc, reader)
	require.Error(t, err)

	var pingErr *errs.PingCLIError
	require.ErrorAs(t, err, &pingErr)
	require.ErrorIs(t, err, promptui.ErrEOF)
}

func TestRunPromptConfirm(t *testing.T) {
	reader := testutils.WriteStringToPipe(t, "y\n")

	parsedInput, err := input.RunPromptConfirm("test", reader)
	require.NoError(t, err)
	require.True(t, parsedInput)
}

func TestRunPromptConfirmNoInput(t *testing.T) {
	reader := testutils.WriteStringToPipe(t, "\n")

	parsedInput, err := input.RunPromptConfirm("test", reader)
	require.NoError(t, err)
	require.False(t, parsedInput)
}

func TestRunPromptConfirmNoInputN(t *testing.T) {
	reader := testutils.WriteStringToPipe(t, "n\n")

	parsedInput, err := input.RunPromptConfirm("test", reader)
	require.NoError(t, err)
	require.False(t, parsedInput)
}

func TestRunPromptConfirmJunkInput(t *testing.T) {
	reader := testutils.WriteStringToPipe(t, "junk\n")

	parsedInput, err := input.RunPromptConfirm("test", reader)
	require.NoError(t, err)
	require.False(t, parsedInput)
}

func TestRunPromptSelect(t *testing.T) {
	testInput := "test-input"
	reader := testutils.WriteStringToPipe(t, fmt.Sprintf("%s\n", testInput))

	parsedInput, err := input.RunPromptSelect("test", []string{testInput}, reader)
	require.NoError(t, err)
	require.Equal(t, testInput, parsedInput)
}

func TestRunPromptSelectError(t *testing.T) {
	reader := testutils.WriteStringToPipe(t, "\x03") // Simulate Ctrl+C

	_, err := input.RunPromptSelect("test", []string{"test-input"}, reader)
	require.Error(t, err)

	var pingErr *errs.PingCLIError
	require.ErrorAs(t, err, &pingErr)
	require.ErrorIs(t, err, promptui.ErrInterrupt)
}
