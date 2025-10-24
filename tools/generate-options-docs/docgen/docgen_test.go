package docgen_test

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/tools/docutil"
	docgen "github.com/pingidentity/pingcli/tools/generate-options-docs/docgen"
)

var update = flag.Bool("update", false, "update golden files for options docs")

// TestOptionsDocGeneration validates both markdown and AsciiDoc outputs against goldens.
func TestOptionsDocGeneration(t *testing.T) {
	flag.Parse()

	md := docgen.GenerateMarkdown()
	adoc := docgen.GenerateAsciiDoc()

	goldenDir := filepath.Join("testdata", "golden")
	if err := os.MkdirAll(goldenDir, 0o750); err != nil { // tighter perms
		t.Fatalf("mkdir golden: %v", err)
	}

	// Normalize dynamic date in AsciiDoc output before storing / comparing.
	adocNorm := docutil.NormalizeForCompare(adoc)

	cases := []struct {
		name    string
		content string
	}{
		{"options.md", md},
		{"options.adoc", adocNorm},
	}

	for _, tc := range cases {
		goldenPath := filepath.Join(goldenDir, tc.name)
		// Validate golden file path remains within goldenDir to mitigate G304
		cleanGolden := filepath.Clean(goldenPath)
		if !strings.HasPrefix(cleanGolden+string(os.PathSeparator), filepath.Clean(goldenDir)+string(os.PathSeparator)) {
			t.Fatalf("invalid golden file path: %s", goldenPath)
		}
		if *update {
			if err := os.WriteFile(cleanGolden, []byte(tc.content), 0o600); err != nil {
				t.Fatalf("write golden %s: %v", tc.name, err)
			}
			t.Logf("updated golden: %s", tc.name)

			continue
		}
		wantBytes, err := os.ReadFile(cleanGolden) // #nosec G304 path validated
		if err != nil {
			t.Fatalf("read golden %s: %v (run with -update to create)", tc.name, err)
		}
		want := string(wantBytes)
		if tc.content != want {
			t.Errorf("mismatch for %s\n--- got ---\n%s\n--- want ---\n%s", tc.name, tc.content, want)
		}
	}
}
