package configuration_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/configuration"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test ValidateViperKey function
func Test_ValidateViperKey(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := configuration.ValidateViperKey("color")
	if err != nil {
		t.Errorf("ValidateViperKey returned error: %v", err)
	}
}

// Test ValidateViperKey function fails with invalid key
func Test_ValidateViperKey_InvalidKey(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^key '.*' is not recognized as a valid configuration key. Valid keys: .*$`
	err := configuration.ValidateViperKey("invalid-key")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ValidateViperKey function fails with empty key
func Test_ValidateViperKey_EmptyKey(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^key '' is not recognized as a valid configuration key. Valid keys: .*$`
	err := configuration.ValidateViperKey("")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ValidateParentViperKey function
func Test_ValidateParentViperKey(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := configuration.ValidateParentViperKey("service")
	if err != nil {
		t.Errorf("ValidateParentViperKey returned error: %v", err)
	}
}

// Test ValidateParentViperKey function fails with invalid key
func Test_ValidateParentViperKey_InvalidKey(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^key '.*' is not recognized as a valid configuration key. Valid keys: .*$`
	err := configuration.ValidateParentViperKey("invalid-key")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ValidateParentViperKey function fails with empty key
func Test_ValidateParentViperKey_EmptyKey(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^key '' is not recognized as a valid configuration key. Valid keys: .*$`
	err := configuration.ValidateParentViperKey("")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test OptionFromViperKey function
func Test_OptionFromViperKey(t *testing.T) {
	testutils_viper.InitVipers(t)

	opt, err := configuration.OptionFromViperKey("color")
	if err != nil {
		t.Errorf("OptionFromViperKey returned error: %v", err)
	}

	if opt.ViperKey != "color" {
		t.Errorf("OptionFromViperKey returned invalid option: %v", opt)
	}
}
