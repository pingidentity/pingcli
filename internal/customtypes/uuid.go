package customtypes

import (
	"fmt"

	"github.com/hashicorp/go-uuid"
	"github.com/spf13/pflag"
)

type UUID string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*UUID)(nil)

func (u *UUID) Set(val string) error {
	if u == nil {
		return fmt.Errorf("failed to set UUID value: %s. UUID is nil", val)
	}

	if val == "" {
		*u = UUID(val)
		return nil
	}

	_, err := uuid.ParseUUID(val)
	if err != nil {
		return err
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
