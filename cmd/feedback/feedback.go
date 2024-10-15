package feedback

import (
	"github.com/pingidentity/pingcli/cmd/common"
	feedback_internal "github.com/pingidentity/pingcli/internal/commands/feedback"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

func NewFeedbackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               `  pingcli feedback`,
		Long:                  "Provides links to report issues and provide feedback on using the CLI to Ping Identity.",
		RunE:                  feedbackRunE,
		Short:                 "Help us improve the CLI. Report issues or send us feedback on using the CLI tool.",
		Use:                   "feedback [flags]",
	}

	return cmd
}

func feedbackRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Running Feedback Subcommand with args %s", args)

	feedback_internal.PrintFeedbackMessage()

	return nil
}
