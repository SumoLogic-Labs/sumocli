package set_max_user_session_timeout_policy

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdSetMaxUserSessionTimeoutPolicy(client *cip.APIClient) *cobra.Command {
	var maxUserSessionTimeout string
	cmd := &cobra.Command{
		Use:   "set-max-user-session-timeout-policy",
		Short: "Set the Max User Session Timeout policy.",
		Long: "Set the Max User Session Timeout policy. When enabled, this policy sets the maximum web session timeout users are able to configure within their user preferences. " +
			"Users preferences will be updated to match this value only if their current preference is set to a higher value.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			setMaxUserSessionTimeoutPolicy(client, maxUserSessionTimeout)
		},
	}
	cmd.Flags().StringVar(&maxUserSessionTimeout, "sessionTimeout", "", "Specify the max user session timeout. "+
		"Valid values are: 5m, 15m, 30m, 1h, 2h, 6h, 12h, 1d, 2d, 3d, 5d, or 7d")
	cmd.MarkFlagRequired("sessionTimeout")
	return cmd
}

func setMaxUserSessionTimeoutPolicy(client *cip.APIClient, maxUserSessionTimeout string) {
	data, response, err := client.SetMaxUserSessionTimeoutPolicy(types.MaxUserSessionTimeoutPolicy{
		MaxUserSessionTimeout: maxUserSessionTimeout,
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
