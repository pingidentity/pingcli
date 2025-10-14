// Copyright © 2025 Ping Identity Corporation

package grpc

import "errors"

var (
	ErrNilConfiguration = errors.New("plugin returned a nil configuration")
)
