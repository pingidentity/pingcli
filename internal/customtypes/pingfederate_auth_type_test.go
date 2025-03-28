// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test PingFederateAuthType Set function
func Test_PingFederateAuthType_Set(t *testing.T) {
	// Create a new PingFederateAuthType
	pingAuthType := new(customtypes.PingFederateAuthenticationType)

	err := pingAuthType.Set(customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Set function fails with invalid value
func Test_PingFederateAuthType_Set_InvalidValue(t *testing.T) {
	pingAuthType := new(customtypes.PingFederateAuthenticationType)

	invalidValue := "invalid"
	expectedErrorPattern := `^unrecognized PingFederate Authentication Type: '.*'\. Must be one of: .*$`
	err := pingAuthType.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Set function fails with nil
func Test_PingFederateAuthType_Set_Nil(t *testing.T) {
	var pingAuthType *customtypes.PingFederateAuthenticationType

	expectedErrorPattern := `^failed to set PingFederate Authentication Type value: .*\. PingFederate Authentication Type is nil$`
	err := pingAuthType.Set(customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test String function
func Test_PingFederateAuthType_String(t *testing.T) {
	pingAuthType := customtypes.PingFederateAuthenticationType(customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC)

	expected := customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC
	actual := pingAuthType.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
