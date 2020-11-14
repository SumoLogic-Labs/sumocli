package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/internal/build"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Short:  "Displays sumocli version",
		Long:   "Displays the version and build number of sumocli.",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Version command started.")
			fmt.Println("Sumocli " + build.Version + " " + build.Build + " " + build.Date)
			logger.Debug().Msg("Version command finished.")
		},
	}

	return cmd
}
