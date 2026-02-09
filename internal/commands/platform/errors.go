// Copyright © 2026 Ping Identity Corporation

package platform_internal

import "errors"

var (
	exportErrorPrefix = "platform export error"

	ErrNilContext                  = errors.New("context is nil")
	ErrReadCaCertPemFile           = errors.New("failed to read CA certificate PEM file")
	ErrAppendToCertPool            = errors.New("failed to append to certificate pool from PEM file")
	ErrBasicAuthEmpty              = errors.New("failed to initialize PingFederate service. Basic authentication username and/or password is not set")
	ErrAccessTokenEmpty            = errors.New("failed to initialize PingFederate service. Access token is not set")
	ErrClientCredentialsEmpty      = errors.New("failed to initialize PingFederate service. Client ID, Client Secret, and/or Token URL is not set")
	ErrPingFederateAuthType        = errors.New("failed to initialize PingFederate service. Unrecognized authentication type")
	ErrPingFederateInit            = errors.New("failed to initialize PingFederate service. Check authentication type and credentials")
	ErrPingFederateContextNil      = errors.New("failed to initialize PingFederate services. context is nil")
	ErrPingFederateCACertParse     = errors.New("failed to parse CA certificate PEM file to certificate pool")
	ErrHttpTransportNil            = errors.New("failed to initialize PingFederate service. HTTP transport is nil")
	ErrHttpsHostEmpty              = errors.New("failed to initialize PingFederate service. HTTPS host is not set")
	ErrRegionCodeRequired          = errors.New("region code is required and must be valid. Please run 'pingcli config set service.pingone.regionCode=<region>'")
	ErrPingOneUnrecognizedAuthType = errors.New("unrecognized or unsupported PingOne authorization grant type")
	ErrPingOneConfigValuesEmpty    = errors.New("failed to initialize pingone API client. one of worker client ID, worker client secret, " +
		"pingone region code, and/or worker environment ID is not set. configure these properties via parameter flags, " +
		"environment variables, or the tool's configuration file (default: $HOME/.pingcli/config.yaml)")
	ErrPingOneInit = errors.New("failed to initialize pingone API client. Check worker client ID, worker client secret," +
		" worker environment ID, and pingone region code configuration values")
	ErrPingOneEnvironmentIDEmpty = errors.New("failed to initialize pingone API client. environment ID is empty. " +
		"configure this property via parameter flags, environment variables, or the tool's configuration file (default: $HOME/.pingcli/config.yaml)")
	ErrOutputDirectoryEmpty       = errors.New("output directory is not set")
	ErrGetPresentWorkingDirectory = errors.New("failed to get present working directory")
	ErrCreateOutputDirectory      = errors.New("failed to create output directory")
	ErrReadOutputDirectory        = errors.New("failed to read contents of output directory")
	ErrOutputDirectoryNotEmpty    = errors.New("output directory is not empty. use '--overwrite' to overwrite existing files and export data")
	ErrDeterminePingOneExportEnv  = errors.New("failed to determine pingone export environment ID")
	ErrPingOneClientNil           = errors.New("pingone API client is nil")
	ErrValidatePingOneEnvId       = errors.New("failed to validate pingone environment ID")
	ErrPingOneEnvNotExist         = errors.New("pingone environment does not exist")
	ErrConnectorListNil           = errors.New("exportable connectors list is nil")
	ErrExportService              = errors.New("failed to export service")
)
