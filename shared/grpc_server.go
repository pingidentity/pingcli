// Copyright © 2025 Ping Identity Corporation

package shared

import (
	"context"

	"github.com/pingidentity/pingcli/shared/proto"
)

var _ proto.ProtoPingCliCommandServer = (*GRPCServer)(nil)

type GRPCServer struct {
	Impl PingCliCommand
	proto.UnimplementedProtoPingCliCommandServer
}

func (s *GRPCServer) Configuration(ctx context.Context, req *proto.Empty) (*proto.ConfigurationResponse, error) {
	cmd, err := s.Impl.Configuration()
	if err != nil {
		return nil, err
	}

	return &proto.ConfigurationResponse{
		Example: cmd.Example,
		Long:    cmd.Long,
		Short:   cmd.Short,
		Use:     cmd.Use,
	}, nil
}

func (s *GRPCServer) Run(ctx context.Context, req *proto.RunRequest) (*proto.Empty, error) {
	err := s.Impl.Run(req.GetArgs())
	return &proto.Empty{}, err
}
