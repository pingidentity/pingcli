// Copyright © 2026 Ping Identity Corporation

package grpc

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/pingidentity/pingcli/internal/proto"
	"google.golang.org/grpc"
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

type LogType int32

const (
	LOG_TYPE_UNSPECIFIED LogType = 0
	LOG_TYPE_DEBUG       LogType = 1
)

type Logger interface {
	Message(message string, fields map[string]string) error
	Success(message string, fields map[string]string) error
	Warn(message string, fields map[string]string) error
	UserError(message string, fields map[string]string) error
	UserFatal(message string, fields map[string]string) error
	PluginError(message string, fields map[string]string) error
}

type PingCliCommand interface {
	Configuration() (*PingCliCommandConfiguration, error)
	Run(args []string, l Logger) error
}

type PingCliCommandGrpcPlugin struct {
	plugin.Plugin
	Impl PingCliCommand
}

func (p *PingCliCommandGrpcPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterPingCliCommandServer(s, &PingCliCommandGRPCServer{
		broker: broker,
		Impl:   p.Impl,
	})

	return nil
}

func (p *PingCliCommandGrpcPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (any, error) {
	return &PingCliCommandGRPCClient{
		broker: broker,
		client: proto.NewPingCliCommandClient(c),
	}, nil
}
