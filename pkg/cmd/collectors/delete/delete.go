package delete

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdCollectorDelete() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Sumo Logic collector",
		Run: func(cmd *cobra.Command, args []string) {
			deleteCollector(id)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the collector to delete")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteCollector(id string) {
	log := logging.GetConsoleLogger()
	requestUrl := "v1/collectors/" + id
	client, request := factory.NewHttpRequest("DELETE", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		log.Info().Msg("collector with id " + id + " has been deleted")
	} else {
		log.Error().Msg("collector with id " + id + " failed to delete, please try again.")
	}
}
