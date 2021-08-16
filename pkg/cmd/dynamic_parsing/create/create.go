package create

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
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
	apiResponse, httpResponse, errorResponse := client.CreateDynamicParsingRule(types.DynamicRuleDefinition{
		Name:    name,
		Scope:   scope,
		Enabled: enabled,
	})
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
