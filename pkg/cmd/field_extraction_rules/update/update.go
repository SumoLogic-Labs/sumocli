package update

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdFieldExtractionRulesUpdate(client *cip.APIClient) *cobra.Command {
	var (
		id              string
		name            string
		scope           string
		parseExpression string
		enabled         bool
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing field extraction rule.",
		Run: func(cmd *cobra.Command, args []string) {
			updateFieldExtractionRule(id, name, scope, parseExpression, enabled, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the field extraction rule")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the field extraction rule")
	cmd.Flags().StringVar(&scope, "scope", "", "Scope of the field extraction rule. "+
		"This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from.")
	cmd.Flags().StringVar(&parseExpression, "parseExpression", "", "Specify the fields to be parsed")
	cmd.Flags().BoolVar(&enabled, "enabled", true, "Set to false if you don't want to enable the rule")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("scope")
	cmd.MarkFlagRequired("parseExpression")
	return cmd
}

func updateFieldExtractionRule(id string, name string, scope string, parseExpression string, enabled bool, client *cip.APIClient) {
	data, response, err := client.UpdateExtractionRule(types.UpdateExtractionRuleDefinition{
		Name:            name,
		Scope:           scope,
		ParseExpression: parseExpression,
		Enabled:         enabled,
	},
		id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
