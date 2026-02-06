// Copyright © 2026 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

const (
	ENUM_PINGONE_REGION_CODE_AP string = "AP"
	ENUM_PINGONE_REGION_CODE_AU string = "AU"
	ENUM_PINGONE_REGION_CODE_CA string = "CA"
	ENUM_PINGONE_REGION_CODE_EU string = "EU"
	ENUM_PINGONE_REGION_CODE_NA string = "NA"
	ENUM_PINGONE_REGION_CODE_SG string = "SG"

	ENUM_PINGONE_TLD_AP string = "asia"
	ENUM_PINGONE_TLD_AU string = "com.au"
	ENUM_PINGONE_TLD_CA string = "ca"
	ENUM_PINGONE_TLD_EU string = "eu"
	ENUM_PINGONE_TLD_NA string = "com"
	ENUM_PINGONE_TLD_SG string = "sg"
)

var (
	pingOneRegionCodeErrorPrefix = "custom type pingone region code error"
)

type PingOneRegionCode string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*PingOneRegionCode)(nil)

// Implement pflag.Value interface for custom type in cobra pingone-region parameter

func (prc *PingOneRegionCode) Set(regionCode string) error {
	if prc == nil {
		return &errs.PingCLIError{Prefix: pingOneRegionCodeErrorPrefix, Err: ErrCustomTypeNil}
	}
	switch {
	case strings.EqualFold(regionCode, ENUM_PINGONE_REGION_CODE_AP):
		*prc = PingOneRegionCode(ENUM_PINGONE_REGION_CODE_AP)
	case strings.EqualFold(regionCode, ENUM_PINGONE_REGION_CODE_AU):
		*prc = PingOneRegionCode(ENUM_PINGONE_REGION_CODE_AU)
	case strings.EqualFold(regionCode, ENUM_PINGONE_REGION_CODE_CA):
		*prc = PingOneRegionCode(ENUM_PINGONE_REGION_CODE_CA)
	case strings.EqualFold(regionCode, ENUM_PINGONE_REGION_CODE_EU):
		*prc = PingOneRegionCode(ENUM_PINGONE_REGION_CODE_EU)
	case strings.EqualFold(regionCode, ENUM_PINGONE_REGION_CODE_NA):
		*prc = PingOneRegionCode(ENUM_PINGONE_REGION_CODE_NA)
	case strings.EqualFold(regionCode, ""):
		*prc = PingOneRegionCode("")
	default:
		return &errs.PingCLIError{Prefix: pingOneRegionCodeErrorPrefix, Err: fmt.Errorf("%w: '%s'. Must be one of: %s", ErrUnrecognizedPingOneRegionCode, regionCode, strings.Join(PingOneRegionCodeValidValues(), ", "))}
	}

	return nil
}

func (prc *PingOneRegionCode) Type() string {
	return "string"
}

func (prc *PingOneRegionCode) String() string {
	if prc == nil {
		return ""
	}

	return string(*prc)
}

func PingOneRegionCodeValidValues() []string {
	pingoneRegionCodes := []string{
		ENUM_PINGONE_REGION_CODE_AP,
		ENUM_PINGONE_REGION_CODE_AU,
		ENUM_PINGONE_REGION_CODE_CA,
		ENUM_PINGONE_REGION_CODE_EU,
		ENUM_PINGONE_REGION_CODE_NA,
		ENUM_PINGONE_REGION_CODE_SG,
	}

	slices.Sort(pingoneRegionCodes)

	return pingoneRegionCodes
}
