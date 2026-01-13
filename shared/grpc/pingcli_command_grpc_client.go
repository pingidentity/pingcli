// Copyright © 2025 Ping Identity Corporation

package grpc

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/pingidentity/pingcli/internal/proto"
	"google.golang.org/grpc"
)

var _ PingCliCommand = (*PingCliCommandGRPCClient)(nil)

type PingCliCommandGRPCClient struct {
	broker *plugin.GRPCBroker
	client proto.PingCliCommandClient
}

func (c *PingCliCommandGRPCClient) Configuration() (*PingCliCommandConfiguration, error) {
	resp, err := c.client.Configuration(context.Background(), &proto.Empty{})
	if err != nil {
		return nil, err
	}

	return &PingCliCommandConfiguration{
		Example: resp.GetExample(),
		Long:    resp.GetLong(),
		Short:   resp.GetShort(),
		Use:     resp.GetUse(),
	}, nil
}

func (m *PingCliCommandGRPCClient) Run(args []string, l Logger, a Authentication) error {
	loggerServer := &LoggerGRPCServer{Impl: l}
	authenticationServer := &AuthenticationGRPCServer{Impl: a}

	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		proto.RegisterLoggerServer(s, loggerServer)
		proto.RegisterAuthenticationServer(s, authenticationServer)
		return s
	}

	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)

	authenticationBrokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(authenticationBrokerID, serverFunc)

	_, err := m.client.Run(context.Background(), &proto.PingCliCommandRunRequest{
		Args:           args,
		Logger:         &brokerID,
		Authentication: &authenticationBrokerID,
	})

	s.Stop()
	return err
}
