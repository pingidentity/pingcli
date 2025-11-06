// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/oauth2"
)

// tokenFileData represents the structure of the credentials file
type tokenFileData struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	Expiry       time.Time `json:"expiry,omitempty"`
}

// getCredentialsFilePath returns the path to the credentials file for a given auth method
func getCredentialsFilePath(authMethod string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	credentialsDir := filepath.Join(homeDir, ".pingcli", "credentials")

	// Create directory if it doesn't exist
	if err := os.MkdirAll(credentialsDir, 0700); err != nil {
		return "", fmt.Errorf("failed to create credentials directory: %w", err)
	}

	// Use auth method as filename
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
		return fmt.Errorf("failed to marshal token data: %w", err)
	}

	// Write to file with restrictive permissions (only owner can read/write)
	if err := os.WriteFile(filePath, jsonData, 0600); err != nil {
		return fmt.Errorf("failed to write token to file: %w", err)
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
	// #nosec G304 -- filePath is constructed from user home dir and auth method
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read credentials file: %w", err)
	}

	// Unmarshal JSON
	var data tokenFileData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal token data: %w", err)
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

// clearTokenFromFile removes the credentials file for a given auth method
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
		return fmt.Errorf("failed to remove credentials file: %w", err)
	}

	return nil
}

// clearAllTokenFilesForGrantType removes all token files for a specific grant type and profile
// This handles cleanup of tokens from old configurations (e.g., when client ID or environment ID changes)
// Pattern: token-*_{grantType}_{profile}.json
func clearAllTokenFilesForGrantType(grantType, profileName string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	credentialsDir := filepath.Join(homeDir, ".pingcli", "credentials")

	// Check if directory exists
	if _, err := os.Stat(credentialsDir); os.IsNotExist(err) {
		// Directory doesn't exist, nothing to clear
		return nil
	}

	// Read all files in credentials directory
	files, err := os.ReadDir(credentialsDir)
	if err != nil {
		return fmt.Errorf("failed to read credentials directory: %w", err)
	}

	// Default profile name if empty
	if profileName == "" {
		profileName = "default"
	}

	var errs []error
	// Look for files matching pattern: token-*_{grantType}_{profile}.json
	// Example: token-a1b2c3d4e5f6g7h8_device_code_production.json
	suffix := fmt.Sprintf("_%s_%s.json", grantType, profileName)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Check if filename matches the pattern for this grant type and profile
		if filepath.Ext(file.Name()) == ".json" && len(file.Name()) > len(suffix) {
			if file.Name()[len(file.Name())-len(suffix):] == suffix {
				filePath := filepath.Join(credentialsDir, file.Name())
				if err := os.Remove(filePath); err != nil {
					errs = append(errs, fmt.Errorf("failed to remove %s: %w", file.Name(), err))
				}
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("failed to clear some token files: %v", errs)
	}

	return nil
}

// clearAllTokenFiles removes all token files from the credentials directory
// This is useful for cleaning up old tokens when configuration changes
func clearAllTokenFiles() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	credentialsDir := filepath.Join(homeDir, ".pingcli", "credentials")

	// Check if directory exists
	if _, err := os.Stat(credentialsDir); os.IsNotExist(err) {
		// Directory doesn't exist, nothing to clear
		return nil
	}

	// Read all files in the credentials directory
	entries, err := os.ReadDir(credentialsDir)
	if err != nil {
		return fmt.Errorf("failed to read credentials directory: %w", err)
	}

	var errs []error
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		// Only remove .json files that look like token files
		filename := entry.Name()
		if filepath.Ext(filename) == ".json" {
			filePath := filepath.Join(credentialsDir, filename)
			if err := os.Remove(filePath); err != nil {
				errs = append(errs, fmt.Errorf("failed to remove %s: %w", filename, err))
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("failed to remove some token files: %v", errs)
	}

	return nil
}
