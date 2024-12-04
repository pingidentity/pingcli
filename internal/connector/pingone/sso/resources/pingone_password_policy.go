package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOnePasswordPolicyResource{}
)

type PingOnePasswordPolicyResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOnePasswordPolicyResource
func PasswordPolicy(clientInfo *connector.PingOneClientInfo) *PingOnePasswordPolicyResource {
	return &PingOnePasswordPolicyResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOnePasswordPolicyResource) ResourceType() string {
	return "pingone_password_policy"
}

func (r *PingOnePasswordPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportPasswordPolicies()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOnePasswordPolicyResource) exportPasswordPolicies() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.PasswordPoliciesApi.ReadAllPasswordPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllPasswordPolicies", r.ResourceType())
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

		for _, passwordPolicy := range embedded.GetPasswordPolicies() {
			passwordPolicyId, passwordPolicyIdOk := passwordPolicy.GetIdOk()
			passwordPolicyName, passwordPolicyNameOk := passwordPolicy.GetNameOk()

			if passwordPolicyIdOk && passwordPolicyNameOk {
				r.addImportBlock(*passwordPolicyId, *passwordPolicyName)
			}
		}
	}

	return nil
}

func (r *PingOnePasswordPolicyResource) addImportBlock(passwordPolicyId, passwordPolicyName string) {
	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Password Policy ID":    passwordPolicyId,
		"Password Policy Name":  passwordPolicyName,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       passwordPolicyName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, passwordPolicyId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
