package config_internal

import "errors"

var (
	ErrUndeterminedProfile = errors.New("unable to determine configuration profile")
)
