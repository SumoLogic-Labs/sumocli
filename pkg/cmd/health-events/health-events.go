package health_events

import (
	"github.com/spf13/cobra"
	NewCmdHealthEventsGet "github.com/wizedkyle/sumocli/pkg/cmd/health-events/get"
	NewCmdHealthEventsList "github.com/wizedkyle/sumocli/pkg/cmd/health-events/list"
)

func NewCmdHealthEvents() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "health-events",
		Short: "Manages health events",
		Long: "Health Events allow you to keep track of the health of your Collectors and Sources. " +
			"You can use them to find and investigate common errors and warnings that are known to cause collection issues.",
	}
	cmd.AddCommand(NewCmdHealthEventsGet.NewCmdHealthEventsGet())
	cmd.AddCommand(NewCmdHealthEventsList.NewCmdHealthEventsList())
	return cmd
}
