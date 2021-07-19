package delete

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"strconv"
)

func NewCmdDeleteSource() *cobra.Command {
	var (
		collectorId int
		sourceId    int
	)

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes the specified Source from the specified Collector.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteSource(collectorId, sourceId)
		},
	}
	cmd.Flags().IntVar(&collectorId, "collectorId", 0, "Specify the collector id the source is associated with")
	cmd.Flags().IntVar(&sourceId, "sourceId", 0, "Specify the source Id to delete")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("sourceId")
	return cmd
}

func deleteSource(collectorId int, sourceId int) {
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/collectors/" + strconv.Itoa(collectorId) + "/sources/" + strconv.Itoa(sourceId)
	client, request := factory.NewHttpRequest("DELETE", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		fmt.Println("source with id " + strconv.Itoa(sourceId) + " failed to delete, please try again.")
	} else {
		fmt.Println("source with id " + strconv.Itoa(sourceId) + " has been deleted")
	}
}
