package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test Request Services Set function
func Test_RequestServices_Set(t *testing.T) {
	rs := new(customtypes.RequestService)

	service := customtypes.ENUM_REQUEST_SERVICE_PINGONE
	err := rs.Set(service)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Set function fails with invalid value
func Test_RequestServices_Set_InvalidValue(t *testing.T) {
	rs := new(customtypes.RequestService)

	invalidValue := "invalid"
	expectedErrorPattern := `^unrecognized Request Service: '.*'\. Must be one of: .*$`
	err := rs.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Set function fails with nil
func Test_RequestServices_Set_Nil(t *testing.T) {
	var rs *customtypes.RequestService

	expectedErrorPattern := `^failed to set RequestService value: .*\. RequestService is nil$`
	err := rs.Set(customtypes.ENUM_REQUEST_SERVICE_PINGONE)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test String function
func Test_RequestServices_String(t *testing.T) {
	rs := customtypes.RequestService(customtypes.ENUM_REQUEST_SERVICE_PINGONE)

	expected := customtypes.ENUM_REQUEST_SERVICE_PINGONE
	actual := rs.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
