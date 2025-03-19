// Copyright © 2025 Ping Identity Corporation

package connector_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
)

// Test sanitization of resource name
func TestSanitize(t *testing.T) {
	sanitizedResourceName := "pingcli--Customer-0020-HTML-0020-Form-0020--0028-PF-0029-"

	importBlock := connector.ImportBlock{
		ResourceName: "Customer HTML Form (PF)",
	}

	importBlock.Sanitize()

	if importBlock.ResourceName != sanitizedResourceName {
		t.Errorf("Sanitize function test failed")
	}
}
