package update_subdomain

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdAccountUpdateSubdomain(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var subdomain string
	cmd := &cobra.Command{
		Use:   "update-subdomain",
		Short: "Update a subdomain. Only the Account Owner can update the subdomain.",
		Run: func(cmd *cobra.Command, args []string) {
			updateSubdomain(subdomain, client, log)
		},
	}
	cmd.Flags().StringVar(&subdomain, "subdomain", "", "Specify a new subdomain (minimum 4 and maximum 63 characters)")
	cmd.MarkFlagRequired("subdomain")
	return cmd
}

func updateSubdomain(subdomain string, client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.UpdateSubdomain(types.ConfigureSubdomainRequest{
		Subdomain: subdomain,
	})
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to update subdomain")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
