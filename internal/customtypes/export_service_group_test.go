package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test ExportServiceGroup Set function
func Test_ExportServiceGroup_Set(t *testing.T) {
	// Create a new ExportServiceGroup
	exportServiceGroup := new(customtypes.ExportServiceGroup)

	err := exportServiceGroup.Set(customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}
}

// Test ExportServiceGroup Set function fails with invalid value
func Test_ExportServiceGroup_Set_InvalidValue(t *testing.T) {
	// Create a new ExportServiceGroup
	exportServiceGroup := new(customtypes.ExportServiceGroup)

	invalidValue := "invalid"

	expectedErrorPattern := `^unrecognized service group '.*'. Must be one of: .*$`
	err := exportServiceGroup.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ExportServiceGroup Set function fails with nil
func Test_ExportServiceGroup_Set_Nil(t *testing.T) {
	var exportServiceGroup *customtypes.ExportServiceGroup

	val := customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE

	expectedErrorPattern := `^failed to set Service Group value: .* Service Group is nil$`
	err := exportServiceGroup.Set(val)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ExportServiceGroup String function
func Test_ExportServiceGroup_String(t *testing.T) {
	exportServiceGroup := customtypes.ExportServiceGroup(customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE)

	expected := customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE
	actual := exportServiceGroup.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}

// Test ExportServiceGroupValidValues
func Test_ExportServiceGroupValidValues(t *testing.T) {
	serviceGroupEnum := customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE

	serviceGroupValidValues := customtypes.ExportServiceGroupValidValues()
	if serviceGroupValidValues[0] != serviceGroupEnum {
		t.Errorf("ExportServiceGroupValidValues returned: %v, expected: %v", serviceGroupValidValues, serviceGroupEnum)
	}
}

// Test ExportServicePingOneValidValues
func Test_ExportServicesPingOneValidValues(t *testing.T) {
	pingOneServiceGroupValidValues := customtypes.ExportServicesPingOneValidValues()
	if len(pingOneServiceGroupValidValues) != 5 {
		t.Errorf("ExportServicesPingOneValidValues returned: %v, expected: %v", len(pingOneServiceGroupValidValues), 5)
	}
}
