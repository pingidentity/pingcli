// Copyright © 2025 Ping Identity Corporation

package connector_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/require"
)

func Test_Sanitize(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                          string
		resourceName                  string
		expectedSanitizedResourceName string
	}{
		{
			name:                          "Happy path - Simple",
			resourceName:                  "Customer",
			expectedSanitizedResourceName: "pingcli__Customer",
		},
		{
			name:                          "Happy path - Alphanumeric",
			resourceName:                  "CustomerHTMLFormPF",
			expectedSanitizedResourceName: "pingcli__CustomerHTMLFormPF",
		},
		{
			name:                          "Happy path - Spaces and Parentheses",
			resourceName:                  "Customer HTML Form (PF)",
			expectedSanitizedResourceName: "pingcli__Customer-0020-HTML-0020-Form-0020--0028-PF-0029-",
		},
		{
			name:                          "Happy path - Special Characters",
			resourceName:                  "Customer@HTML#Form$PF%",
			expectedSanitizedResourceName: "pingcli__Customer-0040-HTML-0023-Form-0024-PF-0025-",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			importBlock := connector.ImportBlock{
				ResourceName: tc.resourceName,
			}

			importBlock.Sanitize()

			require.Equal(t, importBlock.ResourceName, tc.expectedSanitizedResourceName)
		})
	}
}
