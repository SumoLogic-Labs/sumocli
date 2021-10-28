package get

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdPermissionsGet(client *cip.APIClient) *cobra.Command {
	var (
		id           string
		explicitOnly bool
		isAdminMode  bool
	)
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Returns content permissions of a content item with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getPermissions(id, explicitOnly, isAdminMode, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of a content item")
	cmd.Flags().BoolVar(&explicitOnly, "explicitOnly", false, "There are two permission types: explicit and implicit. "+
		"Permissions specifically assigned to the content item are explicit. Permissions derived from a parent content item, like a folder are implicit. "+
		"To return only explicit permissions set this to true.")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getPermissions(id string, explicitOnly bool, isAdminMode bool, client *cip.APIClient) {
	var options *types.GetContentPermissionsOpts
	options.ExplicitOnly = optional.NewBool(explicitOnly)
	options.IsAdminMode = optional.NewString(cmdutils.AdminMode(isAdminMode))
	data, response, err := client.GetContentPermissions(id, options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
