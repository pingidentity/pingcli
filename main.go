// Copyright © 2025 Ping Identity Corporation

package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"slices"

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

	if !slices.Contains(os.Args, "--version") &&
		!slices.Contains(os.Args, "-v") &&
		!slices.Contains(os.Args, "--help") &&
		!slices.Contains(os.Args, "-h") {
		detailedExitCodeWarnLogged, err := output.DetailedExitCodeWarnLogged()
		if err != nil {
			output.UserError(fmt.Sprintf("Failed to execute pingcli: %v", err), nil)
			os.Exit(1)
		}
		if detailedExitCodeWarnLogged {
			os.Exit(2)
		}
	}

	os.Exit(0)
}
