package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test HTTP Method Set function
func Test_HTTPMethod_Set(t *testing.T) {
	// Create a new HTTPMethod
	httpMethod := new(customtypes.HTTPMethod)

	err := httpMethod.Set(customtypes.ENUM_HTTP_METHOD_GET)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}
}

// Test Set function fails with invalid value
func Test_HTTPMethod_Set_InvalidValue(t *testing.T) {
	httpMethod := new(customtypes.HTTPMethod)

	invalidValue := "invalid"
	expectedErrorPattern := `^unrecognized HTTP Method: '.*'. Must be one of: .*$`
	err := httpMethod.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Set function fails with nil
func Test_HTTPMethod_Set_Nil(t *testing.T) {
	var httpMethod *customtypes.HTTPMethod

	expectedErrorPattern := `^failed to set HTTP Method value: .*\. HTTPMethod is nil$`
	err := httpMethod.Set(customtypes.ENUM_HTTP_METHOD_GET)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test String function
func Test_HTTPMethod_String(t *testing.T) {
	httpMethod := customtypes.HTTPMethod(customtypes.ENUM_HTTP_METHOD_GET)

	expected := customtypes.ENUM_HTTP_METHOD_GET
	actual := httpMethod.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
