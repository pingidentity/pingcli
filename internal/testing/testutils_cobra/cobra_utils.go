// Copyright © 2025 Ping Identity Corporation

package testutils_cobra

import (
	"bytes"
	"testing"

	"github.com/pingidentity/pingcli/cmd"
	"github.com/pingidentity/pingcli/internal/configuration"
	testutils_viper "github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// ExecutePingcli executes the pingcli command with the provided arguments
// and returns the error if any
func ExecutePingcli(t *testing.T, args ...string) (err error) {
	t.Helper()

	// Reset options for testing individual executions of pingcli
	configuration.InitAllOptions()

	root := cmd.NewRootCommand("test-version", "test-commit")

	// Add config location to the root command
	configFilepath := testutils_viper.CreateConfigFile(t)
	args = append([]string{"--config", configFilepath}, args...)
	root.SetArgs(args)

	return root.Execute()
}

// ExecutePingcliCaptureCobraOutput executes the pingcli command with
// the provided arguments and returns the output and error if any
// Note: The returned output will only contain cobra module specific output
// such as usage, help, and cobra errors
// It will NOT include internal/output/output.go output
// nor with it contain zerolog logs
func ExecutePingcliCaptureCobraOutput(t *testing.T, args ...string) (output string, err error) {
	t.Helper()

	// Reset options for testing individual executions of pingcli
	configuration.InitAllOptions()

	root := cmd.NewRootCommand("test-version", "test-commit")

	// Add config location to the root command
	configFilepath := testutils_viper.CreateConfigFile(t)
	args = append([]string{"--config", configFilepath}, args...)
	root.SetArgs(args)

	// Create byte buffer to capture output
	var stdout bytes.Buffer
	root.SetOut(&stdout)
	root.SetErr(&stdout)

	return stdout.String(), root.Execute()
}
