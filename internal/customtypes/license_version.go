// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"regexp"

	"github.com/spf13/pflag"
)

type LicenseVersion string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*LicenseVersion)(nil)

// Implement pflag.Value interface for custom type in cobra MultiService parameter
func (lp *LicenseVersion) Set(version string) error {
	if lp == nil {
		return fmt.Errorf("failed to set LicenseVersion value: %s. LicenseVersion is nil", version)
	}

	// The license version must be of the form "major.minor" or empty
	if version == "" {
		*lp = LicenseVersion("")

		return nil
	}

	// Validate the format of the version string via regex
	if !regexp.MustCompile(`^\d+\.\d+$`).MatchString(version) {
		return fmt.Errorf("failed to set LicenseVersion value: %s. Invalid version format, must be 'major.minor'. Example: '12.3'", version)
	}

	*lp = LicenseVersion(version)

	return nil
}

func (lp LicenseVersion) Type() string {
	return "string"
}

func (lp LicenseVersion) String() string {
	return string(lp)
}
