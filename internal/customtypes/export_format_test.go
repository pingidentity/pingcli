// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test ExportFormat Set function
func Test_ExportFormat_Set(t *testing.T) {
	// Create a new ExportFormat
	exportFormat := new(customtypes.ExportFormat)

	err := exportFormat.Set(customtypes.ENUM_EXPORT_FORMAT_HCL)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}
}

// Test Set function fails with invalid value
func Test_ExportFormat_Set_InvalidValue(t *testing.T) {
	// Create a new ExportFormat
	exportFormat := new(customtypes.ExportFormat)

	invalidValue := "invalid"

	expectedErrorPattern := `^unrecognized export format '.*'. Must be one of: .*$`
	err := exportFormat.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Set function fails with nil
func Test_ExportFormat_Set_Nil(t *testing.T) {
	var exportFormat *customtypes.ExportFormat

	val := customtypes.ENUM_EXPORT_FORMAT_HCL

	expectedErrorPattern := `^failed to set Export Format value: .* Export Format is nil$`
	err := exportFormat.Set(val)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test String function
func Test_ExportFormat_String(t *testing.T) {
	exportFormat := customtypes.ExportFormat(customtypes.ENUM_EXPORT_FORMAT_HCL)

	expected := customtypes.ENUM_EXPORT_FORMAT_HCL
	actual := exportFormat.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
