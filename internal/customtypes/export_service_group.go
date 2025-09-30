// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

const (
	ENUM_EXPORT_SERVICE_GROUP_PINGONE string = "pingone"
)

var (
	exportServiceGroupErrorPrefix = "custom type export service group error"
	ErrUnrecognisedServiceGroup   = errors.New("unrecognized service group")
)

type ExportServiceGroup string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*ExportServiceGroup)(nil)

func (esg *ExportServiceGroup) Set(serviceGroup string) error {
	if esg == nil {
		return &errs.PingCLIError{Prefix: exportServiceGroupErrorPrefix, Err: ErrCustomTypeNil}
	}

	if serviceGroup == "" {
		return nil
	}

	// Create a map of valid service groups to check the user provided group against
	validServiceGroups := ExportServiceGroupValidValues()
	validServiceGroupMap := make(map[string]struct{}, len(validServiceGroups))
	for _, s := range validServiceGroups {
		validServiceGroupMap[s] = struct{}{}
	}

	if _, ok := validServiceGroupMap[serviceGroup]; !ok {
		return &errs.PingCLIError{Prefix: exportServiceGroupErrorPrefix, Err: fmt.Errorf("%w '%s': must be one of %s", ErrUnrecognisedServiceGroup, serviceGroup, strings.Join(validServiceGroups, ", "))}
	}

	*esg = ExportServiceGroup(serviceGroup)

	return nil
}

func (esg *ExportServiceGroup) Type() string {
	return "string"
}

func (esg *ExportServiceGroup) String() string {
	if esg == nil {
		return ""
	}
	return string(*esg)
}

func (esg *ExportServiceGroup) GetServicesInGroup() []string {
	if esg == nil {
		return []string{}
	}

	switch esg.String() {
	case ENUM_EXPORT_SERVICE_GROUP_PINGONE:
		return []string{
			ENUM_EXPORT_SERVICE_PINGONE_PLATFORM,
			ENUM_EXPORT_SERVICE_PINGONE_AUTHORIZE,
			ENUM_EXPORT_SERVICE_PINGONE_SSO,
			ENUM_EXPORT_SERVICE_PINGONE_MFA,
			ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
		}
	default:
		return []string{}
	}
}

func ExportServiceGroupValidValues() []string {
	validServiceGroups := []string{
		ENUM_EXPORT_SERVICE_GROUP_PINGONE,
	}

	slices.Sort(validServiceGroups)

	return validServiceGroups
}
