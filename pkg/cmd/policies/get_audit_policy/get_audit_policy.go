package get_audit_policy

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdGetAuditPolicy(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-audit-policy",
		Short: "Get the Audit policy.",
		Long:  "Get the Audit policy. This policy specifies whether audit records for your account are enabled. You can access details about reported account events in the Sumo Logic Audit Index.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getAuditPolicy(client)
		},
	}
	return cmd
}

func getAuditPolicy(client *cip.APIClient) {
	data, response, err := client.GetAuditPolicy()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
