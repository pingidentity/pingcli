// Copyright © 2026 Ping Identity Corporation

package configuration_auth_test

import (
	"testing"

	configuration_auth "github.com/pingidentity/pingcli/internal/configuration/auth"
	"github.com/pingidentity/pingcli/internal/configuration/options"
)

func TestInitAuthOptions(t *testing.T) {
	configuration_auth.InitAuthOptions()

	// Test device-code option
	deviceCodeOption := options.AuthMethodDeviceCodeOption
	if deviceCodeOption.CobraParamName != "device-code" {
		t.Errorf("Expected CobraParamName to be 'device-code', got %q", deviceCodeOption.CobraParamName)
	}
	if deviceCodeOption.Type != options.BOOL {
		t.Errorf("Expected Type to be BOOL, got %v", deviceCodeOption.Type)
	}
	if deviceCodeOption.Sensitive {
		t.Error("Expected Sensitive to be false")
	}
	if deviceCodeOption.Flag.Usage != "Use device authorization flow" {
		t.Errorf("Expected Usage to be 'Use device authorization flow', got %q", deviceCodeOption.Flag.Usage)
	}
	if deviceCodeOption.Flag == nil {
		t.Fatal("Flag should not be nil")
	}

	// Test client-credentials option
	clientCredentialsOption := options.AuthMethodClientCredentialsOption
	if clientCredentialsOption.CobraParamName != "client-credentials" {
		t.Errorf("Expected CobraParamName to be 'client-credentials', got %q", clientCredentialsOption.CobraParamName)
	}
	if clientCredentialsOption.Type != options.BOOL {
		t.Errorf("Expected Type to be BOOL, got %v", clientCredentialsOption.Type)
	}
	if clientCredentialsOption.Sensitive {
		t.Error("Expected Sensitive to be false")
	}
	if clientCredentialsOption.Flag.Usage != "Use client credentials flow" {
		t.Errorf("Expected Usage to be 'Use client credentials flow', got %q", clientCredentialsOption.Flag.Usage)
	}
	if clientCredentialsOption.Flag == nil {
		t.Fatal("Flag should not be nil")
	}

	// Test authorization-code option
	authorizationCodeOption := options.AuthMethodAuthorizationCodeOption
	if authorizationCodeOption.CobraParamName != "authorization-code" {
		t.Errorf("Expected CobraParamName to be 'authorization-code', got %q", authorizationCodeOption.CobraParamName)
	}
	if authorizationCodeOption.Type != options.BOOL {
		t.Errorf("Expected Type to be BOOL, got %v", authorizationCodeOption.Type)
	}
	if authorizationCodeOption.Sensitive {
		t.Error("Expected Sensitive to be false")
	}
	if authorizationCodeOption.Flag.Usage != "Use authorization code flow" {
		t.Errorf("Expected Usage to be 'Use authorization code flow', got %q", authorizationCodeOption.Flag.Usage)
	}
	if authorizationCodeOption.Flag == nil {
		t.Fatal("Flag should not be nil")
	}
}

func TestAuthOptionDefaults(t *testing.T) {
	configuration_auth.InitAuthOptions()

	// All grant type flags should default to false
	deviceCodeOption := options.AuthMethodDeviceCodeOption
	defaultValue := deviceCodeOption.DefaultValue.String()
	if defaultValue != "false" {
		t.Errorf("Expected default value to be 'false', got %q", defaultValue)
	}

	clientCredentialsOption := options.AuthMethodClientCredentialsOption
	defaultValue = clientCredentialsOption.DefaultValue.String()
	if defaultValue != "false" {
		t.Errorf("Expected default value to be 'false', got %q", defaultValue)
	}

	authorizationCodeOption := options.AuthMethodAuthorizationCodeOption
	defaultValue = authorizationCodeOption.DefaultValue.String()
	if defaultValue != "false" {
		t.Errorf("Expected default value to be 'false', got %q", defaultValue)
	}
}

func TestAuthOptionShorthandFlags(t *testing.T) {
	configuration_auth.InitAuthOptions()

	// Test shorthand flags
	deviceCodeOption := options.AuthMethodDeviceCodeOption
	if deviceCodeOption.Flag.Shorthand != "d" {
		t.Errorf("Expected shorthand to be 'd', got %q", deviceCodeOption.Flag.Shorthand)
	}

	clientCredentialsOption := options.AuthMethodClientCredentialsOption
	if clientCredentialsOption.Flag.Shorthand != "c" {
		t.Errorf("Expected shorthand to be 'c', got %q", clientCredentialsOption.Flag.Shorthand)
	}

	authorizationCodeOption := options.AuthMethodAuthorizationCodeOption
	if authorizationCodeOption.Flag.Shorthand != "a" {
		t.Errorf("Expected shorthand to be 'a', got %q", authorizationCodeOption.Flag.Shorthand)
	}
}

func TestAuthOptionBooleanBehavior(t *testing.T) {
	configuration_auth.InitAuthOptions()

	// Test that boolean flags have NoOptDefVal set to "true" for proper boolean behavior
	deviceCodeOption := options.AuthMethodDeviceCodeOption
	if deviceCodeOption.Flag.NoOptDefVal != "true" {
		t.Errorf("Expected NoOptDefVal to be 'true', got %q", deviceCodeOption.Flag.NoOptDefVal)
	}

	clientCredentialsOption := options.AuthMethodClientCredentialsOption
	if clientCredentialsOption.Flag.NoOptDefVal != "true" {
		t.Errorf("Expected NoOptDefVal to be 'true', got %q", clientCredentialsOption.Flag.NoOptDefVal)
	}

	authorizationCodeOption := options.AuthMethodAuthorizationCodeOption
	if authorizationCodeOption.Flag.NoOptDefVal != "true" {
		t.Errorf("Expected NoOptDefVal to be 'true', got %q", authorizationCodeOption.Flag.NoOptDefVal)
	}
}

func TestAllAuthOptionsInitialized(t *testing.T) {
	configuration_auth.InitAuthOptions()

	// Verify all grant type options are properly initialized
	authOptions := []options.Option{
		options.AuthMethodDeviceCodeOption,
		options.AuthMethodClientCredentialsOption,
		options.AuthMethodAuthorizationCodeOption,
	}

	for _, option := range authOptions {
		if option.Flag == nil {
			t.Error("Auth option flag should not be nil")
		}
		if option.CobraParamName == "" {
			t.Error("Auth option should have cobra param name")
		}
		if option.Flag.Usage == "" {
			t.Error("Auth option should have usage description")
		}
		if option.Type != options.BOOL {
			t.Errorf("Auth option should be boolean type, got %v", option.Type)
		}
		if option.Sensitive {
			t.Error("Auth option should not be sensitive")
		}
	}
}
