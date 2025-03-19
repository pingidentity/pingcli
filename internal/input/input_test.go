// Copyright © 2025 Ping Identity Corporation

package input

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func mockValidateFunc(input string) error {
	if input == "invalid" {
		return fmt.Errorf("invalid input")
	}
	return nil
}

// Test RunPrompt function
func TestRunPrompt(t *testing.T) {
	testInput := "test-input"
	reader := testutils.WriteStringToPipe(fmt.Sprintf("%s\n", testInput), t)
	parsedInput, err := RunPrompt("test", nil, reader)
	if err != nil {
		t.Errorf("Error running RunPrompt: %v", err)
	}

	if parsedInput != testInput {
		t.Errorf("Expected '%s', but got '%s'", testInput, parsedInput)
	}
}

// Test RunPrompt function with validation
func TestRunPromptWithValidation(t *testing.T) {
	testInput := "test-input"
	reader := testutils.WriteStringToPipe(fmt.Sprintf("%s\n", testInput), t)
	parsedInput, err := RunPrompt("test", mockValidateFunc, reader)
	if err != nil {
		t.Errorf("Error running RunPrompt: %v", err)
	}

	if parsedInput != testInput {
		t.Errorf("Expected '%s', but got '%s'", testInput, parsedInput)
	}
}

// Test RunPrompt function with validation error
func TestRunPromptWithValidationError(t *testing.T) {
	testInput := "invalid"
	reader := testutils.WriteStringToPipe(fmt.Sprintf("%s\n", testInput), t)
	_, err := RunPrompt("test", mockValidateFunc, reader)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

// Test RunPromptConfirm function
func TestRunPromptConfirm(t *testing.T) {
	reader := testutils.WriteStringToPipe("y\n", t)
	parsedInput, err := RunPromptConfirm("test", reader)
	if err != nil {
		t.Errorf("Error running RunPromptConfirm: %v", err)
	}

	if !parsedInput {
		t.Errorf("Expected true, but got false")
	}
}

// Test RunPromptConfirm function with no input
func TestRunPromptConfirmNoInput(t *testing.T) {
	reader := testutils.WriteStringToPipe("\n", t)
	parsedInput, err := RunPromptConfirm("test", reader)
	if err != nil {
		t.Errorf("Error running RunPromptConfirm: %v", err)
	}

	if parsedInput {
		t.Errorf("Expected false, but got true")
	}
}

// Test RunPromptConfirm function with "n" input
func TestRunPromptConfirmNoInputN(t *testing.T) {
	reader := testutils.WriteStringToPipe("n\n", t)
	parsedInput, err := RunPromptConfirm("test", reader)
	if err != nil {
		t.Errorf("Error running RunPromptConfirm: %v", err)
	}

	if parsedInput {
		t.Errorf("Expected false, but got true")
	}
}

// Test RunPromptConfirm function with junk input
func TestRunPromptConfirmJunkInput(t *testing.T) {
	reader := testutils.WriteStringToPipe("junk\n", t)
	parsedInput, err := RunPromptConfirm("test", reader)
	if err != nil {
		t.Errorf("Error running RunPromptConfirm: %v", err)
	}

	if parsedInput {
		t.Errorf("Expected false, but got true")
	}
}

// Test RunPromptSelect function
func TestRunPromptSelect(t *testing.T) {
	testInput := "test-input"
	reader := testutils.WriteStringToPipe(fmt.Sprintf("%s\n", testInput), t)
	parsedInput, err := RunPromptSelect("test", []string{testInput}, reader)
	if err != nil {
		t.Errorf("Error running RunPromptSelect: %v", err)
	}

	if parsedInput != testInput {
		t.Errorf("Expected '%s', but got '%s'", testInput, parsedInput)
	}
}
