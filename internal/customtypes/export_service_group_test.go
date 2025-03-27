package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test ExportServiceGroup Set function
func Test_ExportServiceGroup_Set(t *testing.T) {
	// Create a new ExportServiceGroup
	esg := new(customtypes.ExportServiceGroup)

	err := esg.Set(customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}
}

// Test ExportServiceGroup Set function fails with invalid value
func Test_ExportServiceGroup_Set_InvalidValue(t *testing.T) {
	// Create a new ExportServiceGroup
	esg := new(customtypes.ExportServiceGroup)

	invalidValue := "invalid"

	expectedErrorPattern := `^unrecognized service group .*\. Must be one of: .*$`
	err := esg.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ExportServiceGroup Set function fails with nil
func Test_ExportServiceGroup_Set_Nil(t *testing.T) {
	var esg *customtypes.ExportServiceGroup

	val := customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE

	expectedErrorPattern := `^failed to set ExportServiceGroup value: .* ExportServiceGroup is nil$`
	err := esg.Set(val)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ExportServiceGroup Valid Values returns expected amount
func Test_ExportServiceGroupValidValues(t *testing.T) {
	serviceGroupEnum := customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE

	serviceGroupValidValues := customtypes.ExportServiceGroupValidValues()
	if serviceGroupValidValues[0] != serviceGroupEnum {
		t.Errorf("ExportServiceGroupValidValues returned: %v, expected: %v", serviceGroupValidValues, serviceGroupEnum)
	}
}
