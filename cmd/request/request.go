// Copyright © 2025 Ping Identity Corporation

package request

import (
	"fmt"

	"github.com/pingidentity/pingcli/cmd/common"
	"github.com/pingidentity/pingcli/internal/autocompletion"
	request_internal "github.com/pingidentity/pingcli/internal/commands/request"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/spf13/cobra"
)

const (
	commandExamples = `  Send a custom API request to the configured PingOne tenant, making a GET request against the /environments endpoint.
    pingcli request --service pingone environments

  Send a custom API request to the configured PingOne tenant, making a GET request to retrieve JSON configuration for a specific environment.
    pingcli request --service pingone --http-method GET --output-format json environments/$MY_ENVIRONMENT_ID

  Send a custom API request to the configured PingOne tenant, making a POST request to create a new environment with JSON data sourced from a file.
    pingcli request --service pingone --http-method POST --data ./my-environment.json environments
	
  Send a custom API request to the configured PingOne tenant, making a POST request to create a new environment using raw JSON data.
    pingcli request --service pingone --http-method POST --data-raw '{"name": "My environment"}' environments

  Send a custom API request to the configured PingOne tenant, making a DELETE request to remove an application attribute mapping.
    pingcli request --service pingone --http-method DELETE environments/$MY_ENVIRONMENT_ID/applications/$MY_APPLICATION_ID/attributes/$MY_ATTRIBUTE_MAPPING_ID`
)

func NewRequestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               commandExamples,
		Long: `Send a custom REST API request to a Ping Service.
		
The custom REST API request is most powerful when product connection details have been configured in the CLI configuration file.
The command offers a cURL-like experience to interact with the Ping platform services, with authentication and environment details dynamically filled by the CLI.`,
		RunE:  requestRunE,
		Short: "Send a custom REST API request to a Ping platform service.",
		Use:   "request [flags] API_URI",
	}

	// --data
	cmd.Flags().AddFlag(options.RequestDataOption.Flag)

	// --data-raw
	cmd.Flags().AddFlag(options.RequestDataRawOption.Flag)

	// --fail, -f
	cmd.Flags().AddFlag(options.RequestFailOption.Flag)

	// --http-method, -m
	cmd.Flags().AddFlag(options.RequestHTTPMethodOption.Flag)
	// auto-completion
	err := cmd.RegisterFlagCompletionFunc(options.RequestHTTPMethodOption.CobraParamName, autocompletion.RequestHTTPMethodFunc)
	if err != nil {
		output.SystemError(fmt.Sprintf("Unable to register auto completion for request flag %s: %v", options.RequestHTTPMethodOption.CobraParamName, err), nil)
	}

	// --service, -s
	cmd.Flags().AddFlag(options.RequestServiceOption.Flag)
	// auto-completion
	err = cmd.RegisterFlagCompletionFunc(options.RequestServiceOption.CobraParamName, autocompletion.RequestServiceFunc)
	if err != nil {
		output.SystemError(fmt.Sprintf("Unable to register auto completion for request flag %s: %v", options.RequestServiceOption.CobraParamName, err), nil)
	}

	return cmd
}

func requestRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Request Subcommand Called.")

	if err := request_internal.RunInternalRequest(args[0]); err != nil {
		return err
	}

	return nil
}
