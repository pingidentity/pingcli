// Copyright © 2026 Ping Identity Corporation

package customtypes

import (
	"strings"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingone-go-client/config"
	"github.com/spf13/pflag"
)

// StorageType is a pflag-compatible wrapper for SDK config.StorageType
type StorageType string

// Verify implements pflag.Value
var _ pflag.Value = (*StorageType)(nil)

const (
	// Values mirror SDK storage types (lowercase)
	ENUM_STORAGE_TYPE_FILE_SYSTEM   string = "file_system"
	ENUM_STORAGE_TYPE_SECURE_LOCAL  string = "secure_local"
	ENUM_STORAGE_TYPE_SECURE_REMOTE string = "secure_remote"
	ENUM_STORAGE_TYPE_NONE          string = "none"
)

var (
	storageTypeErrorPrefix = "custom type storage type error"
)

func (st *StorageType) Set(v string) error {
	if st == nil {
		return &errs.PingCLIError{Prefix: storageTypeErrorPrefix, Err: ErrCustomTypeNil}
	}

	s := strings.TrimSpace(strings.ToLower(v))

	switch s {
	case string(config.StorageTypeFileSystem):
		*st = StorageType(ENUM_STORAGE_TYPE_FILE_SYSTEM)
	case string(config.StorageTypeSecureLocal):
		*st = StorageType(ENUM_STORAGE_TYPE_SECURE_LOCAL)
	case string(config.StorageTypeSecureRemote):
		*st = StorageType(ENUM_STORAGE_TYPE_SECURE_REMOTE)
	case string(config.StorageTypeNone):
		*st = StorageType(ENUM_STORAGE_TYPE_NONE)
	case "":
		// Treat empty as default (secure_local)
		*st = StorageType(ENUM_STORAGE_TYPE_SECURE_LOCAL)
	default:
		return &errs.PingCLIError{Prefix: storageTypeErrorPrefix, Err: ErrUnrecognizedStorageType}
	}

	return nil
}

func (st *StorageType) Type() string {
	return "string"
}

func (st *StorageType) String() string {
	if st == nil {
		return ""
	}

	return string(*st)
}

func StorageTypeValidValues() []string {
	return []string{
		ENUM_STORAGE_TYPE_FILE_SYSTEM,
		ENUM_STORAGE_TYPE_SECURE_LOCAL,
		ENUM_STORAGE_TYPE_SECURE_REMOTE,
		ENUM_STORAGE_TYPE_NONE,
	}
}
