// Copyright © 2026 Ping Identity Corporation

package request_internal

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// Test RunInternalRequest function with fail
func Test_RunInternalRequestWithFail(t *testing.T) {
	if os.Getenv("RUN_INTERNAL_FAIL_TEST") == "true" {
		testutils_koanf.InitKoanfs(t)
		t.Setenv(options.RequestServiceOption.EnvVar, "pingone")
		options.RequestFailOption.Flag.Changed = true
		err := options.RequestFailOption.Flag.Value.Set("true")
		if err != nil {
			t.Fatal(err)
		}
		_ = RunInternalRequest("environments/failTest")
		t.Fatal("This should never run due to internal request resulting in os.Exit(1)")
	} else {
		cmdName := os.Args[0]
		cmd := exec.CommandContext(t.Context(), cmdName, "-test.run=Test_RunInternalRequestWithFail") //#nosec G204 -- This is a test
		cmd.Env = append(os.Environ(), "RUN_INTERNAL_FAIL_TEST=true")
		err := cmd.Run()

		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			if !exitErr.Success() {
				return
			}
		}

		t.Fatalf("The process did not exit with a non-zero: %s", err)
	}
}

// Test RunInternalRequest function with empty service
func Test_RunInternalRequest_EmptyService(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := os.Unsetenv(options.RequestServiceOption.EnvVar)
	if err != nil {
		t.Fatalf("failed to unset environment variable: %v", err)
	}

	err = RunInternalRequest("environments")
	expectedErrorPattern := "service is not set"
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalRequest function with unrecognized service
func Test_RunInternalRequest_UnrecognizedService(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	t.Setenv(options.RequestServiceOption.EnvVar, "invalid-service")

	err := RunInternalRequest("environments")
	expectedErrorPattern := "unrecognized service.*invalid-service"
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test getData function
func Test_getDataRaw(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedData := "{data: 'json'}"
	t.Setenv(options.RequestDataRawOption.EnvVar, expectedData)

	data, err := GetDataRaw()
	testutils.CheckExpectedError(t, err, nil)

	if data != expectedData {
		t.Errorf("expected %s, got %s", expectedData, data)
	}
}

// Test getData function with empty data
func Test_getDataRaw_EmptyData(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	t.Setenv(options.RequestDataRawOption.EnvVar, "")

	data, err := GetDataRaw()
	testutils.CheckExpectedError(t, err, nil)

	if data != "" {
		t.Errorf("expected empty data, got %s", data)
	}
}

// Test getData function with file input
func Test_getDataFile_FileInput(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedData := "{data: 'json from file'}"
	testDir := t.TempDir()
	testFile := testDir + "/test.json"
	err := os.WriteFile(testFile, []byte(expectedData), 0600)
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	t.Setenv(options.RequestDataOption.EnvVar, testFile)

	data, err := GetDataFile()
	testutils.CheckExpectedError(t, err, nil)

	if data != expectedData {
		t.Errorf("expected %s, got %s", expectedData, data)
	}
}

// Test getData function with non-existent file input
func Test_getDataFile_NonExistentFileInput(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	t.Setenv(options.RequestDataOption.EnvVar, "non_existent_file.json")

	_, err := GetDataFile()
	expectedErrorPattern := `^open .*: no such file or directory$`
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
