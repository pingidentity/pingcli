// Copyright © 2025 Ping Identity Corporation

package profiles

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/knadh/koanf/v2"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/errs"
)

var (
	validateErrorPrefix             = "profile validation error"
	ErrValidatePingCLIConfiguration = errors.New("failed to validate Ping CLI configuration")
	ErrInvalidConfigurationKey      = errors.New("invalid configuration key(s) found in profile")
	ErrUnrecognizedVariableType     = errors.New("unrecognized variable type for key")
	ErrValidateBoolean              = errors.New("invalid boolean value")
	ErrValidateUUID                 = errors.New("invalid uuid value")
	ErrValidateOutputFormat         = errors.New("invalid output format value")
	ErrValidatePingOneRegionCode    = errors.New("invalid pingone region code value")
	ErrValidateString               = errors.New("invalid string value")
	ErrValidateStringSlice          = errors.New("invalid string slice value")
	ErrValidateExportServiceGroup   = errors.New("invalid export service group value")
	ErrValidateExportServices       = errors.New("invalid export services value")
	ErrValidateExportFormat         = errors.New("invalid export format value")
	ErrValidateHTTPMethod           = errors.New("invalid http method value")
	ErrValidateRequestService       = errors.New("invalid request service value")
	ErrValidateInt                  = errors.New("invalid int value")
	ErrValidatePingFederateAuthType = errors.New("invalid pingfederate auth type value")
	ErrValidatePingOneAuthType      = errors.New("invalid pingone auth type value")
	ErrValidateLicenseProduct       = errors.New("invalid license product value")
	ErrValidateLicenseVersion       = errors.New("invalid license version value")
)

func Validate() (err error) {
	koanfConfig, err := GetKoanfConfig()
	if err != nil {
		return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: err}
	}

	// Get a slice of all profile names configured in the config.yaml file
	profileNames := koanfConfig.ProfileNames()

	// Validate profile names
	if err = validateProfileNames(profileNames); err != nil {
		return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: err}
	}

	profileName, err := GetOptionValue(options.RootProfileOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: err}
	}

	if profileName != "" {
		// Make sure selected profile is in the configuration file
		if !slices.Contains(profileNames, profileName) {
			return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("%w: '%s' profile not found in configuration "+
				"file %s", ErrValidatePingCLIConfiguration, profileName, koanfConfig.GetKoanfConfigFile())}
		}
	}

	// active profile has no env var or cobra flag, so always get from config file
	activeProfileName, ok, err := KoanfValueFromOption(options.RootActiveProfileOption, "")
	if err != nil {
		return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: err}
	}
	if !ok {
		return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("%w: active profile not set in configuration file %s",
			ErrValidatePingCLIConfiguration, koanfConfig.GetKoanfConfigFile())}
	}

	// Make sure selected active profile is in the configuration file
	if !slices.Contains(profileNames, activeProfileName) {
		return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("%w: active profile '%s' not found in configuration "+
			"file %s", ErrValidatePingCLIConfiguration, activeProfileName, koanfConfig.GetKoanfConfigFile())}
	}

	// for each profile key, validate the profile koanf
	for _, pName := range profileNames {
		subKoanf, err := koanfConfig.GetProfileKoanf(pName)
		if err != nil {
			return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: err}
		}

		if err := validateProfileKeys(pName, subKoanf); err != nil {
			return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: err}
		}

		if err := validateProfileValues(pName, subKoanf); err != nil {
			return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: err}
		}
	}

	return nil
}

func validateProfileNames(profileNames []string) error {
	koanfConfig, err := GetKoanfConfig()
	if err != nil {
		return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: err}
	}

	for _, profileName := range profileNames {
		if err := koanfConfig.ValidateProfileNameFormat(profileName); err != nil {
			return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: err}
		}
	}

	return nil
}

func validateProfileKeys(profileName string, profileKoanf *koanf.Koanf) error {
	validProfileKeys := configuration.KoanfKeys()

	// Get all keys koanf has loaded from config file.
	// If a key found in the config file is not in the koanfKeys list,
	// it is an invalid key.
	var invalidKeys []string
	for key := range profileKoanf.All() {
		if !slices.ContainsFunc(validProfileKeys, func(v string) bool {
			return v == key
		}) {
			invalidKeys = append(invalidKeys, key)
		}
	}

	if len(invalidKeys) > 0 {
		invalidKeysStr := strings.Join(invalidKeys, ", ")
		validKeysStr := strings.Join(validProfileKeys, ", ")

		return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("%w %s: %s\nMust use one of: %s", ErrInvalidConfigurationKey, profileName, invalidKeysStr, validKeysStr)}
	}

	return nil
}

