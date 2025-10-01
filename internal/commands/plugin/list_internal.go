// Copyright © 2025 Ping Identity Corporation

package plugin_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var (
	listErrorPrefix = "failed to list plugins"
)

func RunInternalPluginList() error {
	existingPluginExectuables, ok, err := profiles.KoanfValueFromOption(options.PluginExecutablesOption, "")
	if err != nil {
		return &errs.PingCLIError{Prefix: listErrorPrefix, Err: fmt.Errorf("%w: %w", ErrReadPluginNamesConfig, err)}
	}
	if !ok {
		output.Message("No plugins configured.", nil)

		return nil
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
