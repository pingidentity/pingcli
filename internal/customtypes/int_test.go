package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test Int Set function
func Test_Int_Set(t *testing.T) {
	i := new(customtypes.Int)

	err := i.Set("42")
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}
}

// Test Set function fails with invalid value
func Test_Int_Set_InvalidValue(t *testing.T) {
	i := new(customtypes.Int)

	invalidValue := "invalid"
	expectedErrorPattern := `^strconv.ParseInt: parsing ".*": invalid syntax$`
	err := i.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Set function fails with nil
func Test_Int_Set_Nil(t *testing.T) {
	var i *customtypes.Int
	val := "42"

	expectedErrorPattern := `^failed to set Int value: .* Int is nil$`
	err := i.Set(val)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test String function
func Test_Int_String(t *testing.T) {
	i := customtypes.Int(42)

	expected := "42"
	actual := i.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
