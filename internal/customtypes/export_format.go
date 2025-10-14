// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

const (
	ENUM_EXPORT_FORMAT_HCL string = "HCL"
)

var (
	exportFormatErrorPrefix = "custom type export format error"
)

type ExportFormat string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*ExportFormat)(nil)

// Implement pflag.Value interface for custom type in cobra export-format parameter

func (ef *ExportFormat) Set(format string) error {
	if ef == nil {
		return &errs.PingCLIError{Prefix: exportFormatErrorPrefix, Err: ErrCustomTypeNil}
	}

	switch {
	case strings.EqualFold(format, ENUM_EXPORT_FORMAT_HCL):
		*ef = ExportFormat(ENUM_EXPORT_FORMAT_HCL)
	case strings.EqualFold(format, ""):
		*ef = ExportFormat("")
	default:
		return &errs.PingCLIError{Prefix: exportFormatErrorPrefix, Err: fmt.Errorf("%w '%s': must be one of %s", ErrUnrecognisedFormat, format, strings.Join(ExportFormatValidValues(), ", "))}
	}

	return nil
}

func (ef *ExportFormat) Type() string {
	return "string"
}

func (ef *ExportFormat) String() string {
	if ef == nil {
		return ""
	}

	return string(*ef)
}

func ExportFormatValidValues() []string {
	exportFormats := []string{
		ENUM_EXPORT_FORMAT_HCL,
	}

	slices.Sort(exportFormats)

	return exportFormats
}
