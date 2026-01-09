// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/customtypes"
)

var (
	// Common errors
	ErrUndeterminedProfile = errors.New("unable to determine configuration profile")

	// Add profile errors
	ErrNoProfileProvided    = errors.New("unable to determine profile name")
	ErrSetActiveFlagInvalid = errors.New("invalid value for set-active flag. must be 'true' or 'false'")
	ErrKoanfNotInitialized  = errors.New("koanf configuration not initialized")

	// List keys errors
	ErrRetrieveKeys = errors.New("failed to retrieve configuration keys")
	ErrNestedMap    = errors.New("failed to create nested map for key")
	ErrMarshalKeys  = errors.New("failed to marshal keys to YAML format")

	// Set errors
	ErrEmptyValue                 = errors.New("the set value provided is empty. Use 'pingcli config unset %s' to unset a key's configuration")
	ErrKeyAssignmentFormat        = errors.New("invalid key-value assignment. Expect 'key=value' format")
	ErrActiveProfileAssignment    = errors.New("invalid active profile assignment. Please use the 'pingcli config set-active-profile <profile-name>' command to set the active profile")
	ErrSetKey                     = errors.New("unable to set key in configuration profile")
	ErrMustBeBoolean              = errors.New("the value assignment must be a boolean. Allowed [true, false]")
	ErrMustBeExportFormat         = fmt.Errorf("the value assignment must be a valid export format. Allowed [%s]", strings.Join(customtypes.ExportFormatValidValues(), ", "))
	ErrMustBeExportServiceGroup   = fmt.Errorf("the value assignment must be a valid export service group. Allowed [%s]", strings.Join(customtypes.ExportServiceGroupValidValues(), ", "))
	ErrMustBeExportService        = fmt.Errorf("the value assignment must be valid export service(s). Allowed [%s]", strings.Join(customtypes.ExportServicesValidValues(), ", "))
	ErrMustBeOutputFormat         = fmt.Errorf("the value assignment must be a valid output format. Allowed [%s]", strings.Join(customtypes.OutputFormatValidValues(), ", "))
	ErrMustBePingoneRegionCode    = fmt.Errorf("the value assignment must be a valid PingOne region code. Allowed [%s]", strings.Join(customtypes.PingOneRegionCodeValidValues(), ", "))
	ErrMustBeString               = errors.New("the value assignment must be a string")
	ErrMustBeStringSlice          = errors.New("the value assignment must be a string slice")
	ErrMustBeUUID                 = errors.New("the value assignment must be a valid UUID")
	ErrMustBePingoneAuthType      = fmt.Errorf("the value assignment must be a valid PingOne Authentication Type. Allowed [%s]", strings.Join(customtypes.PingOneAuthenticationTypeValidValues(), ", "))
	ErrMustBePingfederateAuthType = fmt.Errorf("the value assignment must be a valid PingFederate Authentication Type. Allowed [%s]", strings.Join(customtypes.PingFederateAuthenticationTypeValidValues(), ", "))
	ErrMustBeInteger              = errors.New("the value assignment must be an integer")
	ErrMustBeHttpMethod           = fmt.Errorf("the value assignment must be a valid HTTP method. Allowed [%s]", strings.Join(customtypes.HTTPMethodValidValues(), ", "))
	ErrMustBeRequestService       = fmt.Errorf("the value assignment must be a valid request service. Allowed [%s]", strings.Join(customtypes.RequestServiceValidValues(), ", "))
	ErrMustBeLicenseProduct       = fmt.Errorf("must be one of: %s", strings.Join(customtypes.LicenseProductValidValues(), ", "))
	ErrMustBeLicenseVersion       = errors.New("the value assignment must be a valid license version. Must be of the form 'major.minor'")
	ErrMustBeStorageType          = fmt.Errorf("must be one of: %s", strings.Join(customtypes.StorageTypeValidValues(), ", "))

	ErrTypeNotRecognized = errors.New("the variable type for the configuration key is not recognized or supported")
)
