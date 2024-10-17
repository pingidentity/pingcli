package options

import "github.com/spf13/pflag"

type OptionType string

// OptionType enums
const (
	ENUM_BOOL                   OptionType = "ENUM_BOOL"
	ENUM_EXPORT_FORMAT          OptionType = "ENUM_EXPORT_FORMAT"
	ENUM_INT                    OptionType = "ENUM_INT"
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
	Type            OptionType
	ViperKey        string
}

func Options() []Option {
	return []Option{
		PingOneAuthenticationTypeOption,
		PingOneAuthenticationWorkerClientIDOption,
		PingOneAuthenticationWorkerClientSecretOption,
		PingOneAuthenticationWorkerEnvironmentIDOption,
		PingOneRegionCodeOption,

		PlatformExportExportFormatOption,
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

		ConfigProfileOption,
		ConfigNameOption,
		ConfigDescriptionOption,
		ConfigAddProfileDescriptionOption,
		ConfigAddProfileNameOption,
		ConfigAddProfileSetActiveOption,
		ConfigDeleteAutoAcceptOption,
		ConfigGetProfileOption,
		ConfigSetProfileOption,
		ConfigUnsetProfileOption,

		RequestDataOption,
		RequestHTTPMethodOption,
		RequestServiceOption,
		RequestAccessTokenOption,
		RequestAccessTokenExpiryOption,
	}
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
	ConfigProfileOption     Option
	ConfigNameOption        Option
	ConfigDescriptionOption Option

	ConfigAddProfileDescriptionOption Option
	ConfigAddProfileNameOption        Option
	ConfigAddProfileSetActiveOption   Option

	ConfigDeleteAutoAcceptOption Option

	ConfigGetProfileOption Option

	ConfigSetProfileOption Option

	ConfigUnsetProfileOption Option
)

// 'pingcli platform export' command options
var (
	PlatformExportExportFormatOption         Option
	PlatformExportServiceOption              Option
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
	RequestHTTPMethodOption        Option
	RequestServiceOption           Option
	RequestAccessTokenOption       Option
	RequestAccessTokenExpiryOption Option
)
