package ingest_budgets

import (
	"github.com/spf13/cobra"
	NewCmdIngestBudgetsCreate "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets/create"
	NewCmdIngestBudgetsDelete "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets/delete"
	NewCmdIngestBudgetsGet "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets/get"
	NewCmdIngestBudgetsList "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets/list"
	NewCmdIngestBudgetsReset "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets/reset"
)

func NewCmdIngestBudgets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingest-budgets",
		Short: "Manage ingest budgets (v1)",
		Long: "Commands that allow you to manage ingest budgets (v1)." +
			"Ingest Budgets allow you to control the capacity of daily ingestion volume sent to Sumo Logic from Collectors.",
	}
	cmd.AddCommand(NewCmdIngestBudgetsCreate.NewCmdIngestBudgetsCreate())
	cmd.AddCommand(NewCmdIngestBudgetsDelete.NewCmdIngestBudgetsDelete())
	cmd.AddCommand(NewCmdIngestBudgetsGet.NewCmdIngestBudgetsGet())
	cmd.AddCommand(NewCmdIngestBudgetsList.NewCmdIngestBudgetsList())
	cmd.AddCommand(NewCmdIngestBudgetsReset.NewCmdIngestBudgetsReset())
	return cmd
}
