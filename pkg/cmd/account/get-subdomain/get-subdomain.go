package get_subdomain

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAccountGetSubdomain(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-subdomain",
		Short: "Get the configured subdomain.",
		Run: func(cmd *cobra.Command, args []string) {
			getSubdomain(client, log)
		},
	}
	return cmd
}

func getSubdomain(client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.GetSubdomain()
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to get subdomain")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
