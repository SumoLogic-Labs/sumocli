package encryption

import (
	"github.com/rs/zerolog"
	"os"
	"strings"
)

const (
	dbusPath    = "/var/lib/dbus/machine-id"
	dbusPathEtc = "/var/machine-id"
)

func getMachineId(log zerolog.Logger) string {
	machineId, err := os.ReadFile(dbusPath)
	if err != nil {
		machineId, err = os.ReadFile(dbusPathEtc)
	}
	if err != nil {
		log.Error().Err(err).Msg("failed to retrieve machine id")
	}
	return trim(string(machineId))
}

func trim(s string) string {
	return strings.TrimSpace(strings.Trim(s, "\n"))
}
