package update

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdAccessKeysUpdate(client *cip.APIClient) *cobra.Command {
	var (
		corsHeaders []string
		disabled    bool
		id          string
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates the properties of existing accessKey by accessId. It can be used to enable or disable the access key and to update the corsHeaders list.",
		Run: func(cmd *cobra.Command, args []string) {
			updateAccessKey(corsHeaders, disabled, id, client)
		},
	}
	cmd.Flags().StringSliceVar(&corsHeaders, "corsHeaders", []string{}, "Specify cors headers (they need to be comma separated e.g. header1,header2,header3 and be valid URLs e.g. https://test.com)")
	cmd.Flags().BoolVar(&disabled, "disabled", false, "Set to true to disable the access key")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the access key to update")
	cmd.MarkFlagRequired("id")
	return cmd
}

func updateAccessKey(corsHeaders []string, disabled bool, id string, client *cip.APIClient) {
	data, response, err := client.UpdateAccessKey(types.AccessKeyUpdateRequest{
		Disabled:    disabled,
		CorsHeaders: corsHeaders,
	},
		id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
