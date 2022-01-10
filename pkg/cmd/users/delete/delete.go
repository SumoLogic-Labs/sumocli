package delete

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdUserDelete(client *cip.APIClient) *cobra.Command {
	var (
		id         string
		transferTo string
	)
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Sumo Logic user",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			deleteUser(id, transferTo, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user to delete")
	cmd.Flags().StringVar(&transferTo, "transferTo", "", "Specify the id of the user to transfer data to")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteUser(id string, transferTo string, client *cip.APIClient) {
	var options types.DeleteUserOpts
	if transferTo != "" {
		options.TransferTo = optional.NewString(transferTo)
	}
	response, err := client.DeleteUser(id, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "User was deleted successfully.")
	}
}
