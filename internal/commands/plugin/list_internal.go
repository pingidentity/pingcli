package plugin_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalPluginList() error {
	existingPluginExectuables, _, err := profiles.KoanfValueFromOption(options.PluginExecutablesOption, "")
	if err != nil {
		return fmt.Errorf("failed to get existing plugin configuration: %w", err)
	}

	listStr := "Plugins:\n"
	for pluginExecutable := range strings.SplitSeq(existingPluginExectuables, ",") {
		if pluginExecutable == "" {
			continue
		}

		listStr += "- " + pluginExecutable + "\n"
	}

	output.Message(listStr, nil)

	return nil
}
