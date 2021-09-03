package get

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdDashboardsGet(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a dashboard by the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getDashboards(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the dashboard")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getDashboards(id string, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.GetDashboard(id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
