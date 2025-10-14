// Copyright © 2025 Ping Identity Corporation

package plugin_internal

import "errors"

var (
	ErrPluginNameEmpty       = errors.New("plugin executable name is empty")
	ErrReadPluginNamesConfig = errors.New("failed to read configured plugin executable names")
	ErrUndeterminedProfile   = errors.New("unable to determine configuration profile")
	ErrPluginAlreadyExists   = errors.New("plugin executable already exists in configuration")
	ErrPluginNotFound        = errors.New("plugin executable not found in system PATH")
)
