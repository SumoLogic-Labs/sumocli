package logging

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"runtime"
	"strings"
	"time"
)

func LogError(err error, log zerolog.Logger) {
	if err != nil {
		log.Error().Err(err)
	}
}

func GetConsoleLogger() zerolog.Logger {
	useColour := true
	if runtime.GOOS == "windows" {
		useColour = false
	}
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		NoColor:    useColour,
	}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("|  %-6s|", i))
	}
	log := zerolog.New(output).With().Timestamp().Logger()
	return log
}
