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

func (c *PingCliCommandGRPCClient) Run(args []string, l Logger) error {
	loggerServer := &LoggerGRPCServer{
		Impl: l,
	}

	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		proto.RegisterLoggerServer(s, loggerServer)

		return s
	}

	brokerId := c.broker.NextId()
	go c.broker.AcceptAndServe(brokerId, serverFunc)

	_, err := c.client.Run(context.Background(), &proto.PingCliCommandRunRequest{
		Args:   args,
		Logger: &brokerId,
	})

	return err
}
