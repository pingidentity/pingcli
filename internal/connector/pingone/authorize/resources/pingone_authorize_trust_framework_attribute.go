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
	_ connector.ExportableResource = &PingoneAuthorizeTrustFrameworkAttributeResource{}
)

type PingoneAuthorizeTrustFrameworkAttributeResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizeTrustFrameworkAttributeResource
func AuthorizeTrustFrameworkAttribute(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeTrustFrameworkAttributeResource {
	return &PingoneAuthorizeTrustFrameworkAttributeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeTrustFrameworkAttributeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	editorAttributeData, err := r.getEditorAttributeData()
	if err != nil {
		return nil, err
	}

	for editorAttributeId, editorAttributeName := range editorAttributeData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Editor Attribute ID":   editorAttributeId,
			"Editor Attribute Name": editorAttributeName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       editorAttributeName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, editorAttributeId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeTrustFrameworkAttributeResource) getEditorAttributeData() (map[string]string, error) {
	editorAttributeData := make(map[string]string)

	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.AuthorizeEditorAttributesApi.ListAttributes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	editorAttributes, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.AuthorizeEditorDataDefinitionsAttributeDefinitionDTO](iter, "ListAttributes", "GetAuthorizationAttributes", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, editorAttribute := range editorAttributes {

		if me, ok := editorAttribute.GetManagedEntityOk(); ok {
			if restrictions, ok := me.GetRestrictionsOk(); ok {
				if readOnly, ok := restrictions.GetReadOnlyOk(); ok {
					if *readOnly {
						continue
					}
				}
			}
		}

		editorAttributeId, editorAttributeIdOk := editorAttribute.GetIdOk()
		editorAttributeName, editorAttributeNameOk := editorAttribute.GetFullNameOk()

		if editorAttributeIdOk && editorAttributeNameOk {
			editorAttributeData[*editorAttributeId] = *editorAttributeName
		}
	}

	return editorAttributeData, nil
}

func (r *PingoneAuthorizeTrustFrameworkAttributeResource) ResourceType() string {
	return "pingone_authorize_trust_framework_attribute"
}
