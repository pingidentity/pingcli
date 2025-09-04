// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/knadh/koanf/v2"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var (
	ErrEmptyValue                 = errors.New("the set value provided is empty. Use 'pingcli config unset %s' to unset a key's configuration")
	ErrDetermineProfileSet        = errors.New("unable to determine profile to set configuration to")
	ErrKeyAssignmentFormat        = errors.New("invalid key-value assignment. Expect 'key=value' format")
	ErrActiveProfileAssignment    = errors.New("invalid active profile assignment. Please use the 'pingcli config set active-profile <profile-name>' command to set the active profile")
	ErrSetKey                     = errors.New("unable to set key in configuration profile")
	ErrMustBeBoolean              = errors.New("the value assignment must be a boolean. Allowed [true, false]")
	ErrMustBeExportFormat         = errors.New(fmt.Sprintf("the value assignment must be a valid export format. Allowed [%s]", strings.Join(customtypes.ExportFormatValidValues(), ", ")))
	ErrMustBeExportServiceGroup   = errors.New(fmt.Sprintf("the value assignment must be a valid export service group. Allowed [%s]", strings.Join(customtypes.ExportServiceGroupValidValues(), ", ")))
	ErrMustBeExportService        = errors.New(fmt.Sprintf("the value assignment must be valid export service(s). Allowed [%s]", strings.Join(customtypes.ExportServicesValidValues(), ", ")))
	ErrMustBeOutputFormat         = errors.New(fmt.Sprintf("the value assignment must be a valid output format. Allowed [%s]", strings.Join(customtypes.OutputFormatValidValues(), ", ")))
	ErrMustBePingoneRegionCode    = errors.New(fmt.Sprintf("the value assignment must be a valid PingOne region code. Allowed [%s]", strings.Join(customtypes.PingOneRegionCodeValidValues(), ", ")))
	ErrMustBeString               = errors.New("the value assignment must be a string")
	ErrMustBeStringSlice          = errors.New("the value assignment must be a string slice")
	ErrMustBeUUID                 = errors.New("the value assignment must be a valid UUID")
	ErrMustBePingoneAuthType      = errors.New(fmt.Sprintf("the value assignment must be a valid PingOne Authentication Type. Allowed [%s]", strings.Join(customtypes.PingOneAuthenticationTypeValidValues(), ", ")))
	ErrMustBePingfederateAuthType = errors.New(fmt.Sprintf("the value assignment must be a valid PingFederate Authentication Type. Allowed [%s]", strings.Join(customtypes.PingFederateAuthenticationTypeValidValues(), ", ")))
	ErrMustBeInteger              = errors.New("the value assignment must be an integer")
	ErrMustBeHttpMethod           = errors.New(fmt.Sprintf("the value assignment must be a valid HTTP method. Allowed [%s]", strings.Join(customtypes.HTTPMethodValidValues(), ", ")))
	ErrMustBeRequestService       = errors.New(fmt.Sprintf("the value assignment must be a valid request service. Allowed [%s]", strings.Join(customtypes.RequestServiceValidValues(), ", ")))
	ErrMustBeLicenseProduct       = errors.New(fmt.Sprintf("the value assignment must be a valid license product. Allowed [%s]", strings.Join(customtypes.LicenseProductValidValues(), ", ")))
	ErrMustBeLicenseVersion       = errors.New("the value assignment must be a valid license version. Must be of the form 'major.minor'")
	ErrTypeNotRecognized          = errors.New("the variable type for the configuration key is not recognized or supported")
)

type SetError struct {
	Err error
}

func (e *SetError) Error() string {
	var err *SetError
	if errors.As(e.Err, &err) {
		return err.Error()
	}
	return fmt.Sprintf("failed to set configuration: %s", e.Err.Error())
}

func (e *SetError) Unwrap() error {
	var err *SetError
	if errors.As(e.Err, &err) {
		return err.Unwrap()
	}
	return e.Err
}

func RunInternalConfigSet(kvPair string) (err error) {
	pName, vKey, vValue, err := readConfigSetOptions(kvPair)
	if err != nil {
		return &SetError{Err: err}
	}

	if err = configuration.ValidateKoanfKey(vKey); err != nil {
		return &SetError{Err: err}
	}

	// Make sure value is not empty, and suggest unset command if it is
	if vValue == "" {
		return &SetError{Err: ErrEmptyValue}
	}

	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &SetError{Err: err}
	}

	subKoanf, err := koanfConfig.GetProfileKoanf(pName)
	if err != nil {
		return &SetError{Err: err}
	}

	opt, err := configuration.OptionFromKoanfKey(vKey)
	if err != nil {
		return &SetError{Err: err}
	}

	if err = setValue(subKoanf, vKey, vValue, opt.Type); err != nil {
		return &SetError{Err: err}
	}

	if err = koanfConfig.SaveProfile(pName, subKoanf); err != nil {
		return &SetError{Err: err}
	}

	msgStr := "Configuration set successfully:\n"

	vVal, _, err := profiles.KoanfValueFromOption(opt, pName)
	if err != nil {
		return &SetError{Err: err}
	}

	unmaskOptionVal, err := profiles.GetOptionValue(options.ConfigUnmaskSecretValueOption)
	if err != nil {
		unmaskOptionVal = "false"
	}

	if opt.Sensitive && strings.EqualFold(unmaskOptionVal, "false") {
		msgStr += fmt.Sprintf("%s=%s", vKey, profiles.MaskValue(vVal))
	} else {
		msgStr += fmt.Sprintf("%s=%s", vKey, vVal)
	}

	output.Success(msgStr, nil)

	return nil
}

