// Copyright © 2025 Ping Identity Corporation

package plugins

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	hplugin "github.com/hashicorp/go-plugin"
	"github.com/pingidentity/pingcli/shared/grpc"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

var (
	testPluginConfig = &grpc.PingCliCommandConfiguration{
		Use:     "test-plugin",
		Short:   "A test plugin",
		Long:    "A longer description for a test plugin",
		Example: "pingcli test-plugin --flag value",
	}
	testRunError = errors.New("plugin run error")
)

// mockPingCliCommand is a mock implementation of the grpc.PingCliCommand interface for testing.
type mockPingCliCommand struct{}

var mockPlugin = &mockPingCliCommand{}

func (m *mockPingCliCommand) Configuration() (*grpc.PingCliCommandConfiguration, error) {
	if configErr := os.Getenv("PINGCLI_TEST_PLUGIN_CONFIG_ERROR"); configErr != "" {
		return nil, errors.New(configErr)
	}
	return testPluginConfig, nil
}

func (m *mockPingCliCommand) Run(args []string, l grpc.Logger) error {
	if runErr := os.Getenv("PINGCLI_TEST_PLUGIN_RUN_ERROR"); runErr != "" {
		return errors.New(runErr)
	}
	return fmt.Errorf("args: %s", strings.Join(args, ","))
}

func TestMain(m *testing.M) {
	if os.Getenv("PINGCLI_TEST_PLUGIN") == "1" {
		hplugin.Serve(&hplugin.ServeConfig{
			HandshakeConfig: grpc.HandshakeConfig,
			Plugins: map[string]hplugin.Plugin{
				grpc.ENUM_PINGCLI_COMMAND_GRPC: &grpc.PingCliCommandGrpcPlugin{Impl: mockPlugin},
			},
			GRPCServer: hplugin.DefaultGRPCServer,
		})
		return
	}
	os.Exit(m.Run())
}

func setupPluginTest(t *testing.T) string {
	t.Setenv("PINGCLI_TEST_PLUGIN", "1")
	return os.Args[0]
}

func Test_pluginConfiguration(t *testing.T) {
	pluginExec := setupPluginTest(t)
	ctx := context.Background()

	t.Run("Happy path", func(t *testing.T) {
		conf, err := pluginConfiguration(ctx, pluginExec)
		require.NoError(t, err)
		require.NotNil(t, conf)
		require.Equal(t, testPluginConfig.Use, conf.Use)
	})

	t.Run("Plugin returns error", func(t *testing.T) {
		t.Setenv("PINGCLI_TEST_PLUGIN_CONFIG_ERROR", "config error")
		_, err := pluginConfiguration(ctx, pluginExec)
		require.Error(t, err)
		require.ErrorContains(t, err, "config error")
	})

	t.Run("Invalid executable", func(t *testing.T) {
		_, err := pluginConfiguration(ctx, "invalid-executable-path")
		require.Error(t, err)
	})
}

func Test_createCmdRunE(t *testing.T) {
	pluginExec := setupPluginTest(t)
	rootCmd := &cobra.Command{Use: "pingcli"}
	rootCmd.PersistentFlags().String("profile", "", "test profile flag")

	pluginCmd := &cobra.Command{
		Use:                "plugin",
		RunE:               createCmdRunE(pluginExec),
		DisableFlagParsing: true, // Match the real implementation
	}
	rootCmd.AddCommand(pluginCmd)

	testCases := []struct {
		name         string
		runError     string
		args         []string
		expectedArgs []string
		expectError  bool
	}{
		{
			name:         "Happy path",
			args:         []string{"plugin", "--profile", "my-profile", "plugin-arg", "--plugin-flag"},
			expectedArgs: []string{"plugin-arg", "--plugin-flag"},
		},
		{
			name:         "Plugin returns error",
			runError:     "plugin run error",
			args:         []string{"plugin", "arg1"},
			expectedArgs: []string{"arg1"},
			expectError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Unset env var from previous runs
			t.Setenv("PINGCLI_TEST_PLUGIN_RUN_ERROR", "")
			if tc.runError != "" {
				t.Setenv("PINGCLI_TEST_PLUGIN_RUN_ERROR", tc.runError)
			}

			ctx := context.Background()
			rootCmd.SetArgs(tc.args)
			err := rootCmd.ExecuteContext(ctx)

			if tc.expectError {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.runError)
			} else {
				require.Error(t, err)
				expectedErrStr := fmt.Sprintf("args: %s", strings.Join(tc.expectedArgs, ","))
				require.Contains(t, err.Error(), expectedErrStr)
			}
		})
	}
}

func Test_filterRootFlags(t *testing.T) {
	rootCmd := &cobra.Command{Use: "root"}
	rootCmd.PersistentFlags().StringP("profile", "p", "", "profile flag")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose flag")
	subCmd := &cobra.Command{Use: "sub"}
	rootCmd.AddCommand(subCmd)

	testCases := []struct {
		name         string
		args         []string
		expectedArgs []string
	}{
		{"No root flags", []string{"plugin-arg", "--plugin-flag"}, []string{"plugin-arg", "--plugin-flag"}},
		{"Long name root flag", []string{"--profile", "my-profile", "plugin-arg"}, []string{"plugin-arg"}},
		{"Short name root flag", []string{"-p", "my-profile", "plugin-arg"}, []string{"plugin-arg"}},
		{"Root flag with equals", []string{"--profile=my-profile", "plugin-arg"}, []string{"plugin-arg"}},
		{"Boolean root flag", []string{"--verbose", "plugin-arg"}, []string{"plugin-arg"}},
		{"Short boolean root flag", []string{"-v", "plugin-arg"}, []string{"plugin-arg"}},
		{"Mixed flags", []string{"-v", "plugin-arg1", "--profile", "prof", "--plugin-flag", "val"}, []string{"plugin-arg1", "--plugin-flag", "val"}},
		{"No args", []string{}, []string{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			filteredArgs := filterRootFlags(subCmd, tc.args)
			require.Equal(t, tc.expectedArgs, filteredArgs)
		})
	}
}
