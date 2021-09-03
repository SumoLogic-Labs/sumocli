package get

import (
	"fmt"
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdCollectorGet(client *cip.APIClient) *cobra.Command {
	var (
		id   string
		name string
	)
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic collector information",
		Long:  "You can use either the id or the name of the collector to specify the collector to return",
		Run: func(cmd *cobra.Command, args []string) {
			getCollector(id, name, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the collector")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the collector")
	return cmd
}

func getCollector(id string, name string, client *cip.APIClient) {
	if id != "" && name != "" {
		fmt.Println("Please specify and id or name, not both.")
	} else if id != "" {
		apiResponse, httpResponse, errorResponse := client.GetCollectorById(id)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse, errorResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	} else if name != "" {
		apiResponse, httpResponse, errorResponse := client.GetCollectorByName(name)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse, errorResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	}
}
