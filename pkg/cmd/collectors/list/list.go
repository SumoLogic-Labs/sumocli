package list

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"net/url"
)

func NewCmdCollectorList() *cobra.Command {
	var (
		filter  string
		limit   int
		offset  int
		offline bool
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic collectors",
		Run: func(cmd *cobra.Command, args []string) {
			log := logging.GetConsoleLogger()
			listCollectors(filter, limit, offset, offline, log)
		},
	}

	cmd.Flags().StringVar(&filter, "filter", "", "Filters the collectors returned using either installed, hosted, dead or alive")
	cmd.Flags().IntVar(&limit, "limit", 1000, "Maximum number of collectors returned")
	cmd.Flags().IntVar(&offset, "offset", 0, "Offset into the list of collectors")
	cmd.Flags().BoolVar(&offline, "offline", false, "Lists offline collectors")
	return cmd
}

func listCollectors(filter string, limit int, offset int, offline bool, log zerolog.Logger) {
	var requestUrl string
	if offline == true {
		requestUrl = "v1/collectors/offline"
	} else {
		requestUrl = "v1/collectors"
	}

	client, request := factory.NewHttpRequest("GET", requestUrl)
	query := url.Values{}
	if filter != "" {
		if factory.ValidateCollectorFilter(filter) == false {
			log.Fatal().Msg(filter + "is an invalid field to filter by. Available fields are installed, hosted, dead or alive.")
		} else {
			query.Add("filter", filter)
		}
	}
}
