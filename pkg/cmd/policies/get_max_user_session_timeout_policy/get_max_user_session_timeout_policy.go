package get_max_user_session_timeout_policy

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdGetMaxUserSessionTimeoutPolicy(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-max-user-session-timeout-policy",
		Short: "Get the Max User Session Timeout policy.",
		Long: "Get the Max User Session Timeout policy. When enabled, this policy sets the maximum web session timeout users are able to configure within their user preferences. " +
			"Users preferences will be updated to match this value only if their current preference is set to a higher value.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getMaxUserSessionTimeoutPolicy(client)
		},
	}
	return cmd
}

func getMaxUserSessionTimeoutPolicy(client *cip.APIClient) {
	data, response, err := client.GetMaxUserSessionTimeoutPolicy()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
