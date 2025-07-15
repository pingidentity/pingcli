// Copyright © 2025 Ping Identity Corporation

package license_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// Test License Command Executes without issue (with all required flags)
func TestLicenseCmd_Execute(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := testutils_cobra.ExecutePingcli(t, "license",
		"--"+options.LicenseProductOption.CobraParamName, customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE,
		"--"+options.LicenseVersionOption.CobraParamName, "12.0")
	testutils.CheckExpectedError(t, err, nil)
}

// Test License Command fails when provided too many arguments
func TestLicenseCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli license': command accepts 0 arg\(s\), received 1$`
	err := testutils_cobra.ExecutePingcli(t, "license", "extra-arg")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test License Command help flag
func TestLicenseCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "license", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "license", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// Test License Command fails with invalid flag
func TestLicenseCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingcli(t, "license", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test License Command fails when required product flag is missing
func TestLicenseCmd_MissingProductFlag(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `^required flag\(s\) "product" not set$`
	err := testutils_cobra.ExecutePingcli(t, "license",
		"--"+options.LicenseVersionOption.CobraParamName, "12.0")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test License Command fails when required version flag is missing
func TestLicenseCmd_MissingVersionFlag(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `^required flag\(s\) "version" not set$`
	err := testutils_cobra.ExecutePingcli(t, "license",
		"--"+options.LicenseProductOption.CobraParamName, "pingfederate")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test License Command with shorthand flags
func TestLicenseCmd_ShorthandFlags(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := testutils_cobra.ExecutePingcli(t, "license",
		"-"+options.LicenseProductOption.Flag.Shorthand, "pingfederate",
		"-"+options.LicenseVersionOption.Flag.Shorthand, "12.0")
	testutils.CheckExpectedError(t, err, nil)
}

// Test License Command with a profile
func TestLicenseCmd_Profile(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := testutils_cobra.ExecutePingcli(t, "license",
		"--"+options.LicenseProductOption.CobraParamName, "pingfederate",
		"--"+options.LicenseVersionOption.CobraParamName, "12.0",
		"--"+options.RootProfileOption.CobraParamName, "default")
	testutils.CheckExpectedError(t, err, nil)
}
