// Copyright © 2025 Ping Identity Corporation

package configuration_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// Test ValidateKoanfKey function
func Test_ValidateKoanfKey(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := configuration.ValidateKoanfKey("noColor")
	if err != nil {
		t.Errorf("ValidateKoanfKey returned error: %v", err)
	}
}

// Test ValidateKoanfKey function fails with invalid key
func Test_ValidateKoanfKey_InvalidKey(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `^key '.*' is not recognized as a valid configuration key.\s*Use 'pingcli config list-keys' to view all available keys`
	err := configuration.ValidateKoanfKey("invalid-key")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ValidateKoanfKey function fails with empty key
func Test_ValidateKoanfKey_EmptyKey(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `^key '' is not recognized as a valid configuration key.\s*Use 'pingcli config list-keys' to view all available keys`
	err := configuration.ValidateKoanfKey("")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ValidateKoanfKey supports case-insensitive keys
func Test_ValidateKoanfKey_CaseInsensitive(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := configuration.ValidateKoanfKey("NoCoLoR")
	if err != nil {
		t.Errorf("ValidateKoanfKey returned error: %v", err)
	}
}

// Test ValidateParentKoanfKey function
func Test_ValidateParentKoanfKey(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := configuration.ValidateParentKoanfKey("service")
	if err != nil {
		t.Errorf("ValidateParentKoanfKey returned error: %v", err)
	}
}

// Test ValidateParentKoanfKey function fails with invalid key
func Test_ValidateParentKoanfKey_InvalidKey(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `^key '.*' is not recognized as a valid configuration key.\s*Use 'pingcli config list-keys' to view all available keys`
	err := configuration.ValidateParentKoanfKey("invalid-key")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ValidateParentKoanfKey function fails with empty key
func Test_ValidateParentKoanfKey_EmptyKey(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `^key '' is not recognized as a valid configuration key.\s*Use 'pingcli config list-keys' to view all available keys`
	err := configuration.ValidateParentKoanfKey("")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ValidateParentKoanfKey supports case-insensitive keys
func Test_ValidateParentKoanfKey_CaseInsensitive(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := configuration.ValidateParentKoanfKey("SeRvIcE")
	if err != nil {
		t.Errorf("ValidateParentKoanfKey returned error: %v", err)
	}
}

// Test OptionFromKoanfKey function
func Test_OptionFromKoanfKey(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	opt, err := configuration.OptionFromKoanfKey("noColor")
	if err != nil {
		t.Errorf("OptionFromKoanfKey returned error: %v", err)
	}

	if opt.KoanfKey != "noColor" {
		t.Errorf("OptionFromKoanfKey returned invalid option: %v", opt)
	}
}

// Test OptionFromKoanfKey supports case-insensitive keys
func Test_OptionFromKoanfKey_CaseInsensitive(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	opt, err := configuration.OptionFromKoanfKey("NoCoLoR")
	if err != nil {
		t.Errorf("OptionFromKoanfKey returned error: %v", err)
	}

	if opt.KoanfKey != options.RootColorOption.KoanfKey {
		t.Errorf("OptionFromKoanfKey returned invalid option: %v", opt)
	}
}
