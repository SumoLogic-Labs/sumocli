package status

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"net/url"
	"strconv"
)

func NewCmdUpgradableCollectorStatus() *cobra.Command {
	var upgradeTaskId int

	cmd := &cobra.Command{
		Use: "status",
		Long: `Gets the status of a collector upgrade or downgrade.
The status of the upgrade can be one of the following
0 - not started
1 - pending, the upgrade is issued waiting a response from the Collector
2 - succeeded
3 - failed
6 - progressing, the upgrade is running on the Collector`,
		Run: func(cmd *cobra.Command, args []string) {
			upgradableCollectorStatus(upgradeTaskId)
		},
	}
	cmd.Flags().IntVar(&upgradeTaskId, "upgradeTaskId", 0, "Id to the upgrade task")
	cmd.MarkFlagRequired("upgradeTaskId")
	return cmd
}

func upgradableCollectorStatus(upgradeTaskId int) {
	log := logging.GetConsoleLogger()
	var status api.UpgradeTaskStatus
	requestUrl := "v1/collectors/upgrades/" + strconv.Itoa(upgradeTaskId)
	client, request := factory.NewHttpRequest("GET", requestUrl)
	query := url.Values{}
	query.Add("upgradeTaskId", strconv.Itoa(upgradeTaskId))
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("error reading body from request")
	}

	err = json.Unmarshal(responseBody, &status)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshalling response body")
	}
	if response.StatusCode == 200 {
		statusJson, err := json.MarshalIndent(status, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal status")
		}
		fmt.Println(string(statusJson))
	} else if response.StatusCode == 404 {
		fmt.Println("The upgrade task Id was not found")
	}

}
