package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizeTrustFrameworkServiceResource{}
)

type PingoneAuthorizeTrustFrameworkServiceResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizeTrustFrameworkServiceResource
func AuthorizeTrustFrameworkService(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeTrustFrameworkServiceResource {
	return &PingoneAuthorizeTrustFrameworkServiceResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeTrustFrameworkServiceResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	editorServiceData, err := r.getEditorServiceData()
	if err != nil {
		return nil, err
	}

	for editorServiceId, editorServiceName := range editorServiceData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Editor Service ID":     editorServiceId,
			"Editor Service Name":   editorServiceName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       editorServiceName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, editorServiceId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeTrustFrameworkServiceResource) getEditorServiceData() (map[string]string, error) {
	editorServiceData := make(map[string]string)

	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.AuthorizeEditorServicesApi.ListServices(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	editorServices, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.AuthorizeEditorDataDefinitionsServiceDefinitionDTO](iter, "ListServices", "GetAuthorizationServices", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, editorService := range editorServices {

		var (
			editorServiceId     *string
			editorServiceIdOk   bool
			editorServiceName   *string
			editorServiceNameOk bool
		)

		switch t := editorService.GetActualInstance().(type) {
		case *authorize.AuthorizeEditorDataServicesConnectorServiceDefinitionDTO:
			editorServiceId, editorServiceIdOk = t.GetIdOk()
			editorServiceName, editorServiceNameOk = t.GetFullNameOk()
		case *authorize.AuthorizeEditorDataServicesHttpServiceDefinitionDTO:
			editorServiceId, editorServiceIdOk = t.GetIdOk()
			editorServiceName, editorServiceNameOk = t.GetFullNameOk()
		case *authorize.AuthorizeEditorDataServicesNoneServiceDefinitionDTO:
			editorServiceId, editorServiceIdOk = t.GetIdOk()
			editorServiceName, editorServiceNameOk = t.GetFullNameOk()
		default:
			continue
		}

		if editorServiceIdOk && editorServiceNameOk {
			editorServiceData[*editorServiceId] = *editorServiceName
		}
	}

	return editorServiceData, nil
}

func (r *PingoneAuthorizeTrustFrameworkServiceResource) ResourceType() string {
	return "pingone_authorize_trust_framework_service"
}
