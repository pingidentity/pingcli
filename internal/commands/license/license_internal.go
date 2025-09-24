// Copyright © 2025 Ping Identity Corporation

package license_internal

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var (
	ErrLicenseDataEmpty = errors.New("returned license data is empty. please check your request parameters")
	ErrGetProduct       = errors.New("failed to get product option value")
	ErrGetVersion       = errors.New("failed to get version option value")
	ErrGetDevopsUser    = errors.New("failed to get devops user option value")
	ErrGetDevopsKey     = errors.New("failed to get devops key option value")
	ErrRequiredValues   = errors.New("product, version, devops user, and devops key must be specified for license request")
	ErrLicenseRequest   = errors.New("license request failed")
	licenseErrorPrefix  = "failed to run license request"
)

func RunInternalLicense() (err error) {
	product, version, devopsUser, devopsKey, err := readLicenseOptionValues()
	if err != nil {
		return &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: err}
	}

	ctx := context.Background()
	licenseData, err := runLicenseRequest(ctx, product, version, devopsUser, devopsKey)
	if err != nil {
		return &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: err}
	}

	if licenseData == "" {
		return &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: ErrLicenseDataEmpty}
	}

	output.Message(licenseData, nil)

	return nil
}

func readLicenseOptionValues() (product, version, devopsUser, devopsKey string, err error) {
	product, err = profiles.GetOptionValue(options.LicenseProductOption)
	if err != nil {
		return product, version, devopsUser, devopsKey, &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: fmt.Errorf("%w: %w", ErrGetProduct, err)}
	}

	version, err = profiles.GetOptionValue(options.LicenseVersionOption)
	if err != nil {
		return product, version, devopsUser, devopsKey, &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: fmt.Errorf("%w: %w", ErrGetVersion, err)}
	}

	devopsUser, err = profiles.GetOptionValue(options.LicenseDevopsUserOption)
	if err != nil {
		return product, version, devopsUser, devopsKey, &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: fmt.Errorf("%w: %w", ErrGetDevopsUser, err)}
	}

	devopsKey, err = profiles.GetOptionValue(options.LicenseDevopsKeyOption)
	if err != nil {
		return product, version, devopsUser, devopsKey, &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: fmt.Errorf("%w: %w", ErrGetDevopsKey, err)}
	}

	if product == "" || version == "" || devopsUser == "" || devopsKey == "" {
		return product, version, devopsUser, devopsKey, &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: ErrRequiredValues}
	}

	return product, version, devopsUser, devopsKey, nil
}

func runLicenseRequest(ctx context.Context, product, version, devopsUser, devopsKey string) (licenseData string, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://license.pingidentity.com/devops/license", nil)
	if err != nil {
		return licenseData, &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: err}
	}

	req.Header.Set("Devops-User", devopsUser)
	req.Header.Set("Devops-Key", devopsKey)
	req.Header.Set("Devops-App", "Ping CLI")
	req.Header.Set("Devops-Purpose", "download-license")
	req.Header.Set("Product", product)
	req.Header.Set("Version", version)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return licenseData, &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: err}
	}
	defer func() {
		cErr := res.Body.Close()
		err = errors.Join(err, cErr)
		if err != nil {
			err = &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: err}
		}
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return licenseData, &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: err}
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: fmt.Errorf("%w with status %d: %s", ErrLicenseRequest, res.StatusCode, string(body))}
	}

	return string(body), nil
}
