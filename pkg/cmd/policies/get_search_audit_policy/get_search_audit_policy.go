package get_search_audit_policy

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdGetSearchAuditPolicy(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-search-audit-policy",
		Short: "Get the Search Audit policy.",
		Long: "Get the Search Audit policy. This policy specifies whether search records for your account are enabled. " +
			"You can access details about your account's search capacity, queries run by users from the Sumo Search Audit Index.",
		Run: func(cmd *cobra.Command, args []string) {
			getSearchAuditPolicy(client)
		},
	}
	return cmd
}

func getSearchAuditPolicy(client *cip.APIClient) {
	data, response, err := client.GetSearchAuditPolicy()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
