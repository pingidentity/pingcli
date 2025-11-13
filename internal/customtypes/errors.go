// Copyright © 2025 Ping Identity Corporation

package customtypes

import "errors"

var (
	ErrCustomTypeNil                 = errors.New("failed to set value. An internal error occurred")
	ErrParseBool                     = errors.New("failed to parse value as bool")
	ErrParseInt                      = errors.New("failed to parse value as int")
	ErrInvalidUUID                   = errors.New("invalid uuid")
	ErrInvalidHeaderFormat           = errors.New("invalid header format. must be in `key:value` format")
	ErrDisallowedAuthHeader          = errors.New("authorization header is not allowed")
	ErrUnrecognizedMethod            = errors.New("unrecognized http method")
	ErrUnrecognizedService           = errors.New("unrecognized request service")
	ErrUnrecognizedOutputFormat      = errors.New("unrecognized output format")
	ErrUnrecognizedPingOneRegionCode = errors.New("unrecognized pingone region code")
	ErrUnrecognizedPingOneAuth       = errors.New("unrecognized pingone authentication type")
	ErrUnrecognizedPingFederateAuth  = errors.New("unrecognized pingfederate authentication type")
	ErrUnrecognizedProduct           = errors.New("unrecognized license product")
	ErrInvalidVersionFormat          = errors.New("invalid version format, must be 'major.minor'")
	ErrUnrecognizedFormat            = errors.New("unrecognized export format")
	ErrUnrecognizedServiceGroup      = errors.New("unrecognized service group")
	ErrUnrecognizedExportService     = errors.New("unrecognized service")
	ErrUnrecognizedAuthService       = errors.New("unrecognized authentication service")
)
