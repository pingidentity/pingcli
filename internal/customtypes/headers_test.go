// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test Headers Set function
func Test_Headers_Set(t *testing.T) {
	hs := new(customtypes.HeaderSlice)

	service := "key: value"
	err := hs.Set(service)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}
}

// Test Headers Set function with invalid value
func Test_Headers_Set_InvalidValue(t *testing.T) {
	hs := new(customtypes.HeaderSlice)

	invalidValue := "invalid=value"
	expectedErrorPattern := `^failed to set Headers: Invalid header: .*\. Headers must be in the proper format. Expected regex pattern: .*$`
	err := hs.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Headers Set function with nil
func Test_Headers_Set_Nil(t *testing.T) {
	var hs *customtypes.HeaderSlice

	expectedErrorPattern := `^failed to set Headers value: .* Headers is nil$`
	err := hs.Set("key: value")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
