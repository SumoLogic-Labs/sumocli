package update

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdDynamicParsingUpdate(client *cip.APIClient) *cobra.Command {
	var (
		id      string
		name    string
		scope   string
		enabled bool
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing dynamic parsing rule",
		Run: func(cmd *cobra.Command, args []string) {
			updateDynamicParsingRule(id, name, scope, enabled, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the dynamic parsing rule")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the dynamic parsing rule")
	cmd.Flags().StringVar(&scope, "scope", "", "Scope of the dynamic parsing rule. "+
		"This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from.")
	cmd.Flags().BoolVar(&enabled, "enabled", true, "Set to false if you don't want to enable the rule")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("scope")
	return cmd
}

func updateDynamicParsingRule(id string, name string, scope string, enabled bool, client *cip.APIClient) {
	apiResponse, httpResponse, errorResponse := client.UpdateDynamicParsingRule(types.DynamicRuleDefinition{
		Name:    name,
		Scope:   scope,
		Enabled: enabled,
	},
		id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
