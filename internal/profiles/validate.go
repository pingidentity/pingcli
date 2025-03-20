// Copyright © 2025 Ping Identity Corporation

package profiles

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/viper"
)

func Validate() (err error) {
	// Get a slice of all profile names configured in the config.yaml file
	profileNames := GetMainConfig().ProfileNames()

	// Iterate over the profileNames and convert each to lowercase
	for i := range profileNames {
		profileNames[i] = strings.ToLower(profileNames[i])
	}

	// Validate profile names
	if err = validateProfileNames(profileNames); err != nil {
		return err
	}

	// Make sure selected active profile is in the configuration file
	activeProfileName, err := GetOptionValue(options.RootActiveProfileOption)
	activeProfileName = strings.ToLower(activeProfileName)
	if err != nil {
		return fmt.Errorf("failed to validate Ping CLI configuration: %v", err)
	}
	if !slices.Contains(profileNames, activeProfileName) {
		return fmt.Errorf("failed to validate Ping CLI configuration: active profile '%s' not found in configuration "+
			"file %s", activeProfileName, GetMainConfig().ViperInstance().ConfigFileUsed())
	}

	// for each profile key, validate the profile viper
	for _, pName := range profileNames {
		subViper, err := GetMainConfig().GetProfileViper(pName)
		if err != nil {
			return fmt.Errorf("failed to validate Ping CLI configuration: %v", err)
		}

		if err := validateProfileKeys(pName, subViper); err != nil {
			return fmt.Errorf("failed to validate Ping CLI configuration: %v", err)
		}

		if err := validateProfileValues(pName, subViper); err != nil {
			return fmt.Errorf("failed to validate Ping CLI configuration: %v", err)
		}
	}

	return nil
}

func validateProfileNames(profileNames []string) error {
	for _, profileName := range profileNames {
		if err := GetMainConfig().ValidateProfileNameFormat(profileName); err != nil {
			return err
		}
	}
	return nil
}

func validateProfileKeys(profileName string, profileViper *viper.Viper) error {
	validProfileKeys := configuration.ViperKeys()

	// Get all keys viper has loaded from config file.
	// If a key found in the config file is not in the viperKeys list,
	// it is an invalid key.
	var invalidKeys []string
	for _, key := range profileViper.AllKeys() {
		if !slices.ContainsFunc(validProfileKeys, func(v string) bool {
			return strings.EqualFold(v, key)
		}) {
			invalidKeys = append(invalidKeys, key)
		}
	}

	if len(invalidKeys) > 0 {
		invalidKeysStr := strings.Join(invalidKeys, ", ")
		validKeysStr := strings.Join(validProfileKeys, ", ")
		return fmt.Errorf("invalid configuration key(s) found in profile %s: %s\nMust use one of: %s", profileName, invalidKeysStr, validKeysStr)
	}
	return nil
}

