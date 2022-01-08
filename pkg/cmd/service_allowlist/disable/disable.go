package disable

import (
	"fmt"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdServiceAllowlistDisable(client *cip.APIClient) *cobra.Command {
	var (
		login   bool
		content bool
		both    bool
	)
	cmd := &cobra.Command{
		Use:   "disable",
		Short: "Disable service allowlisting functionality for login/API authentication or content sharing for the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			disableServiceAllowlist(login, content, both, client)
		},
	}
	cmd.Flags().BoolVar(&login, "login", false, "Set to true if you want the allowlist to affect logins")
	cmd.Flags().BoolVar(&content, "content", false, "Set to true if you want the allowlist to affect content")
	cmd.Flags().BoolVar(&both, "both", false, "Set to true if you want the allowlist to affect both logins and content")
	return cmd
}

func disableServiceAllowlist(login bool, content bool, both bool, client *cip.APIClient) {
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
	response, err := client.DisableAllowlisting(allowlistType)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "Service allowlisting was disabled successfully.")
	}
}
