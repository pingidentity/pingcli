package config_internal

import (
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"gopkg.in/yaml.v3"
)

func returnKeysYamlString() string {
	inputString := returnKeysString("", " ")
	// Split the input by spaces
	parts := strings.Fields(inputString)

	// Create a nested map based on the period (.) separator
	result := make(map[string]interface{})

	// Flag to indicate if activeProfile is processed
	activeProfileFound := false

	// Iterate through each part
	for _, part := range parts {
		keys := strings.Split(part, ".")

		// Only treat the first occurrence of activeProfile as the root key
		if keys[0] == "activeProfile" && !activeProfileFound {
			// Set activeProfile as the top-level key and initialize it as an empty string
			result["activeProfile"] = ""
			activeProfileFound = true
			continue // Skip further processing for the "activeProfile" key itself
		}

		// Now handle nested elements only under activeProfile
		if activeProfileFound {
			// Ensure we create a map for the activeProfile if it's not already created
			if _, exists := result["activeProfile"].(map[string]interface{}); !exists {
				result["activeProfile"] = make(map[string]interface{})
			}

			// Get the map for activeProfile
			current := result["activeProfile"].(map[string]interface{})

			// Process each key in the current part
			for i, key := range keys {
				if i == len(keys)-1 {
					// Set the value for the last key, which should be an empty string
					current[key] = ""
				} else {
					// Create intermediate maps for nested keys
					if _, exists := current[key]; !exists {
						current[key] = make(map[string]interface{})
					}
					// Move deeper into the nested map
					current = current[key].(map[string]interface{})
				}
			}
		}
	}

	// Marshal the result into YAML
	yamlData, err := yaml.Marshal(result)
	if err != nil {
		output.Warn("Error marshaling keys to YAML format", nil)
		return ""
	}

	return string(yamlData)
}

func returnKeysString(stringPrefix, keyPrefix string) string {
	var err error
	l := logger.Get()
	validKeys := configuration.ViperKeys()

	validKeysJoined := strings.Join(validKeys, keyPrefix)

	if len(validKeys) == 0 {
		l.Err(err).Msg("Unable to retrieve valid keys.")
		return ""
	} else {
		return stringPrefix + validKeysJoined
	}
}

func RunInternalConfigListKeys() (err error) {
	l := logger.Get()
	var outputMessageString string
	if yamlStr, err := profiles.GetOptionValue(options.ConfigListKeysYamlOption); yamlStr == "true" {
		if err != nil {
			l.Err(err).Msg("Unable to get list keys option.")
		}
		// Output the YAML data as a string
		outputMessageString = returnKeysYamlString()
	} else {
		listKeysStr := "Valid Keys:\n- "
		outputMessageString = returnKeysString(listKeysStr, "\n- ")
	}

	output.Message(outputMessageString, nil)

	return nil
}
