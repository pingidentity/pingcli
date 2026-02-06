// Copyright © 2026 Ping Identity Corporation

package config_internal

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"gopkg.in/yaml.v3"
)

var (
	listKeysErrorPrefix = "failed to get configuration keys list"
)

func returnKeysYamlString() (keysYamlStr string, err error) {
	koanfKeys := configuration.KoanfKeys()

	if len(koanfKeys) == 0 {
		return keysYamlStr, &errs.PingCLIError{Prefix: listKeysErrorPrefix, Err: ErrRetrieveKeys}
	}

	// Split the input string into individual keys
	keyMap := make(map[string]interface{})

	// Iterate over each koanf key
	for _, koanfKey := range koanfKeys {
		// Skip the "activeProfile" key
		if koanfKey == options.RootActiveProfileOption.KoanfKey {
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
					wrappedErr := fmt.Errorf("key '%s': %w", koanfKey, ErrNestedMap)

					return keysYamlStr, &errs.PingCLIError{Prefix: listKeysErrorPrefix, Err: wrappedErr}
				}
			}
		}
	}

	// Marshal the result into YAML
	yamlData, err := yaml.Marshal(keyMap)
	if err != nil {
		return keysYamlStr, &errs.PingCLIError{Prefix: listKeysErrorPrefix, Err: ErrMarshalKeys}
	}

	keysYamlStr = string(yamlData)

	return keysYamlStr, nil
}

func returnKeysString() (string, error) {
	validKeys := configuration.KoanfKeys()

	if len(validKeys) == 0 {
		return "", &errs.PingCLIError{Prefix: listKeysErrorPrefix, Err: ErrRetrieveKeys}
	}

	// Remove the "activeProfile" key from the list
	validKeys = slices.DeleteFunc(validKeys, func(s string) bool {
		return s == options.RootActiveProfileOption.KoanfKey
	})

	validKeysJoined := strings.Join(validKeys, "\n- ")

	return "Valid Keys:\n- " + validKeysJoined, nil
}

func RunInternalConfigListKeys() (err error) {
	var outputMessageString string
	yamlFlagStr, err := profiles.GetOptionValue(options.ConfigListKeysYamlOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: listKeysErrorPrefix, Err: err}
	}
	if yamlFlagStr == "true" {
		// Output the YAML data as a string
		outputMessageString, err = returnKeysYamlString()
		if err != nil {
			return &errs.PingCLIError{Prefix: listKeysErrorPrefix, Err: err}
		}
	} else {
		// Output data list string
		outputMessageString, err = returnKeysString()
		if err != nil {
			return &errs.PingCLIError{Prefix: listKeysErrorPrefix, Err: err}
		}
	}

	output.Message(outputMessageString, nil)

	return nil
}
