package create

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdCreateSource() *cobra.Command {
	var collectorId int
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a source on the specified Sumo Logic collector",
		Run: func(cmd *cobra.Command, args []string) {
			log := logging.GetConsoleLogger()
			CreateSource(log)
		},
	}

	cmd.Flags().IntVar(&collectorId, "collectorId", 0, "Specify the collector id to attach the source to")
	cmd.MarkFlagRequired("collectorId")
	return cmd
}

func CreateSource(log zerolog.Logger) string {
	var createSourceResponse api.SourcesResponse

	requestBodySchema := &api.CreateSourcesRequest{
		SourceType: "",
		Name:       "",
		Interval:   0,
		HostName:   "",
		Metrics:    nil,
	}
}
