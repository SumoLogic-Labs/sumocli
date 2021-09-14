package set_search_audit_policy

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdSetSearchAuditPolicy(client *cip.APIClient) *cobra.Command {
	var enabled bool
	cmd := &cobra.Command{
		Use:   "set-search-audit-policy",
		Short: "Set the Search Audit policy.",
		Long: "Set the Search Audit policy. This policy specifies whether search records for your account are enabled. " +
			"You can access details about your account's search capacity, queries run by users from the Sumo Search Audit Index.",
		Run: func(cmd *cobra.Command, args []string) {
			setSearchAuditPolicy(client, enabled)
		},
	}
	cmd.Flags().BoolVar(&enabled, "enabled", true, "If set to false the audit policy is disabled.")
	return cmd
}

func setSearchAuditPolicy(client *cip.APIClient, enabled bool) {
	data, response, err := client.SetSearchAuditPolicy(types.SearchAuditPolicy{
		enabled,
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
