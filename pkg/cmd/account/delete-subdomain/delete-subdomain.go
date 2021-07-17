package delete_subdomain

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAccountDeleteSubdomain(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-subdomain",
		Short: "Delete the configured subdomain.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteSubdomain(client, log)
		},
	}
	return cmd
}

func deleteSubdomain(client *cip.APIClient, log *zerolog.Logger) {
	httpResponse, errorResponse := client.DeleteSubdomain()
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to delete subdomain")
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "The subdomain was successfully deleted.")
	}
}
