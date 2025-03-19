// Copyright © 2025 Ping Identity Corporation

package completion

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	desc = `To load completions:

Bash:

  $ source <(%[1]s completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ %[1]s completion bash > /etc/bash_completion.d/%[1]s
  # macOS:
  $ source <(%[1]s completion zsh)

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ %[1]s completion zsh > "${fpath[1]}/_%[1]s"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ %[1]s completion fish | source

  # To load completions for each session, execute once:
  $ %[1]s completion fish > ~/.config/fish/completions/%[1]s.fish

PowerShell:

  PS> %[1]s completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> %[1]s completion powershell > %[1]s.ps1
  # and source this file from your PowerShell profile.
`
)

func completionCmdRunE(cmd *cobra.Command, args []string) error {
	switch args[0] {
	case "bash":
		_ = cmd.Root().GenBashCompletionV2(cmd.OutOrStdout(), true)
	case "zsh":
		_ = cmd.Root().GenZshCompletion(cmd.OutOrStdout())
	case "fish":
		_ = cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
	case "powershell":
		_ = cmd.Root().GenPowerShellCompletion(cmd.OutOrStdout())
	}

	return nil
}

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "completion [SHELL]",
		Short:     "Prints shell completion scripts",
		Long:      fmt.Sprintf(desc, "pingcli"),
		ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE:      completionCmdRunE,
	}
	return cmd
}
