package ingest_budgets_v2

import (
	"github.com/spf13/cobra"
	NewCmdIngestBudgetsV2Create "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets-v2/create"
	NewCmdIngestBudgetsV2Delete "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets-v2/delete"
	NewCmdIngestBudgetsV2Get "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets-v2/get"
	NewCmdIngestBudgetsV2List "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets-v2/list"
	NewCmdIngestBudgetsV2Reset "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets-v2/reset"
	NewCmdIngestBudgetsV2Update "github.com/wizedkyle/sumocli/pkg/cmd/ingest-budgets-v2/update"
)

func NewCmdIngestBudgetsV2() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingest-budgets-v2",
		Short: "Manage ingest budgets (v2)",
		Long:  "Ingest Budgets V2 provide you the ability to create and assign budgets to your log data by Fields instead of using a Field Value.",
	}
	cmd.AddCommand(NewCmdIngestBudgetsV2Create.NewCmdIngestBudgetsV2Create())
	cmd.AddCommand(NewCmdIngestBudgetsV2Delete.NewCmdIngestBudgetsV2Delete())
	cmd.AddCommand(NewCmdIngestBudgetsV2Get.NewCmdIngestBudgetsV2Get())
	cmd.AddCommand(NewCmdIngestBudgetsV2List.NewCmdIngestBudgetsV2List())
	cmd.AddCommand(NewCmdIngestBudgetsV2Reset.NewCmdIngestBudgetsV2Reset())
	cmd.AddCommand(NewCmdIngestBudgetsV2Update.NewCmdIngestBudgetsV2Update())
	return cmd
}
