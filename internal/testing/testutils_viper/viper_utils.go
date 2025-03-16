package testutils_viper

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/profiles"
)

const (
	outputDirectoryReplacement = "[REPLACE_WITH_OUTPUT_DIRECTORY]"
)

var (
	configFileContents               string
	defaultConfigFileContentsPattern string = `activeProfile: default
default:
    description: "default description"
    noColor: true
    outputFormat: text
    export:
        outputDirectory: %s
    service:
        pingone:
            regionCode: %s
            authentication:
                type: worker
                worker:
                    clientid: %s
                    clientsecret: %s
                    environmentid: %s
        pingfederate:
            adminapipath: %s
            authentication:
                type: basicauth
                basicauth:
                    username: %s
                    password: %s
            httpshost: %s
            insecureTrustAllTLS: true
            xBypassExternalValidationHeader: true
production:
    description: "test profile description"
    noColor: true
    outputFormat: text
    service:
        pingfederate:
            insecureTrustAllTLS: false
            xBypassExternalValidationHeader: false`
)

func CreateConfigFile(t *testing.T) string {
	t.Helper()

	if configFileContents == "" {
		configFileContents = strings.Replace(getDefaultConfigFileContents(), outputDirectoryReplacement, t.TempDir(), 1)
	}

	configFilepath := t.TempDir() + "/config.yaml"
	if err := os.WriteFile(configFilepath, []byte(configFileContents), 0600); err != nil {
		t.Fatalf("Failed to create config file: %s", err)
	}

	return configFilepath
}

func configureMainViper(t *testing.T) {
	t.Helper()

	// Create and write to a temporary config file
	configFilepath := CreateConfigFile(t)
	// Give main viper instance a file location to write to
	mainViper := profiles.GetMainConfig().ViperInstance()
	mainViper.SetConfigFile(configFilepath)
	if err := mainViper.ReadInConfig(); err != nil {
		t.Fatal(err)
	}

	activePName := profiles.GetMainConfig().ViperInstance().GetString(options.RootActiveProfileOption.ViperKey)

	if err := profiles.GetMainConfig().ChangeActiveProfile(activePName); err != nil {
		t.Fatal(err)
	}
}

func InitVipers(t *testing.T) {
	t.Helper()

	configuration.InitAllOptions()

	configFileContents = strings.Replace(getDefaultConfigFileContents(), outputDirectoryReplacement, t.TempDir(), 1)

	configureMainViper(t)
}

func InitVipersCustomFile(t *testing.T, fileContents string) {
	t.Helper()

	configFileContents = fileContents
	configureMainViper(t)
}

func getDefaultConfigFileContents() string {
	return fmt.Sprintf(defaultConfigFileContentsPattern,
		outputDirectoryReplacement,
		os.Getenv(options.PingOneRegionCodeOption.EnvVar),
		os.Getenv(options.PingOneAuthenticationWorkerClientIDOption.EnvVar),
		os.Getenv(options.PingOneAuthenticationWorkerClientSecretOption.EnvVar),
		os.Getenv(options.PingOneAuthenticationWorkerEnvironmentIDOption.EnvVar),
		os.Getenv(options.PingFederateAdminAPIPathOption.EnvVar),
		os.Getenv(options.PingFederateBasicAuthUsernameOption.EnvVar),
		os.Getenv(options.PingFederateBasicAuthPasswordOption.EnvVar),
		os.Getenv(options.PingFederateHTTPSHostOption.EnvVar))
}
