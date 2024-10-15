package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test Bool Set function
func Test_Bool_Set(t *testing.T) {
	b := new(customtypes.Bool)
	val := "true"

	err := b.Set(val)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}

	val = "false"

	err = b.Set(val)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}
}

// Test Set function fails with invalid value
func Test_Bool_Set_InvalidValue(t *testing.T) {
	b := new(customtypes.Bool)
	val := "invalid"

	expectedErrorPattern := `^strconv.ParseBool: parsing ".*": invalid syntax$`
	err := b.Set(val)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Set function fails with nil
func Test_Bool_Set_Nil(t *testing.T) {
	var b *customtypes.Bool
	val := "true"

	expectedErrorPattern := `^failed to set Bool value: .* Bool is nil$`
	err := b.Set(val)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test String function
func Test_Bool_String(t *testing.T) {
	b := customtypes.Bool(true)

	expected := "true"
	actual := b.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}

	b = customtypes.Bool(false)

	expected = "false"
	actual = b.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}

// Test Bool function
func Test_Bool_Bool(t *testing.T) {
	b := customtypes.Bool(true)

	expected := true
	actual := b.Bool()
	if actual != expected {
		t.Errorf("Bool returned: %t, expected: %t", actual, expected)
	}

	b = customtypes.Bool(false)

	expected = false
	actual = b.Bool()
	if actual != expected {
		t.Errorf("Bool returned: %t, expected: %t", actual, expected)
	}
}
