package metrics_query

import (
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdMetricsQuery(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "metrics-query",
		Short: "Execute queries on metrics",
		Long:  "Execute queries on various metrics and retrieve multiple time-series (data-points) over time range(s).",
	}
	return cmd
}
