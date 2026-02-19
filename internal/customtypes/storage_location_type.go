// Copyright © 2026 Ping Identity Corporation

package customtypes

// StorageLocationType defines the type of storage where credentials are saved
type StorageLocationType string

const (
	StorageLocationKeychain StorageLocationType = "keychain"
	StorageLocationFile     StorageLocationType = "file"
)
