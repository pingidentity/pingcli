package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test PingOneRegion Set function
func Test_PingOneRegion_Set(t *testing.T) {
	prc := new(customtypes.PingoneRegionCode)

	err := prc.Set(customtypes.ENUM_PINGONE_REGION_CODE_AP)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}
}

// Test Set function fails with invalid value
func Test_PingOneRegion_Set_InvalidValue(t *testing.T) {
	prc := new(customtypes.PingoneRegionCode)

	invalidValue := "invalid"

	expectedErrorPattern := `^unrecognized Pingone Region Code: '.*'\. Must be one of: .*$`
	err := prc.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Set function fails with nil
func Test_PingOneRegion_Set_Nil(t *testing.T) {
	var prc *customtypes.PingoneRegionCode

	val := customtypes.ENUM_PINGONE_REGION_CODE_AP

	expectedErrorPattern := `^failed to set Pingone Region Code value: .* Pingone Region Code is nil$`
	err := prc.Set(val)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test String function
func Test_PingOneRegion_String(t *testing.T) {
	pingoneRegion := customtypes.PingoneRegionCode(customtypes.ENUM_PINGONE_REGION_CODE_CA)

	expected := customtypes.ENUM_PINGONE_REGION_CODE_CA
	actual := pingoneRegion.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
