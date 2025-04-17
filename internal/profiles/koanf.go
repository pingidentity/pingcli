package profiles

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/v2"
	"github.com/pingidentity/pingcli/internal/configuration/options"
)

var (
	k *KoanfConfig = NewKoanfConfig("")
)

type KoanfConfig struct {
	koanfInstance  *koanf.Koanf
	configFilePath *string
}

func NewKoanfConfig(cnfFilePath string) *KoanfConfig {
	return &KoanfConfig{
		koanfInstance:  koanf.New("."),
		configFilePath: &cnfFilePath,
	}
}

func GetKoanfConfig() *KoanfConfig {
	return k
}

func (k KoanfConfig) GetKoanfConfigFile() string {
	return *k.configFilePath
}

func (k *KoanfConfig) SetKoanfConfigFile(cnfFilePath string) error {
	k.configFilePath = &cnfFilePath
	if k.configFilePath != &cnfFilePath {
		return fmt.Errorf("failed to set koanf config file path: %s", cnfFilePath)
	}

	return nil
}

func (k *KoanfConfig) KoanfInstance() *koanf.Koanf {
	return k.koanfInstance
}

func cobraParamValueFromOption(opt options.Option) (value string, ok bool) {
	if opt.CobraParamValue != nil && opt.Flag.Changed {
		return opt.CobraParamValue.String(), true
	}

	return "", false
}

func GetActiveProfileName(k *koanf.Koanf) string {
	if k.Exists(options.RootActiveProfileOption.CobraParamName) && k.Get(options.RootActiveProfileOption.CobraParamName) != nil {
		activeProfileOption, ok := k.Get(options.RootActiveProfileOption.CobraParamName).(string)
		if ok {
			return activeProfileOption
		}

		return ""
	}

	return ""
}

