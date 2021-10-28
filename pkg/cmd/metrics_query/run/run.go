package run

import (
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdMetricsQueryRun(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}

func run(client *cip.APIClient) {
	data, response, err := client.RunMetricsQueries()
}
