package get

import (
	"fmt"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
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
		data, response, err := client.GetCollectorById(id)
		if err != nil {
			cmdutils.OutputError(response, err)
		} else {
			cmdutils.Output(data, response, err, "")
		}
	} else if name != "" {
		data, response, err := client.GetCollectorByName(name)
		if err != nil {
			cmdutils.OutputError(response, err)
		} else {
			cmdutils.Output(data, response, err, "")
		}
	}
}
