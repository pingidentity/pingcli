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
	BOOL OptionType = iota
	AUTH_SERVICES
	EXPORT_FORMAT
	EXPORT_SERVICE_GROUP
	EXPORT_SERVICES
	HEADER
	INT
	LICENSE_PRODUCT
	LICENSE_VERSION
	OUTPUT_FORMAT
	PINGFEDERATE_AUTH_TYPE
	PINGONE_AUTH_TYPE
	PINGONE_REGION_CODE
	REQUEST_HTTP_METHOD
	REQUEST_SERVICE
	STRING
	STRING_SLICE
	UUID
)

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
		ConfigAddProfileDescriptionOption,
		ConfigAddProfileNameOption,
		ConfigAddProfileSetActiveOption,
		ConfigDeleteAutoAcceptOption,
		ConfigListKeysYamlOption,
		ConfigUnmaskSecretValueOption,

		AuthMethodAuthCodeOption,
		AuthMethodClientCredentialsOption,
		AuthMethodDeviceCodeOption,
		AuthServiceOption,
		AuthFileStorageOption,

		LicenseProductOption,
		LicenseVersionOption,
		LicenseDevopsUserOption,
		LicenseDevopsKeyOption,

		PingFederateAccessTokenAuthAccessTokenOption,
		PingFederateAdminAPIPathOption,
		PingFederateAuthenticationTypeOption,
		PingFederateBasicAuthPasswordOption,
		PingFederateBasicAuthUsernameOption,
		PingFederateCACertificatePemFilesOption,
		PingFederateClientCredentialsAuthClientIDOption,
		PingFederateClientCredentialsAuthClientSecretOption,
		PingFederateClientCredentialsAuthTokenURLOption,
		PingFederateClientCredentialsAuthScopesOption,
		PingFederateHTTPSHostOption,
		PingFederateInsecureTrustAllTLSOption,
		PingFederateXBypassExternalValidationHeaderOption,

		PingOneAuthenticationAuthCodeClientIDOption,
		PingOneAuthenticationAuthCodeEnvironmentIDOption,
		PingOneAuthenticationAuthCodePortOption,
		PingOneAuthenticationAuthCodeRedirectURIPathOption,
		PingOneAuthenticationAuthCodeRedirectURIPortOption,
		PingOneAuthenticationAuthCodeScopesOption,
		PingOneAuthenticationAPIEnvironmentIDOption,
		PingOneAuthenticationClientCredentialsClientIDOption,
		PingOneAuthenticationClientCredentialsClientSecretOption,
		PingOneAuthenticationClientCredentialsEnvironmentIDOption,
		PingOneAuthenticationClientCredentialsScopesOption,
		PingOneAuthenticationDeviceCodeClientIDOption,
		PingOneAuthenticationDeviceCodeEnvironmentIDOption,
		PingOneAuthenticationDeviceCodeScopesOption,
		PingOneAuthenticationTypeOption,
		PingOneAuthenticationWorkerClientIDOption,
		PingOneAuthenticationWorkerClientSecretOption,
		PingOneAuthenticationWorkerEnvironmentIDOption,
		PingOneRegionCodeOption,

		PlatformExportExportFormatOption,
		PlatformExportOutputDirectoryOption,
		PlatformExportOverwriteOption,
		PlatformExportPingOneEnvironmentIDOption,
		PlatformExportServiceGroupOption,
		PlatformExportServiceOption,

		PluginExecutablesOption,

		ProfileDescriptionOption,

		RequestAccessTokenExpiryOption,
		RequestAccessTokenOption,
		RequestDataOption,
		RequestDataRawOption,
		RequestFailOption,
		RequestHeaderOption,
		RequestHTTPMethodOption,
		RequestServiceOption,

		RootActiveProfileOption,
		RootColorOption,
		RootConfigOption,
		RootDetailedExitCodeOption,
		RootOutputFormatOption,
		RootProfileOption,
	}

	// Sort the options list by koanf key
	slices.SortFunc(optList, func(opt1, opt2 Option) int {
		return strings.Compare(opt1.KoanfKey, opt2.KoanfKey)
	})

	return optList
}

