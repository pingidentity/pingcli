package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

const (
	ENUM_EXPORT_SERVICE_GROUP_PINGONE string = "pingone"
)

type ExportServiceGroup string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*ExportServiceGroup)(nil)

func (esg *ExportServiceGroup) Set(serviceGroup string) error {
	if esg == nil {
		return fmt.Errorf("failed to set ExportServiceGroup value: %s. ExportServiceGroup is nil", serviceGroup)
	}

	if serviceGroup == "" {
		return nil
	}

	switch {
	case strings.EqualFold(serviceGroup, ENUM_EXPORT_SERVICE_GROUP_PINGONE):
		*esg = ExportServiceGroup(ENUM_EXPORT_SERVICE_GROUP_PINGONE)
	default:
		return fmt.Errorf("unrecognized service group '%s'. Must be one of: %s", serviceGroup, strings.Join(ExportServiceGroupValidValues(), ", "))
	}
	return nil
}

func (es *ExportServices) SetServiceGroup(serviceGroup string) error {
	if es == nil {
		return fmt.Errorf("failed to set ExportServices value: %s. ExportServices is nil", serviceGroup)
	}

	switch {
	case strings.EqualFold(ENUM_EXPORT_SERVICE_GROUP_PINGONE, serviceGroup):
		*es = append(*es, ExportServicesPingOneValidValues()...)
	}

	slices.Sort(*es)
	return nil
}

func (esg ExportServiceGroup) Type() string {
	return "string"
}

func (esg ExportServiceGroup) String() string {
	return string(esg)
}

func ExportServiceGroupValidValues() []string {
	validServiceGroups := []string{
		ENUM_EXPORT_SERVICE_GROUP_PINGONE,
	}

	slices.Sort(validServiceGroups)

	return validServiceGroups
}
