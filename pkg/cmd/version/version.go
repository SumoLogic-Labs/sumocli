package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/internal/build"
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
