package main

import (
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
	rootCmd := cmd.NewRootCommand(version, commit)

	err := rootCmd.Execute()
	if err != nil {
		output.Print(output.Opts{
			ErrorMessage: err.Error(),
			Message:      "Failed to execute pingcli",
			Result:       output.ENUM_RESULT_FAILURE,
		})
	}
}
