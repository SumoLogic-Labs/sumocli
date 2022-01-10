package get_subdomain

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdAccountGetSubdomain(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-subdomain",
		Short: "Get the configured subdomain.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getSubdomain(client)
		},
	}
	return cmd
}

func getSubdomain(client *cip.APIClient) {
	data, response, err := client.GetSubdomain()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
