// Copyright © 2026 Ping Identity Corporation

package autocompletion

import "errors"

var (
	// Common autocompletion errors
	ErrGetConfiguration = errors.New("unable to get configuration")
	ErrGetActiveProfile = errors.New("unable to get active profile")
)

const (
	autocompletionErrorPrefix = "autocompletion failed"
)
