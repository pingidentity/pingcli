// Copyright © 2025 Ping Identity Corporation

package request_internal

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RunInternalRequest(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	workerEnvId, err := profiles.GetOptionValue(options.PingOneAuthenticationWorkerEnvironmentIDOption)
	require.NoError(t, err)

	defaultService := customtypes.RequestService(customtypes.ENUM_REQUEST_SERVICE_PINGONE)
	defaultHttpMethod := customtypes.HTTPMethod(customtypes.ENUM_HTTP_METHOD_GET)
	defaultRegionCode := customtypes.PingOneRegionCode(customtypes.ENUM_PINGONE_REGION_CODE_NA)

	testCases := []struct {
		name                     string
		uri                      string
		service                  *customtypes.RequestService
		httpMethod               *customtypes.HTTPMethod
		regionCode               *customtypes.PingOneRegionCode
		workerEnvId              *customtypes.String
		workerClientId           *customtypes.String
		runTwiceToSetAccessToken bool
		expectedError            error
	}{
		{
			name: "Happy path - Run internal request",
			uri:  fmt.Sprintf("environments/%s/populations", workerEnvId),
		},
		{
			name:          "Test request with empty service",
			uri:           fmt.Sprintf("environments/%s/populations", workerEnvId),
			service:       utils.Pointer(customtypes.RequestService("")),
			expectedError: ErrServiceEmpty,
		},
		{
			name:          "Test with invalid service",
			uri:           fmt.Sprintf("environments/%s/populations", workerEnvId),
			service:       utils.Pointer(customtypes.RequestService("invalid-service")),
			expectedError: ErrUnrecognizedService,
		},
		{
			name: "Happy Path - Test with invalid URI",
			uri:  "invalid-uri",
		},
		{
			name:          "Test with empty HTTP method",
			uri:           fmt.Sprintf("environments/%s/populations", workerEnvId),
			httpMethod:    utils.Pointer(customtypes.HTTPMethod("")),
			expectedError: ErrHttpMethodEmpty,
		},
		{
			name:          "Test with invalid HTTP method",
			uri:           fmt.Sprintf("environments/%s/populations", workerEnvId),
			httpMethod:    utils.Pointer(customtypes.HTTPMethod("invalid-http-method")),
			expectedError: ErrUnrecognizedHttpMethod,
		},
		{
			name:          "Test with empty pingone region code",
			uri:           fmt.Sprintf("environments/%s/populations", workerEnvId),
			regionCode:    utils.Pointer(customtypes.PingOneRegionCode("")),
			expectedError: ErrPingOneRegionCodeEmpty,
		},
		{
			name:          "Test with invalid pingone region code",
			uri:           fmt.Sprintf("environments/%s/populations", workerEnvId),
			regionCode:    utils.Pointer(customtypes.PingOneRegionCode("invalid-region-code")),
			expectedError: ErrUnrecognizedPingOneRegionCode,
		},
		{
			name:          "Test with empty worker environment ID",
			uri:           fmt.Sprintf("environments/%s/populations", workerEnvId),
			workerEnvId:   utils.Pointer(customtypes.String("")),
			expectedError: ErrPingOneWorkerEnvIDEmpty,
		},
		{
			name:           "Test with empty worker client ID",
			uri:            fmt.Sprintf("environments/%s/populations", workerEnvId),
			workerClientId: utils.Pointer(customtypes.String("")),
			expectedError:  ErrPingOneClientIDAndSecretEmpty,
		},
		{
			name:           "Test with invalid worker client ID",
			uri:            fmt.Sprintf("environments/%s/populations", workerEnvId),
			workerClientId: utils.Pointer(customtypes.String("invalid-client-id")),
			expectedError:  ErrPingOneAuthenticate,
		},
		{
			name:                     "Happy path - Run internal request twice to set access token",
			uri:                      fmt.Sprintf("environments/%s/populations", workerEnvId),
			runTwiceToSetAccessToken: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			options.RequestServiceOption.Flag.Changed = true
			if tc.service != nil {
				options.RequestServiceOption.CobraParamValue = tc.service
			} else {
				options.RequestServiceOption.CobraParamValue = &defaultService
			}

			options.RequestHTTPMethodOption.Flag.Changed = true
			if tc.httpMethod != nil {
				options.RequestHTTPMethodOption.CobraParamValue = tc.httpMethod
			} else {
				options.RequestHTTPMethodOption.CobraParamValue = &defaultHttpMethod
			}

			options.PingOneRegionCodeOption.Flag.Changed = true
			if tc.regionCode != nil {
				options.PingOneRegionCodeOption.CobraParamValue = tc.regionCode
			} else {
				options.PingOneRegionCodeOption.CobraParamValue = &defaultRegionCode
			}

			if tc.workerEnvId != nil {
				options.PingOneAuthenticationWorkerEnvironmentIDOption.Flag.Changed = true
				options.PingOneAuthenticationWorkerEnvironmentIDOption.CobraParamValue = tc.workerEnvId
			}

			if tc.workerClientId != nil {
				options.PingOneAuthenticationWorkerClientIDOption.Flag.Changed = true
				options.PingOneAuthenticationWorkerClientIDOption.CobraParamValue = tc.workerClientId
			}

			err := RunInternalRequest(tc.uri)

			if tc.expectedError != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}

			if tc.runTwiceToSetAccessToken {
				err = RunInternalRequest(tc.uri)

				if tc.expectedError != nil {
					require.Error(t, err)
					assert.ErrorIs(t, err, tc.expectedError)
				} else {
					assert.NoError(t, err)
				}
			}
		})
	}
}

