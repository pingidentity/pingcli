// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"gopkg.in/yaml.v3"
)

var (
	ErrRetrieveKeys = errors.New("failed to retrieve configuration keys")
	ErrNestedMap    = errors.New("failed to create nested map for key")
	ErrMarshalKeys  = errors.New("failed to marshal keys to YAML format")
)

type ListKeysError struct {
	Err error
}

func (e *ListKeysError) Error() string {
	var err *ListKeysError
	if errors.As(e.Err, &err) {
		return err.Error()
	}
	return fmt.Sprintf("failed to get configuration keys list: %s", e.Err.Error())
}

func (e *ListKeysError) Unwrap() error {
	var err *ListKeysError
	if errors.As(e.Err, &err) {
		return err.Unwrap()
	}
	return e.Err
}

func returnKeysYamlString() (keysYamlStr string, err error) {
	koanfKeys := configuration.KoanfKeys()

	if len(koanfKeys) == 0 {
		return keysYamlStr, &ListKeysError{Err: ErrRetrieveKeys}
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
					return keysYamlStr, &ListKeysError{Err: ErrNestedMap}
				}
			}
		}
	}

	// Marshal the result into YAML
	yamlData, err := yaml.Marshal(keyMap)
	if err != nil {
		return keysYamlStr, &ListKeysError{Err: ErrMarshalKeys}
	}

	keysYamlStr = string(yamlData)
	return keysYamlStr, nil
}

func returnKeysString() (string, error) {
	validKeys := configuration.KoanfKeys()

	if len(validKeys) == 0 {
		return "", &ListKeysError{Err: ErrRetrieveKeys}
	} else {
		validKeysJoined := strings.Join(validKeys, "\n- ")

		return "Valid Keys:\n- " + validKeysJoined, nil
	}
}

func RunInternalConfigListKeys() (err error) {
	var outputMessageString string
	yamlFlagStr, err := profiles.GetOptionValue(options.ConfigListKeysYamlOption)
	if err != nil {
		return &ListKeysError{Err: err}
	}
	if yamlFlagStr == "true" {
		// Output the YAML data as a string
		outputMessageString, err = returnKeysYamlString()
		if err != nil {
			return &ListKeysError{Err: err}
		}
	} else {
		// Output data list string
		outputMessageString, err = returnKeysString()
		if err != nil {
			return &ListKeysError{Err: err}
		}
	}

	output.Message(outputMessageString, nil)

	return nil
}
