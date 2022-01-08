package create_subdomain

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdAccountCreateSubdomain(client *cip.APIClient) *cobra.Command {
	var subdomain string
	cmd := &cobra.Command{
		Use:   "create-subdomain",
		Short: "Create a subdomain. Only the Account Owner can create a subdomain.",
		Run: func(cmd *cobra.Command, args []string) {
			createSubdomain(subdomain, client)
		},
	}
	cmd.Flags().StringVar(&subdomain, "subdomain", "", "Specify a subdomain (minimum 4 and maximum 63 characters)")
	cmd.MarkFlagRequired("subdomain")
	return cmd
}

func createSubdomain(subdomain string, client *cip.APIClient) {
	data, response, err := client.CreateSubdomain(types.ConfigureSubdomainRequest{
		Subdomain: subdomain,
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
