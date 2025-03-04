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
	_ connector.ExportableResource = &PingoneAuthorizeTrustFrameworkConditionResource{}
)

type PingoneAuthorizeTrustFrameworkConditionResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizeTrustFrameworkConditionResource
func AuthorizeTrustFrameworkCondition(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeTrustFrameworkConditionResource {
	return &PingoneAuthorizeTrustFrameworkConditionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeTrustFrameworkConditionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	editorConditionData, err := r.getEditorConditionData()
	if err != nil {
		return nil, err
	}

	for editorConditionId, editorConditionName := range editorConditionData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Editor Condition ID":   editorConditionId,
			"Editor Condition Name": editorConditionName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       editorConditionName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, editorConditionId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeTrustFrameworkConditionResource) getEditorConditionData() (map[string]string, error) {
	editorConditionData := make(map[string]string)

	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.AuthorizeEditorConditionsApi.ListConditions(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	editorConditions, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.AuthorizeEditorDataDefinitionsConditionDefinitionDTO](iter, "ListConditions", "GetAuthorizationConditions", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, editorCondition := range editorConditions {

		editorConditionId, editorConditionIdOk := editorCondition.GetIdOk()
		editorConditionName, editorConditionNameOk := editorCondition.GetFullNameOk()

		if editorConditionIdOk && editorConditionNameOk {
			editorConditionData[*editorConditionId] = *editorConditionName
		}
	}

	return editorConditionData, nil
}

func (r *PingoneAuthorizeTrustFrameworkConditionResource) ResourceType() string {
	return "pingone_authorize_trust_framework_condition"
}
