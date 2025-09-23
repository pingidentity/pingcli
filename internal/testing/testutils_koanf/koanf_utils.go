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
        services: ["%s"]
    license:
        devopsUser: %s
        devopsKey: %s
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

	defaultLegacyConfigFileContentsPattern string = `activeprofile: default
default:
    description: "default description"
    nocolor: true
    outputformat: text
    export:
        outputdirectory: %s
        servicegroup: %s
        services: ["%s"]
    service:
        pingone:
            regioncode: %s
            authentication:
                type: worker
                worker:
                    clientid: %s
                    clientsecret: %s
                    environmentid: %s
        pingfederate:
            adminapipath: /pf-admin-api/v1
            authentication:
                type: basicauth
                basicauth:
                    username: Administrator
                    password: 2FederateM0re
            httpshost: https://localhost:9999
            insecuretrustalltls: true
            xbypassexternalvalidationheader: true
production:
    description: "test profile description"
    nocolor: true
    outputformat: text
    service:
        pingfederate:
            insecuretrustalltls: false
            xbypassexternalvalidationheader: false`
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

func configureMainKoanf(t *testing.T) *profiles.KoanfConfig {
	t.Helper()

	configFilePath = CreateConfigFile(t)
	koanfConfig := profiles.NewKoanfConfig(configFilePath)

	if err := koanfConfig.KoanfInstance().Load(file.Provider(configFilePath), yaml.Parser()); err != nil {
		t.Fatalf("Failed to load configuration from file '%s': %v", configFilePath, err)
	}

	return koanfConfig
}

func InitKoanfs(t *testing.T) *profiles.KoanfConfig {
	t.Helper()

	configuration.InitAllOptions()

	configFileContents = strings.Replace(getDefaultConfigFileContents(), outputDirectoryReplacement, t.TempDir()+"/config.yaml", 1)

	return configureMainKoanf(t)
}

func InitKoanfsCustomFile(t *testing.T, fileContents string) {
	t.Helper()

	configFileContents = fileContents
	configureMainKoanf(t)
}

func getDefaultConfigFileContents() string {
	return fmt.Sprintf(defaultConfigFileContentsPattern,
		outputDirectoryReplacement,
		customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
		os.Getenv("TEST_PING_IDENTITY_DEVOPS_USER"),
		os.Getenv("TEST_PING_IDENTITY_DEVOPS_KEY"),
		os.Getenv("TEST_PINGONE_REGION_CODE"),
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID"),
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET"),
		os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
	)
}

func ReturnDefaultLegacyConfigFileContents() string {
	return fmt.Sprintf(defaultLegacyConfigFileContentsPattern,
		outputDirectoryReplacement,
		customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE,
		customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
		os.Getenv("TEST_PINGONE_REGION_CODE"),
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID"),
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET"),
		os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
	)
}
