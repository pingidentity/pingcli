// Copyright © 2025 Ping Identity Corporation

package testutils_koanf

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
)

const (
	outputDirectoryReplacement = "[REPLACE_WITH_OUTPUT_DIRECTORY]"
)

var (
	configFileContents               string
	configFilePath                   string
	defaultConfigFileContentsPattern string = `activeProfile: default
default:
    description: "default description"
    noColor: true
    outputFormat: text
    export:
        outputDirectory: %s
        serviceGroup: %s
        services: ["%s"]
    service:
        pingOne:
            regionCode: %s
            authentication:
                type: worker
                worker:
                    clientID: %s
                    clientSecret: %s
                    environmentID: %s
        pingFederate:
            adminAPIPath: /pf-admin-api/v1
            authentication:
                type: basicAuth
                basicAuth:
                    username: Administrator
                    password: 2FederateM0re
            httpsHost: https://localhost:9999
            insecureTrustAllTLS: true
            xBypassExternalValidationHeader: true
production:
    description: "test profile description"
    noColor: true
    outputFormat: text
    service:
        pingFederate:
            insecureTrustAllTLS: false
            xBypassExternalValidationHeader: false`
)

func CreateConfigFile(t *testing.T) string {
	t.Helper()

	if configFileContents == "" {
		configFileContents = strings.Replace(getDefaultConfigFileContents(), outputDirectoryReplacement, t.TempDir(), 1)
	}

	configFilePath := t.TempDir() + "/config.yaml"
	if err := os.WriteFile(configFilePath, []byte(configFileContents), 0600); err != nil {
		t.Fatalf("Failed to create config file: %s", err)
	}

	return configFilePath
}

func configureMainKoanf(t *testing.T) {
	t.Helper()

	configFilePath = CreateConfigFile(t)
	mainKoanf := profiles.GetKoanfConfig()
	mainKoanf.SetKoanfConfigFile(configFilePath)

	if err := mainKoanf.KoanfInstance().Load(file.Provider(configFilePath), yaml.Parser()); err != nil {
		t.Fatalf("Failed to load configurationhere from file '%s': %v", configFilePath, err)
	}
}

func InitKoanfs(t *testing.T) {
	t.Helper()

	configuration.InitAllOptions()

	configFileContents = strings.Replace(getDefaultConfigFileContents(), outputDirectoryReplacement, t.TempDir()+"/config.yaml", 1)

	configureMainKoanf(t)
}

func InitKoanfsCustomFile(t *testing.T, fileContents string) {
	t.Helper()

	configFileContents = fileContents
	configureMainKoanf(t)
}

func getDefaultConfigFileContents() string {
	return fmt.Sprintf(defaultConfigFileContentsPattern,
		outputDirectoryReplacement,
		customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE,
		customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
		os.Getenv("TEST_PINGONE_REGION_CODE"),
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID"),
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET"),
		os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
	)
}
