// Copyright © 2025 Ping Identity Corporation

package configuration

import "errors"

var (
	ErrInvalidConfigurationKey = errors.New("provided key is not recognized as a valid configuration key.\nuse 'pingcli config list-keys' to view all available keys")
	ErrNoOptionForKey          = errors.New("no option found for the provided configuration key")
	ErrEmptyKeyForOptionSearch = errors.New("empty key provided for option search, too many matches with options not configured with a koanf key")
)
