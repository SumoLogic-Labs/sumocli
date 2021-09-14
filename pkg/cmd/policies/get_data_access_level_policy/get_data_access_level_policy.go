package get_data_access_level_policy

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdGetDataAccessLevelPolicy(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-data-access-level-policy",
		Short: "Get the Data Access Level policy.",
		Long: "Get the Data Access Level policy. When enabled, this policy sets the default data access level for all newly created dashboards to the viewer’s role access filter. " +
			"Otherwise, newly created dashboards will default to the sharer’s role access filter and might display data that viewers’ roles don’t allow them to view.",
		Run: func(cmd *cobra.Command, args []string) {
			getDataAccessLevelPolicy(client)
		},
	}
	return cmd
}

func getDataAccessLevelPolicy(client *cip.APIClient) {
	data, response, err := client.GetDataAccessLevelPolicy()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
