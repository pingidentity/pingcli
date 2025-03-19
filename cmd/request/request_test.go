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
	defer pipeReader.Close()
	os.Stdout = pipeWriter

	err = testutils_cobra.ExecutePingcli(t, "request",
		"--service", "pingone",
		"--http-method", "GET",
		fmt.Sprintf("environments/%s/populations", os.Getenv(options.PingOneAuthenticationWorkerEnvironmentIDOption.EnvVar)),
	)
	testutils.CheckExpectedError(t, err, nil)

	os.Stdout = originalStdout
	pipeWriter.Close()

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
		"--service", "invalid-service",
		"--http-method", "GET",
		fmt.Sprintf("environments/%s/populations", os.Getenv(options.PingOneAuthenticationWorkerEnvironmentIDOption.EnvVar)),
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Request Command with Invalid HTTP Method
func TestRequestCmd_Execute_InvalidHTTPMethod(t *testing.T) {
	expectedErrorPattern := `^invalid argument ".*" for "-m, --http-method" flag: unrecognized HTTP Method: '.*'. Must be one of: .*$`
	err := testutils_cobra.ExecutePingcli(t, "request",
		"--service", "pingone",
		"--http-method", "INVALID",
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
