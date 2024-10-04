package output

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var (
	boldRed = color.New(color.FgRed).Add(color.Bold).SprintfFunc()
	cyan    = color.New(color.FgCyan).SprintfFunc()
	green   = color.New(color.FgGreen).SprintfFunc()
	red     = color.New(color.FgRed).SprintfFunc()
	white   = color.New(color.FgWhite).SprintfFunc()
	yellow  = color.New(color.FgYellow).SprintfFunc()
)

type Result string

type Opts struct {
	Fields       map[string]interface{}
	Message      string
	ErrorMessage string
	FatalMessage string
	Result       Result
}

const (
	ENUM_RESULT_NIL           Result = ""
	ENUM_RESULT_SUCCESS       Result = "Success"
	ENUM_RESULT_NOACTION_OK   Result = "No Action (OK)"
	ENUM_RESULT_NOACTION_WARN Result = "No Action (Warning)"
	ENUM_RESULT_FAILURE       Result = "Failure"
)

func Print(output Opts) {
	colorizeOutput, err := profiles.GetOptionValue(options.RootColorOption)
	if err != nil {
		color.NoColor = false
	} else {
		colorizeOutputBool, err := strconv.ParseBool(colorizeOutput)
		if err != nil {
			color.NoColor = false
		} else {
			color.NoColor = !colorizeOutputBool
		}
	}

	outputFormat, err := profiles.GetOptionValue(options.RootOutputFormatOption)
	if err != nil {
		outputFormat = customtypes.ENUM_OUTPUT_FORMAT_TEXT
	}

	switch outputFormat {
	case customtypes.ENUM_OUTPUT_FORMAT_TEXT:
		printText(output)
	case customtypes.ENUM_OUTPUT_FORMAT_JSON:
		printJson(output)
	default:
		printText(Opts{
			Message: fmt.Sprintf("Output format %q is not recognized. Defaulting to \"text\" output", outputFormat),
			Result:  ENUM_RESULT_NOACTION_WARN,
		})
		printText(output)
	}
}

func printText(opts Opts) {
	l := logger.Get()

	var resultFormat string
	var resultColor func(format string, a ...interface{}) string

	// Determine message color and format based on status
	switch opts.Result {
	case ENUM_RESULT_SUCCESS:
		resultFormat = "%s - %s"
		resultColor = green
	case ENUM_RESULT_NOACTION_OK:
		resultFormat = "%s - %s"
		resultColor = green
	case ENUM_RESULT_NOACTION_WARN:
		resultFormat = "%s - %s"
		resultColor = yellow
	case ENUM_RESULT_FAILURE:
		resultFormat = "%s - %s"
		resultColor = red
	case ENUM_RESULT_NIL:
		resultFormat = "%s%s"
		resultColor = white
	default:
		resultFormat = "%s%s"
		resultColor = white
	}

	// Supply the user a formatted message and a result status if any.
	fmt.Println(resultColor(resultFormat, opts.Message, opts.Result))
	l.Info().Msg(resultColor(resultFormat, opts.Message, opts.Result))

	// Output and log any additional key/value pairs supplied to the user.
	if opts.Fields != nil {
		fmt.Println(cyan("Additional Information:"))
		for k, v := range opts.Fields {
			switch typedValue := v.(type) {
			// If the value is a json.RawMessage, print it as a string
			case json.RawMessage:
				fmt.Println(cyan("%s: %s", k, typedValue))
				l.Info().Msgf("%s: %s", k, typedValue)
			default:
				fmt.Println(cyan("%s: %v", k, v))
				l.Info().Msgf("%s: %v", k, v)
			}
		}
	}

	// Inform the user of an error and log the error
	if opts.ErrorMessage != "" {
		fmt.Println(red("Error: %s", opts.ErrorMessage))
		l.Error().Msg(opts.ErrorMessage)
	}

	// Inform the user of a fatal error and log the fatal error. This exits the program.
	if opts.FatalMessage != "" {
		fmt.Println(boldRed("Fatal: %s", opts.FatalMessage))
		l.Fatal().Msg(opts.FatalMessage)
	}

}

func printJson(opts Opts) {
	l := logger.Get()

	// Convert the CommandOutput struct to JSON
	jsonOut, err := json.MarshalIndent(opts, "", "  ")
	if err != nil {
		l.Error().Err(err).Msgf("Failed to serialize output as JSON")
	}

	// Output the JSON as uncolored string
	fmt.Println(string(jsonOut))

	switch opts.Result {
	case ENUM_RESULT_NOACTION_WARN:
		l.Warn().Msg(string(jsonOut))
	case ENUM_RESULT_FAILURE:
		// Log the error if exists
		if opts.ErrorMessage != "" {
			l.Error().Msg(opts.ErrorMessage)
		}

		// Log the fatal error if exists. This exits the program.
		if opts.FatalMessage != "" {
			l.Fatal().Msg(opts.FatalMessage)
		}
	default: //ENUM_RESULT_SUCCESS, ENUM_RESULT_NIL, ENUM_RESULT_NOACTION_OK
		l.Info().Msg(string(jsonOut))
	}

}
