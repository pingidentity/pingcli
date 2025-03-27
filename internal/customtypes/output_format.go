// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

const (
	ENUM_OUTPUT_FORMAT_TEXT string = "text"
	ENUM_OUTPUT_FORMAT_JSON string = "json"
)

type OutputFormat string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*OutputFormat)(nil)

// Implement pflag.Value interface for custom type in cobra pingcli-output parameter

func (o *OutputFormat) Set(outputFormat string) error {
	if o == nil {
		return fmt.Errorf("failed to set Output Format value: %s. Output Format is nil", outputFormat)
	}

	switch {
	case strings.EqualFold(outputFormat, ENUM_OUTPUT_FORMAT_TEXT):
		*o = OutputFormat(ENUM_OUTPUT_FORMAT_TEXT)
	case strings.EqualFold(outputFormat, ENUM_OUTPUT_FORMAT_JSON):
		*o = OutputFormat(ENUM_OUTPUT_FORMAT_JSON)
	case strings.EqualFold(outputFormat, ""):
		*o = OutputFormat("")
	default:
		return fmt.Errorf("unrecognized Output Format: '%s'. Must be one of: %s", outputFormat, strings.Join(OutputFormatValidValues(), ", "))
	}

	return nil
}

func (o OutputFormat) Type() string {
	return "string"
}

func (o OutputFormat) String() string {
	return string(o)
}

func OutputFormatValidValues() []string {
	outputFormats := []string{
		ENUM_OUTPUT_FORMAT_TEXT,
		ENUM_OUTPUT_FORMAT_JSON,
	}

	slices.Sort(outputFormats)

	return outputFormats
}
