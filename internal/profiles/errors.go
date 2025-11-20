// Copyright © 2025 Ping Identity Corporation

package profiles

import "errors"

var (
	// Validation errors
	ErrValidatePingCLIConfiguration = errors.New("failed to validate Ping CLI configuration")
	ErrInvalidConfigurationKey      = errors.New("invalid configuration key(s) found in profile")
	ErrUnrecognizedVariableType     = errors.New("unrecognized variable type for key")
	ErrValidateBoolean              = errors.New("invalid boolean value")
	ErrValidateUUID                 = errors.New("invalid uuid value")
	ErrValidateOutputFormat         = errors.New("invalid output format value")
	ErrValidatePingOneRegionCode    = errors.New("invalid pingone region code value")
	ErrValidateString               = errors.New("invalid string value")
	ErrValidateStringSlice          = errors.New("invalid string slice value")
	ErrValidateAuthProvider         = errors.New("invalid auth provider value")
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

	// Koanf errors
	ErrNoOptionValue                     = errors.New("no option value found")
	ErrKoanfNotInitialized               = errors.New("koanf instance is not initialized")
	ErrProfileNameEmpty                  = errors.New("invalid profile name: profile name cannot be empty")
	ErrProfileNameFormat                 = errors.New("invalid profile name: profile name must contain only alphanumeric characters, underscores, and dashes")
	ErrProfileNameSameAsActiveProfileKey = errors.New("invalid profile name: profile name cannot be the same as the active profile key")
	ErrSetActiveProfile                  = errors.New("error setting active profile")
	ErrWriteKoanfFile                    = errors.New("failed to write configuration file to disk")
	ErrProfileNameNotExist               = errors.New("invalid profile name: profile name does not exist")
	ErrProfileNameAlreadyExists          = errors.New("invalid profile name: profile name already exists")
	ErrKoanfProfileExtractAndLoad        = errors.New("failed to extract and load profile configuration")
	ErrSetKoanfKeyValue                  = errors.New("failed to set koanf key value")
	ErrMarshalKoanf                      = errors.New("failed to marshal koanf configuration")
	ErrKoanfMerge                        = errors.New("failed to merge koanf configuration")
	ErrDeleteActiveProfile               = errors.New("the active profile cannot be deleted")
	ErrSetKoanfKeyDefaultValue           = errors.New("failed to set koanf key default value")
)
