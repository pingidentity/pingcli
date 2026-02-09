// Copyright © 2026 Ping Identity Corporation

package license_internal

import "errors"

var (
	ErrLicenseDataEmpty = errors.New("returned license data is empty. please check your request parameters")
	ErrGetProduct       = errors.New("failed to get product option value")
	ErrGetVersion       = errors.New("failed to get version option value")
	ErrGetDevopsUser    = errors.New("failed to get devops user option value")
	ErrGetDevopsKey     = errors.New("failed to get devops key option value")
	ErrRequiredValues   = errors.New("product, version, devops user, and devops key must be specified for license request")
	ErrLicenseRequest   = errors.New("license request failed")
)
