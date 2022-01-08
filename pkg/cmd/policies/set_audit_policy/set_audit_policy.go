package set_audit_policy

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdSetAuditPolicy(client *cip.APIClient) *cobra.Command {
	var enabled bool
	cmd := &cobra.Command{
		Use:   "set-audit-policy",
		Short: "Set the Audit policy.",
		Long: "Set the Audit policy. This policy specifies whether audit records for your account are enabled. " +
			"You can access details about reported account events in the Sumo Logic Audit Index.",
		Run: func(cmd *cobra.Command, args []string) {
			setAuditPolicy(client, enabled)
		},
	}
	cmd.Flags().BoolVar(&enabled, "enabled", true, "If set to false the audit policy is disabled.")
	return cmd
}

func setAuditPolicy(client *cip.APIClient, enabled bool) {
	data, response, err := client.SetAuditPolicy(types.AuditPolicy{
		enabled,
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
