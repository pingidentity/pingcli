// Copyright © 2025 Ping Identity Corporation

package grpc

import (
	"context"

	"github.com/pingidentity/pingcli/internal/proto"
)

var _ proto.LoggerServer = (*LoggerGRPCServer)(nil)

type LoggerGRPCServer struct {
	Impl Logger
	proto.UnimplementedLoggerServer
}

func (s *LoggerGRPCServer) Message(ctx context.Context, req *proto.LoggerRequest) (*proto.Empty, error) {
	err := s.Impl.Message(req.GetMessage(), req.GetFields())

	return &proto.Empty{}, err
}

func (s *LoggerGRPCServer) Success(ctx context.Context, req *proto.LoggerRequest) (*proto.Empty, error) {
	err := s.Impl.Success(req.GetMessage(), req.GetFields())

	return &proto.Empty{}, err
}

func (s *LoggerGRPCServer) Warn(ctx context.Context, req *proto.LoggerRequest) (*proto.Empty, error) {
	err := s.Impl.Warn(req.GetMessage(), req.GetFields())

	return &proto.Empty{}, err
}

func (s *LoggerGRPCServer) UserError(ctx context.Context, req *proto.LoggerRequest) (*proto.Empty, error) {
	err := s.Impl.UserError(req.GetMessage(), req.GetFields())

	return &proto.Empty{}, err
}

func (s *LoggerGRPCServer) UserFatal(ctx context.Context, req *proto.LoggerRequest) (*proto.Empty, error) {
	err := s.Impl.UserFatal(req.GetMessage(), req.GetFields())

	return &proto.Empty{}, err
}

func (s *LoggerGRPCServer) PluginError(ctx context.Context, req *proto.LoggerRequest) (*proto.Empty, error) {
	err := s.Impl.PluginError(req.GetMessage(), req.GetFields())

	return &proto.Empty{}, err
}
