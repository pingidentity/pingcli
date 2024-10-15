package logger

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	once   sync.Once
	logger zerolog.Logger
)

// Create a get function for a standardized zerolog logger
func Get() zerolog.Logger {
	once.Do(func() {

		// Viper config is not initialized yet, so read environment variables directly
		logLevelEnv := os.Getenv("PINGCLI_LOG_LEVEL")
		logPathEnv := os.Getenv("PINGCLI_LOG_PATH")

		var logLevel zerolog.Level

		switch logLevelEnv {
		case "PANIC":
			logLevel = zerolog.PanicLevel
		case "FATAL":
			logLevel = zerolog.FatalLevel
		case "ERROR":
			logLevel = zerolog.ErrorLevel
		case "WARN":
			logLevel = zerolog.WarnLevel
		case "INFO":
			logLevel = zerolog.InfoLevel
		case "DEBUG":
			logLevel = zerolog.DebugLevel
		case "TRACE":
			logLevel = zerolog.TraceLevel
		case "NOLEVEL":
			logLevel = zerolog.NoLevel
		default:
			logLevel = zerolog.Disabled
		}

		var output io.Writer

		// Handle log file creation when PINGCLI_LOG_PATH is defined
		if logPathEnv != "" && logLevel != zerolog.Disabled {

			var err error
			output, err = os.Create(logPathEnv)
			if err != nil {
				// Most likely the directory specified for the log path does not exist
				log.Fatal().Err(err).Msgf("Unable to create log file at: %s", logPathEnv)
			}

		} else {
			output = zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC3339,
			}
		}

		logger = zerolog.New(output).
			Level(logLevel).
			With().
			Timestamp().
			Logger()
	})
	return logger
}
