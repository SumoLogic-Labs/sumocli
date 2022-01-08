package install

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdAppsInstall(client *cip.APIClient) *cobra.Command {
	var (
		destinationFolderId string
		description         string
		logSource           string
		name                string
		uuid                string
	)
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Installs the app with given UUID in the folder specified.",
		Run: func(cmd *cobra.Command, args []string) {
			installApp(destinationFolderId, description, logSource, name, uuid, client)
		},
	}
	cmd.Flags().StringVar(&destinationFolderId, "destinationFolderId", "", "Specify the folder id that the app should be installed into")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the app")
	cmd.Flags().StringVar(&logSource, "logSource", "", "Specify a log source name (for example _sourceCategory=test)")
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the app")
	cmd.Flags().StringVar(&uuid, "uuid", "", "Specify the UUID of the app to install")
	cmd.MarkFlagRequired("destinationFolderId")
	cmd.MarkFlagRequired("description")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("uuid")
	return cmd
}

func installApp(destinationFolderId string, description string, logSource string, name string, uuid string, client *cip.APIClient) {
	data, response, err := client.InstallApp(types.AppInstallRequest{
		Name:                name,
		Description:         description,
		DestinationFolderId: destinationFolderId,
		DataSourceValues: map[string]string{
			"logsrc": logSource,
		},
	},
		uuid)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
