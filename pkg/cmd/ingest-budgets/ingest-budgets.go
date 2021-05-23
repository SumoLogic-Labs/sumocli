package ingest_budgets

import (
	"github.com/spf13/cobra"
	NewCmdIngestBudgetsAssignCollector "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets/assign-collector"
	NewCmdIngestBudgetsCreate "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets/create"
	NewCmdIngestBudgetsDelete "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets/delete"
	NewCmdIngestBudgetsGet "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets/get"
	NewCmdIngestBudgetsGetAssociatedCollectors "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets/get-associated-collectors"
	NewCmdIngestBudgetsList "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets/list"
	NewCmdIngestBudgetsRemoveCollector "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets/remove-collector"
	NewCmdIngestBudgetsReset "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets/reset"
	NewCmdIngestBudgetsUpdate "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets/update"
)

func NewCmdIngestBudgets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingest-budgets",
		Short: "Manage ingest budgets (v1)",
		Long: "Commands that allow you to manage ingest budgets (v1)." +
			"Ingest Budgets allow you to control the capacity of daily ingestion volume sent to Sumo Logic from Collectors.",
	}
	cmd.AddCommand(NewCmdIngestBudgetsAssignCollector.NewCmdIngestBudgetsAssignCollector())
	cmd.AddCommand(NewCmdIngestBudgetsCreate.NewCmdIngestBudgetsCreate())
	cmd.AddCommand(NewCmdIngestBudgetsDelete.NewCmdIngestBudgetsDelete())
	cmd.AddCommand(NewCmdIngestBudgetsGet.NewCmdIngestBudgetsGet())
	cmd.AddCommand(NewCmdIngestBudgetsGetAssociatedCollectors.NewCmdIngestBudgetsGetAssociatedCollectors())
	cmd.AddCommand(NewCmdIngestBudgetsList.NewCmdIngestBudgetsList())
	cmd.AddCommand(NewCmdIngestBudgetsRemoveCollector.NewCmIngestBudgetsRemoveAssociatedCollector())
	cmd.AddCommand(NewCmdIngestBudgetsReset.NewCmdIngestBudgetsReset())
	cmd.AddCommand(NewCmdIngestBudgetsUpdate.NewCmdIngestBudgetsUpdate())
	return cmd
}
