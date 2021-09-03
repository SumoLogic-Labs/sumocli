package health_events

import (
	NewCmdHealthEventsGet "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/health_events/get"
	NewCmdHealthEventsList "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/health_events/list"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdHealthEvents(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "health-events",
		Short: "Manages health events",
		Long: "Health Events allow you to keep track of the health of your Collectors and Sources. " +
			"You can use them to find and investigate common errors and warnings that are known to cause collection issues.",
	}
	cmd.AddCommand(NewCmdHealthEventsGet.NewCmdHealthEventsGet(client))
	cmd.AddCommand(NewCmdHealthEventsList.NewCmdHealthEventsList(client))
	return cmd
}
