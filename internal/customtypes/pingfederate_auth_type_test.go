package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test PingfederateAuthType Set function
func Test_PingfederateAuthType_Set(t *testing.T) {
	// Create a new PingfederateAuthType
	pingAuthType := new(customtypes.PingfederateAuthenticationType)

	err := pingAuthType.Set(customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Set function fails with invalid value
func Test_PingfederateAuthType_Set_InvalidValue(t *testing.T) {
	pingAuthType := new(customtypes.PingfederateAuthenticationType)

	invalidValue := "invalid"
	expectedErrorPattern := `^unrecognized Pingfederate Authentication Type: '.*'\. Must be one of: .*$`
	err := pingAuthType.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Set function fails with nil
func Test_PingfederateAuthType_Set_Nil(t *testing.T) {
	var pingAuthType *customtypes.PingfederateAuthenticationType

	expectedErrorPattern := `^failed to set Pingfederate Authentication Type value: .*\. Pingfederate Authentication Type is nil$`
	err := pingAuthType.Set(customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test String function
func Test_PingfederateAuthType_String(t *testing.T) {
	pingAuthType := customtypes.PingfederateAuthenticationType(customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC)

	expected := customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC
	actual := pingAuthType.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
