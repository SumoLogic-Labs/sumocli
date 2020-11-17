package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"time"
)

func LogError(err error, log zerolog.Logger) {
	if err != nil {
		log.Error().Err(err)
	}
}

func LogErrorWithMessage(msg string, err error, log zerolog.Logger) {
	if err != nil {
		log.Error().Err(err).Msg(msg)
	}
}

func GetLoggerForCommand(command *cobra.Command) zerolog.Logger {
	verbose, _ := command.Root().PersistentFlags().GetBool("verbose")
	suppressLogging, _ := command.Root().PersistentFlags().GetBool("quiet")

	if suppressLogging {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	} else {
		if verbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	}
	useColour := true
	if runtime.GOOS == "windows" {
		useColour = false
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		NoColor:    useColour,
	}).With().Caller().Str("command", command.Name()).Logger()
	return log.Logger
}
