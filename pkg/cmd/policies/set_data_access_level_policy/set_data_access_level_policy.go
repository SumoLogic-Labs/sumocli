package set_data_access_level_policy

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdSetDataAccessLevelPolicy(client *cip.APIClient) *cobra.Command {
	var enabled bool
	cmd := &cobra.Command{
		Use:   "set-data-access-level-policy",
		Short: "Set the Data Access Level policy.",
		Long: "Set the Data Access Level policy. When enabled, this policy sets the default data access level for all newly created dashboards to the viewer’s role access filter. " +
			"Otherwise, newly created dashboards will default to the sharer’s role access filter and might display data that viewers’ roles don’t allow them to view.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			setDataAccessLevelPolicy(client, enabled)
		},
	}
	cmd.Flags().BoolVar(&enabled, "enabled", true, "If set to false the audit policy is disabled.")
	return cmd
}

func setDataAccessLevelPolicy(client *cip.APIClient, enabled bool) {
	data, response, err := client.SetDataAccessLevelPolicy(types.DataAccessLevelPolicy{
		enabled,
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
