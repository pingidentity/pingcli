package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var update = flag.Bool("update", false, "update golden files for command docs")

// TestCommandDocGeneration generates documentation for the real root command and compares
// a subset of produced files (root command + nav) against golden fixtures.
func TestCommandDocGeneration(t *testing.T) {
	flag.Parse()

	tmp := t.TempDir()
	// Run the generator with deterministic date so golden files are stable.
	date := "January 2, 2006" // Intentional fixed sample date
	os.Args = []string{"docgen", "-o", tmp, "-date", date}
	main()

	goldenDir := filepath.Join("testdata", "golden")
	if err := os.MkdirAll(goldenDir, 0o750); err != nil {
		t.Fatalf("mkdir golden: %v", err)
	}

	files := []string{"pingcli.adoc", "nav.adoc"}

	for _, f := range files {
		gotPath := filepath.Join(tmp, f)
		// Validate generated file path stays within temporary directory (mitigates G304)
		cleanGot := filepath.Clean(gotPath)
		if !strings.HasPrefix(cleanGot+string(os.PathSeparator), filepath.Clean(tmp)+string(os.PathSeparator)) {
			t.Fatalf("generated file path %s is outside of temp directory", gotPath)
		}
		gotBytes, err := os.ReadFile(cleanGot) // #nosec G304 path validated above
		if err != nil {
			t.Fatalf("read generated %s: %v", f, err)
		}
		got := normalizeDynamic(string(gotBytes))

		goldenPath := filepath.Join(goldenDir, f)
		if *update {
			if err := os.WriteFile(goldenPath, []byte(got), 0o600); err != nil {
				t.Fatalf("write golden %s: %v", f, err)
			}
			t.Logf("updated golden: %s", f)

			continue
		}
		wantBytes, err := os.ReadFile(goldenPath) // #nosec G304 path validated above
		if err != nil {
			t.Fatalf("read golden %s: %v (run with -update to create)", f, err)
		}
		want := normalizeDynamic(string(wantBytes))
		if got != want {
			t.Errorf("mismatch for %s\n--- got ---\n%s\n--- want ---\n%s", f, got, want)
		}
	}
}

// normalizeDynamic strips lines containing created / revision dates to avoid churn.
func normalizeDynamic(s string) string {
	out := make([]string, 0, 64)
	for _, line := range strings.Split(s, "\n") {
		if strings.HasPrefix(line, ":created-date:") || strings.HasPrefix(line, ":revdate:") {
			continue
		}
		out = append(out, line)
	}
	return strings.Join(out, "\n")
}
