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

func (e *PingCLIError) Error() string {
	// Check if the wrapped error is also a PingCLIError to avoid redundant prefixes
	var err *PingCLIError
	if errors.As(e.Err, &err) {
		if strings.EqualFold(err.Prefix, e.Prefix) {
			return err.Error()
		}
	}

	return fmt.Sprintf("%s: %s", e.Prefix, e.Err.Error())
}

func (e *PingCLIError) Unwrap() error {
	return e.Err
}
