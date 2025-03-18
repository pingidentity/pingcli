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
	_ connector.ExportableResource = &PingoneAuthorizePolicyManagementStatementResource{}
)

type PingoneAuthorizePolicyManagementStatementResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizePolicyManagementStatementResource
func AuthorizePolicyManagementStatement(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizePolicyManagementStatementResource {
	return &PingoneAuthorizePolicyManagementStatementResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizePolicyManagementStatementResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	editorStatementData, err := r.getEditorStatementData()
	if err != nil {
		return nil, err
	}

	for editorStatementId, editorStatementName := range editorStatementData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Editor Statement ID":   editorStatementId,
			"Editor Statement Name": editorStatementName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       editorStatementName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, editorStatementId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizePolicyManagementStatementResource) getEditorStatementData() (map[string]string, error) {
	editorStatementData := make(map[string]string)

	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.AuthorizeEditorStatementsApi.ListStatements(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	editorStatements, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.AuthorizeEditorDataStatementsReferenceableStatementDTO](iter, "ListStatements", "GetAuthorizationStatements", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, editorStatement := range editorStatements {
		editorStatementId, editorStatementIdOk := editorStatement.GetIdOk()
		editorStatementName, editorStatementNameOk := editorStatement.GetNameOk()

		if editorStatementIdOk && editorStatementNameOk {
			editorStatementData[*editorStatementId] = *editorStatementName
		}
	}

	return editorStatementData, nil
}

func (r *PingoneAuthorizePolicyManagementStatementResource) ResourceType() string {
	return "pingone_authorize_policy_management_statement"
}
