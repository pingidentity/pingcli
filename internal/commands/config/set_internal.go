// Copyright © 2026 Ping Identity Corporation

package config_internal

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf/v2"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var (
	setErrorPrefix = "failed to set configuration"
)

func RunInternalConfigSet(kvPair string) (err error) {
	pName, vKey, vValue, err := readConfigSetOptions(kvPair)
	if err != nil {
		return &errs.PingCLIError{Prefix: setErrorPrefix, Err: err}
	}

	if err = configuration.ValidateKoanfKey(vKey); err != nil {
		return &errs.PingCLIError{Prefix: setErrorPrefix, Err: err}
	}

	// Make sure value is not empty, and suggest unset command if it is
	if vValue == "" {
		return &errs.PingCLIError{Prefix: setErrorPrefix, Err: ErrEmptyValue}
	}

	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &errs.PingCLIError{Prefix: setErrorPrefix, Err: err}
	}

	subKoanf, err := koanfConfig.GetProfileKoanf(pName)
	if err != nil {
		return &errs.PingCLIError{Prefix: setErrorPrefix, Err: err}
	}

	opt, err := configuration.OptionFromKoanfKey(vKey)
	if err != nil {
		return &errs.PingCLIError{Prefix: setErrorPrefix, Err: err}
	}

	if err = setValue(subKoanf, opt.KoanfKey, vValue, opt.Type); err != nil {
		return &errs.PingCLIError{Prefix: setErrorPrefix, Err: err}
	}

	if err = koanfConfig.SaveProfile(pName, subKoanf); err != nil {
		return &errs.PingCLIError{Prefix: setErrorPrefix, Err: err}
	}

	msgStr := "Configuration set successfully:\n"

	vVal, _, err := profiles.KoanfValueFromOption(opt, pName)
	if err != nil {
		return &errs.PingCLIError{Prefix: setErrorPrefix, Err: err}
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
		return pName, vKey, vValue, &errs.PingCLIError{Prefix: setErrorPrefix, Err: err}
	}

	if vKey, vValue, err = parseKeyValuePair(kvPair); err != nil {
		return pName, vKey, vValue, &errs.PingCLIError{Prefix: setErrorPrefix, Err: err}
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
		return pName, &errs.PingCLIError{Prefix: setErrorPrefix, Err: err}
	}

	if pName == "" {
		return pName, &errs.PingCLIError{Prefix: setErrorPrefix, Err: ErrUndeterminedProfile}
	}

	return pName, nil
}

func parseKeyValuePair(kvPair string) (key string, value string, err error) {
	parsedInput := strings.SplitN(kvPair, "=", 2)
	if len(parsedInput) < 2 {
		return key, value, &errs.PingCLIError{Prefix: setErrorPrefix, Err: ErrKeyAssignmentFormat}
	}

	if strings.EqualFold(parsedInput[0], options.RootActiveProfileOption.KoanfKey) {
		return key, value, &errs.PingCLIError{Prefix: setErrorPrefix, Err: ErrActiveProfileAssignment}
	}

	return parsedInput[0], parsedInput[1], nil
}

func setValue(profileKoanf *koanf.Koanf, vKey, vValue string, valueType options.OptionType) (err error) {
	switch valueType {
	case options.BOOL:
		b := new(customtypes.Bool)
		if err = b.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBeBoolean, err)
		}
		err = profileKoanf.Set(vKey, b)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.EXPORT_FORMAT:
		exportFormat := new(customtypes.ExportFormat)
		if err = exportFormat.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBeExportFormat, err)
		}
		err = profileKoanf.Set(vKey, exportFormat)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.EXPORT_SERVICE_GROUP:
		exportServiceGroup := new(customtypes.ExportServiceGroup)
		if err = exportServiceGroup.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBeExportServiceGroup, err)
		}
		err = profileKoanf.Set(vKey, exportServiceGroup)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.EXPORT_SERVICES:
		exportServices := new(customtypes.ExportServices)
		if err = exportServices.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBeExportService, err)
		}
		err = profileKoanf.Set(vKey, exportServices)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.OUTPUT_FORMAT:
		outputFormat := new(customtypes.OutputFormat)
		if err = outputFormat.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBeOutputFormat, err)
		}
		err = profileKoanf.Set(vKey, outputFormat)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.PINGONE_REGION_CODE:
		region := new(customtypes.PingOneRegionCode)
		if err = region.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBePingoneRegionCode, err)
		}
		err = profileKoanf.Set(vKey, region)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.STRING:
		str := new(customtypes.String)
		if err = str.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBeString, err)
		}
		err = profileKoanf.Set(vKey, str)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.STRING_SLICE:
		strSlice := new(customtypes.StringSlice)
		if err = strSlice.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBeStringSlice, err)
		}
		err = profileKoanf.Set(vKey, strSlice)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.UUID:
		uuid := new(customtypes.UUID)
		if err = uuid.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBeUUID, err)
		}
		err = profileKoanf.Set(vKey, uuid)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.PINGONE_AUTH_TYPE:
		authType := new(customtypes.PingOneAuthenticationType)
		if err = authType.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBePingoneAuthType, err)
		}
		err = profileKoanf.Set(vKey, authType)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.PINGFEDERATE_AUTH_TYPE:
		authType := new(customtypes.PingFederateAuthenticationType)
		if err = authType.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBePingfederateAuthType, err)
		}
		err = profileKoanf.Set(vKey, authType)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.INT:
		intValue := new(customtypes.Int)
		if err = intValue.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBeInteger, err)
		}
		err = profileKoanf.Set(vKey, intValue)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.REQUEST_HTTP_METHOD:
		httpMethod := new(customtypes.HTTPMethod)
		if err = httpMethod.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBeHttpMethod, err)
		}
		err = profileKoanf.Set(vKey, httpMethod)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.REQUEST_SERVICE:
		service := new(customtypes.RequestService)
		if err = service.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBeRequestService, err)
		}
		err = profileKoanf.Set(vKey, service)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.LICENSE_PRODUCT:
		licenseProduct := new(customtypes.LicenseProduct)
		if err = licenseProduct.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBeLicenseProduct, err)
		}
		err = profileKoanf.Set(vKey, licenseProduct)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.LICENSE_VERSION:
		licenseVersion := new(customtypes.LicenseVersion)
		if err = licenseVersion.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBeLicenseVersion, err)
		}
		err = profileKoanf.Set(vKey, licenseVersion)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	case options.STORAGE_TYPE:
		storageType := new(customtypes.StorageType)
		if err = storageType.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrMustBeStorageType, err)
		}
		err = profileKoanf.Set(vKey, storageType)
		if err != nil {
			return fmt.Errorf("value for key '%s': %w: %w", vKey, ErrSetKey, err)
		}
	default:
		return &errs.PingCLIError{Prefix: setErrorPrefix, Err: ErrTypeNotRecognized}
	}

	return nil
}
