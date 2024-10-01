package configuration_profiles

import (
	"github.com/pingidentity/pingctl/internal/configuration/options"
	"github.com/pingidentity/pingctl/internal/customtypes"
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
		Type:            options.ENUM_STRING,
		ViperKey:        "description",
	}
}
