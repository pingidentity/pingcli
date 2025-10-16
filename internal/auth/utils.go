// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingone-go-client/config"
	pingoneoauth2 "github.com/pingidentity/pingone-go-client/oauth2"
)

// applyRegionConfigurationToConfigConfiguration applies the PingOne region configuration to a config.Configuration
func applyRegionConfigurationToConfigConfiguration(configConfiguration *config.Configuration) (*config.Configuration, error) {
	regionCode, err := profiles.GetOptionValue(options.PingOneRegionCodeOption)
	if err != nil {
		return nil, err
	}

	switch regionCode {
	case customtypes.ENUM_PINGONE_REGION_CODE_AP:
		configConfiguration.WithTopLevelDomain(pingoneoauth2.TopLevelDomainAPAC)
	case customtypes.ENUM_PINGONE_REGION_CODE_AU:
		configConfiguration.WithTopLevelDomain(pingoneoauth2.TopLevelDomainAU)
	case customtypes.ENUM_PINGONE_REGION_CODE_CA:
		configConfiguration.WithTopLevelDomain(pingoneoauth2.TopLevelDomainCA)
	case customtypes.ENUM_PINGONE_REGION_CODE_EU:
		configConfiguration.WithTopLevelDomain(pingoneoauth2.TopLevelDomainEU)
	case customtypes.ENUM_PINGONE_REGION_CODE_NA:
		configConfiguration.WithTopLevelDomain(pingoneoauth2.TopLevelDomainNA)
	case customtypes.ENUM_PINGONE_REGION_CODE_SG:
		configConfiguration.WithTopLevelDomain(pingoneoauth2.TopLevelDomainSG)
	default:
		return nil, fmt.Errorf("PingOne region code is required and must be valid")
	}

	return configConfiguration, nil
}

// parseScopesList takes a comma-separated string of scopes and returns a cleaned slice
func parseScopesList(scopesStr string) []string {
	if scopesStr == "" {
		return nil
	}

	var scopesList []string
	for _, scope := range strings.Split(scopesStr, ",") {
		trimmedScope := strings.TrimSpace(scope)
		if trimmedScope != "" {
			scopesList = append(scopesList, trimmedScope)
		}
	}

	return scopesList
}

// applyRegionConfiguration applies the PingOne region configuration to a credentials config.Configuration
func applyRegionConfiguration(cfg *config.Configuration) (*config.Configuration, error) {
	regionCode, err := profiles.GetOptionValue(options.PingOneRegionCodeOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get region code: %w", err)
	}

	switch regionCode {
	case customtypes.ENUM_PINGONE_REGION_CODE_AP:
		cfg = cfg.WithTopLevelDomain(pingoneoauth2.TopLevelDomainAPAC)
	case customtypes.ENUM_PINGONE_REGION_CODE_AU:
		cfg = cfg.WithTopLevelDomain(pingoneoauth2.TopLevelDomainAU)
	case customtypes.ENUM_PINGONE_REGION_CODE_CA:
		cfg = cfg.WithTopLevelDomain(pingoneoauth2.TopLevelDomainCA)
	case customtypes.ENUM_PINGONE_REGION_CODE_EU:
		cfg = cfg.WithTopLevelDomain(pingoneoauth2.TopLevelDomainEU)
	case customtypes.ENUM_PINGONE_REGION_CODE_NA:
		cfg = cfg.WithTopLevelDomain(pingoneoauth2.TopLevelDomainNA)
	case customtypes.ENUM_PINGONE_REGION_CODE_SG:
		cfg = cfg.WithTopLevelDomain(pingoneoauth2.TopLevelDomainSG)
	default:
		return nil, fmt.Errorf("region code is required and must be valid. Please run 'pingcli config set service.pingone.regionCode=<region>'")
	}

	return cfg, nil
}
