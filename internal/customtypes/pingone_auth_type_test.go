package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test Pingone Authentication Type Set function
func Test_PingoneAuthType_Set(t *testing.T) {
	// Create a new PingoneAuthType
	pingAuthType := new(customtypes.PingoneAuthenticationType)

	err := pingAuthType.Set(customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Set function fails with invalid value
func Test_PingoneAuthType_Set_InvalidValue(t *testing.T) {
	pingAuthType := new(customtypes.PingoneAuthenticationType)

	invalidValue := "invalid"
	expectedErrorPattern := `^unrecognized Pingone Authentication Type: '.*'\. Must be one of: .*$`
	err := pingAuthType.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Set function fails with nil
func Test_PingoneAuthType_Set_Nil(t *testing.T) {
	var pingAuthType *customtypes.PingoneAuthenticationType

	expectedErrorPattern := `^failed to set Pingone Authentication Type value: .*\. Pingone Authentication Type is nil$`
	err := pingAuthType.Set(customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test String function
func Test_PingoneAuthType_String(t *testing.T) {
	pingAuthType := customtypes.PingoneAuthenticationType(customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER)

	expected := customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER
	actual := pingAuthType.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
