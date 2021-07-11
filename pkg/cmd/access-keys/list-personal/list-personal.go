package list_personal

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAccessKeysListPersonal(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-personal",
		Short: "List all access keys that belong to your user.",
		Run: func(cmd *cobra.Command, args []string) {
			listPersonalAccessKeys(client, log)
		},
	}
	return cmd
}

func listPersonalAccessKeys(client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.ListPersonalAccessKeys()
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to list personal access keys")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
