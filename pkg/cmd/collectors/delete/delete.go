package delete

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"net/url"
	"strconv"
)

func NewCmdCollectorDelete() *cobra.Command {
	var (
		aliveBeforeDays int
		force           bool
		id              string
		offline         bool
	)

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Sumo Logic collector",
		Run: func(cmd *cobra.Command, args []string) {
			deleteCollector(aliveBeforeDays, force, id, offline)
		},
	}

	cmd.Flags().IntVar(&aliveBeforeDays, "aliveBeforeDays", 100, "Minimum number of days the collectors have been offline")
	cmd.Flags().BoolVar(&force, "force", false, "Forces removal of offline collectors useful for CI/CD pipelines")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the collector to delete")
	cmd.Flags().BoolVar(&offline, "offline", false, "Removes all offline collectors")
	return cmd
}

func deleteCollector(aliveBeforeDays int, force bool, id string, offline bool) {
	log := logging.GetConsoleLogger()
	if offline == true {
		if force == true {
			requestUrl := "v1/collectors/offline"
			client, request := factory.NewHttpRequest("DELETE", requestUrl)
			query := url.Values{}
			query.Add("aliveBeforeDays", strconv.Itoa(aliveBeforeDays))
			request.URL.RawQuery = query.Encode()
			response, err := client.Do(request)
			if err != nil {
				log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
			}
			defer response.Body.Close()
			if response.StatusCode == 200 {
				log.Info().Msg("offline collectors have been deleted")
			} else {
				log.Error().Msg("failed to delete offline collectors")
			}
		} else if force == false {
			promptConfirm := promptui.Prompt{
				Label:     "Confirm that you want to delete all offline collectors?",
				IsConfirm: true,
			}
			_, err := promptConfirm.Run()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to perform confirmation")
			}

			requestUrl := "v1/collectors/offline"
			client, request := factory.NewHttpRequest("DELETE", requestUrl)
			query := url.Values{}
			query.Add("aliveBeforeDays", strconv.Itoa(aliveBeforeDays))
			request.URL.RawQuery = query.Encode()
			response, err := client.Do(request)
			if err != nil {
				log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
			}
			defer response.Body.Close()
			if response.StatusCode == 200 {
				log.Info().Msg("offline collectors have been deleted")
			} else {
				log.Error().Msg("failed to delete offline collectors")
			}
		}
	} else if offline == false {
		if id == "" {
			log.Fatal().Msg("a valid collector id needs to be provided, you can run sumocli collectors list to get ids")
		}

		requestUrl := "v1/collectors/" + id
		client, request := factory.NewHttpRequest("DELETE", requestUrl)
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
		}
		defer response.Body.Close()
		if response.StatusCode == 200 {
			fmt.Println("collector with id " + id + " has been deleted")
		} else {
			fmt.Println("collector with id " + id + " failed to delete, please try again.")
		}
	}
}
