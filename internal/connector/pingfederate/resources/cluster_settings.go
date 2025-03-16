package resources

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
	"golang.org/x/mod/semver"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateClusterSettingsResource{}
)

type PingFederateClusterSettingsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateClusterSettingsResource
func ClusterSettings(clientInfo *connector.ClientInfo) *PingFederateClusterSettingsResource {
	return &PingFederateClusterSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateClusterSettingsResource) ResourceType() string {
	return "pingfederate_cluster_settings"
}

func (r *PingFederateClusterSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	valid, err := r.ValidPingFederateVersion()
	if err != nil {
		return nil, err
	}
	if !valid {
		l.Warn().Msgf("'%s' Resource is not supported in the version of PingFederate used. Skipping export.", r.ResourceType())
		return &importBlocks, nil
	}

	clusterSettingsId := "cluster_settings_singleton_id"
	clusterSettingsName := "Cluster Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       clusterSettingsName,
		ResourceID:         clusterSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}

func (r *PingFederateClusterSettingsResource) ValidPingFederateVersion() (bool, error) {
	versionObj, response, err := r.clientInfo.PingFederateApiClient.VersionAPI.GetVersion(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetVersion", r.ResourceType())
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	version, versionOk := versionObj.GetVersionOk()
	if !versionOk {
		return false, common.DataNilError(r.ResourceType(), response)
	}

	semVer := (*version)[:strings.LastIndex(*version, ".")]
	compareResult := semver.Compare(fmt.Sprintf("v%s", semVer), "v12.0.0")
	return compareResult >= 0, nil
}