// Test RunInternalRequest function with fail
func Test_RunInternalRequestWithFail(t *testing.T) {
	if os.Getenv("RUN_INTERNAL_FAIL_TEST") == "true" {
		testutils_koanf.InitKoanfs(t)

		service := customtypes.RequestService(customtypes.ENUM_REQUEST_SERVICE_PINGONE)
		fail := customtypes.String("true")

		options.RequestServiceOption.Flag.Changed = true
		options.RequestServiceOption.CobraParamValue = &service

		options.RequestFailOption.Flag.Changed = true
		options.RequestFailOption.CobraParamValue = &fail

		_ = RunInternalRequest("environments/failTest")
		t.Fatal("This should never run due to internal request resulting in os.Exit(1)")
	} else {
		cmdName := os.Args[0]
		cmd := exec.CommandContext(t.Context(), cmdName, "-test.run=Test_RunInternalRequestWithFail") //#nosec G204 -- This is a test
		cmd.Env = append(os.Environ(), "RUN_INTERNAL_FAIL_TEST=true")
		output, err := cmd.CombinedOutput()

		require.Contains(t, string(output), "ERROR: Failed Custom Request")
		require.NotContains(t, string(output), "This should never run due to internal request resulting in os.Exit(1)")

		var exitErr *exec.ExitError
		require.ErrorAs(t, err, &exitErr)
		require.False(t, exitErr.Success(), "Process should exit with a non-zero")
	}
}

func Test_getData(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	dataFileContents := `{data: 'json from file'}`
	dataRawContents := `{data: 'json from raw'}`

	dataFile := createDataJSONFile(t, dataFileContents)

	testCases := []struct {
		name          string
		rawData       *customtypes.String
		dataFile      *customtypes.String
		expectedError error
	}{
		{
			name:    "Happy path - get data from rawData",
			rawData: utils.Pointer(customtypes.String(dataRawContents)),
		},
		{
			name:     "Happy path - get data from dataFile",
			dataFile: utils.Pointer(customtypes.String(dataFile)),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			require.True(t, (tc.rawData != nil) != (tc.dataFile != nil), "Either rawData or dataFile must be set, but not both")

			var (
				dataStr string
				err     error
			)

			if tc.rawData != nil {
				options.RequestDataRawOption.Flag.Changed = true
				options.RequestDataRawOption.CobraParamValue = tc.rawData

				dataStr, err = getDataRaw()

				require.Equal(t, dataStr, dataRawContents)
			}

			if tc.dataFile != nil {
				options.RequestDataOption.Flag.Changed = true
				options.RequestDataOption.CobraParamValue = tc.dataFile

				dataStr, err = getDataFile()

				require.Equal(t, dataStr, dataFileContents)
			}

			if tc.expectedError != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func createDataJSONFile(t *testing.T, data string) string {
	t.Helper()

	file, err := os.CreateTemp(t.TempDir(), "data-*.json")
	require.NoError(t, err)

	_, err = file.WriteString(data)
	require.NoError(t, err)

	err = file.Close()
	require.NoError(t, err)

	return file.Name()
}
