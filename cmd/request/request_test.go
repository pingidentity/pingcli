// Copyright © 2026 Ping Identity Corporation

package request_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/cmd/common"
	auth_internal "github.com/pingidentity/pingcli/internal/commands/auth"
	request_internal "github.com/pingidentity/pingcli/internal/commands/request"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RequestCommand_Validation(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                string
		args                []string
		expectErr           bool
		expectedErrIs       error
		expectedErrContains string
	}{
		{
			name: "Happy Path - with header",
			args: []string{
				"--" + options.RequestServiceOption.CobraParamName, "pingone",
				"--" + options.RequestHTTPMethodOption.CobraParamName, "GET",
				"--" + options.RequestHeaderOption.CobraParamName, "Content-Type: application/json",
				fmt.Sprintf("environments/%s/users", os.Getenv("TEST_PINGONE_ENVIRONMENT_ID")),
			},
			expectErr: false,
		},
		{
			name:      "Happy Path - help",
			args:      []string{"--help"},
			expectErr: false,
		},
		{
			name:          "Too many arguments",
			args:          []string{"arg1", "arg2"},
			expectErr:     true,
			expectedErrIs: common.ErrExactArgs,
		},
		{
			name:                "Invalid flag",
			args:                []string{"--invalid-flag"},
			expectErr:           true,
			expectedErrContains: "unknown flag",
		},
		{
			name: "Invalid service",
			args: []string{
				"--" + options.RequestServiceOption.CobraParamName, "invalid-service",
				"some/path",
			},
			expectErr:     true,
			expectedErrIs: customtypes.ErrUnrecognizedService,
		},
		{
			name: "Invalid HTTP Method",
			args: []string{
				"--" + options.RequestServiceOption.CobraParamName, "pingone",
				"--" + options.RequestHTTPMethodOption.CobraParamName, "INVALID",
				"some/path",
			},
			expectErr:     true,
			expectedErrIs: customtypes.ErrUnrecognizedMethod,
		},
		{
			name:          "Missing required service flag",
			args:          []string{"some/path"},
			expectErr:     true,
			expectedErrIs: request_internal.ErrServiceEmpty,
		},
		{
			name: "Invalid header format",
			args: []string{
				"--" + options.RequestServiceOption.CobraParamName, "pingone",
				"--" + options.RequestHeaderOption.CobraParamName, "invalid=header",
				"some/path",
			},
			expectErr:     true,
			expectedErrIs: nil,
		},
		{
			name: "Disallowed Authorization header",
			args: []string{
				"--" + options.RequestServiceOption.CobraParamName, "pingone",
				"--" + options.RequestHeaderOption.CobraParamName, "Authorization: Bearer token",
				"some/path",
			},
			expectErr:     true,
			expectedErrIs: customtypes.ErrDisallowedAuthHeader,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			if tc.name == "Happy Path - with header" && (os.Getenv("CI") != "" || os.Getenv("GITHUB_ACTIONS") != "") {
				t.Skip("Skipping live request test in CI environment due to keychain warning")
			}

			err := testutils_cobra.ExecutePingcli(t, append([]string{"request"}, tc.args...)...)

			if !tc.expectErr {
				require.NoError(t, err)

				return
			}

			assert.Error(t, err)
			if tc.expectedErrIs != nil {
				assert.ErrorIs(t, err, tc.expectedErrIs)
			}
			if tc.expectedErrContains != "" {
				assert.ErrorContains(t, err, tc.expectedErrContains)
			}
		})
	}
}

// Test_RequestCommand_E2E performs an end-to-end test of the request command,
// making a real API call and validating the JSON output.
func Test_RequestCommand_E2E(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	t.Setenv(options.AuthStorageOption.EnvVar, customtypes.ENUM_STORAGE_TYPE_NONE)

	originalStdout := os.Stdout
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w

	err = testutils_cobra.ExecutePingcli(t, "request",
		"--"+options.RequestServiceOption.CobraParamName, "pingone",
		"--"+options.RequestHTTPMethodOption.CobraParamName, "GET",
		fmt.Sprintf("environments/%s/populations", os.Getenv("TEST_PINGONE_ENVIRONMENT_ID")),
	)
	require.NoError(t, err)

	os.Stdout = originalStdout
	require.NoError(t, w.Close())

	outputBytes, err := io.ReadAll(r)
	require.NoError(t, err)
	require.NoError(t, r.Close())

	// Capture response json body
	re := regexp.MustCompile(`(?s)response:\s+(\{.*\})`)
	matches := re.FindSubmatch(outputBytes)
	require.Len(t, matches, 2, "Failed to capture JSON body from command output")

	bodyJSON := matches[1]
	assert.NotEmpty(t, bodyJSON, "Response JSON body is empty")
	assert.True(t, json.Valid(bodyJSON), "Output JSON is not valid")
}

func Test_RequestCommand_HelpIncludesPingOneEnvironmentFlag(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	output, err := testutils_cobra.ExecutePingcliCaptureCobraOutput(t, "request", "--help")
	require.NoError(t, err)
	assert.Contains(t, output, "--"+options.PingOneAuthenticationAPIEnvironmentIDOption.CobraParamName)
}

func Test_RequestCommand_ClientCredentialsEnvironmentErrorUsesSupportedConfigKey(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	t.Setenv(options.PingOneAuthenticationAPIEnvironmentIDOption.EnvVar, "")
	t.Setenv(options.PingOneAuthenticationClientCredentialsClientIDOption.EnvVar, "00000000-0000-0000-0000-000000000001")
	t.Setenv(options.PingOneAuthenticationClientCredentialsClientSecretOption.EnvVar, "test-secret")

	koanfCfg, err := profiles.GetKoanfConfig()
	require.NoError(t, err)
	require.NoError(t, koanfCfg.KoanfInstance().Set("default."+options.PingOneAuthenticationAPIEnvironmentIDOption.KoanfKey, ""))

	message := auth_internal.ErrClientCredentialsEnvironmentIDNotConfigured.Error()
	assert.Contains(t, message, "service.pingOne.authentication.environmentID")
	assert.NotContains(t, message, "service.pingone.authentication.clientCredentials.environmentID")

	_, err = auth_internal.GetClientCredentialsConfiguration()
	require.Error(t, err)
	assert.ErrorContains(t, err, "service.pingOne.authentication.environmentID")
	assert.False(t, strings.Contains(err.Error(), "service.pingone."))
}
