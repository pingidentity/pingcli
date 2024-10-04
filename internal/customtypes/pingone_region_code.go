package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

const (
	ENUM_PINGONE_REGION_CODE_AP string = "AP"
	ENUM_PINGONE_REGION_CODE_AU string = "AU"
	ENUM_PINGONE_REGION_CODE_CA string = "CA"
	ENUM_PINGONE_REGION_CODE_EU string = "EU"
	ENUM_PINGONE_REGION_CODE_NA string = "NA"

	ENUM_PINGONE_TLD_AP string = "asia"
	ENUM_PINGONE_TLD_AU string = "com.au"
	ENUM_PINGONE_TLD_CA string = "ca"
	ENUM_PINGONE_TLD_EU string = "eu"
	ENUM_PINGONE_TLD_NA string = "com"
)

type PingoneRegionCode string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*PingoneRegionCode)(nil)

// Implement pflag.Value interface for custom type in cobra pingone-region parameter

func (prc *PingoneRegionCode) Set(regionCode string) error {
	if prc == nil {
		return fmt.Errorf("failed to set Pingone Region Code value: %s. Pingone Region Code is nil", regionCode)
	}
	switch {
	case strings.EqualFold(regionCode, ENUM_PINGONE_REGION_CODE_AP):
		*prc = PingoneRegionCode(ENUM_PINGONE_REGION_CODE_AP)
	case strings.EqualFold(regionCode, ENUM_PINGONE_REGION_CODE_AU):
		*prc = PingoneRegionCode(ENUM_PINGONE_REGION_CODE_AU)
	case strings.EqualFold(regionCode, ENUM_PINGONE_REGION_CODE_CA):
		*prc = PingoneRegionCode(ENUM_PINGONE_REGION_CODE_CA)
	case strings.EqualFold(regionCode, ENUM_PINGONE_REGION_CODE_EU):
		*prc = PingoneRegionCode(ENUM_PINGONE_REGION_CODE_EU)
	case strings.EqualFold(regionCode, ENUM_PINGONE_REGION_CODE_NA):
		*prc = PingoneRegionCode(ENUM_PINGONE_REGION_CODE_NA)
	case strings.EqualFold(regionCode, ""):
		*prc = PingoneRegionCode("")
	default:
		return fmt.Errorf("unrecognized Pingone Region Code: '%s'. Must be one of: %s", regionCode, strings.Join(PingoneRegionCodeValidValues(), ", "))
	}
	return nil
}

func (prc PingoneRegionCode) Type() string {
	return "string"
}

func (prc PingoneRegionCode) String() string {
	return string(prc)
}

func PingoneRegionCodeValidValues() []string {
	pingoneRegionCodes := []string{
		ENUM_PINGONE_REGION_CODE_AP,
		ENUM_PINGONE_REGION_CODE_AU,
		ENUM_PINGONE_REGION_CODE_CA,
		ENUM_PINGONE_REGION_CODE_EU,
		ENUM_PINGONE_REGION_CODE_NA,
	}

	slices.Sort(pingoneRegionCodes)

	return pingoneRegionCodes
}
