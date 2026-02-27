// Copyright © 2026 Ping Identity Corporation

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
	licenseErrorPrefix = "failed to run license request"
)

type licenseOptions struct {
	product    string
	version    string
	devopsUser string
	devopsKey  string
}

func RunInternalLicense() (err error) {
	opts, err := readLicenseOptionValues()
	if err != nil {
		return &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: err}
	}

	ctx := context.Background()
	licenseData, err := runLicenseRequest(ctx, opts.product, opts.version, opts.devopsUser, opts.devopsKey)
	if err != nil {
		return &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: err}
	}

	if licenseData == "" {
		return &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: ErrLicenseDataEmpty}
	}

	output.Message(licenseData, nil)

	return nil
}

func readLicenseOptionValues() (*licenseOptions, error) {
	opts := &licenseOptions{}
	var err error

	opts.product, err = profiles.GetOptionValue(options.LicenseProductOption)
	if err != nil {
		return nil, &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: fmt.Errorf("%w: %w", ErrGetProduct, err)}
	}

	opts.version, err = profiles.GetOptionValue(options.LicenseVersionOption)
	if err != nil {
		return nil, &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: fmt.Errorf("%w: %w", ErrGetVersion, err)}
	}

	opts.devopsUser, err = profiles.GetOptionValue(options.LicenseDevopsUserOption)
	if err != nil {
		return nil, &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: fmt.Errorf("%w: %w", ErrGetDevopsUser, err)}
	}

	opts.devopsKey, err = profiles.GetOptionValue(options.LicenseDevopsKeyOption)
	if err != nil {
		return nil, &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: fmt.Errorf("%w: %w", ErrGetDevopsKey, err)}
	}

	if opts.product == "" || opts.version == "" || opts.devopsUser == "" || opts.devopsKey == "" {
		return nil, &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: ErrRequiredValues}
	}

	return opts, nil
}

func runLicenseRequest(ctx context.Context, product, version, devopsUser, devopsKey string) (licenseData string, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://license.pingidentity.com/devops/license", nil)
	if err != nil {
		return licenseData, &errs.PingCLIError{Prefix: licenseErrorPrefix, Err: err}
	}

	req.Header.Set("Devops-User", devopsUser)
	req.Header.Set("Devops-Key", devopsKey)
	req.Header.Set("Devops-App", "PingCLI")
	req.Header.Set("Devops-Purpose", "get-license")
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
