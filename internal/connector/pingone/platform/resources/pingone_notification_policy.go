package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneNotificationPolicyResource{}
)

type PingOneNotificationPolicyResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneNotificationPolicyResource
func NotificationPolicy(clientInfo *connector.PingOneClientInfo) *PingOneNotificationPolicyResource {
	return &PingOneNotificationPolicyResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneNotificationPolicyResource) ResourceType() string {
	return "pingone_notification_policy"
}

func (r *PingOneNotificationPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportNotificationPolicies()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneNotificationPolicyResource) exportNotificationPolicies() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.NotificationsPoliciesApi.ReadAllNotificationsPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllNotificationsPolicies", r.ResourceType())
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

		for _, notificationPolicy := range embedded.GetNotificationsPolicies() {
			notificationPolicyId, notificationPolicyIdOk := notificationPolicy.GetIdOk()
			notificationPolicyName, notificationPolicyNameOk := notificationPolicy.GetNameOk()

			if notificationPolicyIdOk && notificationPolicyNameOk {
				r.addImportBlock(*notificationPolicyId, *notificationPolicyName)
			}
		}
	}

	return nil
}

func (r *PingOneNotificationPolicyResource) addImportBlock(notificationPolicyId, notificationPolicyName string) {
	commentData := map[string]string{
		"Export Environment ID":    r.clientInfo.ExportEnvironmentID,
		"Notification Policy ID":   notificationPolicyId,
		"Notification Policy Name": notificationPolicyName,
		"Resource Type":            r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       notificationPolicyName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, notificationPolicyId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
