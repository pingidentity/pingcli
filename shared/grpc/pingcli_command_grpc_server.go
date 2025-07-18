// Copyright © 2025 Ping Identity Corporation

package grpc

import (
	"context"
	"errors"

	"github.com/hashicorp/go-plugin"
	"github.com/pingidentity/pingcli/internal/proto"
)

var _ proto.PingCliCommandServer = (*PingCliCommandGRPCServer)(nil)

type PingCliCommandGRPCServer struct {
	Impl   PingCliCommand
	broker *plugin.GRPCBroker
	proto.UnimplementedPingCliCommandServer
}

func (s *PingCliCommandGRPCServer) Configuration(ctx context.Context, req *proto.Empty) (*proto.PingCliCommandConfigurationResponse, error) {
	cmd, err := s.Impl.Configuration()
	if err != nil {
		return nil, err
	}

	return &proto.PingCliCommandConfigurationResponse{
		Example: &cmd.Example,
		Long:    &cmd.Long,
		Short:   &cmd.Short,
		Use:     &cmd.Use,
	}, nil
}

func (s *PingCliCommandGRPCServer) Run(ctx context.Context, req *proto.PingCliCommandRunRequest) (em *proto.Empty, err error) {
	conn, err := s.broker.Dial(req.GetLogger())
	if err != nil {
		return nil, err
	}
	defer func() {
		cErr := conn.Close()
		if cErr != nil {
			err = errors.Join(err, cErr)
		}
	}()

	loggerClient := &LoggerGRPCClient{
		proto.NewLoggerClient(conn),
	}

	err = s.Impl.Run(req.GetArgs(), loggerClient)

	return &proto.Empty{}, err
}
