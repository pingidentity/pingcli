// Copyright © 2025 Ping Identity Corporation

package errs

import (
	"errors"
	"fmt"
	"strings"
)

type PingCLIError struct {
	Err    error
	Prefix string
}

var (
	ErrInvalidInput = errors.New("invalid input")
)

func (e *PingCLIError) Error() string {
	if e == nil || e.Err == nil {
		return ""
	}

	// Check if the wrapped error is also a PingCLIError to avoid redundant prefixes
	var pingErr *PingCLIError
	if errors.As(e.Err, &pingErr) {
		if strings.EqualFold(pingErr.Prefix, e.Prefix) {
			return pingErr.Error()
		}
	}

	// if both are empty this just returns an empty string anyways
	errMsg := e.Err.Error()
	if e.Prefix == "" {
		return errMsg
	}
	if errMsg == "" {
		return e.Prefix
	}

	return fmt.Sprintf("%s: %s", e.Prefix, errMsg)
}

func (e *PingCLIError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Err
}
