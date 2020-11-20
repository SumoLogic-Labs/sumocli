package disable

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdUserDisableMFA() *cobra.Command {
	var (
		id       string
		email    string
		password string
	)

	cmd := &cobra.Command{
		Use:   "disable mfa",
		Short: "Disables MFA for a Sumo Logic user.",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("User disable mfa request started.")
			userDisableMFA(id, email, password, logger)
			logger.Debug().Msg("User disable mfa request finished.")
		},
	}

	return cmd
}

func userDisableMFA(id string, email string, password string, logger zerolog.Logger) {

}
