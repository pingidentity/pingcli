// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"regexp"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

var (
	licenseVersionErrorPrefix = "custom type license version error"
)

type LicenseVersion string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*LicenseVersion)(nil)

// Implement pflag.Value interface for custom type in cobra MultiService parameter
func (lv *LicenseVersion) Set(version string) error {
	if lv == nil {
		return &errs.PingCLIError{Prefix: licenseVersionErrorPrefix, Err: ErrCustomTypeNil}
	}

	// The license version must be of the form "major.minor" or empty
	if version == "" {
		*lv = LicenseVersion("")

		return nil
	}

	// Validate the format of the version string via regex
	if !regexp.MustCompile(`^\d+\.\d+$`).MatchString(version) {
		return &errs.PingCLIError{Prefix: licenseVersionErrorPrefix, Err: fmt.Errorf("%w: %s. Example: '12.3'", ErrInvalidVersionFormat, version)}
	}

	*lv = LicenseVersion(version)

	return nil
}

func (lv *LicenseVersion) Type() string {
	return "string"
}

func (lv *LicenseVersion) String() string {
	if lv == nil {
		return ""
	}

	return string(*lv)
}
