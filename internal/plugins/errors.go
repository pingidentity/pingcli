// Copyright © 2025 Ping Identity Corporation

package plugins

import "errors"

var (
	ErrGetPluginExecutables = errors.New("failed to get configured plugin executables")
	ErrCreateRPCClient      = errors.New("failed to create plugin rpc client")
	ErrDispensePlugin       = errors.New("the rpc client failed to dispense plugin executable")
	ErrCastPluginInterface  = errors.New("failed to cast plugin executable to grpc.PingCliCommand interface")
	ErrPluginConfiguration  = errors.New("failed to get plugin configuration")
	ErrExecutePlugin        = errors.New("failed to execute plugin command")
)
