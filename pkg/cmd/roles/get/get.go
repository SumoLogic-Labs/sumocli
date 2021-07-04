package get

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdRoleGet(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic role information",
		Run: func(cmd *cobra.Command, args []string) {
			getRole(client, id, log)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to get")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getRole(client *cip.APIClient, id string, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.GetRole(id)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to get role")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
