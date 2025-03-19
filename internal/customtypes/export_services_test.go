// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

// Test ExportServices Set function
func Test_ExportServices_Set(t *testing.T) {
	es := new(customtypes.ExportServices)

	service := customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA
	err := es.Set(service)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}

	services := es.GetServices()
	if len(services) != 1 {
		t.Errorf("GetServices returned: %v, expected: %v", services, service)
	}

	if services[0] != service {
		t.Errorf("GetServices returned: %v, expected: %v", services, service)
	}
}

// Test ExportServices Set function with invalid value
func Test_ExportServices_Set_InvalidValue(t *testing.T) {
	es := new(customtypes.ExportServices)

	invalidValue := "invalid"
	expectedErrorPattern := `^failed to set ExportServices: Invalid service: .*\. Allowed services: .*$`
	err := es.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ExportServices Set function with nil
func Test_ExportServices_Set_Nil(t *testing.T) {
	var es *customtypes.ExportServices

	service := customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA
	expectedErrorPattern := `^failed to set ExportServices value: .* ExportServices is nil$`
	err := es.Set(service)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ExportServices ContainsPingOneService function
func Test_ExportServices_ContainsPingOneService(t *testing.T) {
	es := new(customtypes.ExportServices)

	service := customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA
	err := es.Set(service)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}

	if !es.ContainsPingOneService() {
		t.Errorf("ContainsPingOneService returned false, expected true")
	}
}

// Test ExportServices ContainsPingFederateService function
func Test_ExportServices_ContainsPingFederateService(t *testing.T) {
	es := new(customtypes.ExportServices)

	service := customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE
	err := es.Set(service)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}

	if !es.ContainsPingFederateService() {
		t.Errorf("ContainsPingFederateService returned false, expected true")
	}
}

// Test ExportServices String function
func Test_ExportServices_String(t *testing.T) {
	es := new(customtypes.ExportServices)

	service := customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA
	err := es.Set(service)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}

	expected := service
	actual := es.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
