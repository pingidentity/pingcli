package feedback_internal

import (
	"github.com/pingidentity/pingcli/internal/output"
)

const FeedbackMessage string = `Thank you for participating in early adoption of the refreshed Ping Identity universal CLI!

We appreciate your feedback and suggestions for improvement regarding your experiences with the CLI.

Please visit the following URL in your browser to fill out a short, anonymous survey that will help guide our development efforts and improve the CLI for all users:

	https://forms.gle/xLz6ao4Ts86Zn2yt9

If you encounter any bugs while using the tool, please open an issue on the project's GitHub repository's issue tracker:

	https://github.com/pingidentity/pingcli/issues/new

`

// Print the feedback message
func PrintFeedbackMessage() {
	output.Print(output.Opts{
		Message: FeedbackMessage,
		Result:  output.ENUM_RESULT_NIL,
	})
}
