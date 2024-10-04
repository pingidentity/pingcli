package main

import (
	"github.com/pingidentity/pingcli/cmd"
	"github.com/pingidentity/pingcli/internal/output"
)

func main() {
	rootCmd := cmd.NewRootCommand()

	err := rootCmd.Execute()
	if err != nil {
		output.Print(output.Opts{
			ErrorMessage: err.Error(),
			Message:      "Failed to execute pingcli",
			Result:       output.ENUM_RESULT_FAILURE,
		})
	}
}
