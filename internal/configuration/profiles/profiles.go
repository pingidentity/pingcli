package configuration_profiles

import (
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
)

func InitProfilesOptions() {
	initDescriptionOption()
}

func initDescriptionOption() {
	options.ProfileDescriptionOption = options.Option{
		CobraParamName:  "",  // No cobra param name
		CobraParamValue: nil, // No cobra param value
		DefaultValue:    new(customtypes.String),
		EnvVar:          "",  // No environment variable
		Flag:            nil, // No flag
		Sensitive:       false,
		Type:            options.ENUM_STRING,
		ViperKey:        "description",
	}
}
