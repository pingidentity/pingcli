// Copyright © 2026 Ping Identity Corporation

package pingone

import "errors"

var (
	ErrUnknownExtractionFunction = errors.New("failed to find extraction function")
	ErrEmbeddedEmpty             = errors.New("failed to get reflect value from embedded. embedded is empty")
	ErrCastReflectValue          = errors.New("failed to cast reflect value")
)
