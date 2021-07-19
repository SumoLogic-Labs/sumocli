package start

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"strings"
)

func NewCmdUpgradeStart() *cobra.Command {
	var (
		collectorId int
		toVersion   string
	)

	cmd := &cobra.Command{
		Use:   "start",
		Short: "Starts an upgrade or downgrade of an existing installed collector",
		Run: func(cmd *cobra.Command, args []string) {
			upgradeStart(collectorId, toVersion)
		},
	}

	cmd.Flags().IntVar(&collectorId, "collectorId", 0, "Id of the collector to upgrade")
	cmd.Flags().StringVar(&toVersion, "version", "", "Version to upgrade or downgrade the collector to")
	cmd.MarkFlagRequired("collectorId")
	return cmd
}

func upgradeStart(collectorId int, toVersion string) {
	log := logging.GetConsoleLogger()
	var upgradeInfo api.UpgradeTaskResponse
	requestBodySchema := &api.UpgradeTask{
		CollectorId: collectorId,
		ToVersion:   toVersion,
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response body")
	}

	requestUrl := "/v1/collectors/upgrades"
	client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("error reading response body from request")
	}

	err = json.Unmarshal(responseBody, &upgradeInfo)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshalling response body")
	}
	if response.StatusCode == 202 {
		id := strings.Split(upgradeInfo.Link.Href, "v1/collectors/upgrades/")
		fmt.Println(id[1])
	} else if response.StatusCode == 400 {
		log.Error().Msg("Bad request please try again")
	}
}
