// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

const (
	ENUM_EXPORT_FORMAT_HCL string = "HCL"
)

type ExportFormat string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*ExportFormat)(nil)

// Implement pflag.Value interface for custom type in cobra export-format parameter

func (ef *ExportFormat) Set(format string) error {
	if ef == nil {
		return fmt.Errorf("failed to set Export Format value: %s. Export Format is nil", format)
	}

	switch {
	case strings.EqualFold(format, ENUM_EXPORT_FORMAT_HCL):
		*ef = ExportFormat(ENUM_EXPORT_FORMAT_HCL)
	case strings.EqualFold(format, ""):
		*ef = ExportFormat("")
	default:
		return fmt.Errorf("unrecognized export format '%s'. Must be one of: %s", format, strings.Join(ExportFormatValidValues(), ", "))
	}
	return nil
}

func (ef ExportFormat) Type() string {
	return "string"
}

func (ef ExportFormat) String() string {
	return string(ef)
}

func ExportFormatValidValues() []string {
	exportFormats := []string{
		ENUM_EXPORT_FORMAT_HCL,
	}

	slices.Sort(exportFormats)

	return exportFormats
}
