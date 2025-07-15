// Copyright © 2025 Ping Identity Corporation

package license_internal

import (
	"os"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// setLicenseProductAndVersion sets up product and version options
func setLicenseProductAndVersion(product, version string) {
	if product != "" {
		productVal := customtypes.LicenseProduct(product)
		options.LicenseProductOption.CobraParamValue = &productVal
		options.LicenseProductOption.Flag.Changed = true
	}

	if version != "" {
		versionVal := customtypes.LicenseVersion(version)
		options.LicenseVersionOption.CobraParamValue = &versionVal
		options.LicenseVersionOption.Flag.Changed = true
	}
}

// Test RunInternalLicense function with valid options
func Test_RunInternalLicense_Success(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	setLicenseProductAndVersion(customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE, "13.0")

	err := RunInternalLicense()
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalLicense with missing product option
func Test_RunInternalLicense_MissingProduct(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Set up test data with missing product
	setLicenseProductAndVersion("", "13.0")

	// Run the function
	expectedErrorPattern := `^failed to run license request: product, version, devops user, and devops key must be specified for license request$`
	err := RunInternalLicense()
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalLicense with missing version option
func Test_RunInternalLicense_MissingVersion(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	setLicenseProductAndVersion(customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE, "")

	// Run the function
	expectedErrorPattern := `^failed to run license request: product, version, devops user, and devops key must be specified for license request$`
	err := RunInternalLicense()
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test readLicenseOptionValues function
func Test_readLicenseOptionValues(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Set up test data with all options
	setLicenseProductAndVersion(customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE, "13.0")

	// Run the function
	product, version, devopsUser, devopsKey, err := readLicenseOptionValues()

	testutils.CheckExpectedError(t, err, nil)
	if product != customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE {
		t.Errorf("expected product %q, got %q", customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE, product)
	}
	if version != "13.0" {
		t.Errorf("expected version %q, got %q", "13.0", version)
	}
	if devopsUser == "" {
		t.Error("expected devops user to be set, but it was empty")
	}
	if devopsKey == "" {
		t.Error("expected devops key to be set, but it was empty")
	}
}

func Test_readLicenseOptionValues_EmptyValues(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	setLicenseProductAndVersion(customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE, "")

	expectedErrorPattern := `^product, version, devops user, and devops key must be specified for license request$`
	_, _, _, _, err := readLicenseOptionValues()
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test runLicenseRequest function success
func Test_runLicenseRequest_Success(t *testing.T) {
	licenseData, err := runLicenseRequest(
		t.Context(),
		customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE,
		"13.0",
		os.Getenv("TEST_PINGCLI_DEVOPS_USER"),
		os.Getenv("TEST_PINGCLI_DEVOPS_KEY"))

	testutils.CheckExpectedError(t, err, nil)
	if licenseData == "" {
		t.Error("expected license data to be non-empty, but it was empty")
	}
}

// Test runLicenseRequest with an invalid devops key
func Test_runLicenseRequest_InvalidDevopsKey(t *testing.T) {
	licenseData, err := runLicenseRequest(
		t.Context(),
		customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE,
		"13.0",
		os.Getenv("TEST_PINGCLI_DEVOPS_USER"),
		"invalid-key")

	expectedErrorPattern := `^license request failed with status 401\: \{ "error"\: "Invalid devops-key header" \}$`
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
	if licenseData != "" {
		t.Error("expected license data to be empty, but it was not")
	}
}
