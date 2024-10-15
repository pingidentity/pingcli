package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test OutputFormat Set function
func Test_OutputFormat_Set(t *testing.T) {
	outputFormat := new(customtypes.OutputFormat)

	err := outputFormat.Set(customtypes.ENUM_OUTPUT_FORMAT_JSON)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}
}

// Test Set function fails with invalid value
func Test_OutputFormat_Set_InvalidValue(t *testing.T) {
	outputFormat := new(customtypes.OutputFormat)

	invalidValue := "invalid"

	expectedErrorPattern := `^unrecognized Output Format: '.*'\. Must be one of: .*$`
	err := outputFormat.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Set function fails with nil
func Test_OutputFormat_Set_Nil(t *testing.T) {
	var outputFormat *customtypes.OutputFormat

	val := customtypes.ENUM_OUTPUT_FORMAT_JSON

	expectedErrorPattern := `^failed to set Output Format value: .* Output Format is nil$`
	err := outputFormat.Set(val)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test String function
func Test_OutputFormat_String(t *testing.T) {
	outputFormat := customtypes.OutputFormat(customtypes.ENUM_OUTPUT_FORMAT_JSON)

	expected := customtypes.ENUM_OUTPUT_FORMAT_JSON
	actual := outputFormat.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
