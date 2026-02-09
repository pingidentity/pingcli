// Copyright © 2026 Ping Identity Corporation

package auth_internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pingidentity/pingcli/internal/constants"
	"github.com/pingidentity/pingcli/internal/errs"
	"golang.org/x/oauth2"
)

// tokenFileData represents the structure of the credentials file
type tokenFileData struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	Expiry       time.Time `json:"expiry,omitempty"`
}

// getCredentialsFilePath returns the path to the credentials file for a given grant type
func getCredentialsFilePath(authMethod string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", &errs.PingCLIError{
			Prefix: "failed to get home directory",
			Err:    err,
		}
	}

	credentialsDir := filepath.Join(homeDir, constants.PingCliDirName, constants.CredentialsDirName)

	// Create directory if it doesn't exist
	if err := os.MkdirAll(credentialsDir, 0700); err != nil {
		return "", &errs.PingCLIError{
			Prefix: "failed to create credentials directory",
			Err:    err,
		}
	}

	// Use grant type as filename
	filename := fmt.Sprintf("%s.json", authMethod)

	return filepath.Join(credentialsDir, filename), nil
}

var (
	// ErrNilToken is returned when attempting to save a nil token
	ErrNilToken = fmt.Errorf("token cannot be nil")
	// ErrCredentialsFileNotExist is returned when credentials file doesn't exist
	ErrCredentialsFileNotExist = fmt.Errorf("credentials file does not exist")
)

// saveTokenToFile saves an OAuth2 token to the credentials file
func saveTokenToFile(token *oauth2.Token, authMethod string) error {
	if token == nil {
		return ErrNilToken
	}

	filePath, err := getCredentialsFilePath(authMethod)
	if err != nil {
		return err
	}

	// Convert token to file format
	data := tokenFileData{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return &errs.PingCLIError{
			Prefix: "failed to marshal token data",
			Err:    err,
		}
	}

	// Write to file with restrictive permissions (only owner can read/write)
	if err := os.WriteFile(filePath, jsonData, 0600); err != nil {
		return &errs.PingCLIError{
			Prefix: "failed to write token to file",
			Err:    err,
		}
	}

	return nil
}

// loadTokenFromFile loads an OAuth2 token from the credentials file
func loadTokenFromFile(authMethod string) (*oauth2.Token, error) {
	filePath, err := getCredentialsFilePath(authMethod)
	if err != nil {
		return nil, err
	}

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, ErrCredentialsFileNotExist
	}

	// Read file
	// #nosec G304 -- filePath is constructed from user home dir and grant type
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: "failed to read credentials file",
			Err:    err,
		}
	}

	// Unmarshal JSON
	var data tokenFileData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, &errs.PingCLIError{
			Prefix: "failed to unmarshal token data",
			Err:    err,
		}
	}

	// Convert to oauth2.Token
	token := &oauth2.Token{
		AccessToken:  data.AccessToken,
		TokenType:    data.TokenType,
		RefreshToken: data.RefreshToken,
		Expiry:       data.Expiry,
	}

	return token, nil
}

// clearTokenFromFile removes the credentials file for a given grant type
func clearTokenFromFile(authMethod string) error {
	filePath, err := getCredentialsFilePath(authMethod)
	if err != nil {
		return err
	}

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// File doesn't exist, nothing to clear
		return nil
	}

	// Remove file
	if err := os.Remove(filePath); err != nil {
		return &errs.PingCLIError{
			Prefix: "failed to remove credentials file",
			Err:    err,
		}
	}

	return nil
}

// clearAllTokenFilesForGrantType removes all token files for a specific provider, grant type and profile
// This handles cleanup of tokens from old configurations (e.g., when client ID or environment ID changes)
// Pattern: token-*_{service}_{grantType}_{profile}.json
func clearAllTokenFilesForGrantType(providerName, grantType, profileName string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return &errs.PingCLIError{
			Prefix: "failed to get home directory",
			Err:    err,
		}
	}

	credentialsDir := filepath.Join(homeDir, constants.PingCliDirName, constants.CredentialsDirName)

	// Check if directory exists
	if _, err := os.Stat(credentialsDir); os.IsNotExist(err) {
		// Directory doesn't exist, nothing to clear
		return nil
	}

	// Read all files in credentials directory
	files, err := os.ReadDir(credentialsDir)
	if err != nil {
		return &errs.PingCLIError{
			Prefix: "failed to read credentials directory",
			Err:    err,
		}
	}

	// Default values if empty
	if providerName == "" {
		providerName = "pingone"
	}
	if profileName == "" {
		profileName = "default"
	}

	var errList []error
	// Look for files matching pattern: token-*_{service}_{grantType}_{profile}.json
	// Example: token-a1b2c3d4e5f6g7h8_pingone_device_code_production.json
	suffix := fmt.Sprintf("_%s_%s_%s.json", providerName, grantType, profileName)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Check if filename matches the pattern for this provider, grant type and profile
		if filepath.Ext(file.Name()) == ".json" && len(file.Name()) > len(suffix) {
			if file.Name()[len(file.Name())-len(suffix):] == suffix {
				filePath := filepath.Join(credentialsDir, file.Name())
				if err := os.Remove(filePath); err != nil {
					errList = append(errList, &errs.PingCLIError{
						Prefix: fmt.Sprintf("failed to remove %s", file.Name()),
						Err:    err,
					})
				}
			}
		}
	}

	if len(errList) > 0 {
		return &errs.PingCLIError{
			Prefix: "failed to clear some token files",
			Err:    errors.Join(errList...),
		}
	}

	return nil
}

// clearAllCredentialFiles removes all internal credential files
// This is used for a full logout/cleanup options
func clearAllCredentialFiles() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return &errs.PingCLIError{
			Prefix: "failed to get home directory",
			Err:    err,
		}
	}

	credentialsDir := filepath.Join(homeDir, constants.PingCliDirName, constants.CredentialsDirName)

	// Check if directory exists
	if _, err := os.Stat(credentialsDir); os.IsNotExist(err) {
		return nil
	}

	// Read all files in credentials directory
	files, err := os.ReadDir(credentialsDir)
	if err != nil {
		return &errs.PingCLIError{
			Prefix: "failed to read credentials directory",
			Err:    err,
		}
	}

	var errList []error
	for _, file := range files {
		// Only remove files, leave directories if any (though typically there aren't any)
		if !file.IsDir() {
			filePath := filepath.Join(credentialsDir, file.Name())
			if err := os.Remove(filePath); err != nil {
				errList = append(errList, &errs.PingCLIError{
					Prefix: fmt.Sprintf("failed to remove %s", file.Name()),
					Err:    err,
				})
			}
		}
	}

	if len(errList) > 0 {
		return &errs.PingCLIError{
			Prefix: "failed to clear some token files",
			Err:    errors.Join(errList...),
		}
	}

	return nil
}
