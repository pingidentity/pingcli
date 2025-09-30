package customtypes

import "errors"

var (
	ErrCustomTypeNil = errors.New("failed to set value. custom type is nil")
)
