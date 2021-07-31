package list_personal

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
