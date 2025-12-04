// Copyright © 2025 Ping Identity Corporation

package input

import (
	"errors"
	"io"

	"github.com/manifoldco/promptui"
	"github.com/pingidentity/pingcli/internal/errs"
)

var (
	inputPromptErrorPrefix = "input prompt error"
)

func RunPrompt(message string, validateFunc func(string) error, rc io.ReadCloser) (string, error) {
	// Submit-only validation: run prompt without live Validate, then validate after submit.
	for {
		p := promptui.Prompt{
			Label: message,
			Stdin: rc,
		}

		userInput, err := p.Run()
		if err != nil {
			return "", &errs.PingCLIError{Prefix: inputPromptErrorPrefix, Err: err}
		}

		if validateFunc != nil {
			if vErr := validateFunc(userInput); vErr != nil {
				//nolint:all
				_ = vErr
				continue
			}
		}

		return userInput, nil
	}
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

		return false, &errs.PingCLIError{Prefix: inputPromptErrorPrefix, Err: err}
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
		return "", &errs.PingCLIError{Prefix: inputPromptErrorPrefix, Err: err}
	}

	return selection, nil
}
