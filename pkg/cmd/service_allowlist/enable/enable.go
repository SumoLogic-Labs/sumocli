package enable

import (
	"fmt"
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdServiceAllowListEnable(client *cip.APIClient) *cobra.Command {
	var (
		login   bool
		content bool
		both    bool
	)
	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable service allowlisting functionality for the organization.",
		Long: "The service allowlisting can be for 1. Login: If enabled, access to Sumo Logic is granted only to CIDRs/IP addresses that are allowlisted. " +
			"2. Content: If enabled, dashboards can be shared with users connecting from CIDRs/IP addresses that are allowlisted without logging in.",
		Run: func(cmd *cobra.Command, args []string) {
			enableServiceAllowlist(login, content, both, client)
		},
	}
	cmd.Flags().BoolVar(&login, "login", false, "Set to true if you want the allowlist to affect logins")
	cmd.Flags().BoolVar(&content, "content", false, "Set to true if you want the allowlist to affect content")
	cmd.Flags().BoolVar(&both, "both", false, "Set to true if you want the allowlist to affect both logins and content")
	return cmd
}

func enableServiceAllowlist(login bool, content bool, both bool, client *cip.APIClient) {
	var allowlistType string
	if login == true {
		allowlistType = "login"
	} else if content == true {
		allowlistType = "content"
	} else if both == true {
		allowlistType = "both"
	} else if login == true && content == true && both == true {
		fmt.Println("Please select either login, content, or both.")
		os.Exit(1)
	}
	httpResponse, errorResponse := client.EnableAllowlisting(allowlistType)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Service allowlisting was enabled successfully.")
	}
}
