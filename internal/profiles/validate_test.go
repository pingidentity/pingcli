// Copyright © 2025 Ping Identity Corporation

package profiles_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// Test Validate function
func TestValidate(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

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
    pingOne:
        export:
            environmentID: "invalid"`

	testutils_koanf.InitKoanfsCustomFile(t, fileContents)

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
    pingOne:
        region: "invalid"`

	testutils_koanf.InitKoanfsCustomFile(t, fileContents)

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
    pingcli:
        noColor: invalid`

	testutils_koanf.InitKoanfsCustomFile(t, fileContents)

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
    pingcli:
        outputFormat: invalid`

	testutils_koanf.InitKoanfsCustomFile(t, fileContents)

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

	testutils_koanf.InitKoanfsCustomFile(t, fileContents)

	err := profiles.Validate()
	if err == nil {
		t.Errorf("Validate returned nil, expected error")
	}
}
