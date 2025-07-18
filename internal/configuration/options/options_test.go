// Copyright © 2025 Ping Identity Corporation

package options_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

func Test_outputOptionsMDInfo(t *testing.T) {
	// Skip this test. Use only to generate markdown table for documentation
	t.SkipNow()

	testutils_koanf.InitKoanfs(t)

	propertyCategoryInformation := make(map[string][]string)

	for _, option := range options.Options() {
		if option.KoanfKey == "" || option.Flag == nil {
			continue
		}

		var flagInfo string
		if option.Flag.Shorthand != "" {
			flagInfo = fmt.Sprintf("--%s / -%s", option.CobraParamName, option.Flag.Shorthand)
		} else {
			flagInfo = fmt.Sprintf("--%s", option.CobraParamName)
		}

		usageString := option.Flag.Usage
		// Replace newlines with '<br><br>'
		usageString = strings.ReplaceAll(usageString, "\n", "<br><br>")

		if !strings.Contains(option.KoanfKey, ".") {
			propertyCategoryInformation["general"] = append(propertyCategoryInformation["general"], fmt.Sprintf("| %s | %d | %s | %s |", option.KoanfKey, option.Type, flagInfo, usageString))
		} else {
			rootKey := strings.Split(option.KoanfKey, ".")[0]
			propertyCategoryInformation[rootKey] = append(propertyCategoryInformation[rootKey], fmt.Sprintf("| %s | %d | %s | %s |", option.KoanfKey, option.Type, flagInfo, usageString))
		}
	}

	var outputString string
	for category, properties := range propertyCategoryInformation {
		outputString += fmt.Sprintf("#### %s Properties\n\n", category)

		outputString += "| Config File Property | Type | Equivalent Parameter | Purpose |\n"
		outputString += "|---|---|---|---|\n"

		slices.Sort(properties)

		for _, property := range properties {
			outputString += property + "\n"
		}

		outputString += "\n"
	}

	fmt.Println(outputString)
}
