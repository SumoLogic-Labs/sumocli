package get_user_concurrent_sessions_limit_policy

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdGetUserConcurrentSessionsLimitPolicy(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-user-concurrent-sessions-limit-policy",
		Short: "Get the User Concurrent Sessions Limit policy.",
		Long: "Get the User Concurrent Sessions Limit policy. When enabled, the number of concurrent sessions a user may have is limited to the value entered. " +
			"If a user exceeds the allowed number of sessions, the user's oldest session will be logged out to accommodate the new one. Disabling this policy means a user may have an unlimited number of concurrent sessions.",
		Run: func(cmd *cobra.Command, args []string) {
			getUserConcurrentSessionsLimitPolicy(client)
		},
	}
	return cmd
}

func getUserConcurrentSessionsLimitPolicy(client *cip.APIClient) {
	data, response, err := client.GetUserConcurrentSessionsLimitPolicy()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
