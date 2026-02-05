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
        outputDirectory: "%s"
        services: ["%s"]
    license:
        devopsUser: %s
        devopsKey: %s
    service:
        pingOne:
            regionCode: %s
            authentication:
                type: client_credentials
                environmentID: %s
                worker:
                    clientID: %s
                    clientSecret: %s
                clientCredentials:
                    clientID: %s
                    clientSecret: %s
                authorizationCode:
                    clientID: %s
                deviceCode:
                    clientID: %s
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
    export:
        outputDirectory: "%s"
        services: ["%s"]
    license:
        devopsUser: %s
        devopsKey: %s
    service:
        pingOne:
            regionCode: %s
            authentication:
                type: client_credentials
                environmentID: %s
                worker:
                    clientID: %s
                    clientSecret: %s
                clientCredentials:
                    clientID: %s
                    clientSecret: %s
                authorizationCode:
                    clientID: %s
                deviceCode:
                    clientID: %s
        pingFederate:
            adminAPIPath: /pf-admin-api/v1
            authentication:
                type: basicAuth
                basicAuth:
                    username: Administrator
                    password: 2FederateM0re
            httpsHost: https://localhost:9999
            insecureTrustAllTLS: false
            xBypassExternalValidationHeader: false`

	defaultLegacyConfigFileContentsPattern string = `activeprofile: default
default:
    description: "default description"
    nocolor: true
    outputformat: text
    export:
        outputdirectory: "%s"
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
		configFileContents = strings.ReplaceAll(GetDefaultConfigFileContents(), outputDirectoryReplacement, t.TempDir())
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
	mainKoanf := profiles.NewKoanfConfig(configFilePath)

	if err := mainKoanf.KoanfInstance().Load(file.Provider(configFilePath), yaml.Parser()); err != nil {
		t.Fatalf("Failed to load configuration from file '%s': %v", configFilePath, err)
	}
}

func InitKoanfs(t *testing.T) {
	t.Helper()

	configuration.InitAllOptions()

	configFileContents = strings.ReplaceAll(GetDefaultConfigFileContents(), outputDirectoryReplacement, t.TempDir()+"/config.yaml")

	configureMainKoanf(t)
}

func InitKoanfsCustomFile(t *testing.T, fileContents string) {
	t.Helper()

	configFileContents = fileContents
	configureMainKoanf(t)
}

func getEnvFallback(keys ...string) string {
	for _, key := range keys {
		val := os.Getenv(key)
		if val != "" {
			return val
		}
	}

	return ""
}

func GetDefaultConfigFileContents() string {
	return fmt.Sprintf(defaultConfigFileContentsPattern,
		outputDirectoryReplacement,                                                                                 // default export outputDirectory
		customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT,                                                            // default export services
		os.Getenv("TEST_PING_IDENTITY_DEVOPS_USER"),                                                                // default license devopsUser
		os.Getenv("TEST_PING_IDENTITY_DEVOPS_KEY"),                                                                 // default license devopsKey
		os.Getenv("TEST_PINGONE_REGION_CODE"),                                                                      // default service pingOne regionCode
		os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),                                                                   // default service pingOne authentication environmentID
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID"),                                                                 // default service pingOne worker clientID
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET"),                                                             // default service pingOne worker clientSecret
		getEnvFallback("TEST_PINGONE_CLIENT_ID", "PINGONE_CLIENT_ID", "TEST_PINGONE_WORKER_CLIENT_ID"),             // default service pingOne clientCredentials clientID: utilizes the client credentials or worker app credentials
		getEnvFallback("TEST_PINGONE_CLIENT_SECRET", "PINGONE_CLIENT_SECRET", "TEST_PINGONE_WORKER_CLIENT_SECRET"), // default service pingOne clientCredentials clientSecret: utilizes the client credentials or worker app credentials
		os.Getenv("TEST_PINGONE_AUTHORIZATION_CODE_CLIENT_ID"),                                                     // default service pingOne authorizationCode clientID
		os.Getenv("TEST_PINGONE_DEVICE_CODE_CLIENT_ID"),                                                            // default service pingOne deviceCode clientID
		outputDirectoryReplacement,                                                                                 // production export outputDirectory
		customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT,                                                            // production export services
		os.Getenv("TEST_PING_IDENTITY_DEVOPS_USER"),                                                                // production license devopsUser
		os.Getenv("TEST_PING_IDENTITY_DEVOPS_KEY"),                                                                 // production license devopsKey
		os.Getenv("TEST_PINGONE_REGION_CODE"),                                                                      // production service pingOne regionCode
		os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),                                                                   // production service pingOne authentication environmentID
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID"),                                                                 // production service pingOne worker clientID
		os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET"),                                                             // production service pingOne worker clientSecret
		getEnvFallback("TEST_PINGONE_CLIENT_ID", "PINGONE_CLIENT_ID", "TEST_PINGONE_WORKER_CLIENT_ID"),             // production service pingOne clientCredentials clientID: utilizes the client credentials or worker app credentials
		getEnvFallback("TEST_PINGONE_CLIENT_SECRET", "PINGONE_CLIENT_SECRET", "TEST_PINGONE_WORKER_CLIENT_SECRET"), // production service pingOne clientCredentials clientSecret: utilizes the client credentials or worker app credentials
		os.Getenv("TEST_PINGONE_AUTHORIZATION_CODE_CLIENT_ID"),                                                     // production service pingOne authorizationCode clientID
		os.Getenv("TEST_PINGONE_DEVICE_CODE_CLIENT_ID"),                                                            // production service pingOne deviceCode clientID
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
