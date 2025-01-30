package testutils_resource

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
)

type ResourceCreationInfoType string

// OptionType enums
const (
	// General ID and Name enums for most resource creation
	ENUM_ID   ResourceCreationInfoType = "ENUM_ID"
	ENUM_NAME ResourceCreationInfoType = "ENUM_NAME"

	// Mapping Info for Mapping resources
	ENUM_SOURCE_ID ResourceCreationInfoType = "ENUM_SOURCE_ID"
	ENUM_TARGET_ID ResourceCreationInfoType = "ENUM_TARGET_ID"

	// Certificate Info for Certificate resources
	ENUM_ISSUER_DN     ResourceCreationInfoType = "ENUM_ISSUER_DN"
	ENUM_SERIAL_NUMBER ResourceCreationInfoType = "ENUM_SERIAL_NUMBER"

	// Miscellaneous Info for resources that don't fit the above
	ENUM_TYPE          ResourceCreationInfoType = "ENUM_TYPE"
	ENUM_CONTEXT_TYPE  ResourceCreationInfoType = "ENUM_CONTEXT_TYPE"
	ENUM_CREDENTIAL    ResourceCreationInfoType = "ENUM_CREDENTIAL"
	ENUM_SOURCE_REF_ID ResourceCreationInfoType = "ENUM_SOURCE_REF_ID"
)

type ResourceCreationInfo map[ResourceCreationInfoType]string

type TestResource struct {
	// Resources required to be created before this resource can be created
	Dependencies []TestResource

	// Creation function for this resource
	CreateFunc func(*testing.T, *connector.ClientInfo, ...string) ResourceCreationInfo

	// Deletion function for this resource
	DeleteFunc func(*testing.T, *connector.ClientInfo, string, string)

	CreationInfo ResourceCreationInfo
}

type TestableResource struct {
	ClientInfo         *connector.ClientInfo
	ExportableResource connector.ExportableResource
	TestResource       TestResource
}

func (tr *TestableResource) CreateResource(t *testing.T, testResource TestResource) ResourceCreationInfo {
	t.Helper()

	createFuncInfo := []string{tr.ExportableResource.ResourceType()}
	for _, dependency := range testResource.Dependencies {
		creationInfo := tr.CreateResource(t, dependency)
		depId, ok := creationInfo[ENUM_ID]
		if !ok {
			t.Fatalf("Failed to get ID from dependency: %v", dependency)
		}

		createFuncInfo = append(createFuncInfo, depId)
	}

	testResource.CreationInfo = testResource.CreateFunc(t, tr.ClientInfo, createFuncInfo...)

	return testResource.CreationInfo
}

func (tr *TestableResource) DeleteResource(t *testing.T, testResource TestResource) {
	t.Helper()

	if testResource.DeleteFunc == nil {
		return
	}

	testResource.DeleteFunc(t, tr.ClientInfo, tr.ExportableResource.ResourceType(), testResource.CreationInfo[ENUM_ID])

	for _, dependency := range testResource.Dependencies {
		tr.DeleteResource(t, dependency)
	}
}
