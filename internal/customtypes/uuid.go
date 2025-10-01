// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-uuid"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

var (
	uuidErrorPrefix = "custom type uuid error"
	ErrInvalidUUID  = errors.New("invalid uuid")
)

type UUID string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*UUID)(nil)

func (u *UUID) Set(val string) error {
	if u == nil {
		return &errs.PingCLIError{Prefix: uuidErrorPrefix, Err: ErrCustomTypeNil}
	}

	if val == "" {
		*u = UUID(val)

		return nil
	}

	_, err := uuid.ParseUUID(val)
	if err != nil {
		return &errs.PingCLIError{Prefix: uuidErrorPrefix, Err: fmt.Errorf("%w '%s': %w", ErrInvalidUUID, val, err)}
	}

	*u = UUID(val)

	return nil
}

func (u *UUID) Type() string {
	return "string"
}

func (u *UUID) String() string {
	if u == nil {
		return ""
	}

	return string(*u)
}
