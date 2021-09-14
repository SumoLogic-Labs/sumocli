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
	data, response, err := client.ListPersonalAccessKeys()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
