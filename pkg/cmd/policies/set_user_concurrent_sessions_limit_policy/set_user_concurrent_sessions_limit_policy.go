package set_user_concurrent_sessions_limit_policy

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdSetUserConcurrentSessionsLimitPolicy(client *cip.APIClient) *cobra.Command {
	var (
		enabled               bool
		maxConcurrentSessions int32
	)
	cmd := &cobra.Command{
		Use:   "set-user-concurrent-sessions-limit-policy",
		Short: "Set the User Concurrent Sessions Limit policy.",
		Long: "Set the User Concurrent Sessions Limit policy. When enabled, the number of concurrent sessions a user may have is limited to the value entered. If a user exceeds the allowed number of sessions, the user's oldest session will be logged out to accommodate the new one. " +
			"Disabling this policy means a user may have an unlimited number of concurrent sessions.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			setUserConcurrentSessionsLimitPolicy(client, enabled, maxConcurrentSessions)
		},
	}
	cmd.Flags().BoolVar(&enabled, "enabled", true, "If set to false the audit policy is disabled.")
	cmd.Flags().Int32Var(&maxConcurrentSessions, "concurrentSessions", 100, "Specify the number of concurrent sessions.")
	return cmd
}

func setUserConcurrentSessionsLimitPolicy(client *cip.APIClient, enabled bool, maxConcurrentSessions int32) {
	data, httpRespone, err := client.SetUserConcurrentSessionsLimitPolicy(types.UserConcurrentSessionsLimitPolicy{
		Enabled:               enabled,
		MaxConcurrentSessions: maxConcurrentSessions,
	})
	if err != nil {
		cmdutils.OutputError(httpRespone, err)
	} else {
		cmdutils.Output(data, httpRespone, err, "")
	}
}
