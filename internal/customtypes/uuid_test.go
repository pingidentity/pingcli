// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test UUID Set function
func Test_UUID_Set(t *testing.T) {
	uuid := new(customtypes.UUID)

	val := "123e4567-e89b-12d3-a456-426614174000"
	err := uuid.Set(val)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}
}

// Test Set function fails with invalid value
func Test_UUID_Set_InvalidValue(t *testing.T) {
	uuid := new(customtypes.UUID)

	invalidValue := "invalid"

	expectedErrorPattern := `^uuid string is wrong length$`
	err := uuid.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Set function fails with nil
func Test_UUID_Set_Nil(t *testing.T) {
	var uuid *customtypes.UUID

	val := "123e4567-e89b-12d3-a456-426614174000"

	expectedErrorPattern := `^failed to set UUID value: .* UUID is nil$`
	err := uuid.Set(val)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test String function
func Test_UUID_String(t *testing.T) {
	uuid := customtypes.UUID("123e4567-e89b-12d3-a456-426614174000")

	expected := "123e4567-e89b-12d3-a456-426614174000"
	actual := uuid.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
