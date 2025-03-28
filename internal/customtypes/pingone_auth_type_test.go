// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test PingOne Authentication Type Set function
func Test_PingOneAuthType_Set(t *testing.T) {
	// Create a new PingOneAuthType
	pingAuthType := new(customtypes.PingOneAuthenticationType)

	err := pingAuthType.Set(customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Set function fails with invalid value
func Test_PingOneAuthType_Set_InvalidValue(t *testing.T) {
	pingAuthType := new(customtypes.PingOneAuthenticationType)

	invalidValue := "invalid"
	expectedErrorPattern := `^unrecognized PingOne Authentication Type: '.*'\. Must be one of: .*$`
	err := pingAuthType.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Set function fails with nil
func Test_PingOneAuthType_Set_Nil(t *testing.T) {
	var pingAuthType *customtypes.PingOneAuthenticationType

	expectedErrorPattern := `^failed to set PingOne Authentication Type value: .*\. PingOne Authentication Type is nil$`
	err := pingAuthType.Set(customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test String function
func Test_PingOneAuthType_String(t *testing.T) {
	pingAuthType := customtypes.PingOneAuthenticationType(customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER)

	expected := customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER
	actual := pingAuthType.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
