// Copyright © 2025 Ping Identity Corporation

package profiles

/* The main viper instance should ONLY interact with the configuration file
on disk. No viper overrides, environment variable bindings, or pflag
bindings should be used with this viper instance. This keeps the config
file as the ONLY source of truth for the main viper instance, and prevents
profile drift, as well as active profile drift and other niche bugs. As a
result, much of the logic in this file avoids the use of mainViper.Set(), and
goes out of the way to modify the config file.*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type MainConfig struct {
	viperInstance *viper.Viper
}

var (
	mainViper *MainConfig = NewMainConfig()
)

// Returns a new MainViper instance
func NewMainConfig() (newMainViper *MainConfig) {
	newMainViper = &MainConfig{
		viperInstance: viper.New(),
	}

	return newMainViper
}

// Returns the MainViper struct
func GetMainConfig() *MainConfig {
	return mainViper
}

func (m MainConfig) ViperInstance() *viper.Viper {
	return m.viperInstance
}

func (m *MainConfig) ChangeActiveProfile(pName string) (err error) {
	if err = m.ValidateExistingProfileName(pName); err != nil {
		return err
	}

	tempViper := viper.New()
	tempViper.SetConfigFile(m.ViperInstance().ConfigFileUsed())
	if err := tempViper.ReadInConfig(); err != nil {
		return err
	}

	tempViper.Set(options.RootActiveProfileOption.ViperKey, pName)

	if err = tempViper.WriteConfig(); err != nil {
		return err
	}

	if err = m.ViperInstance().ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (m MainConfig) ChangeProfileName(oldPName, newPName string) (err error) {
	if oldPName == newPName {
		return nil
	}

	err = m.ValidateExistingProfileName(oldPName)
	if err != nil {
		return err
	}

	err = m.ValidateNewProfileName(newPName)
	if err != nil {
		return err
	}

	subViper, err := m.GetProfileViper(oldPName)
	if err != nil {
		return err
	}

	if err = m.DeleteProfile(oldPName); err != nil {
		return err
	}

	if err = m.SaveProfile(newPName, subViper); err != nil {
		return err
	}

	return nil
}

func (m MainConfig) ChangeProfileDescription(pName, description string) (err error) {
	if err = m.ValidateExistingProfileName(pName); err != nil {
		return err
	}

	subViper, err := m.GetProfileViper(pName)
	if err != nil {
		return err
	}

	subViper.Set(options.ProfileDescriptionOption.ViperKey, description)

	if err = m.SaveProfile(pName, subViper); err != nil {
		return err
	}

	return nil
}

func (m MainConfig) GetProfileViper(pName string) (subViper *viper.Viper, err error) {
	if err = m.ValidateExistingProfileName(pName); err != nil {
		return nil, err
	}

	subViper = m.ViperInstance().Sub(pName)
	if subViper == nil {
		return nil, fmt.Errorf("failed to get profile viper: profile '%s' does not exist", pName)
	}

	return subViper, nil
}

// Viper gives no built-in delete or unset method for keys
// Using this "workaround" described here: https://github.com/spf13/viper/issues/632
func (m MainConfig) DeleteProfile(pName string) (err error) {
	if err = m.ValidateExistingProfileName(pName); err != nil {
		return err
	}

	activeProfileName, err := GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		return err
	}

	if strings.EqualFold(activeProfileName, pName) {
		return fmt.Errorf("'%s' is the active profile and cannot be deleted", pName)
	}

	mainViperConfigMap := m.ViperInstance().AllSettings()
	delete(mainViperConfigMap, pName)

	encodedConfig, err := json.MarshalIndent(mainViperConfigMap, "", " ")
	if err != nil {
		return err
	}

	err = m.ViperInstance().ReadConfig(bytes.NewReader(encodedConfig))
	if err != nil {
		return err
	}

	err = m.ViperInstance().WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

// Get all profile names from config.yaml configuration file
// Returns a sorted slice of profile names
func (m MainConfig) ProfileNames() (profileNames []string) {
	keySet := make(map[string]struct{})
	mainViperKeys := m.ViperInstance().AllKeys()
	for _, key := range mainViperKeys {
		//Do not add Active profile viper key to profileNames
		if strings.EqualFold(key, options.RootActiveProfileOption.ViperKey) {
			continue
		}

		pName := strings.Split(key, ".")[0]
		if _, ok := keySet[pName]; !ok {
			keySet[pName] = struct{}{}
			profileNames = append(profileNames, pName)
		}
	}

	slices.Sort(profileNames)

	return profileNames
}

func (m MainConfig) SaveProfile(pName string, subViper *viper.Viper) (err error) {
	mainViperConfigMap := m.ViperInstance().AllSettings()
	subViperConfigMap := subViper.AllSettings()

	mainViperConfigMap[pName] = subViperConfigMap

	encodedConfig, err := json.MarshalIndent(mainViperConfigMap, "", " ")
	if err != nil {
		return err
	}

	err = m.ViperInstance().ReadConfig(bytes.NewReader(encodedConfig))
	if err != nil {
		return err
	}

	err = m.ViperInstance().WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

// The profile name must exist
func (m MainConfig) ValidateExistingProfileName(pName string) (err error) {
	if pName == "" {
		return fmt.Errorf("invalid profile name: profile name cannot be empty")
	}

	pNames := m.ProfileNames()

	if !slices.ContainsFunc(pNames, func(n string) bool {
		return strings.EqualFold(n, pName)
	}) {
		return fmt.Errorf("invalid profile name: '%s' profile does not exist", pName)
	}

	return nil
}

// The profile name format must be valid
// The new profile name must be unique
func (m MainConfig) ValidateNewProfileName(pName string) (err error) {
	if err = m.ValidateProfileNameFormat(pName); err != nil {
		return err
	}

	pNames := m.ProfileNames()
	if slices.ContainsFunc(pNames, func(n string) bool {
		return strings.EqualFold(n, pName)
	}) {
		return fmt.Errorf("invalid profile name: '%s'. profile already exists", pName)
	}

	return nil
}

// The profile name must contain only alphanumeric characters, underscores, and dashes
// The profile name cannot be empty
func (m MainConfig) ValidateProfileNameFormat(pName string) (err error) {
	if pName == "" {
		return fmt.Errorf("invalid profile name: profile name cannot be empty")
	}

	re := regexp.MustCompile(`^[a-z0-9\_\-]+$`)
	if !re.MatchString(pName) {
		return fmt.Errorf("invalid profile name: '%s'. name must be lowercase and contain only alphanumeric characters, underscores, and dashes", pName)
	}

	return nil
}

// If the new profile name is the same as the existing profile name, that is valid
// Otherwise treat newPName as a new profile name and validate it
func (m MainConfig) ValidateUpdateExistingProfileName(ePName, newPName string) (err error) {
	if ePName == newPName {
		return nil
	}

	if err = m.ValidateNewProfileName(newPName); err != nil {
		return err
	}

	return nil
}

func (m MainConfig) ProfileToString(pName string) (yamlStr string, err error) {
	if err = m.ValidateExistingProfileName(pName); err != nil {
		return "", err
	}

	subViper, err := m.GetProfileViper(pName)
	if err != nil {
		return "", err
	}

	yaml, err := yaml.Marshal(subViper.AllSettings())
	if err != nil {
		return "", fmt.Errorf("failed to yaml marshal active profile: %v", err)
	}

	return string(yaml), nil
}

func (m MainConfig) ProfileViperValue(pName, viperKey string) (yamlStr string, err error) {
	if err = m.ValidateExistingProfileName(pName); err != nil {
		return "", err
	}

	subViper, err := m.GetProfileViper(pName)
	if err != nil {
		return "", err
	}

	if !subViper.IsSet(viperKey) {
		return "", fmt.Errorf("configuration key '%s' is not set in profile '%s'", viperKey, pName)
	}

	yaml, err := yaml.Marshal(subViper.Get(viperKey))
	if err != nil {
		return "", fmt.Errorf("failed to yaml marshal configuration value from key '%s': %v", viperKey, err)
	}

	return string(yaml), nil
}

func (m MainConfig) DefaultMissingViperKeys() (err error) {
	// For each profile, if a viper key from an option doesn't exist, set it to the default value
	for _, pName := range m.ProfileNames() {
		subViper, err := m.GetProfileViper(pName)
		if err != nil {
			return err
		}

		for _, opt := range options.Options() {
			if opt.ViperKey == "" || opt.ViperKey == options.RootActiveProfileOption.ViperKey {
				continue
			}
			if !subViper.IsSet(opt.ViperKey) {
				subViper.Set(opt.ViperKey, opt.DefaultValue)
			}
		}
		err = m.SaveProfile(pName, subViper)
		if err != nil {
			return fmt.Errorf("Failed to save profile '%s': %v", pName, err)
		}
	}

	return nil
}

func GetOptionValue(opt options.Option) (pFlagValue string, err error) {
	// 1st priority: cobra param flag value
	cobraParamValue, ok := cobraParamValueFromOption(opt)
	if ok {
		return cobraParamValue, nil
	}

	// 2nd priority: environment variable value
	pFlagValue = os.Getenv(opt.EnvVar)
	if pFlagValue != "" {
		return pFlagValue, nil
	}

	// 3rd priority: viper value
	viperValue, ok, err := ViperValueFromOption(opt)
	if err != nil {
		return "", err
	}
	if ok {
		return viperValue, nil
	}

	// 4th priority: default value
	if opt.DefaultValue != nil {
		pFlagValue = opt.DefaultValue.String()
		return pFlagValue, nil
	}

	// This is a error, as it means the option is not configured internally to contain one of the 4 values above.
	// This should never happen, as all options should at least have a default value.
	return "", fmt.Errorf("failed to get option value: no value found: %v", opt)
}

func MaskValue(value string) string {
	if value == "" {
		return ""
	}

	// Mask all values to the same asterisk length
	// providing no additional information about the value when logged.
	return strings.Repeat("*", 8)
}

func cobraParamValueFromOption(opt options.Option) (value string, ok bool) {
	if opt.CobraParamValue != nil && opt.Flag.Changed {
		return opt.CobraParamValue.String(), true
	}

	return "", false
}

func ViperValueFromOption(opt options.Option) (value string, ok bool, err error) {
	mainConfig := GetMainConfig()
	if opt.ViperKey != "" && mainConfig != nil {
		var (
			vValue            any
			mainViperInstance = mainConfig.ViperInstance()
		)

		// Case 1: Viper Key is the ActiveProfile Key, get value from main viper instance
		if opt.ViperKey == options.RootActiveProfileOption.ViperKey && mainViperInstance != nil {
			vValue = mainViperInstance.Get(opt.ViperKey)
		} else {
			// Case 2: --profile flag has been set, get value from set profile viper instance
			// Case 3: no --profile flag set, get value from active profile viper instance defined in main viper instance
			// This recursive call is safe, as options.RootProfileOption.ViperKey is not set
			pName, err := GetOptionValue(options.RootProfileOption)
			if err != nil {
				return "", false, err
			}
			if pName == "" {
				pName, err = GetOptionValue(options.RootActiveProfileOption)
				if err != nil {
					return "", false, err
				}
			}

			subViper, err := mainConfig.GetProfileViper(pName)
			if err != nil {
				return "", false, err
			}

			vValue = subViper.Get(opt.ViperKey)
		}

		switch typedValue := vValue.(type) {
		case nil:
			return "", false, nil
		case string:
			return typedValue, true, nil
		case []string:
			return strings.Join(typedValue, ","), true, nil
		case []any:
			strSlice := []string{}
			for _, v := range typedValue {
				strSlice = append(strSlice, fmt.Sprintf("%v", v))
			}
			return strings.Join(strSlice, ","), true, nil
		default:
			return fmt.Sprintf("%v", typedValue), true, nil
		}
	}

	return "", false, nil
}
