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

	return cfg, nil
}

// parseScopesList takes a space-separated string of scopes and returns a cleaned slice
func parseScopesList(scopesStr string) []string {
	if scopesStr == "" {
		return nil
	}

	var scopesList []string
	for scope := range strings.SplitSeq(scopesStr, " ") {
		if scope != "" {
			scopesList = append(scopesList, scope)
		}
	}

	return scopesList
}
