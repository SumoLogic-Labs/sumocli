package encryption

import (
	"github.com/rs/zerolog"
	"golang.org/x/sys/windows/registry"
)

func getMachineId(log zerolog.Logger) string {
	registryKey, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.QUERY_VALUE|registry.WOW64_64KEY)
	if err != nil {
		log.Error().Err(err).Msg("failed to open registry item")
	}
	defer registryKey.Close()
	machineId, _, err := registryKey.GetStringValue("MachineGuid")
	if err != nil {
		log.Error().Err(err).Msg("failed to get MachineGuid")
	}
	return machineId
}
