package create

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdAccessKeysCreate(client *cip.APIClient) *cobra.Command {
	var (
		name        string
		corsHeaders []string
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new access ID and key pair. The new access key can be used from the domains specified in corsHeaders field.",
		Run: func(cmd *cobra.Command, args []string) {
			createAccessKey(name, corsHeaders, client)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the access key")
	cmd.Flags().StringSliceVar(&corsHeaders, "corsHeaders", []string{}, "Specify cors headers (they need to be comma separated e.g. header1,header2,header3 and be valid URLs e.g. https://test.com)")
	cmd.MarkFlagRequired("name")
	return cmd
}

func createAccessKey(name string, corsHeaders []string, client *cip.APIClient) {
	data, response, err := client.CreateAccessKey(types.AccessKeyCreateRequest{
		Label:       name,
		CorsHeaders: corsHeaders,
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
