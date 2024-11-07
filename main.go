package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/pingidentity/pingcli/cmd"
	"github.com/pingidentity/pingcli/internal/output"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"
	commit  string = "dev"
)

func main() {
	// Try to get the commit hash from the build info if it wasn't set by goreleaser
	if commit == "dev" {
		if info, ok := debug.ReadBuildInfo(); ok {
			for _, setting := range info.Settings {
				if setting.Key == "vcs.revision" {
					commit = setting.Value
					break
				}
			}
		}
	}

	rootCmd := cmd.NewRootCommand(version, commit)

	err := rootCmd.Execute()
	if err != nil {
		output.UserError(fmt.Sprintf("Failed to execute pingcli: %v", err), nil)
		os.Exit(1)
	}
}
