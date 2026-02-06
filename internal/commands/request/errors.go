// Copyright © 2026 Ping Identity Corporation

package request_internal

import "errors"

var (
	ErrServiceEmpty                  = errors.New("service is not set")
	ErrUnrecognizedService           = errors.New("unrecognized service")
	ErrHttpMethodEmpty               = errors.New("http method is not set")
	ErrUnrecognizedHttpMethod        = errors.New("unrecognized http method")
	ErrPingOneRegionCodeEmpty        = errors.New("PingOne region code is not set")
	ErrUnrecognizedPingOneRegionCode = errors.New("unrecognized PingOne region code")
	ErrPingOneWorkerEnvIDEmpty       = errors.New("PingOne worker environment ID is not set")
	ErrPingOneClientIDAndSecretEmpty = errors.New("PingOne client ID and/or client secret is not set")
	ErrPingOneAuthenticate           = errors.New("failed to authenticate with PingOne")
)
