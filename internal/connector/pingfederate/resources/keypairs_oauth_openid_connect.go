// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateKeypairsOauthOpenidConnectResource{}
)

type PingFederateKeypairsOauthOpenidConnectResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateKeypairsOauthOpenidConnectResource
func KeypairsOauthOpenidConnect(clientInfo *connector.ClientInfo) *PingFederateKeypairsOauthOpenidConnectResource {
	return &PingFederateKeypairsOauthOpenidConnectResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateKeypairsOauthOpenidConnectResource) ResourceType() string {
	return "pingfederate_keypairs_oauth_openid_connect"
}

func (r *PingFederateKeypairsOauthOpenidConnectResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	keypairsOauthOpenidConnectId := "keypairs_oauth_openid_connect_singleton_id"
	keypairsOauthOpenidConnectName := "Keypairs Oauth Openid Connect"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       keypairsOauthOpenidConnectName,
		ResourceID:         keypairsOauthOpenidConnectId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
