// Copyright © 2025 Ping Identity Corporation

package common

import "errors"

var (
	ErrExactArgs = errors.New("exact number of arguments not provided")
	ErrRangeArgs = errors.New("argument count not in valid range")
)
