package ingest_budgets

import (
	NewCmdIngestBudgetsAssignCollector "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets/assign_collector"
	NewCmdIngestBudgetsCreate "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets/create"
	NewCmdIngestBudgetsDelete "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets/delete"
	NewCmdIngestBudgetsGet "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets/get"
	NewCmdIngestBudgetsGetAssociatedCollectors "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets/get_associated_collectors"
	NewCmdIngestBudgetsList "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets/list"
	NewCmdIngestBudgetsRemoveCollector "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets/remove_collector"
	NewCmdIngestBudgetsReset "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets/reset"
	NewCmdIngestBudgetsUpdate "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets/update"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdIngestBudgets(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingest-budgets",
		Short: "Manage ingest budgets (v1)",
		Long: "Commands that allow you to manage ingest budgets (v1)." +
			"Ingest Budgets allow you to control the capacity of daily ingestion volume sent to Sumo Logic from Collectors.",
	}
	cmd.AddCommand(NewCmdIngestBudgetsAssignCollector.NewCmdIngestBudgetsAssignCollector(client))
	cmd.AddCommand(NewCmdIngestBudgetsCreate.NewCmdIngestBudgetsCreate(client))
	cmd.AddCommand(NewCmdIngestBudgetsDelete.NewCmdIngestBudgetsDelete(client))
	cmd.AddCommand(NewCmdIngestBudgetsGet.NewCmdIngestBudgetsGet(client))
	cmd.AddCommand(NewCmdIngestBudgetsGetAssociatedCollectors.NewCmdIngestBudgetsGetAssociatedCollectors(client))
	cmd.AddCommand(NewCmdIngestBudgetsList.NewCmdIngestBudgetsList(client))
	cmd.AddCommand(NewCmdIngestBudgetsRemoveCollector.NewCmIngestBudgetsRemoveAssociatedCollector(client))
	cmd.AddCommand(NewCmdIngestBudgetsReset.NewCmdIngestBudgetsReset(client))
	cmd.AddCommand(NewCmdIngestBudgetsUpdate.NewCmdIngestBudgetsUpdate(client))
	return cmd
}
