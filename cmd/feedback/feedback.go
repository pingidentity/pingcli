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
		Long: `A command to provide the user information
on how to give feedback or get help with the tool
through the use of the GitHub repository's issue tracker.`,
		RunE:  feedbackRunE,
		Short: "Information on tool feedback",
		Use:   "feedback [flags]",
	}

	return cmd
}

func feedbackRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Running Feedback Subcommand with args %s", args)

	feedback_internal.PrintFeedbackMessage()

	return nil
}