func KoanfValueFromOption(opt options.Option) (value string, ok bool, err error) {
	if opt.KoanfKey != "" {
		var (
			kValue            any
			mainKoanfInstance = GetKoanfConfig()
		)

		// Case 1: Koanf Key is the ActiveProfile Key, get value from main koanf instance
		if opt.KoanfKey != "" && opt.KoanfKey == options.RootActiveProfileOption.KoanfKey && mainKoanfInstance != nil {
			kValue = mainKoanfInstance.KoanfInstance().Get(opt.KoanfKey)
		} else {
			// // Case 2: --profile flag has been set, get value from set profile koanf instance
			// // Case 3: no --profile flag set, get value from active profile koanf instance defined in main koanf instance
			// // This recursive call is safe, as options.RootProfileOption.KoanfKey is not set
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

			// Get the sub koanf instance for the profile
			subKoanf, err := mainKoanfInstance.GetProfileKoanf(pName)
			if err != nil {
				return "", false, err
			}

			kValue = subKoanf.Get(opt.KoanfKey)
		}

		switch typedValue := kValue.(type) {
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

// Get all profile names from config.yaml configuration file
// Returns a sorted slice of profile names
func (k KoanfConfig) ProfileNames() (profileNames []string) {
	keySet := make(map[string]struct{})
	mainKoanfKeys := k.KoanfInstance().All()
	for key := range mainKoanfKeys {
		// Do not add Active profile koanf key to profileNames
		if key == options.RootActiveProfileOption.KoanfKey {
			continue
		}

		pName := strings.Split(key, ".")[0]
		if _, ok := keySet[pName]; !ok {
			keySet[pName] = struct{}{}
			profileNames = append(profileNames, pName)
		}
	}

	return profileNames
}

// The profile name must contain only alphanumeric characters, underscores, and dashes
// The profile name cannot be empty
func (k KoanfConfig) ValidateProfileNameFormat(pName string) (err error) {
	if pName == "" {
		return fmt.Errorf("invalid profile name: profile name cannot be empty")
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9\_\-]+$`)
	if !re.MatchString(pName) {
		return fmt.Errorf("invalid profile name: '%s'. name must contain only alphanumeric characters, underscores, and dashes", pName)
	}

	return nil
}

func (k KoanfConfig) ChangeActiveProfile(pName string) (err error) {
	if err = k.ValidateExistingProfileName(pName); err != nil {
		return err
	}

	err = k.KoanfInstance().Set(options.RootActiveProfileOption.KoanfKey, pName)
	if err != nil {
		return fmt.Errorf("error setting active profile: %w", err)
	}

	return nil
}

// The profile name must exist
func (k KoanfConfig) ValidateExistingProfileName(pName string) (err error) {
	if pName == "" {
		return fmt.Errorf("invalid profile name: profile name cannot be empty")
	}

	pNames := k.ProfileNames()

	if !slices.ContainsFunc(pNames, func(n string) bool {
		return n == pName
	}) {
		return fmt.Errorf("invalid profile name: '%s' profile does not exist", pName)
	}

	return nil
}

// The profile name format must be valid
// The new profile name must be unique
func (k KoanfConfig) ValidateNewProfileName(pName string) (err error) {
	if err = k.ValidateProfileNameFormat(pName); err != nil {
		return err
	}

	pNames := k.ProfileNames()

	if slices.ContainsFunc(pNames, func(n string) bool {
		return n == pName
	}) {
		return fmt.Errorf("invalid profile name: '%s'. profile already exists", pName)
	}

	return nil
}

func (k KoanfConfig) GetProfileKoanf(pName string) (subKoanf *koanf.Koanf, err error) {
	if err = k.ValidateExistingProfileName(pName); err != nil {
		return nil, err
	}

	// Create a new koanf instance for the profile
	subKoanf = koanf.New(".")
	err = subKoanf.Load(confmap.Provider(k.KoanfInstance().Cut(pName).All(), "."), nil)
	if err != nil {
		return nil, fmt.Errorf("error marshalling koanf: %w", err)
	}

	return subKoanf, nil
}

func (k KoanfConfig) WriteFile() (err error) {
	encodedConfig, err := k.KoanfInstance().Marshal(yaml.Parser())
	if err != nil {
		return fmt.Errorf("error marshalling koanf: %w", err)
	}

	f, err := os.OpenFile(k.GetKoanfConfigFile(), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file (%s): %w", k.GetKoanfConfigFile(), err)
	}

	defer f.Close()

	_, err = f.Write(encodedConfig)
	if err != nil {
		return fmt.Errorf("error writing to file (%s): %w", k.GetKoanfConfigFile(), err)
	}

	err = f.Sync()
	if err != nil {
		return fmt.Errorf("error syncing file (%s): %w", k.GetKoanfConfigFile(), err)
	}

	return nil
}

func (k KoanfConfig) SaveProfile(pName string, subKoanf *koanf.Koanf) (err error) {
	err = k.KoanfInstance().MergeAt(subKoanf, pName)
	if err != nil {
		return fmt.Errorf("error merging koanf: %w", err)
	}

	err = k.WriteFile()
	if err != nil {
		return fmt.Errorf("failed to save profile '%s': %w", pName, err)
	}

	return nil
}

func (k KoanfConfig) DeleteProfile(pName string) (err error) {
	if err = k.ValidateExistingProfileName(pName); err != nil {
		return err
	}

	activeProfileName, err := GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		return err
	}

	if activeProfileName == pName {
		return fmt.Errorf("'%s' is the active profile and cannot be deleted", pName)
	}

	// Delete the profile from the main koanf
	k.KoanfInstance().Delete(pName)

	err = k.WriteFile()
	if err != nil {
		return fmt.Errorf("failed to delete profile '%s': %w", pName, err)
	}

	return nil
}

func (k KoanfConfig) DefaultMissingKoanfKeys() (err error) {
	// For each profile, if a koanf key from an option doesn't exist, set it to the default value
	for _, pName := range k.ProfileNames() {
		subKoanf, err := k.GetProfileKoanf(pName)
		if err != nil {
			return err
		}

		for _, opt := range options.Options() {
			if opt.KoanfKey == "" || opt.KoanfKey == options.RootActiveProfileOption.KoanfKey {
				continue
			}

			if !subKoanf.Exists(opt.KoanfKey) {
				err = subKoanf.Set(opt.KoanfKey, opt.DefaultValue)
				if err != nil {
					return fmt.Errorf("error setting default value for koanf key %s: %w", opt.KoanfKey, err)
				}
			}
		}
		err = k.SaveProfile(pName, subKoanf)
		if err != nil {
			return fmt.Errorf("failed to save profile '%s': %w", pName, err)
		}
	}

	return nil
}

func GetOptionValue(opt options.Option) (string, error) {
	// 1st priority: cobra param flag value
	if cobraParamValue, ok := cobraParamValueFromOption(opt); ok {
		return cobraParamValue, nil
	}

	// 2nd priority: environment variable value
	if pFlagValue := os.Getenv(opt.EnvVar); pFlagValue != "" {
		return pFlagValue, nil
	}

	// 3rd priority: koanf value
	koanfValue, ok, _ := KoanfValueFromOption(opt)
	if ok {
		return koanfValue, nil
	}

	// 4th priority: default value
	if opt.DefaultValue != nil {
		return opt.DefaultValue.String(), nil
	}

	// This is a error, as it means the option is not configured internally to contain one of the 4 values above.
	// This should never happen, as all options should at least have a default value.
	return "", fmt.Errorf("failed to get option value: no value found: %v", opt)
}

func MaskValue(value any) string {
	stringValue, ok := value.(string)
	if ok && stringValue == "" {
		return stringValue
	}

	// Mask all values to the same asterisk length
	// providing no additional information about the value when logged.
	return strings.Repeat("*", 8)
}
