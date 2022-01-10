package create

import (
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdCollectorCreate(client *cip.APIClient) *cobra.Command {
	var (
		name        string
		description string
		category    string
		fieldNames  []string
		fieldValues []string
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a Sumo Logic Hosted Collector",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			Collector(name, description, category, fieldNames, fieldValues, client)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the collector")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the collector")
	cmd.Flags().StringVar(&category, "category", "", "sourceCategory for the collector, this will overwrite the categories on configured sources")
	cmd.Flags().StringSliceVar(&fieldNames, "fieldNames", []string{}, "Specify the names of fields to add to the collector "+
		"{names need to be comma separated e.g. field1,field2")
	cmd.Flags().StringSliceVar(&fieldValues, "fieldValues", []string{}, "Specify the values of fields to add to the collector "+
		"(values need to be comma separated e.g. value1,value2")

	cmd.MarkFlagRequired("name")
	return cmd
}

func Collector(name string, description string, category string, fieldNames []string, fieldValues []string, client *cip.APIClient) {
	fields := cmdutils.GenerateFieldsMap(fieldNames, fieldValues)
	data, response, err := client.CreateCollector(types.CreateCollectorRequest{
		Collector: types.CreateCollectorRequestDefinition{
			CollectorType: "Hosted",
			Name:          name,
			Description:   description,
			Category:      category,
			Fields:        fields,
		},
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
