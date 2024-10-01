package profiles_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test Validate function
func TestValidate(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := profiles.Validate()
	if err != nil {
		t.Errorf("Validate returned error: %v", err)
	}
}

// Test Validate function with invalid uuid
func TestValidateInvalidProfile(t *testing.T) {
	fileContents := `activeProfile: default
default:
    description: "default description"
    pingone:
        export:
            environmentid: "invalid"`

	testutils_viper.InitVipersCustomFile(t, fileContents)

	err := profiles.Validate()
	if err == nil {
		t.Errorf("Validate returned nil, expected error")
	}
}

// Test Validate function with invalid region
func TestValidateInvalidRegion(t *testing.T) {
	fileContents := `activeProfile: default
default:
    description: "default description"
    pingone:
        region: "invalid"`

	testutils_viper.InitVipersCustomFile(t, fileContents)

	err := profiles.Validate()
	if err == nil {
		t.Errorf("Validate returned nil, expected error")
	}
}

// Test Validate function with invalid bool
func TestValidateInvalidBool(t *testing.T) {
	fileContents := `activeProfile: default
default:
    description: "default description"
    pingctl:
        color: invalid`

	testutils_viper.InitVipersCustomFile(t, fileContents)

	err := profiles.Validate()
	if err == nil {
		t.Errorf("Validate returned nil, expected error")
	}
}

// Test Validate function with invalid output format
func TestValidateInvalidOutputFormat(t *testing.T) {
	fileContents := `activeProfile: default
default:
    description: "default description"
    pingctl:
        outputFormat: invalid`

	testutils_viper.InitVipersCustomFile(t, fileContents)

	err := profiles.Validate()
	if err == nil {
		t.Errorf("Validate returned nil, expected error")
	}
}

// Test Validate function with invalid profile name
func TestValidateInvalidProfileName(t *testing.T) {
	fileContents := `activeProfile: default
default:
    description: "default description"
invalid(&*^&*^&*^**$):
    description: "default description"`

	testutils_viper.InitVipersCustomFile(t, fileContents)

	err := profiles.Validate()
	if err == nil {
		t.Errorf("Validate returned nil, expected error")
	}
}
