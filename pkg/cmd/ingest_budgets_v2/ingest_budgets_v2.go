package ingest_budgets_v2

import (
	NewCmdIngestBudgetsV2Create "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets_v2/create"
	NewCmdIngestBudgetsV2Delete "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets_v2/delete"
	NewCmdIngestBudgetsV2Get "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets_v2/get"
	NewCmdIngestBudgetsV2List "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets_v2/list"
	NewCmdIngestBudgetsV2Reset "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets_v2/reset"
	NewCmdIngestBudgetsV2Update "github.com/SumoLogic-Labs/sumocli/pkg/cmd/ingest_budgets_v2/update"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdIngestBudgetsV2(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingest-budgets-v2",
		Short: "Manage ingest budgets (v2)",
		Long:  "Ingest Budgets V2 provide you the ability to create and assign budgets to your log data by Fields instead of using a Field Value.",
	}
	cmd.AddCommand(NewCmdIngestBudgetsV2Create.NewCmdIngestBudgetsV2Create(client))
	cmd.AddCommand(NewCmdIngestBudgetsV2Delete.NewCmdIngestBudgetsV2Delete(client))
	cmd.AddCommand(NewCmdIngestBudgetsV2Get.NewCmdIngestBudgetsV2Get(client))
	cmd.AddCommand(NewCmdIngestBudgetsV2List.NewCmdIngestBudgetsV2List(client))
	cmd.AddCommand(NewCmdIngestBudgetsV2Reset.NewCmdIngestBudgetsV2Reset(client))
	cmd.AddCommand(NewCmdIngestBudgetsV2Update.NewCmdIngestBudgetsV2Update(client))
	return cmd
}
