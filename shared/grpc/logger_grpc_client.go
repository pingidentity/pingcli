// Copyright © 2026 Ping Identity Corporation

package grpc

import (
	"context"

	"github.com/pingidentity/pingcli/internal/proto"
)

var _ Logger = (*LoggerGRPCClient)(nil)

type LoggerGRPCClient struct {
	client proto.LoggerClient
}

func (c LoggerGRPCClient) Message(message string, fields map[string]string) error {
	_, err := c.client.Message(context.Background(), &proto.LoggerRequest{
		Message: &message,
		Fields:  fields,
	})

	return err
}

func (c LoggerGRPCClient) Success(message string, fields map[string]string) error {
	_, err := c.client.Success(context.Background(), &proto.LoggerRequest{
		Message: &message,
		Fields:  fields,
	})

	return err
}

func (c LoggerGRPCClient) Warn(message string, fields map[string]string) error {
	_, err := c.client.Warn(context.Background(), &proto.LoggerRequest{
		Message: &message,
		Fields:  fields,
	})

	return err
}

func (c LoggerGRPCClient) UserError(message string, fields map[string]string) error {
	_, err := c.client.UserError(context.Background(), &proto.LoggerRequest{
		Message: &message,
		Fields:  fields,
	})

	return err
}

func (c LoggerGRPCClient) UserFatal(message string, fields map[string]string) error {
	_, err := c.client.UserFatal(context.Background(), &proto.LoggerRequest{
		Message: &message,
		Fields:  fields,
	})

	return err
}

func (c LoggerGRPCClient) PluginError(message string, fields map[string]string) error {
	_, err := c.client.PluginError(context.Background(), &proto.LoggerRequest{
		Message: &message,
		Fields:  fields,
	})

	return err
}
