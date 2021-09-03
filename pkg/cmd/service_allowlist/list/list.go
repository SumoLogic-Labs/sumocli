package list

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
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
	apiResponse, httpResponse, errorResponse := client.ListAllowlistedCidrs()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
