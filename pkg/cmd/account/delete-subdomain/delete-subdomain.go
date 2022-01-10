package delete_subdomain

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdAccountDeleteSubdomain(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-subdomain",
		Short: "Delete the configured subdomain.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			deleteSubdomain(client)
		},
	}
	return cmd
}

func deleteSubdomain(client *cip.APIClient) {
	response, err := client.DeleteSubdomain()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "The subdomain was successfully deleted.")
	}
}
