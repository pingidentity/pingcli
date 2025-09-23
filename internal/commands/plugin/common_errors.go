package plugin_internal

import "errors"

var (
	ErrPluginNameEmpty       = errors.New("plugin executable name is empty")
	ErrReadPluginNamesConfig = errors.New("failed to read configured plugin executable names")
	ErrUndeterminedProfile   = errors.New("unable to determine configuration profile")
)
