package configuration_platform

import (
	"fmt"
	"os"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/pflag"
)

func InitPlatformExportOptions() {
	initFormatOption()
	initServicesOption()
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
			Usage:     fmt.Sprintf("Specifies export format\nAllowed: [%s]. Also configurable via environment variable %s", strings.Join(customtypes.ExportFormatValidValues(), ", "), envVar),
			Value:     cobraValue,
			DefValue:  customtypes.ENUM_EXPORT_FORMAT_HCL,
		},
		Type:     options.ENUM_STRING,
		ViperKey: "export.format",
	}
}

func initServicesOption() {
	cobraParamName := "services"
	cobraValue := new(customtypes.ExportServices)
	defaultValue := customtypes.ExportServices(customtypes.ExportServicesValidValues())
	envVar := "PINGCLI_EXPORT_SERVICES"

	options.PlatformExportServiceOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "s",
			Usage:     fmt.Sprintf("Specifies service(s) to export. Accepts comma-separated string to delimit multiple services. Allowed: [%s]. Also configurable via environment variable %s", strings.Join(customtypes.ExportServicesValidValues(), ", "), envVar),
			Value:     cobraValue,
			DefValue:  strings.Join(customtypes.ExportServicesValidValues(), ", "),
		},
		Type:     options.ENUM_EXPORT_SERVICES,
		ViperKey: "export.services",
	}
}

func initOutputDirectoryOption() {
	cobraParamName := "output-directory"
	cobraValue := new(customtypes.String)
	defaultValue := getDefaultExportDir()
	envVar := "PINGCLI_EXPORT_OUTPUT_DIRECTORY"

	options.PlatformExportOutputDirectoryOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "d",
			Usage:     fmt.Sprintf("Specifies output directory for export. Also configurable via environment variable %s", envVar),
			Value:     cobraValue,
			DefValue:  "$(pwd)/export",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "export.outputDirectory",
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
			Usage:     "Overwrite existing generated exports in output directory.",
			Value:     cobraValue,
			DefValue:  "false",
		},
		Type:     options.ENUM_BOOL,
		ViperKey: "export.overwrite",
	}
}

func getDefaultExportDir() (defaultExportDir *customtypes.String) {
	l := logger.Get()
	pwd, err := os.Getwd()
	if err != nil {
		l.Err(err).Msg("Failed to determine current working directory")
		return nil
	}

	defaultExportDir = new(customtypes.String)

	err = defaultExportDir.Set(fmt.Sprintf("%s/export", pwd))
	if err != nil {
		l.Err(err).Msg("Failed to set default export directory")
		return nil
	}

	return defaultExportDir
}

func initPingOneEnvironmentIDOption() {
	cobraParamName := "pingone-export-environment-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCLI_PINGONE_EXPORT_ENVIRONMENT_ID"

	options.PlatformExportPingoneEnvironmentIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The ID of the Pingone environment to export. Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_UUID,
		ViperKey: "export.pingone.environmentID",
	}
}
