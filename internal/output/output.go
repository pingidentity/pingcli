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
	"github.com/rs/zerolog"
)

var (
	boldRed = color.New(color.FgRed).Add(color.Bold).SprintfFunc()
	cyan    = color.New(color.FgCyan).SprintfFunc()
	green   = color.New(color.FgGreen).SprintfFunc()
	red     = color.New(color.FgRed).SprintfFunc()
	white   = color.New(color.FgWhite).SprintfFunc()
	yellow  = color.New(color.FgYellow).SprintfFunc()
)

// Set the faith color option based on user configuration
func SetColorize() {
	disableColorOutput, err := profiles.GetOptionValue(options.RootColorOption)
	if err != nil {
		color.NoColor = false
	} else {
		disableColorOutputBool, err := strconv.ParseBool(disableColorOutput)
		if err != nil {
			color.NoColor = false
		} else {
			color.NoColor = disableColorOutputBool
		}
	}
}

// This function outputs white text to supply information to the user.
func Message(message string, fields map[string]interface{}) {
	l := logger.Get()

	print(message, fields, white, l.Info)
}

// This function outputs green text to inform the user of success
func Success(message string, fields map[string]interface{}) {
	l := logger.Get()

	print(fmt.Sprintf("SUCCESS: %s", message), fields, green, l.Info)
}

// This function outputs yellow text to inform the user of a warning
func Warn(message string, fields map[string]interface{}) {
	l := logger.Get()

	print(fmt.Sprintf("WARNING: %s", message), fields, yellow, l.Warn)
}

// This functions is used to inform the user their configuration
// or input to pingcli has caused an error.
func UserError(message string, fields map[string]interface{}) {
	l := logger.Get()
	print(fmt.Sprintf("ERROR: %s", message), fields, red, l.Error)
}

// This functions is used to inform the user their configuration
// or input to pingcli has caused an fatal error that exits the program immediately.
func UserFatal(message string, fields map[string]interface{}) {
	l := logger.Get()
	print(fmt.Sprintf("FATAL: %s", message), fields, boldRed, l.Fatal)
}

// This function is used to inform the user a system-level error
// has occurred. These errors should indicate a bug or bad behavior
// of the tool.
func SystemError(message string, fields map[string]interface{}) {
	l := logger.Get()
	systemMsg := fmt.Sprintf(`FATAL: %s
		
This is a bug in the Ping CLI and needs reporting to our team.
		
Please raise an issue at https://github.com/pingidentity/pingcli`,
		message)

	// l.Fatal() exits the program prematurely before the message is printed
	// pass nil to print the message before exiting
	print(systemMsg, fields, boldRed, l.Fatal)
}

func print(message string, fields map[string]interface{}, colorFunc func(format string, a ...interface{}) string,
	logEventFunc func() *zerolog.Event) {
	SetColorize()

	outputFormat, err := profiles.GetOptionValue(options.RootOutputFormatOption)
	if err != nil {
		outputFormat = customtypes.ENUM_OUTPUT_FORMAT_TEXT
	}

	switch outputFormat {
	case customtypes.ENUM_OUTPUT_FORMAT_TEXT:
		printText(message, fields, colorFunc, logEventFunc)
	case customtypes.ENUM_OUTPUT_FORMAT_JSON:
		printJson(message, fields, logEventFunc)
	default:
		l := logger.Get()
		printText(fmt.Sprintf("Output format %q is not recognized. Defaulting to \"text\" output", outputFormat), nil, yellow, l.Warn)
		printText(message, fields, colorFunc, logEventFunc)
	}

}

func printText(message string, fields map[string]interface{}, colorFunc func(format string, a ...interface{}) string,
	logEventFunc func() *zerolog.Event) {
	l := logger.Get()

	for k, v := range fields {
		switch typedValue := v.(type) {
		// If the value is a json.RawMessage, print it as a string
		case json.RawMessage:
			fmt.Println(cyan("%s: \n%s", k, typedValue))
			l.Info().Msg(cyan("%s: %s", k, typedValue))
		default:
			fmt.Println(cyan("%s: \n%v", k, v))
			l.Info().Msg(cyan("%s: %v", k, v))
		}
	}

	fmt.Println(colorFunc(message))
	logEventFunc().Msg(colorFunc(message))
}

func printJson(message string, fields map[string]interface{}, logEventFunc func() *zerolog.Event) {
	l := logger.Get()

	if fields == nil {
		fields = map[string]interface{}{
			"message": message,
		}
	}

	if fields["message"] == nil {
		fields["message"] = message
	}

	// Convert the CommandOutput struct to JSON
	jsonOut, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		l.Error().Err(err).Msgf("Failed to serialize output as JSON")
		return
	}

	// Output the JSON as uncolored string
	fmt.Println(string(jsonOut))
	logEventFunc().Msg(string(jsonOut))
}
