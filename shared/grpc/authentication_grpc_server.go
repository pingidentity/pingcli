// Copyright © 2025 Ping Identity Corporation

package grpc

import (
	"context"

	"github.com/pingidentity/pingcli/internal/proto"
)

var _ proto.AuthenticationServer = (*AuthenticationGRPCServer)(nil)

type AuthenticationGRPCServer struct {
	Impl Authentication
	proto.UnimplementedAuthenticationServer
}

func (s *AuthenticationGRPCServer) GetToken(ctx context.Context, req *proto.Empty) (*proto.AuthenticationToken, error) {
	token, err := s.Impl.GetToken()
	if err != nil {
		return nil, err
	}

	return &proto.AuthenticationToken{
		AccessToken: &token,
	}, nil
}
