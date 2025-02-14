# Autocompletion Configuration
## To load autocompletion for select command flags, run the applicable command for your environment.

## Bash:
  ### To load completions for one session, execute:
  `$ source <(pingcli completion bash)`

  ### To load completions for every future session, execute once:
  #### Linux:
  `$ pingcli completion bash > /etc/bash_completion.d/pingcli`
  #### macOS:
  `$ pingcli completion bash > $(brew --prefix)/etc/bash_completion.d/pingcli`

## Zsh:
  ### If shell completion is not already enabled in your environment,
  ### you will need to enable it.  You can execute the following once:
  `$ echo "autoload -U compinit; compinit" >> ~/.zshrc`

  ### To load completions for every future session, execute once:
  `$ pingcli completion zsh > "${fpath[1]}/_pingcli"`

  ### You will need to start a new shell for this setup to take effect.

## fish:
  ### To load completions for one session, execute:
  `$ pingcli completion fish | source`

  ### To load completions for every future session, execute once:
  `$ pingcli completion fish > ~/.config/fish/completions/pingcli.fish`

## PowerShell:
  ### To load completions for one session, execute:
  `PS> pingcli completion powershell | Out-String | Invoke-Expression`

  ### To load completions for every future session:
  `PS> pingcli completion powershell > pingcli.ps1`
  ### then source this file from your PowerShell profile