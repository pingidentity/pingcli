package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneApplicationResourceGrantResource{}
)

type PingOneApplicationResourceGrantResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneApplicationResourceGrantResource
func ApplicationResourceGrant(clientInfo *connector.PingOneClientInfo) *PingOneApplicationResourceGrantResource {
	return &PingOneApplicationResourceGrantResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneApplicationResourceGrantResource) ResourceType() string {
	return "pingone_application_resource_grant"
}

func (r *PingOneApplicationResourceGrantResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportApplicationResourceGrants()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneApplicationResourceGrantResource) exportApplicationResourceGrants() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllApplications", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, app := range embedded.GetApplications() {
			var (
				appId     *string
				appIdOk   bool
				appName   *string
				appNameOk bool
			)

			switch {
			case app.ApplicationPingOnePortal != nil:
				appId, appIdOk = app.ApplicationPingOnePortal.GetIdOk()
				appName, appNameOk = app.ApplicationPingOnePortal.GetNameOk()
			case app.ApplicationPingOneSelfService != nil:
				appId, appIdOk = app.ApplicationPingOneSelfService.GetIdOk()
				appName, appNameOk = app.ApplicationPingOneSelfService.GetNameOk()
			case app.ApplicationExternalLink != nil:
				appId, appIdOk = app.ApplicationExternalLink.GetIdOk()
				appName, appNameOk = app.ApplicationExternalLink.GetNameOk()
			default:
				continue
			}

			if appIdOk && appNameOk {
				err := r.exportApplicationResourceGrantsByApplication(*appId, *appName)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *PingOneApplicationResourceGrantResource) exportApplicationResourceGrantsByApplication(appId, appName string) error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationResourceGrantsApi.ReadAllApplicationGrants(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, appId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllApplicationGrants", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, grant := range embedded.GetGrants() {
			grantId, grantIdOk := grant.GetIdOk()
			grantResource, grantResourceOk := grant.GetResourceOk()

			if grantIdOk && grantResourceOk {
				grantResourceId, grantResourceIdOk := grantResource.GetIdOk()

				if grantResourceIdOk {
					err := r.exportApplicationResourceGrantsByResource(appId, appName, *grantId, *grantResourceId)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (r *PingOneApplicationResourceGrantResource) exportApplicationResourceGrantsByResource(appId, appName, grantId, grantResourceId string) error {
	resource, response, err := r.clientInfo.ApiClient.ManagementAPIClient.ResourcesApi.ReadOneResource(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, grantResourceId).Execute()
	err = common.HandleClientResponse(response, err, "ReadOneResource", r.ResourceType())
	if err != nil {
		return err
	}

	if resource != nil {
		resourceName, resourceNameOk := resource.GetNameOk()
		if resourceNameOk {
			r.addImportBlock(appId, appName, grantId, *resourceName)
		}
	}

	return nil
}

func (r *PingOneApplicationResourceGrantResource) addImportBlock(appId, appName, grantId, resourceName string) {
	commentData := map[string]string{
		"Application ID":                appId,
		"Application Name":              appName,
		"Application Resource Grant ID": grantId,
		"Application Resource Name":     resourceName,
		"Export Environment ID":         r.clientInfo.ExportEnvironmentID,
		"Resource Type":                 r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_%s", appName, resourceName),
		ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, appId, grantId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
