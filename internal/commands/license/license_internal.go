package license_internal

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

func RunInternalLicense() (err error) {
	product, version, devopsUser, devopsKey, err := readLicenseOptionValues()
	if err != nil {
		return fmt.Errorf("failed to run license request: %w", err)
	}

	ctx := context.Background()
	licenseData, err := runLicenseRequest(ctx, product, version, devopsUser, devopsKey)
	if err != nil {
		return fmt.Errorf("failed to run license request: %w", err)
	}

	if licenseData == "" {
		return fmt.Errorf("failed to run license request: returned license data is empty, please check your request parameters")
	}

	output.Message(licenseData, nil)

	return nil
}

func readLicenseOptionValues() (product, version, devopsUser, devopsKey string, err error) {
	product, err = profiles.GetOptionValue(options.LicenseProductOption)
	if err != nil {
		return "", "", "", "", fmt.Errorf("failed to get product option: %w", err)
	}

	version, err = profiles.GetOptionValue(options.LicenseVersionOption)
	if err != nil {
		return "", "", "", "", fmt.Errorf("failed to get version option: %w", err)
	}

	devopsUser, err = profiles.GetOptionValue(options.LicenseDevopsUserOption)
	if err != nil {
		return "", "", "", "", fmt.Errorf("failed to get devops user option: %w", err)
	}

	devopsKey, err = profiles.GetOptionValue(options.LicenseDevopsKeyOption)
	if err != nil {
		return "", "", "", "", fmt.Errorf("failed to get devops key option: %w", err)
	}

	if product == "" || version == "" || devopsUser == "" || devopsKey == "" {
		return "", "", "", "", fmt.Errorf("product, version, devops user, and devops key must be specified for license request")
	}

	return product, version, devopsUser, devopsKey, nil
}

func runLicenseRequest(ctx context.Context, product, version, devopsUser, devopsKey string) (licenseData string, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://license.pingidentity.com/devops/license", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create license request: %w", err)
	}

	req.Header.Set("Devops-User", devopsUser)
	req.Header.Set("Devops-Key", devopsKey)
	req.Header.Set("Devops-App", "PingCLI")
	req.Header.Set("Devops-Purpose", "download-license")
	req.Header.Set("Product", product)
	req.Header.Set("Version", version)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute license request: %w", err)
	}
	defer func() {
		cErr := res.Body.Close()
		err = errors.Join(err, cErr)
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", fmt.Errorf("license request failed with status %d: %s", res.StatusCode, string(body))
	}

	return string(body), nil
}
