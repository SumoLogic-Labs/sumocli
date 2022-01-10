package version

import (
	"fmt"
	"github.com/SumoLogic-Labs/sumocli/internal/build"
	"github.com/spf13/cobra"
)

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Displays sumocli version",
		Long:  "Displays the version and build number of sumocli.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Sumocli " + build.Version + " " + build.Build + " " + build.Date)
		},
	}
	return cmd
}
