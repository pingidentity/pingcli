// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

const (
	ENUM_LICENSE_PRODUCT_PING_ACCESS                  string = "pingaccess"
	ENUM_LICENSE_PRODUCT_PING_AUTHORIZE               string = "pingauthorize"
	ENUM_LICENSE_PRODUCT_PING_AUTHORIZE_POLICY_EDITOR string = "pingauthorize-policy-editor"
	ENUM_LICENSE_PRODUCT_PING_CENTRAL                 string = "pingcentral"
	ENUM_LICENSE_PRODUCT_PING_DIRECTORY               string = "pingdirectory"
	ENUM_LICENSE_PRODUCT_PING_DIRECTORY_PROXY         string = "pingdirectoryproxy"
	ENUM_LICENSE_PRODUCT_PING_FEDERATE                string = "pingfederate"
)

type LicenseProduct string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*LicenseProduct)(nil)

// Implement pflag.Value interface for custom type in cobra MultiService parameter
func (lp *LicenseProduct) Set(product string) error {
	if lp == nil {
		return fmt.Errorf("failed to set LicenseProduct value: %s. LicenseProduct is nil", product)
	}

	switch {
	case strings.EqualFold(product, ENUM_LICENSE_PRODUCT_PING_ACCESS):
		*lp = LicenseProduct(ENUM_LICENSE_PRODUCT_PING_ACCESS)
	case strings.EqualFold(product, ENUM_LICENSE_PRODUCT_PING_AUTHORIZE):
		*lp = LicenseProduct(ENUM_LICENSE_PRODUCT_PING_AUTHORIZE)
	case strings.EqualFold(product, ENUM_LICENSE_PRODUCT_PING_AUTHORIZE_POLICY_EDITOR):
		*lp = LicenseProduct(ENUM_LICENSE_PRODUCT_PING_AUTHORIZE_POLICY_EDITOR)
	case strings.EqualFold(product, ENUM_LICENSE_PRODUCT_PING_CENTRAL):
		*lp = LicenseProduct(ENUM_LICENSE_PRODUCT_PING_CENTRAL)
	case strings.EqualFold(product, ENUM_LICENSE_PRODUCT_PING_DIRECTORY):
		*lp = LicenseProduct(ENUM_LICENSE_PRODUCT_PING_DIRECTORY)
	case strings.EqualFold(product, ENUM_LICENSE_PRODUCT_PING_DIRECTORY_PROXY):
		*lp = LicenseProduct(ENUM_LICENSE_PRODUCT_PING_DIRECTORY_PROXY)
	case strings.EqualFold(product, ENUM_LICENSE_PRODUCT_PING_FEDERATE):
		*lp = LicenseProduct(ENUM_LICENSE_PRODUCT_PING_FEDERATE)
	case strings.EqualFold(product, ""): // Allow empty string to be set
		*lp = LicenseProduct("")
	default:
		return fmt.Errorf("unrecognized License Product: '%s'. Must be one of: %s", product, strings.Join(LicenseProductValidValues(), ", "))
	}

	return nil
}

func (lp LicenseProduct) Type() string {
	return "string"
}

func (lp LicenseProduct) String() string {
	return string(lp)
}

func LicenseProductValidValues() []string {
	products := []string{
		ENUM_LICENSE_PRODUCT_PING_ACCESS,
		ENUM_LICENSE_PRODUCT_PING_AUTHORIZE,
		ENUM_LICENSE_PRODUCT_PING_AUTHORIZE_POLICY_EDITOR,
		ENUM_LICENSE_PRODUCT_PING_CENTRAL,
		ENUM_LICENSE_PRODUCT_PING_DIRECTORY,
		ENUM_LICENSE_PRODUCT_PING_DIRECTORY_PROXY,
		ENUM_LICENSE_PRODUCT_PING_FEDERATE,
	}

	slices.Sort(products)

	return products
}
