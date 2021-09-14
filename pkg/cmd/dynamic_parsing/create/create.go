package create

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdDynamicParsingCreate(client *cip.APIClient) *cobra.Command {
	var (
		name    string
		scope   string
		enabled bool
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new dynamic parsing rule.",
		Run: func(cmd *cobra.Command, args []string) {
			createDynamicParsingRule(name, scope, enabled, client)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the dynamic parsing rule")
	cmd.Flags().StringVar(&scope, "scope", "", "Scope of the dynamic parsing rule. "+
		"This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from.")
	cmd.Flags().BoolVar(&enabled, "enabled", true, "Set to false if you don't want to enable the rule")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("scope")
	return cmd
}

func createDynamicParsingRule(name string, scope string, enabled bool, client *cip.APIClient) {
	data, response, err := client.CreateDynamicParsingRule(types.DynamicRuleDefinition{
		Name:    name,
		Scope:   scope,
		Enabled: enabled,
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
