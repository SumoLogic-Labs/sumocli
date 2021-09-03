package list_personal

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdAccessKeysListPersonal(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-personal",
		Short: "List all access keys that belong to your user.",
		Run: func(cmd *cobra.Command, args []string) {
			listPersonalAccessKeys(client)
		},
	}
	return cmd
}

func listPersonalAccessKeys(client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.ListPersonalAccessKeys()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
