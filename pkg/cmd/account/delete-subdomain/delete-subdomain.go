package delete_subdomain

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdAccountDeleteSubdomain(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-subdomain",
		Short: "Delete the configured subdomain.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteSubdomain(client)
		},
	}
	return cmd
}

func deleteSubdomain(client *cip.APIClient) {
	httpResponse, errorResponse := client.DeleteSubdomain()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "The subdomain was successfully deleted.")
	}
}
