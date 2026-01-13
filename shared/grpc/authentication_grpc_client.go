// Copyright © 2025 Ping Identity Corporation

package grpc

import (
	"context"

	"github.com/pingidentity/pingcli/internal/proto"
)

var _ Authentication = (*AuthenticationGRPCClient)(nil)

type AuthenticationGRPCClient struct {
	client proto.AuthenticationClient
}

func (c AuthenticationGRPCClient) GetToken() (string, error) {
	resp, err := c.client.GetToken(context.Background(), &proto.Empty{})
	if err != nil {
		return "", err
	}
	return resp.GetAccessToken(), nil
}
