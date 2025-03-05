package options

import (
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

type OptionType string

// OptionType enums
const (
	ENUM_BOOL                   OptionType = "ENUM_BOOL"
	ENUM_EXPORT_FORMAT          OptionType = "ENUM_EXPORT_FORMAT"
	ENUM_INT                    OptionType = "ENUM_INT"
	ENUM_EXPORT_SERVICE_GROUP   OptionType = "ENUM_EXPORT_SERVICE_GROUP"
	ENUM_EXPORT_SERVICES        OptionType = "ENUM_EXPORT_SERVICES"
	ENUM_OUTPUT_FORMAT          OptionType = "ENUM_OUTPUT_FORMAT"
	ENUM_PINGFEDERATE_AUTH_TYPE OptionType = "ENUM_PINGFEDERATE_AUTH_TYPE"
	ENUM_PINGONE_AUTH_TYPE      OptionType = "ENUM_PINGONE_AUTH_TYPE"
	ENUM_PINGONE_REGION_CODE    OptionType = "ENUM_PINGONE_REGION_CODE"
	ENUM_REQUEST_HTTP_METHOD    OptionType = "ENUM_REQUEST_HTTP_METHOD"
	ENUM_REQUEST_SERVICE        OptionType = "ENUM_REQUEST_SERVICE"
	ENUM_STRING                 OptionType = "ENUM_STRING"
	ENUM_STRING_SLICE           OptionType = "ENUM_STRING_SLICE"
	ENUM_UUID                   OptionType = "ENUM_UUID"
)

type Option struct {
	CobraParamName  string
	CobraParamValue pflag.Value
	DefaultValue    pflag.Value
	EnvVar          string
	Flag            *pflag.Flag
	Sensitive       bool
	Type            OptionType
	ViperKey        string
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
		RequestHTTPMethodOption,
		RequestServiceOption,
		RequestAccessTokenOption,
		RequestAccessTokenExpiryOption,
		RequestFailOption,
	}

	// Sort the options list by viper key
	slices.SortFunc(optList, func(opt1, opt2 Option) int {
		return strings.Compare(opt1.ViperKey, opt2.ViperKey)
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

// Generic viper profile options
var (
	ProfileDescriptionOption Option
)

// Root Command Options
var (
	RootActiveProfileOption Option
	RootProfileOption       Option
	RootColorOption         Option
	RootConfigOption        Option
	RootOutputFormatOption  Option
)

// 'pingcli request' command options
var (
	RequestDataOption              Option
	RequestDataRawOption           Option
	RequestHTTPMethodOption        Option
	RequestServiceOption           Option
	RequestAccessTokenOption       Option
	RequestAccessTokenExpiryOption Option
	RequestFailOption              Option
)
