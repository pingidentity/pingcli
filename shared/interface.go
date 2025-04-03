// Copyright © 2025 Ping Identity Corporation

package shared

import (
	"context"

	"google.golang.org/grpc"

	"github.com/hashicorp/go-plugin"
	"github.com/pingidentity/pingcli/shared/proto"
)

var (
	HandshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "PINGCLI_COMMAND_PLUGIN_KEY",
		MagicCookieValue: "8b8a9351-bef8-42fc-a642-7cb10f12a49c",
	}

	ENUM_PINGCLI_COMMAND_GRPC string = "pingcli_command_grpc"

	PluginMap = map[string]plugin.Plugin{
		ENUM_PINGCLI_COMMAND_GRPC: &PingCliCommandGrpcPlugin{},
	}
)

type PingCliCommandConfiguration struct {
	Example string
	Long    string
	Short   string
	Use     string
}

type PingCliCommand interface {
	Configuration() (*PingCliCommandConfiguration, error)
	Run([]string) error
}

type PingCliCommandGrpcPlugin struct {
	plugin.Plugin
	Impl PingCliCommand
}

func (p *PingCliCommandGrpcPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterProtoPingCliCommandServer(s, &GRPCServer{
		Impl: p.Impl,
	})
	return nil
}

func (p *PingCliCommandGrpcPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		client: proto.NewProtoPingCliCommandClient(c),
	}, nil
}
