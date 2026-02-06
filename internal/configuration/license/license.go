// Copyright © 2026 Ping Identity Corporation

package configuration_license

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitLicenseOptions() {
	initProductOption()
	initVersionOption()
	initDevopsUserOption()
	initDevopsKeyOption()
}

func initProductOption() {
	cobraParamName := "product"
	cobraValue := new(customtypes.LicenseProduct)
	defaultValue := customtypes.LicenseProduct("")

	options.LicenseProductOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "p",
			Usage: fmt.Sprintf(
				"The product for which to request a license. "+
					"\nOptions are: %s."+
					"\nExample: '%s'",
				strings.Join(customtypes.LicenseProductValidValues(), ", "),
				customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE,
			),
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.LICENSE_PRODUCT,
		KoanfKey:  "", // No koanf key
	}
}

func initVersionOption() {
	cobraParamName := "version"
	cobraValue := new(customtypes.LicenseVersion)
	defaultValue := customtypes.LicenseVersion("")

	options.LicenseVersionOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "v",
			Usage: "The version of the product for which to request a license. Must be of the form 'major.minor'. " +
				"\nExample: '12.3'",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.LICENSE_VERSION,
		KoanfKey:  "", // No koanf key
	}
}

func initDevopsUserOption() {
	cobraParamName := "devops-user"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	options.LicenseDevopsUserOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "PINGCLI_LICENSE_DEVOPS_USER",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "u",
			Usage: "The DevOps user for the license request. " +
				"\n See https://developer.pingidentity.com/devops/how-to/devopsRegistration.html on how to register a DevOps user. " +
				"\n You can save the DevOps user and key in your profile using the 'pingcli config' commands.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.STRING,
		KoanfKey:  "license.devopsUser",
	}
}

func initDevopsKeyOption() {
	cobraParamName := "devops-key"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	options.LicenseDevopsKeyOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "PINGCLI_LICENSE_DEVOPS_KEY",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "k",
			Usage: "The DevOps key for the license request. " +
				"\n See https://developer.pingidentity.com/devops/how-to/devopsRegistration.html on how to register a DevOps user. " +
				"\n You can save the DevOps user and key in your profile using the 'pingcli config' commands.",
			Value: cobraValue,
		},
		Sensitive: true,
		Type:      options.STRING,
		KoanfKey:  "license.devopsKey",
	}
}
