// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf/v2"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigSet(kvPair string) (err error) {
	pName, vKey, vValue, err := readConfigSetOptions(kvPair)
	if err != nil {
		return fmt.Errorf("failed to set configuration: %w", err)
	}

	if err = configuration.ValidateKoanfKey(vKey); err != nil {
		return fmt.Errorf("failed to set configuration: %w", err)
	}

	// Make sure value is not empty, and suggest unset command if it is
	if vValue == "" {
		return fmt.Errorf("failed to set configuration: value for key '%s' is empty. Use 'pingcli config unset %s' to unset the key", vKey, vKey)
	}

	subKoanf, err := profiles.GetKoanfConfig().GetProfileKoanf(pName)
	if err != nil {
		return fmt.Errorf("failed to set configuration: %w", err)
	}

	opt, err := configuration.OptionFromKoanfKey(vKey)
	if err != nil {
		return fmt.Errorf("failed to set configuration: %w", err)
	}

	if err = setValue(subKoanf, opt.KoanfKey, vValue, opt.Type); err != nil {
		return fmt.Errorf("failed to set configuration: %w", err)
	}

	if err = profiles.GetKoanfConfig().SaveProfile(pName, subKoanf); err != nil {
		return fmt.Errorf("failed to set configuration: %w", err)
	}

	msgStr := "Configuration set successfully:\n"

	vVal, _, err := profiles.KoanfValueFromOption(opt, pName)
	if err != nil {
		return fmt.Errorf("failed to set configuration: %w", err)
	}

	unmaskOptionVal, err := profiles.GetOptionValue(options.ConfigUnmaskSecretValueOption)
	if err != nil {
		unmaskOptionVal = "false"
	}

	if opt.Sensitive && strings.EqualFold(unmaskOptionVal, "false") {
		msgStr += fmt.Sprintf("%s=%s", opt.KoanfKey, profiles.MaskValue(vVal))
	} else {
		msgStr += fmt.Sprintf("%s=%s", opt.KoanfKey, vVal)
	}

	output.Success(msgStr, nil)

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
	if !options.RootProfileOption.Flag.Changed {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
	} else {
		pName, err = profiles.GetOptionValue(options.RootProfileOption)
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

	if strings.EqualFold(parsedInput[0], options.RootActiveProfileOption.KoanfKey) {
		return "", "", fmt.Errorf("invalid assignment. Please use the 'pingcli config set active-profile <profile-name>' command to set the active profile")
	}

	return parsedInput[0], parsedInput[1], nil
}

func setValue(profileKoanf *koanf.Koanf, vKey, vValue string, valueType options.OptionType) (err error) {
	switch valueType {
	case options.BOOL:
		b := new(customtypes.Bool)
		if err = b.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a boolean. Allowed [true, false]: %w", vKey, err)
		}
		err = profileKoanf.Set(vKey, b)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.EXPORT_FORMAT:
		exportFormat := new(customtypes.ExportFormat)
		if err = exportFormat.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid export format. Allowed [%s]: %w", vKey, strings.Join(customtypes.ExportFormatValidValues(), ", "), err)
		}
		err = profileKoanf.Set(vKey, exportFormat)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.EXPORT_SERVICE_GROUP:
		exportServiceGroup := new(customtypes.ExportServiceGroup)
		if err = exportServiceGroup.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be valid export service group. Allowed [%s]: %w", vKey, strings.Join(customtypes.ExportServiceGroupValidValues(), ", "), err)
		}
		err = profileKoanf.Set(vKey, exportServiceGroup)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.EXPORT_SERVICES:
		exportServices := new(customtypes.ExportServices)
		if err = exportServices.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be valid export service(s). Allowed [%s]: %w", vKey, strings.Join(customtypes.ExportServicesValidValues(), ", "), err)
		}
		err = profileKoanf.Set(vKey, exportServices)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.OUTPUT_FORMAT:
		outputFormat := new(customtypes.OutputFormat)
		if err = outputFormat.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid output format. Allowed [%s]: %w", vKey, strings.Join(customtypes.OutputFormatValidValues(), ", "), err)
		}
		err = profileKoanf.Set(vKey, outputFormat)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.PINGONE_REGION_CODE:
		region := new(customtypes.String)
		if err = region.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid PingOne Region Code. Allowed [%s]: %w", vKey, strings.Join(customtypes.PingOneRegionCodeValidValues(), ", "), err)
		}
		err = profileKoanf.Set(vKey, region)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.STRING:
		str := new(customtypes.String)
		if err = str.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a string: %w", vKey, err)
		}
		err = profileKoanf.Set(vKey, str)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.STRING_SLICE:
		strSlice := new(customtypes.StringSlice)
		if err = strSlice.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a string slice: %w", vKey, err)
		}
		err = profileKoanf.Set(vKey, strSlice)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.UUID:
		uuid := new(customtypes.UUID)
		if err = uuid.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid UUID: %w", vKey, err)
		}
		err = profileKoanf.Set(vKey, uuid)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.PINGONE_AUTH_TYPE:
		authType := new(customtypes.PingOneAuthenticationType)
		if err = authType.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid PingOne Authentication Type. Allowed [%s]: %w", vKey, strings.Join(customtypes.PingOneAuthenticationTypeValidValues(), ", "), err)
		}
		err = profileKoanf.Set(vKey, authType)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.PINGFEDERATE_AUTH_TYPE:
		authType := new(customtypes.PingFederateAuthenticationType)
		if err = authType.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid PingFederate Authentication Type. Allowed [%s]: %w", vKey, strings.Join(customtypes.PingFederateAuthenticationTypeValidValues(), ", "), err)
		}
		err = profileKoanf.Set(vKey, authType)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.INT:
		intValue := new(customtypes.Int)
		if err = intValue.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be an integer: %w", vKey, err)
		}
		err = profileKoanf.Set(vKey, intValue)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.REQUEST_HTTP_METHOD:
		httpMethod := new(customtypes.HTTPMethod)
		if err = httpMethod.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid HTTP method. Allowed [%s]: %w", vKey, strings.Join(customtypes.HTTPMethodValidValues(), ", "), err)
		}
		err = profileKoanf.Set(vKey, httpMethod)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.REQUEST_SERVICE:
		service := new(customtypes.RequestService)
		if err = service.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid request service. Allowed [%s]: %w", vKey, strings.Join(customtypes.RequestServiceValidValues(), ", "), err)
		}
		err = profileKoanf.Set(vKey, service)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.LICENSE_PRODUCT:
		licenseProduct := new(customtypes.LicenseProduct)
		if err = licenseProduct.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid license product. Allowed [%s]: %w", vKey, strings.Join(customtypes.LicenseProductValidValues(), ", "), err)
		}
		err = profileKoanf.Set(vKey, licenseProduct)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	case options.LICENSE_VERSION:
		licenseVersion := new(customtypes.LicenseVersion)
		if err = licenseVersion.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid license version. Must be of the form 'major.minor': %w", vKey, err)
		}
		err = profileKoanf.Set(vKey, licenseVersion)
		if err != nil {
			return fmt.Errorf("unable to set key '%w' in koanf profile: ", err)
		}
	default:
		return fmt.Errorf("failed to set configuration: variable type for key '%s' is not recognized", vKey)
	}

	return nil
}
