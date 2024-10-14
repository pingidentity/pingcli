package request

import (
	"github.com/pingidentity/pingcli/cmd/common"
	request_internal "github.com/pingidentity/pingcli/internal/commands/request"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	commandExamples = `  Send a custom API request to the configured PingOne tenant, making a GET request against the /environments endpoint.
    pingcli request --service pingone environments

  Send a custom API request to the configured PingOne tenant, making a GET request to retrieve JSON configuration for a specific environment.
    pingcli request --service pingone --http-method GET environments/$MY_ENVIRONMENT_ID

  Send a custom API request to the configured PingOne tenant, making a POST request to create a new environment with raw JSON data.
    pingcli request --service pingone --http-method POST --data '{"name": "My environment"}' environments

  Send a custom API request to the configured PingOne tenant, making a POST request to create a new environment with JSON data sourced from a file.
    pingcli request --service pingone --http-method POST --data @./my-environment.json environments

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

	cmd.Flags().AddFlag(options.RequestHTTPMethodOption.Flag)
	cmd.Flags().AddFlag(options.RequestServiceOption.Flag)
	cmd.Flags().AddFlag(options.RequestDataOption.Flag)

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
