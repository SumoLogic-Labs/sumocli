package list

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdServiceAllowlistList(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all allowlisted CIDR notations and/or IP addresses for the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			listServiceAllowlist(client)
		},
	}
	return cmd
}

func listServiceAllowlist(client *cip.APIClient) {
	data, response, err := client.ListAllowlistedCidrs()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
