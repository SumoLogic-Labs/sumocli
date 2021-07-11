package update

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdAccessKeysUpdate(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var (
		corsHeaders []string
		disabled    bool
		id          string
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates the properties of existing accessKey by accessId. It can be used to enable or disable the access key and to update the corsHeaders list.",
		Run: func(cmd *cobra.Command, args []string) {
			updateAccessKey(corsHeaders, disabled, id, client, log)
		},
	}
	cmd.Flags().StringSliceVar(&corsHeaders, "corsHeaders", []string{}, "Specify cors headers (they need to be comma separated e.g. header1,header2,header3 and be valid URLs e.g. https://test.com)")
	cmd.Flags().BoolVar(&disabled, "disabled", false, "Set to true to disable the access key")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the access key to update")
	cmd.MarkFlagRequired("id")
	return cmd
}

func updateAccessKey(corsHeaders []string, disabled bool, id string, client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.UpdateAccessKey(types.AccessKeyUpdateRequest{
		Disabled:    disabled,
		CorsHeaders: corsHeaders,
	},
		id)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to update access key")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