func readConfigSetOptions(kvPair string) (pName string, vKey string, vValue string, err error) {
	if pName, err = readConfigSetProfileName(); err != nil {
		return pName, vKey, vValue, &SetError{Err: err}
	}

	if vKey, vValue, err = parseKeyValuePair(kvPair); err != nil {
		return pName, vKey, vValue, &SetError{Err: err}
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
		return pName, &SetError{Err: err}
	}

	if pName == "" {
		return pName, &SetError{Err: ErrDetermineProfileSet}
	}

	return pName, nil
}

func parseKeyValuePair(kvPair string) (key string, value string, err error) {
	parsedInput := strings.SplitN(kvPair, "=", 2)
	if len(parsedInput) < 2 {
		return key, value, &SetError{Err: ErrKeyAssignmentFormat}
	}

	if strings.EqualFold(parsedInput[0], options.RootActiveProfileOption.KoanfKey) {
		return key, value, &SetError{Err: ErrActiveProfileAssignment}
	}

	return parsedInput[0], parsedInput[1], nil
}

func setValue(profileKoanf *koanf.Koanf, vKey, vValue string, valueType options.OptionType) (err error) {
	switch valueType {
	case options.BOOL:
		b := new(customtypes.Bool)
		if err = b.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBeBoolean, err)
		}
		err = profileKoanf.Set(vKey, b)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.EXPORT_FORMAT:
		exportFormat := new(customtypes.ExportFormat)
		if err = exportFormat.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBeExportFormat, err)
		}
		err = profileKoanf.Set(vKey, exportFormat)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.EXPORT_SERVICE_GROUP:
		exportServiceGroup := new(customtypes.ExportServiceGroup)
		if err = exportServiceGroup.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBeExportServiceGroup, err)
		}
		err = profileKoanf.Set(vKey, exportServiceGroup)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.EXPORT_SERVICES:
		exportServices := new(customtypes.ExportServices)
		if err = exportServices.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBeExportService, err)
		}
		err = profileKoanf.Set(vKey, exportServices)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.OUTPUT_FORMAT:
		outputFormat := new(customtypes.OutputFormat)
		if err = outputFormat.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBeOutputFormat, err)
		}
		err = profileKoanf.Set(vKey, outputFormat)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.PINGONE_REGION_CODE:
		region := new(customtypes.PingOneRegionCode)
		if err = region.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBePingoneRegionCode, err)
		}
		err = profileKoanf.Set(vKey, region)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.STRING:
		str := new(customtypes.String)
		if err = str.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBeString, err)
		}
		err = profileKoanf.Set(vKey, str)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.STRING_SLICE:
		strSlice := new(customtypes.StringSlice)
		if err = strSlice.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBeStringSlice, err)
		}
		err = profileKoanf.Set(vKey, strSlice)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.UUID:
		uuid := new(customtypes.UUID)
		if err = uuid.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBeUUID, err)
		}
		err = profileKoanf.Set(vKey, uuid)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.PINGONE_AUTH_TYPE:
		authType := new(customtypes.PingOneAuthenticationType)
		if err = authType.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBePingoneAuthType, err)
		}
		err = profileKoanf.Set(vKey, authType)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.PINGFEDERATE_AUTH_TYPE:
		authType := new(customtypes.PingFederateAuthenticationType)
		if err = authType.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBePingfederateAuthType, err)
		}
		err = profileKoanf.Set(vKey, authType)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.INT:
		intValue := new(customtypes.Int)
		if err = intValue.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBeInteger, err)
		}
		err = profileKoanf.Set(vKey, intValue)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.REQUEST_HTTP_METHOD:
		httpMethod := new(customtypes.HTTPMethod)
		if err = httpMethod.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBeHttpMethod, err)
		}
		err = profileKoanf.Set(vKey, httpMethod)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.REQUEST_SERVICE:
		service := new(customtypes.RequestService)
		if err = service.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBeRequestService, err)
		}
		err = profileKoanf.Set(vKey, service)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.LICENSE_PRODUCT:
		licenseProduct := new(customtypes.LicenseProduct)
		if err = licenseProduct.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBeLicenseProduct, err)
		}
		err = profileKoanf.Set(vKey, licenseProduct)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	case options.LICENSE_VERSION:
		licenseVersion := new(customtypes.LicenseVersion)
		if err = licenseVersion.Set(vValue); err != nil {
			return fmt.Errorf("%w: %v", ErrMustBeLicenseVersion, err)
		}
		err = profileKoanf.Set(vKey, licenseVersion)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSetKey, err)
		}
	default:
		return &SetError{Err: ErrTypeNotRecognized}
	}

	return nil
}
