// Copyright © 2025 Ping Identity Corporation

package input

import (
	"errors"
	"fmt"
	"io"

	"github.com/manifoldco/promptui"
)

type InputPromptError struct {
	Err error
}

func (e *InputPromptError) Error() string {
	return fmt.Sprintf("input prompt failed: %s", e.Err.Error())
}

func (e *InputPromptError) Unwrap() error {
	return e.Err
}

func RunPrompt(message string, validateFunc func(string) error, rc io.ReadCloser) (string, error) {
	p := promptui.Prompt{
		Label:    message,
		Validate: validateFunc,
		Stdin:    rc,
	}

	userInput, err := p.Run()
	if err != nil {
		return "", &InputPromptError{Err: err}
	}

	return userInput, nil
}

func RunPromptConfirm(message string, rc io.ReadCloser) (bool, error) {
	p := promptui.Prompt{
		Label:     message,
		IsConfirm: true,
		Stdin:     rc,
	}

	// This is odd behavior discussed in https://github.com/manifoldco/promptui/issues/81
	// If err is type promptui.ErrAbort, the user can be assumed to have responded "No"
	_, err := p.Run()
	if err != nil {
		if errors.Is(err, promptui.ErrAbort) {
			return false, nil
		}

		return false, &InputPromptError{Err: err}
	}

	return true, nil
}

func RunPromptSelect(message string, items []string, rc io.ReadCloser) (selection string, err error) {
	p := promptui.Select{
		Label: message,
		Items: items,
		Size:  len(items),
		Stdin: rc,
	}

	_, selection, err = p.Run()
	if err != nil {
		return "", &InputPromptError{Err: err}
	}

	return selection, nil
}