func validateProfileValues(pName string, profileKoanf *koanf.Koanf) (err error) {
	for key := range profileKoanf.All() {
		opt, err := configuration.OptionFromKoanfKey(key)
		if err != nil {
			return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: err}
		}

		vValue := profileKoanf.Get(key)

		switch opt.Type {
		case options.BOOL:
			switch typedValue := vValue.(type) {
			case *customtypes.Bool:
				continue
			case string:
				b := new(customtypes.Bool)
				if err = b.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateBoolean, typedValue, err)}
				}
			case bool:
				continue
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateBoolean, typedValue, typedValue)}
			}
		case options.UUID:
			switch typedValue := vValue.(type) {
			case *customtypes.UUID:
				continue
			case string:
				u := new(customtypes.UUID)
				if err = u.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateUUID, typedValue, err)}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateUUID, typedValue, typedValue)}
			}
		case options.OUTPUT_FORMAT:
			switch typedValue := vValue.(type) {
			case *customtypes.OutputFormat:
				continue
			case string:
				o := new(customtypes.OutputFormat)
				if err = o.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateOutputFormat, typedValue, err)}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateOutputFormat, typedValue, typedValue)}
			}
		case options.PINGONE_REGION_CODE:
			switch typedValue := vValue.(type) {
			case *customtypes.PingOneRegionCode:
				continue
			case string:
				prc := new(customtypes.PingOneRegionCode)
				if err = prc.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidatePingOneRegionCode, typedValue, err)}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidatePingOneRegionCode, typedValue, typedValue)}
			}
		case options.STRING:
			switch typedValue := vValue.(type) {
			case *customtypes.String:
				continue
			case string:
				s := new(customtypes.String)
				if err = s.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateString, typedValue, err)}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateString, typedValue, typedValue)}
			}
		case options.STRING_SLICE:
			switch typedValue := vValue.(type) {
			case *customtypes.StringSlice:
				continue
			case string:
				ss := new(customtypes.StringSlice)
				if err = ss.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateStringSlice, typedValue, err)}
				}
			case []any:
				ss := new(customtypes.StringSlice)
				for _, v := range typedValue {
					switch innerTypedValue := v.(type) {
					case string:
						if err = ss.Set(innerTypedValue); err != nil {
							return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateStringSlice, typedValue, err)}
						}
					default:
						return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateStringSlice, typedValue, typedValue)}
					}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateStringSlice, typedValue, typedValue)}
			}
		case options.EXPORT_SERVICE_GROUP:
			switch typedValue := vValue.(type) {
			case *customtypes.ExportServiceGroup:
				continue
			case string:
				esg := new(customtypes.ExportServiceGroup)
				if err = esg.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateExportServiceGroup, typedValue, err)}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateExportServiceGroup, typedValue, typedValue)}
			}
		case options.EXPORT_SERVICES:
			switch typedValue := vValue.(type) {
			case *customtypes.ExportServices:
				continue
			case string:
				es := new(customtypes.ExportServices)
				if err = es.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateExportServices, typedValue, err)}
				}
			case []any:
				es := new(customtypes.ExportServices)
				for _, v := range typedValue {
					switch innerTypedValue := v.(type) {
					case string:
						if err = es.Set(innerTypedValue); err != nil {
							return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateExportServices, typedValue, err)}
						}
					default:
						return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateExportServices, typedValue, typedValue)}
					}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateExportServices, typedValue, typedValue)}
			}
		case options.EXPORT_FORMAT:
			switch typedValue := vValue.(type) {
			case *customtypes.ExportFormat:
				continue
			case string:
				ef := new(customtypes.ExportFormat)
				if err = ef.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateExportFormat, typedValue, err)}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateExportFormat, typedValue, typedValue)}
			}
		case options.REQUEST_HTTP_METHOD:
			switch typedValue := vValue.(type) {
			case *customtypes.HTTPMethod:
				continue
			case string:
				hm := new(customtypes.HTTPMethod)
				if err = hm.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateHTTPMethod, typedValue, err)}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateHTTPMethod, typedValue, typedValue)}
			}
		case options.REQUEST_SERVICE:
			switch typedValue := vValue.(type) {
			case *customtypes.RequestService:
				continue
			case string:
				rs := new(customtypes.RequestService)
				if err = rs.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateRequestService, typedValue, err)}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateRequestService, typedValue, typedValue)}
			}
		case options.INT:
			switch typedValue := vValue.(type) {
			case *customtypes.Int:
				continue
			case int:
				continue
			case int64:
				continue
			case string:
				i := new(customtypes.Int)
				if err = i.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateInt, typedValue, err)}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateInt, typedValue, typedValue)}
			}
		case options.PINGFEDERATE_AUTH_TYPE:
			switch typedValue := vValue.(type) {
			case *customtypes.PingFederateAuthenticationType:
				continue
			case string:
				pfa := new(customtypes.PingFederateAuthenticationType)
				if err = pfa.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidatePingFederateAuthType, typedValue, err)}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidatePingFederateAuthType, typedValue, typedValue)}
			}
		case options.PINGONE_AUTH_TYPE:
			switch typedValue := vValue.(type) {
			case *customtypes.PingOneAuthenticationType:
				continue
			case string:
				pat := new(customtypes.PingOneAuthenticationType)
				if err = pat.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidatePingOneAuthType, typedValue, err)}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidatePingOneAuthType, typedValue, typedValue)}
			}
		case options.LICENSE_PRODUCT:
			switch typedValue := vValue.(type) {
			case *customtypes.LicenseProduct:
				continue
			case string:
				lp := new(customtypes.LicenseProduct)
				if err = lp.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateLicenseProduct, typedValue, err)}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateLicenseProduct, typedValue, typedValue)}
			}
		case options.LICENSE_VERSION:
			switch typedValue := vValue.(type) {
			case *customtypes.LicenseVersion:
				continue
			case string:
				lv := new(customtypes.LicenseVersion)
				if err = lv.Set(typedValue); err != nil {
					return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s': %w", pName, ErrValidateLicenseVersion, typedValue, err)}
				}
			default:
				return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("profile '%s': %w '%s' of type '%T'", pName, ErrValidateLicenseVersion, typedValue, typedValue)}
			}
		default:
			return &errs.PingCLIError{Prefix: validateErrorPrefix, Err: fmt.Errorf("%w: %d", ErrUnrecognizedVariableType, opt.Type)}
		}
	}

	return nil
}
