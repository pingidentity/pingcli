// Copyright © 2025 Ping Identity Corporation

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
	koanfKeys := configuration.KoanfKeys()

	if len(koanfKeys) == 0 {
		return "", fmt.Errorf("unable to retrieve valid keys")
	}

	// Split the input string into individual keys
	keyMap := make(map[string]interface{})

	// Iterate over each koanf key
	for _, koanfKey := range koanfKeys {
		// Skip the "activeProfile" key
		if koanfKey == "activeProfile" {
			continue
		}

		// Create a nested map for each yaml key
		var (
			currentMap   = keyMap
			currentMapOk bool
		)
		yamlKeys := strings.Split(koanfKey, ".")
		for i, k := range yamlKeys {
			// If it's the last yaml key, set an empty map
			if i == len(yamlKeys)-1 {
				currentMap[k] = ""
			} else {
				// Otherwise, create or navigate to the next level
				if _, exists := currentMap[k]; !exists {
					currentMap[k] = make(map[string]interface{})
				}
				currentMap, currentMapOk = currentMap[k].(map[string]interface{})
				if !currentMapOk {
					return "", fmt.Errorf("failed to get configuration keys list: error creating nested map for key %s", koanfKey)
				}
			}
		}
	}

	// Marshal the result into YAML
	yamlData, err := yaml.Marshal(keyMap)
	if err != nil {
		return "", fmt.Errorf("error marshaling keys to YAML format")
	}

	return string(yamlData), nil
}

func returnKeysString() (string, error) {
	// var err error
	validKeys := configuration.KoanfKeys()

	if len(validKeys) == 0 {
		return "", fmt.Errorf("unable to retrieve valid keys")
	} else {
		validKeysJoined := strings.Join(validKeys, "\n- ")

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
