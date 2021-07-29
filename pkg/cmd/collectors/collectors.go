package collectors

import (
	"github.com/spf13/cobra"
	cmdCollectorCreate "github.com/wizedkyle/sumocli/pkg/cmd/collectors/create"
	cmdCollectorDelete "github.com/wizedkyle/sumocli/pkg/cmd/collectors/delete"
	cmdCollectorGet "github.com/wizedkyle/sumocli/pkg/cmd/collectors/get"
	cmdCollectorList "github.com/wizedkyle/sumocli/pkg/cmd/collectors/list"
	cmdCollectorUpdate "github.com/wizedkyle/sumocli/pkg/cmd/collectors/update"
	cmdCollectorUpgrade "github.com/wizedkyle/sumocli/pkg/cmd/collectors/upgrade"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdCollectors(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collectors <command>",
		Short: "Manages collectors",
	}

	cmd.AddCommand(cmdCollectorCreate.NewCmdCollectorCreate(client))
	cmd.AddCommand(cmdCollectorDelete.NewCmdCollectorDelete(client))
	cmd.AddCommand(cmdCollectorGet.NewCmdCollectorGet(client))
	cmd.AddCommand(cmdCollectorList.NewCmdCollectorList(client))
	cmd.AddCommand(cmdCollectorUpdate.NewCmdCollectorUpdate())
	cmd.AddCommand(cmdCollectorUpgrade.NewCmdUpgradeCollectors(client))
	return cmd
}
