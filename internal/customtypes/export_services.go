// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

const (
	ENUM_EXPORT_SERVICE_PINGONE_PLATFORM  string = "pingone-platform"
	ENUM_EXPORT_SERVICE_PINGONE_AUTHORIZE string = "pingone-authorize"
	ENUM_EXPORT_SERVICE_PINGONE_SSO       string = "pingone-sso"
	ENUM_EXPORT_SERVICE_PINGONE_MFA       string = "pingone-mfa"
	ENUM_EXPORT_SERVICE_PINGONE_PROTECT   string = "pingone-protect"
	ENUM_EXPORT_SERVICE_PINGFEDERATE      string = "pingfederate"
)

var (
	exportServicesErrorPrefix = "custom type export services error"
)

type ExportServices []string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*ExportServices)(nil)

// Implement pflag.Value interface for custom type in cobra MultiService parameter
func (es *ExportServices) GetServices() []string {
	if es == nil {
		return []string{}
	}

	return []string(*es)
}

func (es *ExportServices) Set(servicesStr string) error {
	if es == nil {
		return &errs.PingCLIError{Prefix: exportServicesErrorPrefix, Err: ErrCustomTypeNil}
	}

	if servicesStr == "" || servicesStr == "[]" {
		return nil
	}

	// Create a map of valid service values to check against user-provided services
	validServiceMap := ExportServicesValidValuesMap()

	// Create a map of existing services set in the ExportServices object
	existingServices := make(map[string]struct{}, len(*es))
	for _, s := range *es {
		existingServices[s] = struct{}{}
	}

	// Loop through user-provided services
	// check for valid value in map
	// check the service does not already exist
	for service := range strings.SplitSeq(servicesStr, ",") {
		service = strings.ToLower(strings.TrimSpace(service))

		enumService, ok := validServiceMap[service]
		if !ok {
			return &errs.PingCLIError{Prefix: exportServicesErrorPrefix, Err: fmt.Errorf("%w '%s': must be one of %s", ErrUnrecognizedExportService, service, strings.Join(ExportServicesValidValues(), ", "))}
		}

		if _, ok := existingServices[enumService]; ok {
			continue
		}

		*es = append(*es, enumService)
	}

	slices.Sort(*es)

	return nil
}

func (es *ExportServices) SetServicesByServiceGroup(serviceGroup *ExportServiceGroup) error {
	if es == nil {
		return &errs.PingCLIError{Prefix: exportServicesErrorPrefix, Err: ErrCustomTypeNil}
	}

	if serviceGroup == nil || serviceGroup.String() == "" {
		return nil
	}

	return es.Set(strings.Join(serviceGroup.GetServicesInGroup(), ","))
}

func (es *ExportServices) ContainsPingOneService() bool {
	if es == nil || len(*es) == 0 {
		return false
	}

	esg := ExportServiceGroup(ENUM_EXPORT_SERVICE_GROUP_PINGONE)
	servicesInGroup := esg.GetServicesInGroup()

	for _, service := range *es {
		if slices.ContainsFunc(servicesInGroup, func(s string) bool {
			return strings.EqualFold(s, service)
		}) {
			return true
		}
	}

	return false
}

func (es *ExportServices) ContainsPingFederateService() bool {
	if es == nil || len(*es) == 0 {
		return false
	}

	return slices.ContainsFunc(*es, func(s string) bool {
		return strings.EqualFold(s, ENUM_EXPORT_SERVICE_PINGFEDERATE)
	})
}

func (es *ExportServices) Type() string {
	return "[]string"
}

func (es *ExportServices) String() string {
	if es == nil {
		return ""
	}

	slices.Sort(*es)

	return strings.Join(*es, ",")
}

func (es *ExportServices) Merge(es2 *ExportServices) error {
	if es == nil {
		return &errs.PingCLIError{Prefix: exportServicesErrorPrefix, Err: ErrCustomTypeNil}
	}

	if es2 == nil {
		return nil
	}

	mergedServices := []string{}

	for _, service := range append(es.GetServices(), es2.GetServices()...) {
		if !slices.Contains(mergedServices, service) {
			mergedServices = append(mergedServices, service)
		}
	}

	slices.Sort(mergedServices)

	return es.Set(strings.Join(mergedServices, ","))
}

func ExportServicesValidValues() []string {
	allServices := []string{
		ENUM_EXPORT_SERVICE_PINGFEDERATE,
		ENUM_EXPORT_SERVICE_PINGONE_PLATFORM,
		ENUM_EXPORT_SERVICE_PINGONE_AUTHORIZE,
		ENUM_EXPORT_SERVICE_PINGONE_SSO,
		ENUM_EXPORT_SERVICE_PINGONE_MFA,
		ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
	}

	slices.Sort(allServices)

	return allServices
}

// ExportServicesValidValuesMap returns a map of valid export service values with lowercase keys
func ExportServicesValidValuesMap() map[string]string {
	validServices := ExportServicesValidValues()
	validServiceMap := make(map[string]string, len(validServices))
	for _, s := range validServices {
		validServiceMap[strings.ToLower(s)] = s
	}

	return validServiceMap
}
