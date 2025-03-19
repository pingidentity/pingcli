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

// Test ExportServiceGroup Set function
func Test_ExportServiceGroup_Set(t *testing.T) {
	// Create a new ExportServiceGroup
	exportServiceGroup := new(customtypes.ExportServices)

	err := exportServiceGroup.SetServiceGroup(customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}
}

// Test ExportServiceGroup Set function fails with invalid value
func Test_ExportServiceGroup_Set_InvalidValue(t *testing.T) {
	// Create a new ExportServiceGroup
	exportServiceGroup := new(customtypes.ExportServices)

	invalidValue := "invalid"

	expectedErrorPattern := `^failed to set ExportServices: Invalid service group: .*\. Allowed service group\(s\): .*$`
	err := exportServiceGroup.SetServiceGroup(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ExportServiceGroup Set function fails with nil
func Test_ExportServiceGroup_Set_Nil(t *testing.T) {
	var exportServiceGroup *customtypes.ExportServices

	val := customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE

	expectedErrorPattern := `^failed to set ExportServices group value: .* ExportServices is nil$`
	err := exportServiceGroup.SetServiceGroup(val)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test ExportServiceGroupValidValues
func Test_ExportServiceGroupValidValues(t *testing.T) {
	serviceGroupEnum := customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE

	serviceGroupValidValues := customtypes.ExportServiceGroupValidValues()
	if serviceGroupValidValues[0] != serviceGroupEnum {
		t.Errorf("ExportServiceGroupValidValues returned: %v, expected: %v", serviceGroupValidValues, serviceGroupEnum)
	}
}

// Test ExportServicePingOneValidValues
func Test_ExportServicesPingOneValidValues(t *testing.T) {
	pingOneServiceGroupValidValues := customtypes.ExportServicesPingOneValidValues()
	if len(pingOneServiceGroupValidValues) != 5 {
		t.Errorf("ExportServicesPingOneValidValues returned: %v, expected: %v", len(pingOneServiceGroupValidValues), 5)
	}
}
