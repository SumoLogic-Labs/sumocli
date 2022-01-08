package get

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdTransformationRulesGet(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a transformation rule with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			get(client, id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the identifier of the transformation rule to retrieve.")
	cmd.MarkFlagRequired("id")
	return cmd
}

func get(client *cip.APIClient, id string) {
	data, response, err := client.GetTransformationRule(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
