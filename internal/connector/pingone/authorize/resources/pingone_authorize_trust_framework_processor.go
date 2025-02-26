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
	_ connector.ExportableResource = &PingoneAuthorizeTrustFrameworkProcessorResource{}
)

type PingoneAuthorizeTrustFrameworkProcessorResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizeTrustFrameworkProcessorResource
func AuthorizeTrustFrameworkProcessor(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeTrustFrameworkProcessorResource {
	return &PingoneAuthorizeTrustFrameworkProcessorResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeTrustFrameworkProcessorResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	editorProcessorData, err := r.getEditorProcessorData()
	if err != nil {
		return nil, err
	}

	for editorProcessorId, editorProcessorName := range editorProcessorData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Editor Processor ID":   editorProcessorId,
			"Editor Processor Name": editorProcessorName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       editorProcessorName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, editorProcessorId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeTrustFrameworkProcessorResource) getEditorProcessorData() (map[string]string, error) {
	editorProcessorData := make(map[string]string)

	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.AuthorizeEditorProcessorsApi.ListProcessors(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	editorProcessors, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.AuthorizeEditorDataDefinitionsProcessorDefinitionDTO](iter, "ListProcessors", "GetAuthorizationProcessors", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, editorProcessor := range editorProcessors {

		editorProcessorId, editorProcessorIdOk := editorProcessor.GetIdOk()
		editorProcessorName, editorProcessorNameOk := editorProcessor.GetFullNameOk()

		if editorProcessorIdOk && editorProcessorNameOk {
			editorProcessorData[*editorProcessorId] = *editorProcessorName
		}
	}

	return editorProcessorData, nil
}

func (r *PingoneAuthorizeTrustFrameworkProcessorResource) ResourceType() string {
	return "pingone_authorize_trust_framework_processor"
}