// 'pingcli config' command options
var (
	ConfigAddProfileDescriptionOption Option
	ConfigAddProfileNameOption        Option
	ConfigAddProfileSetActiveOption   Option
	ConfigDeleteAutoAcceptOption      Option
	ConfigListKeysYamlOption          Option
	ConfigUnmaskSecretValueOption     Option
)

// 'pingcli login' command options
var (
	AuthMethodAuthCodeOption          Option
	AuthMethodClientCredentialsOption Option
	AuthMethodDeviceCodeOption        Option
	AuthServiceOption                 Option
	AuthFileStorageOption             Option
)

// License options
var (
	LicenseProductOption    Option
	LicenseVersionOption    Option
	LicenseDevopsUserOption Option
	LicenseDevopsKeyOption  Option
)

// pingfederate service options
var (
	PingFederateAccessTokenAuthAccessTokenOption        Option
	PingFederateAdminAPIPathOption                      Option
	PingFederateAuthenticationTypeOption                Option
	PingFederateBasicAuthPasswordOption                 Option
	PingFederateBasicAuthUsernameOption                 Option
	PingFederateCACertificatePemFilesOption             Option
	PingFederateClientCredentialsAuthClientIDOption     Option
	PingFederateClientCredentialsAuthClientSecretOption Option
	PingFederateClientCredentialsAuthScopesOption       Option
	PingFederateClientCredentialsAuthTokenURLOption     Option
	PingFederateHTTPSHostOption                         Option
	PingFederateInsecureTrustAllTLSOption               Option
	PingFederateXBypassExternalValidationHeaderOption   Option
)

// pingone service options
var (
	PingOneAuthenticationAPIEnvironmentIDOption               Option
	PingOneAuthenticationAuthCodeClientIDOption               Option
	PingOneAuthenticationAuthCodeEnvironmentIDOption          Option
	PingOneAuthenticationAuthCodePortOption                   Option
	PingOneAuthenticationAuthCodeRedirectURIPathOption        Option
	PingOneAuthenticationAuthCodeRedirectURIPortOption        Option
	PingOneAuthenticationAuthCodeScopesOption                 Option
	PingOneAuthenticationClientCredentialsClientIDOption      Option
	PingOneAuthenticationClientCredentialsClientSecretOption  Option
	PingOneAuthenticationClientCredentialsEnvironmentIDOption Option
	PingOneAuthenticationClientCredentialsScopesOption        Option
	PingOneAuthenticationDeviceCodeClientIDOption             Option
	PingOneAuthenticationDeviceCodeEnvironmentIDOption        Option
	PingOneAuthenticationDeviceCodeScopesOption               Option
	PingOneAuthenticationTypeOption                           Option
	PingOneAuthenticationWorkerClientIDOption                 Option
	PingOneAuthenticationWorkerClientSecretOption             Option
	PingOneAuthenticationWorkerEnvironmentIDOption            Option
	PingOneRegionCodeOption                                   Option
)

// 'pingcli platform export' command options
var (
	PlatformExportExportFormatOption         Option
	PlatformExportOutputDirectoryOption      Option
	PlatformExportOverwriteOption            Option
	PlatformExportPingOneEnvironmentIDOption Option
	PlatformExportServiceGroupOption         Option
	PlatformExportServiceOption              Option
)

// 'pingcli plugin' command options
var (
	PluginExecutablesOption Option
)

// Generic koanf profile options
var (
	ProfileDescriptionOption Option
)

// Root Command Options
var (
	RootActiveProfileOption    Option
	RootColorOption            Option
	RootConfigOption           Option
	RootDetailedExitCodeOption Option
	RootOutputFormatOption     Option
	RootProfileOption          Option
)

// 'pingcli request' command options
var (
	RequestAccessTokenExpiryOption Option
	RequestAccessTokenOption       Option
	RequestDataOption              Option
	RequestDataRawOption           Option
	RequestFailOption              Option
	RequestHeaderOption            Option
	RequestHTTPMethodOption        Option
	RequestServiceOption           Option
)
