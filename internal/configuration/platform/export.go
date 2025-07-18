// Copyright © 2025 Ping Identity Corporation

package configuration_platform

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitPlatformExportOptions() {
	initFormatOption()
	initServicesOption()
	initServiceGroupOption()
	initOutputDirectoryOption()
	initOverwriteOption()
	initPingOneEnvironmentIDOption()
}

func initFormatOption() {
	cobraParamName := "format"
	cobraValue := new(customtypes.ExportFormat)
	defaultValue := customtypes.ExportFormat(customtypes.ENUM_EXPORT_FORMAT_HCL)
	envVar := "PINGCLI_EXPORT_FORMAT"

	options.PlatformExportExportFormatOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "f",
			Usage: fmt.Sprintf(
				"Specifies the export format. (default %s)"+
					"\nOptions are: %s.",
				customtypes.ENUM_EXPORT_FORMAT_HCL,
				strings.Join(customtypes.ExportFormatValidValues(), ", "),
			),
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.EXPORT_FORMAT,
		KoanfKey:  "export.format",
	}
}

func initServiceGroupOption() {
	cobraParamName := "service-group"
	cobraValue := new(customtypes.ExportServiceGroup)
	defaultValue := customtypes.ExportServiceGroup("")
	envVar := "PINGCLI_EXPORT_SERVICE_GROUP"
	options.PlatformExportServiceGroupOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "g",
			Usage: fmt.Sprintf(
				"Specifies the service group to export. "+
					"\nOptions are: %s."+
					"\nExample: '%s'",
				strings.Join(customtypes.ExportServiceGroupValidValues(), ", "),
				customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE,
			),
			Value: cobraValue,
		},
		Sensitive: false,
		KoanfKey:  "export.serviceGroup",
		Type:      options.EXPORT_SERVICE_GROUP,
	}
}

func initServicesOption() {
	cobraParamName := "services"
	cobraValue := new(customtypes.ExportServices)
	defaultValue := customtypes.ExportServices([]string{})
	envVar := "PINGCLI_EXPORT_SERVICES"
	options.PlatformExportServiceOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "s",
			Usage: fmt.Sprintf(
				"Specifies the service(s) to export. Accepts a comma-separated string to delimit multiple services. "+
					"\nOptions are: %s."+
					"\nExample: '%s,%s,%s'",
				strings.Join(customtypes.ExportServicesValidValues(), ", "),
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO,
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA,
				customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
			),
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.EXPORT_SERVICES,
		KoanfKey:  "export.services",
	}
}

func initOutputDirectoryOption() {
	cobraParamName := "output-directory"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_EXPORT_OUTPUT_DIRECTORY"

	options.PlatformExportOutputDirectoryOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "d",
			Usage: "Specifies the output directory for export. Can be an absolute filepath or a relative filepath of" +
				" the present working directory. " +
				"\nExample: '/Users/example/pingcli-export'" +
				"\nExample: 'pingcli-export'",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.STRING,
		KoanfKey:  "export.outputDirectory",
	}
}

func initOverwriteOption() {
	cobraParamName := "overwrite"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)

	options.PlatformExportOverwriteOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "PINGCLI_EXPORT_OVERWRITE",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "o",
			Usage: "Overwrites the existing generated exports in output directory. " +
				"(default false)",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make this flag a boolean flag
		},
		Sensitive: false,
		KoanfKey:  "export.overwrite",
		Type:      options.BOOL,
	}
}

func initPingOneEnvironmentIDOption() {
	cobraParamName := "pingone-export-environment-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCLI_PINGONE_EXPORT_ENVIRONMENT_ID"

	options.PlatformExportPingOneEnvironmentIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The ID of the PingOne environment to export. Must be a valid PingOne UUID.",
			Value: cobraValue,
		},
		KoanfKey: "export.pingOne.environmentID",
		Type:     options.UUID,
	}
}
