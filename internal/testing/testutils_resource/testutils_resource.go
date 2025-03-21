// Copyright © 2025 Ping Identity Corporation

package testutils_resource

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
)

type ResourceCreationInfoType string

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

	// Language Info for Language resources
	ENUM_LOCALE ResourceCreationInfoType = "ENUM_LOCALE"

	// Template Info for Template resources
	ENUM_TEMPLATE_VARIANT         ResourceCreationInfoType = "ENUM_TEMPLATE_VARIANT"
	ENUM_TEMPLATE_DELIVERY_METHOD ResourceCreationInfoType = "ENUM_TEMPLATE_DELIVERY_METHOD"

	// Miscellaneous Info for resources that don't fit the above
	ENUM_TYPE          ResourceCreationInfoType = "ENUM_TYPE"
	ENUM_CONTEXT_TYPE  ResourceCreationInfoType = "ENUM_CONTEXT_TYPE"
	ENUM_CREDENTIAL    ResourceCreationInfoType = "ENUM_CREDENTIAL" //#nosec G101 -- This is not hard-coded credentials
	ENUM_SOURCE_REF_ID ResourceCreationInfoType = "ENUM_SOURCE_REF_ID"
)

type ResourceCreationInfo struct {
	SelfInfo map[ResourceCreationInfoType]string
	DepIds   []string
}

// The TestableResource struct is used to create and delete resources in a test, without prior configuration needed
// on a service. This allows different developers and contributors to provide their own test service credentials,
// which would consistently create and clean configuration needed for testing without requiring shared credentials on
// a central test service.
//
// Further, this struct is notably decoupled from resource unit tests and service connector integration test. This
// allows for both tests to leverage the same struct, without worrying about setup and cleanup. Golang 'defer' applies
// to the current scope, so this allows for each test to have its own setup and cleanup, without knowing which test
// ran first or if the resource is still needed for subsequent tests.
//
// Finally, this struct allows the integration test to initialize terraform only once, which makes the terraform
// --generate-config-out testing almost an order of magnitude faster.
type TestableResource struct {
	// SDK client used in creation and deletion of this TestableResource
	ClientInfo *connector.ClientInfo

	// Creation function for this TestableResources
	CreateFunc func(*testing.T, *connector.ClientInfo, string, ...string) ResourceCreationInfo

	// TestableResource information like ID, Name, etc.
	CreationInfo ResourceCreationInfo

	// Deletion function for this TestableResources
	DeleteFunc func(*testing.T, *connector.ClientInfo, string, ...string)

	// TestableResources required to be created before this TestableResource can be created
	Dependencies []*TestableResource

	// ExportableResource that this TestableResource is testing
	ExportableResource connector.ExportableResource
}

func (tr *TestableResource) CreateResource(t *testing.T) ResourceCreationInfo {
	t.Helper()

	// Some resources like out_of_band_auth_plugins do not implement ExportableResource
	resourceType := "<nil>"
	if tr.ExportableResource != nil {
		resourceType = tr.ExportableResource.ResourceType()
	}

	createdDepIds := []string{}
	for _, dependency := range tr.Dependencies {
		// Recursively create dependencies
		dependency.CreationInfo = dependency.CreateResource(t)
		depId, ok := dependency.CreationInfo.SelfInfo[ENUM_ID]
		if !ok {
			t.Fatalf("Failed to get ID from dependency: %v", dependency)
		}

		if len(dependency.CreationInfo.DepIds) > 0 {
			createdDepIds = append(createdDepIds, dependency.CreationInfo.DepIds...)
		}
		createdDepIds = append(createdDepIds, depId)
	}

	if tr.CreateFunc != nil {
		tr.CreationInfo = tr.CreateFunc(t, tr.ClientInfo, resourceType, createdDepIds...)
	}

	return tr.CreationInfo
}

func (tr *TestableResource) DeleteResource(t *testing.T) {
	t.Helper()

	resourceType := "<nil>"
	if tr.ExportableResource != nil {
		resourceType = tr.ExportableResource.ResourceType()
	}

	ids := append(tr.CreationInfo.DepIds, tr.CreationInfo.SelfInfo[ENUM_ID])

	if tr.DeleteFunc != nil {
		tr.DeleteFunc(t, tr.ClientInfo, resourceType, ids...)
	}

	for _, dependency := range tr.Dependencies {
		dependency.DeleteResource(t)
	}
}
