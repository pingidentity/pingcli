// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"errors"
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
)

func Test_RunInternalConfigListKeys(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError error
	}{
		{
			name: "Get List of Keys",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := RunInternalConfigListKeys()

			if tc.expectedError != nil {
				assert.Error(t, err)
				var listKeysErr *ListKeysError
				if errors.As(err, &listKeysErr) {
					assert.ErrorIs(t, listKeysErr.Unwrap(), tc.expectedError)
				} else {
					assert.Fail(t, "Expected error to be of type ListKeysError")
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
