package update

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
	"os"
)

func NewCmdDashboardsUpdate(client *cip.APIClient) *cobra.Command {
	var (
		id   string
		file string
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a dashboard by the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			updateDashboard(file, id, client)
		},
	}
	cmd.Flags().StringVar(&file, "file", "", "Specify the full file path to a json file containing a dashboard definition."+
		"The definition can be retrieved from running sumocli dashboards get or from exporting the dashboard in the UI."+
		"If you have exported the dashboard definition you may need to modify the panel ids before updating.")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the dashboard to update")
	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("id")
	return cmd
}

func updateDashboard(file string, id string, client *cip.APIClient) {
	var dashboardDefinition types.DashboardRequest
	fileData, err := os.ReadFile(file)
	if err != nil {
		log.Error().Err(err).Msg("failed to read file " + file)
	}
	err = json.Unmarshal(fileData, &dashboardDefinition)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal file")
	}
	apiResponse, httpResponse, errorResponse := client.UpdateDashboard(dashboardDefinition, id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
