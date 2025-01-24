package config_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"gopkg.in/yaml.v3"
)

func returnKeysYamlString() (string, error) {
	var err error
	validKeys := configuration.ViperKeys()

	validKeysJoined := strings.Join(validKeys, " ")

	if len(validKeys) == 0 {
		return "", fmt.Errorf("unable to retrieve valid keys")
	}

	// Split the input string into individual keys
	parts := strings.Split(validKeysJoined, " ")
	result := make(map[string]interface{})

	// Iterate over each part
	for _, part := range parts {
		// Skip the "activeProfile" key
		if part == "activeProfile" {
			continue
		}

		// Create a nested map for each part
		currentMap := result
		keys := strings.Split(part, ".")
		for i, key := range keys {
			// If it's the last key, set an empty map
			if i == len(keys)-1 {
				currentMap[key] = ""
			} else {
				// Otherwise, create or navigate to the next level
				if _, exists := currentMap[key]; !exists {
					currentMap[key] = make(map[string]interface{})
				}
				currentMap = currentMap[key].(map[string]interface{})
			}
		}
	}

	// Marshal the result into YAML
	yamlData, err := yaml.Marshal(result)
	if err != nil {
		return "", fmt.Errorf("error marshaling keys to YAML format")
	}

	return string(yamlData), nil
}

func returnKeysString() (string, error) {
	// var err error
	validKeys := configuration.ViperKeys()

	validKeysJoined := strings.Join(validKeys, "\n- ")

	if len(validKeys) == 0 {
		return "", fmt.Errorf("unable to retrieve valid keys")
	} else {
		return "Valid Keys:\n- " + validKeysJoined, nil
	}
}

func RunInternalConfigListKeys() (err error) {
	var outputMessageString string
	yamlFlagStr, err := profiles.GetOptionValue(options.ConfigListKeysYamlOption)
	if err != nil {
		return err
	}
	if yamlFlagStr == "true" {
		// Output the YAML data as a string
		outputMessageString, err = returnKeysYamlString()
		if err != nil {
			return err
		}
	} else {
		// Output data list string
		outputMessageString, err = returnKeysString()
		if err != nil {
			return err
		}
	}

	output.Message(outputMessageString, nil)

	return nil
}
