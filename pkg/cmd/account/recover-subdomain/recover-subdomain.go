package recover_subdomain

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdAccountRecoverSubdomain(client *cip.APIClient) *cobra.Command {
	var email string
	cmd := &cobra.Command{
		Use:   "recover-subdomain",
		Short: "Send an email with the subdomain information for a user with the given email address.",
		Run: func(cmd *cobra.Command, args []string) {
			recoverSubdomain(email, client)
		},
	}
	cmd.Flags().StringVar(&email, "email", "", "Specify an email address of the user to get subdomain information")
	cmd.MarkFlagRequired("email")
	return cmd
}

func recoverSubdomain(email string, client *cip.APIClient) {
	httpResponse, errorResponse := client.RecoverSubdomains(email)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "An email containing information about associated subdomains for the given email was sent.")
	}
}
