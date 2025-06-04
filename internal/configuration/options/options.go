// Copyright © 2025 Ping Identity Corporation

package options

import (
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

type OptionType int

// OptionType enums
const (
	ENUM_BOOL OptionType = iota
	ENUM_EXPORT_FORMAT
	ENUM_HEADER
	ENUM_INT
	ENUM_EXPORT_SERVICE_GROUP
	ENUM_EXPORT_SERVICES
	ENUM_OUTPUT_FORMAT
	ENUM_PINGFEDERATE_AUTH_TYPE
	ENUM_PINGONE_AUTH_TYPE
	ENUM_PINGONE_REGION_CODE
	ENUM_REQUEST_HTTP_METHOD
	ENUM_REQUEST_SERVICE
	ENUM_STRING
	ENUM_STRING_SLICE
	ENUM_UUID
)

var optionTypeString = map[OptionType]string{
	ENUM_BOOL:                   "ENUM_BOOL",
	ENUM_EXPORT_FORMAT:          "ENUM_EXPORT_FORMAT",
	ENUM_HEADER:                 "ENUM_HEADER",
	ENUM_INT:                    "ENUM_INT",
	ENUM_EXPORT_SERVICE_GROUP:   "ENUM_EXPORT_SERVICE_GROUP",
	ENUM_EXPORT_SERVICES:        "ENUM_EXPORT_SERVICES",
	ENUM_OUTPUT_FORMAT:          "ENUM_OUTPUT_FORMAT",
	ENUM_PINGFEDERATE_AUTH_TYPE: "ENUM_PINGFEDERATE_AUTH_TYPE",
	ENUM_PINGONE_AUTH_TYPE:      "ENUM_PINGONE_AUTH_TYPE",
	ENUM_PINGONE_REGION_CODE:    "ENUM_PINGONE_REGION_CODE",
	ENUM_REQUEST_HTTP_METHOD:    "ENUM_REQUEST_HTTP_METHOD",
	ENUM_REQUEST_SERVICE:        "ENUM_REQUEST_SERVICE",
	ENUM_STRING:                 "ENUM_STRING",
	ENUM_STRING_SLICE:           "ENUM_STRING_SLICE",
	ENUM_UUID:                   "ENUM_UUID",
}

var optionTypeFriendlyString = map[OptionType]string{
	ENUM_BOOL:                   "Boolean",
	ENUM_EXPORT_FORMAT:          "String (enum)",
	ENUM_HEADER:                 "Boolean",
	ENUM_INT:                    "Number",
	ENUM_EXPORT_SERVICE_GROUP:   "String (enum)",
	ENUM_EXPORT_SERVICES:        "String Array (enum)",
	ENUM_OUTPUT_FORMAT:          "String (enum)",
	ENUM_PINGFEDERATE_AUTH_TYPE: "String (enum)",
	ENUM_PINGONE_AUTH_TYPE:      "String (enum)",
	ENUM_PINGONE_REGION_CODE:    "String (enum)",
	ENUM_REQUEST_HTTP_METHOD:    "String (enum)",
	ENUM_REQUEST_SERVICE:        "String (enum)",
	ENUM_STRING:                 "String",
	ENUM_STRING_SLICE:           "String Array",
	ENUM_UUID:                   "String (UUID Format)",
}

func (e OptionType) String() string {
	if s, ok := optionTypeString[e]; ok {
		return s
	}

	return "ENUM_UNKNOWN"
}

func (e OptionType) FriendlyString() string {
	if s, ok := optionTypeFriendlyString[e]; ok {
		return s
	}

	return "Unknown"
}

type Option struct {
	CobraParamName  string
	CobraParamValue pflag.Value
	DefaultValue    pflag.Value
	EnvVar          string
	Flag            *pflag.Flag
	Sensitive       bool
	Type            OptionType
	KoanfKey        string
}

func Options() []Option {
	optList := []Option{
		PingOneAuthenticationTypeOption,
		PingOneAuthenticationWorkerClientIDOption,
		PingOneAuthenticationWorkerClientSecretOption,
		PingOneAuthenticationWorkerEnvironmentIDOption,
		PingOneRegionCodeOption,

		PlatformExportExportFormatOption,
		PlatformExportServiceGroupOption,
		PlatformExportServiceOption,
		PlatformExportOutputDirectoryOption,
		PlatformExportOverwriteOption,
		PlatformExportPingOneEnvironmentIDOption,

		PingFederateHTTPSHostOption,
		PingFederateAdminAPIPathOption,
		PingFederateXBypassExternalValidationHeaderOption,
		PingFederateCACertificatePemFilesOption,
		PingFederateInsecureTrustAllTLSOption,
		PingFederateBasicAuthUsernameOption,
		PingFederateBasicAuthPasswordOption,
		PingFederateAccessTokenAuthAccessTokenOption,
		PingFederateClientCredentialsAuthClientIDOption,
		PingFederateClientCredentialsAuthClientSecretOption,
		PingFederateClientCredentialsAuthTokenURLOption,
		PingFederateClientCredentialsAuthScopesOption,
		PingFederateAuthenticationTypeOption,

		RootActiveProfileOption,
		RootProfileOption,
		RootColorOption,
		RootConfigOption,
		RootDetailedExitCodeOption,
		RootOutputFormatOption,

		ProfileDescriptionOption,

		ConfigAddProfileDescriptionOption,
		ConfigAddProfileNameOption,
		ConfigAddProfileSetActiveOption,
		ConfigDeleteAutoAcceptOption,
		ConfigListKeysYamlOption,
		ConfigUnmaskSecretValueOption,

		RequestDataOption,
		RequestDataRawOption,
		RequestHeaderOption,
		RequestHTTPMethodOption,
		RequestServiceOption,
		RequestAccessTokenOption,
		RequestAccessTokenExpiryOption,
		RequestFailOption,
	}

	// Sort the options list by koanf key
	slices.SortFunc(optList, func(opt1, opt2 Option) int {
		return strings.Compare(opt1.KoanfKey, opt2.KoanfKey)
	})

	return optList
}

// pingone service options
var (
	PingOneAuthenticationTypeOption                Option
	PingOneAuthenticationWorkerClientIDOption      Option
	PingOneAuthenticationWorkerClientSecretOption  Option
	PingOneAuthenticationWorkerEnvironmentIDOption Option
	PingOneRegionCodeOption                        Option
)

// pingfederate service options
var (
	PingFederateHTTPSHostOption                         Option
	PingFederateAdminAPIPathOption                      Option
	PingFederateXBypassExternalValidationHeaderOption   Option
	PingFederateCACertificatePemFilesOption             Option
	PingFederateInsecureTrustAllTLSOption               Option
	PingFederateBasicAuthUsernameOption                 Option
	PingFederateBasicAuthPasswordOption                 Option
	PingFederateAccessTokenAuthAccessTokenOption        Option
	PingFederateClientCredentialsAuthClientIDOption     Option
	PingFederateClientCredentialsAuthClientSecretOption Option
	PingFederateClientCredentialsAuthTokenURLOption     Option
	PingFederateClientCredentialsAuthScopesOption       Option
	PingFederateAuthenticationTypeOption                Option
)

// 'pingcli config' command options
var (
	ConfigAddProfileDescriptionOption Option
	ConfigAddProfileNameOption        Option
	ConfigAddProfileSetActiveOption   Option

	ConfigListKeysYamlOption Option

	ConfigDeleteAutoAcceptOption Option

	ConfigUnmaskSecretValueOption Option
)

// 'pingcli platform export' command options
var (
	PlatformExportExportFormatOption         Option
	PlatformExportServiceOption              Option
	PlatformExportServiceGroupOption         Option
	PlatformExportOutputDirectoryOption      Option
	PlatformExportOverwriteOption            Option
	PlatformExportPingOneEnvironmentIDOption Option
)

// Generic koanf profile options
var (
	ProfileDescriptionOption Option
)

// Root Command Options
var (
	RootActiveProfileOption    Option
	RootDetailedExitCodeOption Option
	RootProfileOption          Option
	RootColorOption            Option
	RootConfigOption           Option
	RootOutputFormatOption     Option
)

// 'pingcli request' command options
var (
	RequestDataOption              Option
	RequestDataRawOption           Option
	RequestHeaderOption            Option
	RequestHTTPMethodOption        Option
	RequestServiceOption           Option
	RequestAccessTokenOption       Option
	RequestAccessTokenExpiryOption Option
	RequestFailOption              Option
)
