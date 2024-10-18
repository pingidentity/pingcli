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

type ActiveProfile struct {
	name          string
	viperInstance *viper.Viper
}

type MainConfig struct {
	viperInstance *viper.Viper
	activeProfile *ActiveProfile
}

var (
	mainViper *MainConfig = NewMainConfig()
)

// Returns a new MainViper instance
func NewMainConfig() (newMainViper *MainConfig) {
	newMainViper = &MainConfig{
		viperInstance: viper.New(),
		activeProfile: nil,
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

func (m MainConfig) ActiveProfile() *ActiveProfile {
	return m.activeProfile
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

	m.activeProfile = &ActiveProfile{
		name:          pName,
		viperInstance: m.ViperInstance().Sub(pName),
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

	subViper := m.ViperInstance().Sub(oldPName)

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

	subViper := m.ViperInstance().Sub(pName)
	subViper.Set(options.ProfileDescriptionOption.ViperKey, description)

	if err = m.SaveProfile(pName, subViper); err != nil {
		return err
	}

	return nil
}

// Viper gives no built-in delete or unset method for keys
// Using this "workaround" described here: https://github.com/spf13/viper/issues/632
func (m MainConfig) DeleteProfile(pName string) (err error) {
	if err = m.ValidateExistingProfileName(pName); err != nil {
		return err
	}

	if pName == m.ActiveProfile().Name() {
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

// The profile name format must be valid
// The profile name must exist
func (m MainConfig) ValidateExistingProfileName(pName string) (err error) {
	if err := m.ValidateProfileNameFormat(pName); err != nil {
		return err
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

	re := regexp.MustCompile(`^[a-zA-Z0-9\_\-]+$`)
	if !re.MatchString(pName) {
		return fmt.Errorf("invalid profile name: '%s'. name must contain only alphanumeric characters, underscores, and dashes", pName)
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

	subViper := m.ViperInstance().Sub(pName)

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

	subViper := m.ViperInstance().Sub(pName)

	if !subViper.IsSet(viperKey) {
		return "", fmt.Errorf("configuration key '%s' is not set in profile '%s'", viperKey, pName)
	}

	yaml, err := yaml.Marshal(subViper.Get(viperKey))
	if err != nil {
		return "", fmt.Errorf("failed to yaml marshal configuration value from key '%s': %v", viperKey, err)
	}

	return string(yaml), nil
}

func (a ActiveProfile) ViperInstance() *viper.Viper {
	return a.viperInstance
}

func (a ActiveProfile) Name() string {
	return a.name
}

func GetOptionValue(opt options.Option) (pFlagValue string, err error) {
	if opt.CobraParamValue != nil && opt.Flag.Changed {
		pFlagValue = opt.CobraParamValue.String()
		return pFlagValue, nil
	}

	pFlagValue = os.Getenv(opt.EnvVar)
	if pFlagValue != "" {
		return pFlagValue, nil
	}

	mainConfig := GetMainConfig()
	if opt.ViperKey != "" && mainConfig != nil {
		var vValue any

		mainViperInstance := mainConfig.ViperInstance()
		// This recursive call is safe, as options.RootProfileOption.ViperKey is not set
		definedProfileName, err := GetOptionValue(options.RootProfileOption)
		if err != nil {
			return "", err
		}

		// 3 Cases:
		// - 1) Viper Key is the ActiveProfile Key, get value from main viper instance
		// - 2) --profile flag has been set, get value from set profile viper instance
		// - 3) no --profile flag set, get value from active profile viper instance defined in main viper instance

		if opt.ViperKey == options.RootActiveProfileOption.ViperKey && mainViperInstance != nil {
			// Case 1
			vValue = mainViperInstance.Get(opt.ViperKey)
		} else if definedProfileName != "" {
			// Case 2
			profileViperInstance := mainViperInstance.Sub(definedProfileName)
			if profileViperInstance != nil {
				vValue = profileViperInstance.Get(opt.ViperKey)
			}
		} else {
			// Case 3
			activeProfile := mainConfig.ActiveProfile()
			if activeProfile != nil {
				profileViperInstance := activeProfile.ViperInstance()
				if profileViperInstance != nil {
					vValue = profileViperInstance.Get(opt.ViperKey)
				}
			}
		}

		switch typedValue := vValue.(type) {
		case nil:
			// Do nothing
		case string:
			return typedValue, nil
		case []string:
			return strings.Join(typedValue, ","), nil
		case []any:
			strSlice := []string{}
			for _, v := range typedValue {
				strSlice = append(strSlice, fmt.Sprintf("%v", v))
			}
			return strings.Join(strSlice, ","), nil
		default:
			return fmt.Sprintf("%v", typedValue), nil
		}
	}

	if opt.DefaultValue != nil {
		pFlagValue = opt.DefaultValue.String()
		return pFlagValue, nil
	}

	return pFlagValue, fmt.Errorf("failed to get option value: no value found: %v", opt)
}
