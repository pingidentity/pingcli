// Copyright © 2025 Ping Identity Corporation

package shared

import (
	"context"

	"github.com/pingidentity/pingcli/shared/proto"
)

var _ PingCliCommand = (*GRPCClient)(nil)

type GRPCClient struct {
	client proto.ProtoPingCliCommandClient
}

func (c *GRPCClient) Configuration() (*PingCliCommandConfiguration, error) {
	resp, err := c.client.Configuration(context.Background(), &proto.Empty{})
	if err != nil {
		return nil, err
	}

	return &PingCliCommandConfiguration{
		Example: resp.Example,
		Long:    resp.Long,
		Short:   resp.Short,
		Use:     resp.Use,
	}, nil

}

func (c *GRPCClient) Run(args []string) error {
	_, err := c.client.Run(context.Background(), &proto.RunRequest{
		Args: args,
	})

	return err
}
