package request

import (
	"fmt"

	"github.com/pingidentity/pingcli/cmd/common"
	request_internal "github.com/pingidentity/pingcli/internal/commands/request"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	commandExamples = `  pingcli request --service pingone environments
  pingcli request --service pingone --http-method GET environments/{{environmentID}}
  pingcli request --service pingone --http-method POST --data {{raw-data}} environments
  pingcli request --service pingone --http-method POST --data @{{filepath}} environments
  pingcli request --service pingone --http-method DELETE environments/{{environmentID}}`

	profileConfigurationFormat = `Profile Configuration Format:
request:
    data: @<Filepath> OR <RawData>
    http-method: <Method>
    service: <Service>
service:
    pingone:
        regionCode: <Code>
        authentication:
            type: <Type>
            worker:
                clientID: <ID>
                clientSecret: <Secret>
                environmentID: <ID>`
)

func NewRequestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               fmt.Sprintf("%s\n\n%s", commandExamples, profileConfigurationFormat),
		Long:                  `Send a custom request to a Ping Service.`,
		RunE:                  requestRunE,
		Short:                 "Send a custom request to a Ping Service.",
		Use:                   "request [flags] API_URI",
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
