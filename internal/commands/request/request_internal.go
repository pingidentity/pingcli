// Copyright © 2025 Ping Identity Corporation

package request_internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalRequest(uri string) (err error) {
	service, err := profiles.GetOptionValue(options.RequestServiceOption)
	if err != nil {
		return fmt.Errorf("failed to send custom request: %w", err)
	}

	if service == "" {
		return fmt.Errorf("failed to send custom request: service is required")
	}

	switch service {
	case customtypes.ENUM_REQUEST_SERVICE_PINGONE:
		err = runInternalPingOneRequest(uri)
		if err != nil {
			return fmt.Errorf("failed to send custom request: %w", err)
		}
	default:
		return fmt.Errorf("failed to send custom request: unrecognized service '%s'", service)
	}

	return nil
}

func GetDataFile() (data string, err error) {
	dataFilepath, err := profiles.GetOptionValue(options.RequestDataOption)
	if err != nil {
		return "", err
	}

	if dataFilepath != "" {
		dataFilepath = filepath.Clean(dataFilepath)
		contents, err := os.ReadFile(dataFilepath)
		if err != nil {
			return "", err
		}

		return string(contents), nil
	}

	return "", nil
}

func GetDataRaw() (data string, err error) {
	data, err = profiles.GetOptionValue(options.RequestDataRawOption)
	if err != nil {
		return "", err
	}

	return data, nil
}
