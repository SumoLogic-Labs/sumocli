package create

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdFieldExtractionRulesCreate(client *cip.APIClient) *cobra.Command {
	var (
		name            string
		scope           string
		parseExpression string
		enabled         bool
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new field extraction rule.",
		Run: func(cmd *cobra.Command, args []string) {
			createFieldExtractionRule(name, scope, parseExpression, enabled, client)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the field extraction rule")
	cmd.Flags().StringVar(&scope, "scope", "", "Scope of the field extraction rule. "+
		"This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from.")
	cmd.Flags().StringVar(&parseExpression, "parseExpression", "", "Specify the fields to be parsed")
	cmd.Flags().BoolVar(&enabled, "enabled", true, "Set to false if you don't want to enable the rule")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("scope")
	cmd.MarkFlagRequired("parseExpression")
	return cmd
}

func createFieldExtractionRule(name string, scope string, parseExpression string, enabled bool, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.CreateExtractionRule(types.ExtractionRuleDefinition{
		Name:            name,
		Scope:           scope,
		ParseExpression: parseExpression,
		Enabled:         enabled,
	})
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
