package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test StringSlice Set function
func Test_StringSlice_Set(t *testing.T) {
	ss := new(customtypes.StringSlice)

	val := "value1,value2"
	err := ss.Set(val)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}
}

// Test Set function fails with nil
func Test_StringSlice_Set_Nil(t *testing.T) {
	var ss *customtypes.StringSlice

	val := "value1,value2"
	expectedErrorPattern := `^failed to set StringSlice value: .* StringSlice is nil$`
	err := ss.Set(val)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test String function
func Test_StringSlice_String(t *testing.T) {
	ss := customtypes.StringSlice([]string{"value1", "value2"})

	expected := "value1,value2"
	actual := ss.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}

// Test StringSlice String function with empty slice
func Test_StringSlice_String_Empty(t *testing.T) {
	ss := customtypes.StringSlice([]string{})

	expected := ""
	actual := ss.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
