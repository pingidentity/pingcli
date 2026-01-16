// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingone-go-client/config"
)

// applyRegionConfiguration applies the PingOne region configuration to a config.Configuration
func applyRegionConfiguration(cfg *config.Configuration) (*config.Configuration, error) {
	regionCode, err := profiles.GetOptionValue(options.PingOneRegionCodeOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get region code: %w", err)
	}

	switch regionCode {
	case customtypes.ENUM_PINGONE_REGION_CODE_AP:
		cfg = cfg.WithTopLevelDomain(config.TopLevelDomainAPAC)
	case customtypes.ENUM_PINGONE_REGION_CODE_AU:
		cfg = cfg.WithTopLevelDomain(config.TopLevelDomainAU)
	case customtypes.ENUM_PINGONE_REGION_CODE_CA:
		cfg = cfg.WithTopLevelDomain(config.TopLevelDomainCA)
	case customtypes.ENUM_PINGONE_REGION_CODE_EU:
		cfg = cfg.WithTopLevelDomain(config.TopLevelDomainEU)
	case customtypes.ENUM_PINGONE_REGION_CODE_NA:
		cfg = cfg.WithTopLevelDomain(config.TopLevelDomainNA)
	case customtypes.ENUM_PINGONE_REGION_CODE_SG:
		cfg = cfg.WithTopLevelDomain(config.TopLevelDomainSG)
	default:
		return nil, &errs.PingCLIError{
			Prefix: fmt.Sprintf("invalid region code '%s'", regionCode),
			Err:    ErrRegionCodeRequired,
		}
	}

	// Get and set the environment ID for API endpoints
	// Prefer the environment ID already present on cfg; fallback to profile values.
	var endpointsEnvironmentID string
	if cfg.Endpoint.EnvironmentID == nil || strings.TrimSpace(*cfg.Endpoint.EnvironmentID) == "" {
		// Primary: general environment ID
		endpointsEnvironmentID, err = profiles.GetOptionValue(options.PingOneAuthenticationAPIEnvironmentIDOption)
		if err != nil {
			return nil, fmt.Errorf("failed to get endpoints environment ID: %w", err)
		}
		if strings.TrimSpace(endpointsEnvironmentID) == "" {
			return nil, &errs.PingCLIError{
				Prefix: "endpoints environment ID is not configured",
				Err:    ErrEnvironmentIDNotConfigured,
			}
		}
		cfg = cfg.WithEnvironmentID(endpointsEnvironmentID)
	}

	return cfg, nil
}

// formatStorageLocation returns a human-friendly message for where credentials were cleared
// based on StorageLocation.
func formatStorageLocation(location StorageLocation) string {
	var locs []string
	if location.Keychain {
		locs = append(locs, "keychain")
	}
	if location.File {
		locs = append(locs, "file storage")
	}

	if len(locs) == 0 {
		return "storage"
	}

	return strings.Join(locs, " and ")
}
