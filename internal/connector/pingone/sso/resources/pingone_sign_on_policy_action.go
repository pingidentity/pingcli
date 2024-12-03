package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneSignOnPolicyActionResource{}
)

type PingOneSignOnPolicyActionResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneSignOnPolicyActionResource
func SignOnPolicyAction(clientInfo *connector.PingOneClientInfo) *PingOneSignOnPolicyActionResource {
	return &PingOneSignOnPolicyActionResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneSignOnPolicyActionResource) ResourceType() string {
	return "pingone_sign_on_policy_action"
}

func (r *PingOneSignOnPolicyActionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportSignOnPolicyActions()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneSignOnPolicyActionResource) exportSignOnPolicyActions() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.SignOnPoliciesApi.ReadAllSignOnPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllSignOnPolicies", r.ResourceType())
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

		for _, signOnPolicy := range embedded.GetSignOnPolicies() {
			signOnPolicyId, signOnPolicyIdOk := signOnPolicy.GetIdOk()
			signOnPolicyName, signOnPolicyNameOk := signOnPolicy.GetNameOk()

			if signOnPolicyIdOk && signOnPolicyNameOk {
				err := r.exportSignOnPolicyActionsBySignOnPolicy(*signOnPolicyId, *signOnPolicyName)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *PingOneSignOnPolicyActionResource) exportSignOnPolicyActionsBySignOnPolicy(signOnPolicyId, signOnPolicyName string) error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.SignOnPolicyActionsApi.ReadAllSignOnPolicyActions(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, signOnPolicyId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllSignOnPolicyActions", r.ResourceType())
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

		for _, action := range embedded.GetActions() {
			var (
				actionId     *string
				actionIdOk   bool
				actionType   *management.EnumSignOnPolicyType
				actionTypeOk bool
			)

			switch {
			case action.SignOnPolicyActionAgreement != nil:
				actionId, actionIdOk = action.SignOnPolicyActionAgreement.GetIdOk()
				actionType, actionTypeOk = action.SignOnPolicyActionAgreement.GetTypeOk()
			case action.SignOnPolicyActionCommon != nil:
				actionId, actionIdOk = action.SignOnPolicyActionCommon.GetIdOk()
				actionType, actionTypeOk = action.SignOnPolicyActionCommon.GetTypeOk()
			case action.SignOnPolicyActionIDFirst != nil:
				actionId, actionIdOk = action.SignOnPolicyActionIDFirst.GetIdOk()
				actionType, actionTypeOk = action.SignOnPolicyActionIDFirst.GetTypeOk()
			case action.SignOnPolicyActionIDP != nil:
				actionId, actionIdOk = action.SignOnPolicyActionIDP.GetIdOk()
				actionType, actionTypeOk = action.SignOnPolicyActionIDP.GetTypeOk()
			case action.SignOnPolicyActionLogin != nil:
				actionId, actionIdOk = action.SignOnPolicyActionLogin.GetIdOk()
				actionType, actionTypeOk = action.SignOnPolicyActionLogin.GetTypeOk()
			case action.SignOnPolicyActionMFA != nil:
				actionId, actionIdOk = action.SignOnPolicyActionMFA.GetIdOk()
				actionType, actionTypeOk = action.SignOnPolicyActionMFA.GetTypeOk()
			case action.SignOnPolicyActionPingIDWinLoginPasswordless != nil:
				actionId, actionIdOk = action.SignOnPolicyActionPingIDWinLoginPasswordless.GetIdOk()
				actionType, actionTypeOk = action.SignOnPolicyActionPingIDWinLoginPasswordless.GetTypeOk()
			default:
				continue
			}

			if actionIdOk && actionTypeOk {
				r.addImportBlock(signOnPolicyId, signOnPolicyName, *actionId, string(*actionType))
			}
		}
	}

	return nil
}

func (r *PingOneSignOnPolicyActionResource) addImportBlock(signOnPolicyId, signOnPolicyName, actionId, actionType string) {
	commentData := map[string]string{
		"Export Environment ID":      r.clientInfo.ExportEnvironmentID,
		"Resource Type":              r.ResourceType(),
		"Sign-On Policy Action ID":   actionId,
		"Sign-On Policy Action Type": actionType,
		"Sign-On Policy ID":          signOnPolicyId,
		"Sign-On Policy Name":        signOnPolicyName,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_%s", signOnPolicyName, actionType),
		ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, signOnPolicyId, actionId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
