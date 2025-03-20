// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"slices"
	"strings"

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

type ExportServices []string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*ExportServices)(nil)

// Implement pflag.Value interface for custom type in cobra MultiService parameter
func (es ExportServices) GetServices() []string {
	return []string(es)
}

func (es *ExportServices) Set(services string) error {
	if es == nil {
		return fmt.Errorf("failed to set ExportServices value: %s. ExportServices is nil", services)
	}

	if services == "" || services == "[]" {
		*es = ExportServices([]string{})
		return nil
	}

	validServices := ExportServicesValidValues()
	serviceList := strings.Split(services, ",")
	returnServiceList := []string{}

	for _, service := range serviceList {
		if !slices.ContainsFunc(validServices, func(validService string) bool {
			if strings.EqualFold(validService, service) {
				if !slices.Contains(returnServiceList, validService) {
					returnServiceList = append(returnServiceList, validService)
				}
				return true
			}
			return false
		}) {
			return fmt.Errorf("failed to set ExportServices: Invalid service: %s. Allowed services: %s", service, strings.Join(validServices, ", "))
		}
	}

	slices.Sort(returnServiceList)

	*es = ExportServices(returnServiceList)
	return nil
}

func (es *ExportServices) SetServicesByServiceGroup(serviceGroup *ExportServiceGroup) error {
	if es == nil {
		return fmt.Errorf("failed to set ExportServices value: %s. ExportServices is nil", serviceGroup)
	}

	switch {
	case strings.EqualFold(ENUM_EXPORT_SERVICE_GROUP_PINGONE, serviceGroup.String()):
		return es.Set(strings.Join(ExportServicesPingOneValidValues(), ","))
	default:
		return fmt.Errorf("failed to SetServicesByServiceGroup: Invalid service group: %s. Allowed services: %s", serviceGroup.String(), strings.Join(ExportServiceGroupValidValues(), ", "))
	}
}

func (es ExportServices) ContainsPingOneService() bool {
	if es == nil {
		return false
	}

	pingoneServices := ExportServicesPingOneValidValues()

	for _, service := range es {
		if slices.ContainsFunc(pingoneServices, func(s string) bool {
			return strings.EqualFold(s, service)
		}) {
			return true
		}
	}

	return false
}

func (es ExportServices) ContainsPingFederateService() bool {
	if es == nil {
		return false
	}

	return slices.Contains(es, ENUM_EXPORT_SERVICE_PINGFEDERATE)
}

func (es ExportServices) Type() string {
	return "[]string"
}

func (es ExportServices) String() string {
	return strings.Join(es, ",")
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

func ExportServicesPingOneValidValues() []string {
	pingOneServices := []string{
		ENUM_EXPORT_SERVICE_PINGONE_PLATFORM,
		ENUM_EXPORT_SERVICE_PINGONE_AUTHORIZE,
		ENUM_EXPORT_SERVICE_PINGONE_SSO,
		ENUM_EXPORT_SERVICE_PINGONE_MFA,
		ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
	}

	slices.Sort(pingOneServices)

	return pingOneServices
}

func (es *ExportServices) Merge(es2 ExportServices) error {
	mergedServices := []string{}

	for _, service := range append(es.GetServices(), es2.GetServices()...) {
		if !slices.Contains(mergedServices, service) {
			mergedServices = append(mergedServices, service)
		}
	}

	slices.Sort(mergedServices)
	return es.Set(strings.Join(mergedServices, ","))
}
