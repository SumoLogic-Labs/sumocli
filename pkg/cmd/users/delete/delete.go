package delete

import (
	"github.com/antihax/optional"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdUserDelete(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var (
		id         string
		transferTo string
	)
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Sumo Logic user",
		Run: func(cmd *cobra.Command, args []string) {
			deleteUser(id, transferTo, client, log)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user to delete")
	cmd.Flags().StringVar(&transferTo, "transferTo", "", "Specify the id of the user to transfer data to")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteUser(id string, transferTo string, client *cip.APIClient, log *zerolog.Logger) {
	var options types.UserManagementApiDeleteUserOpts
	if transferTo != "" {
		options.TransferTo = optional.NewString(transferTo)
	}
	httpResponse, errorResponse := client.DeleteUser(id, &options)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to delete user")
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "User was deleted successfully.")
	}
}