func validateProfileValues(pName string, profileViper *viper.Viper) (err error) {
	for _, key := range profileViper.AllKeys() {
		opt, err := configuration.OptionFromViperKey(key)
		if err != nil {
			return err
		}

		vValue := profileViper.Get(key)

		switch opt.Type {
		case options.ENUM_BOOL:
			switch typedValue := vValue.(type) {
			case *customtypes.Bool:
				continue
			case string:
				b := new(customtypes.Bool)
				if err = b.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a boolean value: %v", pName, typedValue, key, err)
				}
			case bool:
				continue
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a boolean value", pName, typedValue, key)
			}
		case options.ENUM_UUID:
			switch typedValue := vValue.(type) {
			case *customtypes.UUID:
				continue
			case string:
				u := new(customtypes.UUID)
				if err = u.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a UUID value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a UUID value", pName, typedValue, key)
			}
		case options.ENUM_OUTPUT_FORMAT:
			switch typedValue := vValue.(type) {
			case *customtypes.OutputFormat:
				continue
			case string:
				o := new(customtypes.OutputFormat)
				if err = o.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not an output format value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not an output format value", pName, typedValue, key)
			}
		case options.ENUM_PINGONE_REGION_CODE:
			switch typedValue := vValue.(type) {
			case *customtypes.PingOneRegionCode:
				continue
			case string:
				prc := new(customtypes.PingOneRegionCode)
				if err = prc.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a PingOne Region Code value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a PingOne Region Code value", pName, typedValue, key)
			}
		case options.ENUM_STRING:
			switch typedValue := vValue.(type) {
			case *customtypes.String:
				continue
			case string:
				s := new(customtypes.String)
				if err = s.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a string value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a string value", pName, typedValue, key)
			}
		case options.ENUM_STRING_SLICE:
			switch typedValue := vValue.(type) {
			case *customtypes.StringSlice:
				continue
			case string:
				ss := new(customtypes.StringSlice)
				if err = ss.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a string slice value: %v", pName, typedValue, key, err)
				}
			case []any:
				ss := new(customtypes.StringSlice)
				for _, v := range typedValue {
					switch innerTypedValue := v.(type) {
					case string:
						if err = ss.Set(innerTypedValue); err != nil {
							return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a string slice value: %v", pName, typedValue, key, err)
						}
					default:
						return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a string slice value", pName, typedValue, key)
					}
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a string slice value", pName, typedValue, key)
			}
		case options.ENUM_EXPORT_SERVICES:
			switch typedValue := vValue.(type) {
			case *customtypes.ExportServices:
				continue
			case string:
				es := new(customtypes.ExportServices)
				if err = es.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a export service value: %v", pName, typedValue, key, err)
				}
			case []any:
				es := new(customtypes.ExportServices)
				for _, v := range typedValue {
					switch innerTypedValue := v.(type) {
					case string:
						if err = es.Set(innerTypedValue); err != nil {
							return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a export service value: %v", pName, typedValue, key, err)
						}
					default:
						return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a export service value", pName, typedValue, key)
					}

				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a export service value", pName, typedValue, key)
			}
		case options.ENUM_EXPORT_FORMAT:
			switch typedValue := vValue.(type) {
			case *customtypes.ExportFormat:
				continue
			case string:
				ef := new(customtypes.ExportFormat)
				if err = ef.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not an export format value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not an export format value", pName, typedValue, key)
			}
		case options.ENUM_REQUEST_HTTP_METHOD:
			switch typedValue := vValue.(type) {
			case *customtypes.HTTPMethod:
				continue
			case string:
				hm := new(customtypes.HTTPMethod)
				if err = hm.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not an HTTP method value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not an HTTP method value", pName, typedValue, key)
			}
		case options.ENUM_REQUEST_SERVICE:
			switch typedValue := vValue.(type) {
			case *customtypes.RequestService:
				continue
			case string:
				rs := new(customtypes.RequestService)
				if err = rs.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a request service value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a request service value", pName, typedValue, key)
			}
		case options.ENUM_INT:
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
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not an int value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not an int value", pName, typedValue, key)
			}
		case options.ENUM_PINGFEDERATE_AUTH_TYPE:
			switch typedValue := vValue.(type) {
			case *customtypes.PingFederateAuthenticationType:
				continue
			case string:
				pfa := new(customtypes.PingFederateAuthenticationType)
				if err = pfa.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a PingFederate Authentication Type value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a PingFederate Authentication Type value", pName, typedValue, key)
			}
		case options.ENUM_PINGONE_AUTH_TYPE:
			switch typedValue := vValue.(type) {
			case *customtypes.PingOneAuthenticationType:
				continue
			case string:
				pat := new(customtypes.PingOneAuthenticationType)
				if err = pat.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a PingOne Authentication Type value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a PingOne Authentication Type value", pName, typedValue, key)
			}
		default:
			return fmt.Errorf("profile '%s': variable type '%s' for key '%s' is not recognized", pName, opt.Type, key)
		}
	}

	return nil
}
