// Copyright © 2026 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

const (
	ENUM_OUTPUT_FORMAT_TEXT string = "text"
	ENUM_OUTPUT_FORMAT_JSON string = "json"
)

var (
	outputFormatErrorPrefix = "custom type output format error"
)

type OutputFormat string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*OutputFormat)(nil)

// Implement pflag.Value interface for custom type in cobra pingcli-output parameter

func (o *OutputFormat) Set(outputFormat string) error {
	if o == nil {
		return &errs.PingCLIError{Prefix: outputFormatErrorPrefix, Err: ErrCustomTypeNil}
	}

	switch {
	case strings.EqualFold(outputFormat, ENUM_OUTPUT_FORMAT_TEXT):
		*o = OutputFormat(ENUM_OUTPUT_FORMAT_TEXT)
	case strings.EqualFold(outputFormat, ENUM_OUTPUT_FORMAT_JSON):
		*o = OutputFormat(ENUM_OUTPUT_FORMAT_JSON)
	case strings.EqualFold(outputFormat, ""):
		*o = OutputFormat("")
	default:
		return &errs.PingCLIError{Prefix: outputFormatErrorPrefix, Err: fmt.Errorf("%w: '%s'. Must be one of: %s", ErrUnrecognizedOutputFormat, outputFormat, strings.Join(OutputFormatValidValues(), ", "))}
	}

	return nil
}

func (o *OutputFormat) Type() string {
	return "string"
}

func (o *OutputFormat) String() string {
	if o == nil {
		return ""
	}

	return string(*o)
}

func OutputFormatValidValues() []string {
	outputFormats := []string{
		ENUM_OUTPUT_FORMAT_TEXT,
		ENUM_OUTPUT_FORMAT_JSON,
	}

	slices.Sort(outputFormats)

	return outputFormats
}
