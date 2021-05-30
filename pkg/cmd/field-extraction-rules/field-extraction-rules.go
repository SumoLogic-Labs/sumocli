package field_extraction_rules

import (
	"github.com/spf13/cobra"
	NewCmdFieldExtractionRulesCreate "github.com/wizedkyle/sumocli/pkg/cmd/field-extraction-rules/create"
	NewCmdFieldExtractionRulesDelete "github.com/wizedkyle/sumocli/pkg/cmd/field-extraction-rules/delete"
	NewCmdFieldExtractionRulesGet "github.com/wizedkyle/sumocli/pkg/cmd/field-extraction-rules/get"
	NewCmdFieldExtractionRulesList "github.com/wizedkyle/sumocli/pkg/cmd/field-extraction-rules/list"
	NewCmdFieldExtractionRulesUpdate "github.com/wizedkyle/sumocli/pkg/cmd/field-extraction-rules/update"
)

func NewCmdFieldExtractionRules() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "field-extraction-rules",
		Short: "Manage field extraction rules",
		Long:  "Field Extraction Rules allow you to parse fields from your log messages at the time the messages are ingested eliminating the need to parse fields in your query.",
	}
	cmd.AddCommand(NewCmdFieldExtractionRulesCreate.NewCmdFieldExtractionRulesCreate())
	cmd.AddCommand(NewCmdFieldExtractionRulesDelete.NewCmdFieldExtractionRulesDelete())
	cmd.AddCommand(NewCmdFieldExtractionRulesGet.NewCmdFieldExtractionRulesGet())
	cmd.AddCommand(NewCmdFieldExtractionRulesList.NewCmdFieldExtractionRulesList())
	cmd.AddCommand(NewCmdFieldExtractionRulesUpdate.NewCmdFieldExtractionRulesUpdate())
	return cmd
}
