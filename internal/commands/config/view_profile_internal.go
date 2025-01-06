package config_internal

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalConfigViewProfile(args []string) (err error) {
	var pName string
	if len(args) == 1 {
		pName = args[0]
	} else {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
		if err != nil {
			return fmt.Errorf("failed to view profile: %v", err)
		}
	}

	msgStr := fmt.Sprintf("Configuration for profile '%s':\n", pName)

	// Sort the options list by viper key
	optList := options.Options()
	slices.SortFunc(optList, func(opt1, opt2 options.Option) int {
		return strings.Compare(opt1.ViperKey, opt2.ViperKey)
	})

	for _, opt := range optList {
		if opt.ViperKey == "" {
			continue
		}

		if opt.Sensitive {
			optVal, err := profiles.GetSensitiveOptionValue(opt, true)
			if err != nil {
				return fmt.Errorf("failed to view profile: %v", err)
			}

			msgStr += fmt.Sprintf("%s=%s\n", opt.ViperKey, optVal)
		} else {
			optVal, err := profiles.GetOptionValue(opt)
			if err != nil {
				return fmt.Errorf("failed to view profile: %v", err)
			}

			msgStr += fmt.Sprintf("%s=%s\n", opt.ViperKey, optVal)
		}
	}

	output.Message(msgStr, nil)

	return nil
}
