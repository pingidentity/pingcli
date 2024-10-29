package config_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/spf13/viper"
)

func RunInternalConfigSet(kvPair string) (err error) {
	pName, vKey, vValue, err := readConfigSetOptions(kvPair)
	if err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	if err = configuration.ValidateViperKey(vKey); err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	// Make sure value is not empty, and suggest unset command if it is
	if vValue == "" {
		return fmt.Errorf("failed to set configuration: value for key '%s' is empty. Use 'pingcli config unset %s' to unset the key", vKey, vKey)
	}

	if err = profiles.GetMainConfig().ValidateExistingProfileName(pName); err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	subViper := profiles.GetMainConfig().ViperInstance().Sub(pName)

	opt, err := configuration.OptionFromViperKey(vKey)
	if err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	if err = setValue(subViper, vKey, vValue, opt.Type); err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	if err = profiles.GetMainConfig().SaveProfile(pName, subViper); err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	yamlStr, err := profiles.GetMainConfig().ProfileToString(pName)
	if err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	output.Success("Configuration set successfully", nil)
	output.Message(yamlStr, nil)

	return nil
}

func readConfigSetOptions(kvPair string) (pName string, vKey string, vValue string, err error) {
	if pName, err = readConfigSetProfileName(); err != nil {
		return pName, vKey, vValue, err
	}

	if vKey, vValue, err = parseKeyValuePair(kvPair); err != nil {
		return pName, vKey, vValue, err
	}

	return pName, vKey, vValue, nil
}

func readConfigSetProfileName() (pName string, err error) {
	if !options.ConfigSetProfileOption.Flag.Changed {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
	} else {
		pName, err = profiles.GetOptionValue(options.ConfigSetProfileOption)
	}

	if err != nil {
		return pName, err
	}

	if pName == "" {
		return pName, fmt.Errorf("unable to determine profile to set configuration to")
	}

	return pName, nil
}

func parseKeyValuePair(kvPair string) (string, string, error) {
	parsedInput := strings.SplitN(kvPair, "=", 2)
	if len(parsedInput) < 2 {
		return "", "", fmt.Errorf("invalid assignment format '%s'. Expect 'key=value' format", kvPair)
	}

	return parsedInput[0], parsedInput[1], nil
}

func setValue(profileViper *viper.Viper, vKey, vValue string, valueType options.OptionType) (err error) {
	switch valueType {
	case options.ENUM_BOOL:
		bool := new(customtypes.Bool)
		if err = bool.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a boolean. Allowed [true, false]: %v", vKey, err)
		}
		profileViper.Set(vKey, bool)
	case options.ENUM_EXPORT_FORMAT:
		exportFormat := new(customtypes.ExportFormat)
		if err = exportFormat.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid export format. Allowed [%s]: %v", vKey, strings.Join(customtypes.ExportFormatValidValues(), ", "), err)
		}
		profileViper.Set(vKey, exportFormat)
	case options.ENUM_EXPORT_SERVICES:
		exportServices := new(customtypes.ExportServices)
		if err = exportServices.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be valid export service(s). Allowed [%s]: %v", vKey, strings.Join(customtypes.ExportServicesValidValues(), ", "), err)
		}
		profileViper.Set(vKey, exportServices)
	case options.ENUM_OUTPUT_FORMAT:
		outputFormat := new(customtypes.OutputFormat)
		if err = outputFormat.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid output format. Allowed [%s]: %v", vKey, strings.Join(customtypes.OutputFormatValidValues(), ", "), err)
		}
		profileViper.Set(vKey, outputFormat)
	case options.ENUM_PINGONE_REGION_CODE:
		region := new(customtypes.PingOneRegionCode)
		if err = region.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid PingOne Region Code. Allowed [%s]: %v", vKey, strings.Join(customtypes.PingOneRegionCodeValidValues(), ", "), err)
		}
		profileViper.Set(vKey, region)
	case options.ENUM_STRING:
		str := new(customtypes.String)
		if err = str.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a string: %v", vKey, err)
		}
		profileViper.Set(vKey, str)
	case options.ENUM_STRING_SLICE:
		strSlice := new(customtypes.StringSlice)
		if err = strSlice.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a string slice: %v", vKey, err)
		}
		profileViper.Set(vKey, strSlice)
	case options.ENUM_UUID:
		uuid := new(customtypes.UUID)
		if err = uuid.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid UUID: %v", vKey, err)
		}
		profileViper.Set(vKey, uuid)
	case options.ENUM_PINGONE_AUTH_TYPE:
		authType := new(customtypes.PingOneAuthenticationType)
		if err = authType.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid PingOne Authentication Type. Allowed [%s]: %v", vKey, strings.Join(customtypes.PingOneAuthenticationTypeValidValues(), ", "), err)
		}
		profileViper.Set(vKey, authType)
	case options.ENUM_PINGFEDERATE_AUTH_TYPE:
		authType := new(customtypes.PingFederateAuthenticationType)
		if err = authType.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid PingFederate Authentication Type. Allowed [%s]: %v", vKey, strings.Join(customtypes.PingFederateAuthenticationTypeValidValues(), ", "), err)
		}
		profileViper.Set(vKey, authType)
	case options.ENUM_INT:
		intValue := new(customtypes.Int)
		if err = intValue.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be an integer: %v", vKey, err)
		}
		profileViper.Set(vKey, intValue)
	case options.ENUM_REQUEST_HTTP_METHOD:
		httpMethod := new(customtypes.HTTPMethod)
		if err = httpMethod.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid HTTP method. Allowed [%s]: %v", vKey, strings.Join(customtypes.HTTPMethodValidValues(), ", "), err)
		}
		profileViper.Set(vKey, httpMethod)
	case options.ENUM_REQUEST_SERVICE:
		service := new(customtypes.RequestService)
		if err = service.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid request service. Allowed [%s]: %v", vKey, strings.Join(customtypes.RequestServiceValidValues(), ", "), err)
		}
		profileViper.Set(vKey, service)
	default:
		return fmt.Errorf("failed to set configuration: variable type for key '%s' is not recognized", vKey)
	}

	return nil
}
