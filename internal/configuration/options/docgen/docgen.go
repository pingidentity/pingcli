package docgen

// Utility to generate markdown documentation for configuration options.
// Extracted from the previous ad-hoc test (Test_outputOptionsMDInfo) so that
// documentation generation can be invoked via a Makefile target instead of a skipped test.

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
)

// Markdown renders the options documentation markdown table sections.
// It ensures all options are initialized by calling configuration.InitAllOptions().
func Markdown() string {
	// Ensure options are initialized (idempotent call)
	configuration.InitAllOptions()

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
		// Replace newlines with '<br><br>' so GitHub markdown renders intentional paragraph breaks.
		usageString = strings.ReplaceAll(usageString, "\n", "<br><br>")

		category := "general"
		if strings.Contains(option.KoanfKey, ".") {
			category = strings.Split(option.KoanfKey, ".")[0]
		}

		propertyCategoryInformation[category] = append(
			propertyCategoryInformation[category],
			fmt.Sprintf("| %s | %d | %s | %s |", option.KoanfKey, option.Type, flagInfo, usageString),
		)
	}

	var outputBuilder strings.Builder
	// Deterministic ordering of categories
	cats := make([]string, 0, len(propertyCategoryInformation))
	for k := range propertyCategoryInformation {
		cats = append(cats, k)
	}
	slices.Sort(cats)

	for _, category := range cats {
		properties := propertyCategoryInformation[category]
		slices.Sort(properties)

		outputBuilder.WriteString(fmt.Sprintf("#### %s Properties\n\n", category))
		outputBuilder.WriteString("| Config File Property | Type | Equivalent Parameter | Purpose |\n")
		outputBuilder.WriteString("|---|---|---|---|\n")
		for _, property := range properties {
			outputBuilder.WriteString(property + "\n")
		}
		outputBuilder.WriteString("\n")
	}

	return outputBuilder.String()
}
