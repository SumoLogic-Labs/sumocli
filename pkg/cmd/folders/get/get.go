package get

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdGet(client *cip.APIClient) *cobra.Command {
	var (
		id          string
		isAdminMode bool
	)
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a folder with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			get(id, isAdminMode, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the identifier of the folder")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("id")
	return cmd
}

func get(id string, isAdminMode bool, client *cip.APIClient) {
	adminMode := cmdutils.AdminMode(isAdminMode)
	apiResponse, httpResponse, errorResponse := client.GetFolder(id, &types.FolderOpts{
		IsAdminMode: optional.NewString(adminMode),
	})
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
