package request_internal

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// Test RunInternalRequest function
func Test_RunInternalRequest(t *testing.T) {
	testutils_viper.InitVipers(t)

	t.Setenv(options.RequestServiceOption.EnvVar, "pingone")

	err := RunInternalRequest(fmt.Sprintf("environments/%s/populations", os.Getenv(options.PingOneAuthenticationWorkerEnvironmentIDOption.EnvVar)))
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalRequest function with fail
func Test_RunInternalRequestWithFail(t *testing.T) {

	if os.Getenv("RUN_INTERNAL_FAIL_TEST") == "true" {
		testutils_viper.InitVipers(t)
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
		cmd := exec.Command(cmdName, "-test.run=Test_RunInternalRequestWithFail")
		cmd.Env = append(os.Environ(), "RUN_INTERNAL_FAIL_TEST=true")
		err := cmd.Run()
		if exitError, ok := err.(*exec.ExitError); ok && !exitError.Success() {
			return
		}
		t.Fatalf("The process did not exit with a non-zero: %s", err)
	}
}

// Test RunInternalRequest function with empty service
func Test_RunInternalRequest_EmptyService(t *testing.T) {
	testutils_viper.InitVipers(t)

	os.Unsetenv(options.RequestServiceOption.EnvVar)

	err := RunInternalRequest("environments")
	expectedErrorPattern := "failed to send custom request: service is required"
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalRequest function with unrecognized service
func Test_RunInternalRequest_UnrecognizedService(t *testing.T) {
	testutils_viper.InitVipers(t)

	t.Setenv(options.RequestServiceOption.EnvVar, "invalid-service")

	err := RunInternalRequest("environments")
	expectedErrorPattern := "failed to send custom request: unrecognized service 'invalid-service'"
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalRequest function with valid service but invalid URI
// This should not error, but rather print a failure message with Body and status of response
func Test_RunInternalRequest_ValidService_InvalidURI(t *testing.T) {
	testutils_viper.InitVipers(t)

	t.Setenv(options.RequestServiceOption.EnvVar, "pingone")

	err := RunInternalRequest("invalid-uri")
	testutils.CheckExpectedError(t, err, nil)
}

// Test runInternalPingOneRequest function
func Test_runInternalPingOneRequest(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := runInternalPingOneRequest("environments")
	testutils.CheckExpectedError(t, err, nil)
}

// Test runInternalPingOneRequest function with invalid URI
// This should not error, but rather print a failure message with Body and status of response
func Test_runInternalPingOneRequest_InvalidURI(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := runInternalPingOneRequest("invalid-uri")
	testutils.CheckExpectedError(t, err, nil)
}

// Test getTopLevelDomain function
func Test_getTopLevelDomain(t *testing.T) {
	testutils_viper.InitVipers(t)

	t.Setenv(options.PingOneRegionCodeOption.EnvVar, customtypes.ENUM_PINGONE_REGION_CODE_CA)

	domain, err := getTopLevelDomain()
	testutils.CheckExpectedError(t, err, nil)

	expectedDomain := customtypes.ENUM_PINGONE_TLD_CA
	if domain != expectedDomain {
		t.Errorf("expected %s, got %s", expectedDomain, domain)
	}
}

// Test getTopLevelDomain function with invalid region code
func Test_getTopLevelDomain_InvalidRegionCode(t *testing.T) {
	testutils_viper.InitVipers(t)

	t.Setenv(options.PingOneRegionCodeOption.EnvVar, "invalid-region")

	_, err := getTopLevelDomain()
	expectedErrorPattern := "unrecognized PingOne region code: 'invalid-region'"
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test pingoneAccessToken function
func Test_pingoneAccessToken(t *testing.T) {
	testutils_viper.InitVipers(t)

	firstToken, err := pingoneAccessToken()
	testutils.CheckExpectedError(t, err, nil)

	// Run the function again to test caching
	secondToken, err := pingoneAccessToken()
	testutils.CheckExpectedError(t, err, nil)

	if firstToken != secondToken {
		t.Errorf("expected access token to be cached, got different tokens: %s and %s", firstToken, secondToken)
	}
}

// Test pingoneAuth function
func Test_pingoneAuth(t *testing.T) {
	testutils_viper.InitVipers(t)

	firstToken, err := pingoneAuth()
	testutils.CheckExpectedError(t, err, nil)

	// Check token was cached
	secondToken, err := pingoneAccessToken()
	testutils.CheckExpectedError(t, err, nil)

	if firstToken != secondToken {
		t.Errorf("expected access token to be cached, got different tokens: %s and %s", firstToken, secondToken)
	}
}

// Test pingoneAuth function with invalid credentials
func Test_pingoneAuth_InvalidCredentials(t *testing.T) {
	testutils_viper.InitVipers(t)

	t.Setenv(options.PingOneAuthenticationWorkerClientIDOption.EnvVar, "invalid")

	_, err := pingoneAuth()
	expectedErrorPattern := `(?s)^failed to authenticate with PingOne: Response Status 401 Unauthorized: Response Body .*$`
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test getData function
func Test_getData(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedData := "{data: 'json'}"
	t.Setenv(options.RequestDataOption.EnvVar, expectedData)

	data, err := getData()
	testutils.CheckExpectedError(t, err, nil)

	if data != expectedData {
		t.Errorf("expected %s, got %s", expectedData, data)
	}
}

// Test getData function with empty data
func Test_getData_EmptyData(t *testing.T) {
	testutils_viper.InitVipers(t)

	t.Setenv(options.RequestDataOption.EnvVar, "")

	data, err := getData()
	testutils.CheckExpectedError(t, err, nil)

	if data != "" {
		t.Errorf("expected empty data, got %s", data)
	}
}

// Test getData function with file input
func Test_getData_FileInput(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedData := "{data: 'json from file'}"
	testDir := t.TempDir()
	testFile := testDir + "/test.json"
	err := os.WriteFile(testFile, []byte(expectedData), 0600)
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	t.Setenv(options.RequestDataOption.EnvVar, "@"+testFile)

	data, err := getData()
	testutils.CheckExpectedError(t, err, nil)

	if data != expectedData {
		t.Errorf("expected %s, got %s", expectedData, data)
	}
}

// Test getData function with non-existent file input
func Test_getData_NonExistentFileInput(t *testing.T) {
	testutils_viper.InitVipers(t)

	t.Setenv(options.RequestDataOption.EnvVar, "@non_existent_file.json")

	_, err := getData()
	expectedErrorPattern := `^open .*: no such file or directory$`
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
