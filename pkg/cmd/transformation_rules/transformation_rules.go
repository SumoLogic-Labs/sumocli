package transformation_rules

import (
	NewCmdTransformationRulesCreate "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/transformation_rules/create"
	NewCmdTransformationRulesDelete "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/transformation_rules/delete"
	NewCmdTransformationRulesGet "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/transformation_rules/get"
	NewCmdTransformationRulesList "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/transformation_rules/list"
	NewCmdTransformationRulesUpdate "github.com/SumoLogic-Incubator/sumocli/pkg/cmd/transformation_rules/update"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdTransformationRules(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transformation-rules",
		Short: "Manage Transformation Rules",
		Long:  "Commands that allow you to manage Transformation Rules in your Sumo Logic tenant.",
	}
	cmd.AddCommand(NewCmdTransformationRulesCreate.NewCmdTransformationRulesCreate(client))
	cmd.AddCommand(NewCmdTransformationRulesDelete.NewCmdTransformationRulesDelete(client))
	cmd.AddCommand(NewCmdTransformationRulesGet.NewCmdTransformationRulesGet(client))
	cmd.AddCommand(NewCmdTransformationRulesList.NewCmdTransformationRulesList(client))
	cmd.AddCommand(NewCmdTransformationRulesUpdate.NewCmdTransformationRulesUpdate(client))
	return cmd
}
