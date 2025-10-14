// Copyright © 2025 Ping Identity Corporation

package common

import "errors"

var (
	ErrParseHCLTemplate         = errors.New("failed to parse HCL import block template")
	ErrExportResource           = errors.New("failed to export resource")
	ErrFileAlreadyExists        = errors.New("generated import file already exists. use --overwrite to overwrite existing export data")
	ErrCreateExportFile         = errors.New("failed to create export file")
	ErrWriteTemplateToFile      = errors.New("failed to write import block template to file")
	ErrUnrecognizedExportFormat = errors.New("unrecognized export format")
	ErrResourceRequestFailed    = errors.New("resource request was not successful")
	ErrExportResources          = errors.New("failed to export resource")
)
