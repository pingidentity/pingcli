package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneSignOnPolicyResource{}
)

type PingOneSignOnPolicyResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneSignOnPolicyResource
func SignOnPolicy(clientInfo *connector.PingOneClientInfo) *PingOneSignOnPolicyResource {
	return &PingOneSignOnPolicyResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneSignOnPolicyResource) ResourceType() string {
	return "pingone_sign_on_policy"
}

func (r *PingOneSignOnPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportSignOnPolicies()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneSignOnPolicyResource) exportSignOnPolicies() error {
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
				r.addImportBlock(*signOnPolicyId, *signOnPolicyName)
			}
		}
	}

	return nil
}

func (r *PingOneSignOnPolicyResource) addImportBlock(signOnPolicyId, signOnPolicyName string) {
	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Resource Type":         r.ResourceType(),
		"Sign-On Policy ID":     signOnPolicyId,
		"Sign-On Policy Name":   signOnPolicyName,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       signOnPolicyName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, signOnPolicyId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
