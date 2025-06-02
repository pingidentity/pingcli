// Copyright © 2025 Ping Identity Corporation

package feedback_internal

import (
	"github.com/pingidentity/pingcli/internal/output"
)

const FeedbackMessage string = `We want to hear your feedback!

New features you'd like to see?
Things you like or dislike?
Bugs you've encountered?

We're looking for your comments and suggestions regarding your experiences with the CLI to make the tool better for our developer community.

Please visit the following URL in your browser to fill out a short, anonymous survey that will help us understand more about your experiences with the CLI.
Most fields are optional, and you can fill out the form as many times as you like, so please do provide feedback whenever you have something to share.

	https://forms.gle/xLz6ao4Ts86Zn2yt9

If you encounter any bugs while using the tool, please report it to us.
Open an issue on the project's GitHub repository's issue tracker:

	https://github.com/pingidentity/pingcli/issues/new

`

// Print the feedback message
func PrintFeedbackMessage() {
	output.Message(FeedbackMessage, nil)
}
