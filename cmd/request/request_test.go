// Copyright © 2025 Ping Identity Corporation

package request_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Request Command Executes without issue
func TestRequestCmd_Execute(t *testing.T) {
	originalStdout := os.Stdout
	pipeReader, pipeWriter, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	defer func() {
		err := pipeReader.Close()
		if err != nil {
			t.Fatalf("Failed to close pipe: %v", err)
		}
	}()
	os.Stdout = pipeWriter

	err = testutils_cobra.ExecutePingcli(t, "request",
		"--"+options.RequestServiceOption.CobraParamName, "pingone",
		"--"+options.RequestHTTPMethodOption.CobraParamName, "GET",
		fmt.Sprintf("environments/%s/populations", os.Getenv("TEST_PINGONE_ENVIRONMENT_ID")),
	)
	testutils.CheckExpectedError(t, err, nil)

	os.Stdout = originalStdout
	err = pipeWriter.Close()
	if err != nil {
		t.Fatalf("Failed to close pipe: %v", err)
	}

	pipeReaderOut, err := io.ReadAll(pipeReader)
	if err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}

	// Capture response json body
	captureGroupName := "BodyJSON"
	re := regexp.MustCompile(fmt.Sprintf(`(?s)^.*response:\s+(?P<%s>\{.*\}).*$`, captureGroupName))
	matchData := re.FindSubmatch(pipeReaderOut)

	for index, name := range re.SubexpNames() {
		if name == captureGroupName {
			if len(matchData) <= index {
				t.Fatalf("Failed to capture JSON body: %v", matchData)
			}
			bodyJSON := matchData[index]

			// Check for valid JSON
			if !json.Valid(bodyJSON) {
				t.Errorf("Invalid JSON: %s", bodyJSON)
			}
		}
	}
}

// Test Request Command fails when provided too many arguments
func TestRequestCmd_Execute_TooManyArguments(t *testing.T) {
	expectedErrorPattern := `accepts 1 arg\(s\), received 2`
	err := testutils_cobra.ExecutePingcli(t, "request", "arg1", "arg2")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Request Command fails when provided invalid flag
func TestRequestCmd_Execute_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `unknown flag: --invalid`
	err := testutils_cobra.ExecutePingcli(t, "request", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Request Command --help, -h flag
func TestRequestCmd_Execute_Help(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "request", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "request", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Request Command with Invalid Service
func TestRequestCmd_Execute_InvalidService(t *testing.T) {
	expectedErrorPattern := `^invalid argument ".*" for "-s, --service" flag: unrecognized Request Service: '.*'. Must be one of: .*$`
	err := testutils_cobra.ExecutePingcli(t, "request",
		"--"+options.RequestServiceOption.CobraParamName, "invalid-service",
		"--"+options.RequestHTTPMethodOption.CobraParamName, "GET",
		fmt.Sprintf("environments/%s/populations", os.Getenv(options.PingOneAuthenticationWorkerEnvironmentIDOption.EnvVar)),
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Request Command with Invalid HTTP Method
func TestRequestCmd_Execute_InvalidHTTPMethod(t *testing.T) {
	expectedErrorPattern := `^invalid argument ".*" for "-m, --http-method" flag: unrecognized HTTP Method: '.*'. Must be one of: .*$`
	err := testutils_cobra.ExecutePingcli(t, "request",
		"--"+options.RequestServiceOption.CobraParamName, "pingone",
		"--"+options.RequestHTTPMethodOption.CobraParamName, "INVALID",
		fmt.Sprintf("environments/%s/populations", os.Getenv(options.PingOneAuthenticationWorkerEnvironmentIDOption.EnvVar)),
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Request Command with Missing Required Service Flag
func TestRequestCmd_Execute_MissingRequiredServiceFlag(t *testing.T) {
	expectedErrorPattern := `failed to send custom request: service is required`
	err := testutils_cobra.ExecutePingcli(t, "request", fmt.Sprintf("environments/%s/populations", os.Getenv(options.PingOneAuthenticationWorkerEnvironmentIDOption.EnvVar)))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Request Command with Header Flag
func TestRequestCmd_Execute_HeaderFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "request",
		"--"+options.RequestServiceOption.CobraParamName, "pingone",
		"--"+options.RequestHTTPMethodOption.CobraParamName, "GET",
		"--"+options.RequestHeaderOption.CobraParamName, "Content-Type: application/vnd.pingidentity.user.import+json",
		fmt.Sprintf("environments/%s/users", os.Getenv(options.PingOneAuthenticationWorkerEnvironmentIDOption.EnvVar)),
	)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Request Command with Header Flag with and without spacing
func TestRequestCmd_Execute_HeaderFlagSpacing(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "request",
		"--"+options.RequestServiceOption.CobraParamName, "pingone",
		"--"+options.RequestHTTPMethodOption.CobraParamName, "GET",
		"--"+options.RequestHeaderOption.CobraParamName, "Test-Header:TestValue",
		"--"+options.RequestHeaderOption.CobraParamName, "Test-Header-Two:\tTestValue",
		fmt.Sprintf("environments/%s/users", os.Getenv(options.PingOneAuthenticationWorkerEnvironmentIDOption.EnvVar)),
	)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Request Command with invalid Header Flag
func TestRequestCmd_Execute_InvalidHeaderFlag(t *testing.T) {
	expectedErrorPattern := `^invalid argument ".*" for "-r, --header" flag: failed to set Headers: Invalid header: invalid=header. Headers must be in the proper format. Expected regex pattern: .*$`
	err := testutils_cobra.ExecutePingcli(t, "request",
		"--"+options.RequestServiceOption.CobraParamName, "pingone",
		"--"+options.RequestHeaderOption.CobraParamName, "invalid=header",
		fmt.Sprintf("environments/%s/populations", os.Getenv(options.PingOneAuthenticationWorkerEnvironmentIDOption.EnvVar)),
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Request Command with disallowed Authorization Header Flag
func TestRequestCmd_Execute_DisallowedAuthorizationFlag(t *testing.T) {
	expectedErrorPattern := `^invalid argument ".*" for "-r, --header" flag: failed to set Headers: Invalid header: Authorization. Authorization header is not allowed$`
	err := testutils_cobra.ExecutePingcli(t, "request",
		"--"+options.RequestServiceOption.CobraParamName, "pingone",
		"--"+options.RequestHeaderOption.CobraParamName, "Authorization: Bearer token",
		fmt.Sprintf("environments/%s/populations", os.Getenv(options.PingOneAuthenticationWorkerEnvironmentIDOption.EnvVar)),
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
