package recover_subdomain

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAccountRecoverSubdomain(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var email string
	cmd := &cobra.Command{
		Use:   "recover-subdomain",
		Short: "Send an email with the subdomain information for a user with the given email address.",
		Run: func(cmd *cobra.Command, args []string) {
			recoverSubdomain(email, client, log)
		},
	}
	cmd.Flags().StringVar(&email, "email", "", "Specify an email address of the user to get subdomain information")
	cmd.MarkFlagRequired("email")
	return cmd
}

func recoverSubdomain(email string, client *cip.APIClient, log *zerolog.Logger) {
	httpResponse, errorResponse := client.RecoverSubdomains(email)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to recover subdomain")
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "An email containing information about associated subdomains for the given email was sent.")
	}
}
