package ingest_budgets_v2

import (
	"github.com/spf13/cobra"
	NewCmdIngestBudgetsV2List "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets-v2/list"
)

func NewCmdIngestBudgetsV2() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingest-budgets-v2",
		Short: "Manage ingest budgets (v2)",
		Long:  "Ingest Budgets V2 provide you the ability to create and assign budgets to your log data by Fields instead of using a Field Value.",
	}
	cmd.AddCommand(NewCmdIngestBudgetsV2List.NewCmdIngestBudgetsV2List())
	return cmd
}
