// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneApplicationResourceGrantResource{}
)

type PingOneApplicationResourceGrantResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneApplicationResourceGrantResource
func ApplicationResourceGrant(clientInfo *connector.ClientInfo) *PingOneApplicationResourceGrantResource {
	return &PingOneApplicationResourceGrantResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneApplicationResourceGrantResource) ResourceType() string {
	return "pingone_application_resource_grant"
}

func (r *PingOneApplicationResourceGrantResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationData, err := r.getApplicationData()
	if err != nil {
		return nil, err
	}

	for applicationId, applicationName := range applicationData {
		applicationResourceGrantData, err := r.getApplicationResourceGrantData(applicationId)
		if err != nil {
			return nil, err
		}

		for applicationResourceGrantId, applicationResourceGrantResourceId := range applicationResourceGrantData {
			resourceName, resourceNameOk, err := r.getGrantResourceName(applicationResourceGrantResourceId)
			if err != nil {
				return nil, err
			}

			if !resourceNameOk {
				continue
			}

			commentData := map[string]string{
				"Application ID":                           applicationId,
				"Application Name":                         applicationName,
				"Application Resource Grant ID":            applicationResourceGrantId,
				"Application Resource Grant Resource Name": resourceName,
				"Export Environment ID":                    r.clientInfo.PingOneExportEnvironmentID,
				"Resource Type":                            r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", applicationName, resourceName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, applicationId, applicationResourceGrantId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationResourceGrantResource) getApplicationData() (map[string]string, error) {
	applicationData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.ReadOneApplication200Response](iter, "ReadAllApplications", "GetApplications", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, innerObj := range apiObjs {
		var (
			applicationId     *string
			applicationIdOk   bool
			applicationName   *string
			applicationNameOk bool
		)

		switch {
		case innerObj.ApplicationPingOnePortal != nil:
			applicationId, applicationIdOk = innerObj.ApplicationPingOnePortal.GetIdOk()
			applicationName, applicationNameOk = innerObj.ApplicationPingOnePortal.GetNameOk()
		case innerObj.ApplicationPingOneSelfService != nil:
			applicationId, applicationIdOk = innerObj.ApplicationPingOneSelfService.GetIdOk()
			applicationName, applicationNameOk = innerObj.ApplicationPingOneSelfService.GetNameOk()
		case innerObj.ApplicationExternalLink != nil:
			applicationId, applicationIdOk = innerObj.ApplicationExternalLink.GetIdOk()
			applicationName, applicationNameOk = innerObj.ApplicationExternalLink.GetNameOk()
		case innerObj.ApplicationOIDC != nil:
			applicationId, applicationIdOk = innerObj.ApplicationOIDC.GetIdOk()
			applicationName, applicationNameOk = innerObj.ApplicationOIDC.GetNameOk()
		case innerObj.ApplicationSAML != nil:
			applicationId, applicationIdOk = innerObj.ApplicationSAML.GetIdOk()
			applicationName, applicationNameOk = innerObj.ApplicationSAML.GetNameOk()
		default:
			continue
		}

		if applicationIdOk && applicationNameOk {
			applicationData[*applicationId] = *applicationName
		}
	}

	return applicationData, nil
}

func (r *PingOneApplicationResourceGrantResource) getApplicationResourceGrantData(applicationId string) (map[string]string, error) {
	applicationResourceGrantData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationResourceGrantsApi.ReadAllApplicationGrants(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, applicationId).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.ApplicationResourceGrant](iter, "ReadAllApplicationGrants", "GetGrants", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, applicationResourceGrant := range apiObjs {
		applicationResourceGrantId, applicationResourceGrantIdOk := applicationResourceGrant.GetIdOk()
		applicationResourceGrantResource, applicationResourceGrantResourceOk := applicationResourceGrant.GetResourceOk()

		if applicationResourceGrantIdOk && applicationResourceGrantResourceOk {
			applicationResourceGrantResourceId, applicationResourceGrantResourceIdOk := applicationResourceGrantResource.GetIdOk()

			if applicationResourceGrantResourceIdOk {
				applicationResourceGrantData[*applicationResourceGrantId] = *applicationResourceGrantResourceId
			}
		}
	}

	return applicationResourceGrantData, nil
}

func (r *PingOneApplicationResourceGrantResource) getGrantResourceName(grantResourceId string) (string, bool, error) {
	resource, response, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.ResourcesApi.ReadOneResource(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, grantResourceId).Execute()
	ok, err := common.HandleClientResponse(response, err, "ReadOneResource", r.ResourceType())
	if err != nil {
		return "", false, err
	}
	if !ok {
		return "", false, nil
	}

	if resource != nil {
		resourceName, resourceNameOk := resource.GetNameOk()
		if resourceNameOk {
			return *resourceName, true, nil
		}
	}

	return "", false, fmt.Errorf("unable to get resource name for grant resource ID: %s", grantResourceId)
}
